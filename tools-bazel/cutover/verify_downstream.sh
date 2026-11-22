#!/usr/bin/env bash
# Builds one Maven consumer, unchanged, in one of two modes:
#   --mode bazel  the cut-over repo's Bazel-built artifacts are installed into a
#                 fresh local repository at version V (nano far above any
#                 published build), so the consumer's version ranges pick them
#   --mode maven  same build against what is published today (the control)
# A failure counts against Bazel only if it happens in bazel mode and not in maven mode.
#
# In bazel mode the script then proves the build really used V: if any of the
# repo's artifacts resolved at another version (an exact pin, or a range that
# excludes V), the result is NOT-INJECTED, not PASS. "Resolved" means on a
# classpath, a file a plugin used (jar, zip, tarball), or a POM in the consumer's
# own parent chain. Other-version POMs that only parent a dependency's POM are
# listed, not counted: the consumer's own dependency management decides versions.
#
# Usage:
#   bash verify_downstream.sh --consumer org/repo --mode bazel|maven --version V \
#     --artifacts-repo DIR --group io.confluent [--branch B] [--cmd "mvn ..."] [--work DIR] [--label L]
#   --label: name for result files and the RESULT line (default: the mode), e.g.
#            "mixed" for a bazel-mode run against a mixed repository
#   --artifacts-repo: file:// repository holding the Bazel-built artifacts at V
#                     (dual_publish.sh's m2-bazel, or the CI publish step's output)
#   --expect: every artifactId the cut-over repo publishes (comma-separated), so an
#             artifact Bazel doesn't build yet (e.g. a parent POM) that the consumer
#             resolved from CodeArtifact is caught too. Default: what --artifacts-repo holds.
#   MAVEN_TAIL=~/.m2/repository (env, local runs only): reuse downloads, read-only.
#   --cmd: default is the consumer's cp CI build: mvn <service.yml extra_build_args>
#          -U -B -ntp clean verify. Use e.g. "mvn -B -ntp clean install -DskipTests"
#          for a compile-only tier.
# Prints one line: RESULT: PASS|FAIL|NOT-INJECTED|NOT-CONSUMED <repo> <label> rc=<n> <seconds>s
set -uo pipefail

CONSUMER="" MODE="" VERSION="" ART_REPO="" GROUP="" BRANCH="" CMD="" EXPECT="" LABEL="" WORK="${PWD}/downstream-work"
while [ $# -gt 0 ]; do
  case "$1" in
    --consumer) CONSUMER="$2"; shift 2;;
    --mode) MODE="$2"; shift 2;;
    --version) VERSION="$2"; shift 2;;
    --artifacts-repo) ART_REPO="$2"; shift 2;;
    --group) GROUP="$2"; shift 2;;
    --branch) BRANCH="$2"; shift 2;;
    --cmd) CMD="$2"; shift 2;;
    --expect) EXPECT="$2"; shift 2;;
    --work) WORK="$2"; shift 2;;
    --label) LABEL="$2"; shift 2;;
    *) echo "unknown arg: $1" >&2; exit 2;;
  esac
done
[ -n "$CONSUMER" ] && [ -n "$MODE" ] && [ -n "$VERSION" ] && [ -n "$ART_REPO" ] && [ -n "$GROUP" ] || {
  echo "usage: verify_downstream.sh --consumer org/repo --mode bazel|maven --version V --artifacts-repo DIR --group G [--branch B] [--cmd ...]" >&2; exit 2; }
case "$MODE" in bazel|maven) ;; *) echo "--mode must be bazel or maven" >&2; exit 2;; esac

NAME="${CONSUMER##*/}"; LABEL="${LABEL:-$MODE}"
mkdir -p "$WORK/results"; WORK="$(cd "$WORK" && pwd)"; ART_REPO="$(cd "$ART_REPO" && pwd)"
SRC="$WORK/src/$NAME-$LABEL" LR="$WORK/lr/$NAME-$LABEL" LOG="$WORK/results/$NAME-$LABEL.log"
rm -rf "$SRC" "$LR"; mkdir -p "$LR" "$(dirname "$SRC")"; : > "$LOG"
GROUP_PATH="$(echo "$GROUP" | tr . /)"
# Artifacts the cut-over repo publishes = what the file:// repository holds at V.
ARTIFACTS="$(cd "$ART_REPO/$GROUP_PATH" && for a in */; do [ -d "$a$VERSION" ] && echo "${a%/}"; done)"
[ -n "$ARTIFACTS" ] || { echo "!! no $GROUP artifacts at $VERSION in $ART_REPO" >&2; exit 2; }
CHECK="$( (echo "$ARTIFACTS"; echo "$EXPECT" | tr ',' '\n') | grep -v '^$' | sort -u)"
TAILARG=""; [ -n "${MAVEN_TAIL:-}" ] && TAILARG="-Dmaven.repo.local.tail=$MAVEN_TAIL"

git clone -q --depth 1 ${BRANCH:+-b "$BRANCH"} "https://github.com/$CONSUMER.git" "$SRC" \
  || { echo "RESULT: FAIL $CONSUMER $LABEL rc=clone 0s"; exit 1; }

install_file() {  # quiet install:install-file into $LR
  mvn -B -q -ntp org.apache.maven.plugins:maven-install-plugin:3.1.4:install-file \
    -Dmaven.repo.local="$LR" "$@" >> "$LOG" 2>&1
}
if [ "$MODE" = bazel ]; then
  for a in $ARTIFACTS; do
    d="$ART_REPO/$GROUP_PATH/$a/$VERSION"; pom="$d/$a-$VERSION.pom"
    if [ -f "$d/$a-$VERSION.jar" ]; then
      install_file -Dfile="$d/$a-$VERSION.jar" -DpomFile="$pom"
    else
      install_file -Dfile="$pom" -DpomFile="$pom" -Dpackaging=pom
    fi
    # Every attached file: <a>-<V>-<classifier>.<ext> (sources/tests jars, resources zips, tarballs).
    for f in "$d/$a-$VERSION"-*; do
      [ -e "$f" ] || continue
      case "$f" in *.sha1|*.sha256|*.sha512|*.md5|*.asc) continue;; esac
      rest="${f#"$d/$a-$VERSION"-}"
      case "$rest" in *.tar.gz) ext="tar.gz"; c="${rest%.tar.gz}";; *) ext="${rest##*.}"; c="${rest%.*}";; esac
      install_file -Dfile="$f" -DgroupId="$GROUP" -DartifactId="$a" -Dversion="$VERSION" \
        -Dpackaging="$ext" -Dclassifier="$c" -DgeneratePom=false
    done
  done
fi

if [ -z "$CMD" ]; then
  EXTRA="$(sed -n 's/^[[:space:]]*extra_build_args:[[:space:]]*"\{0,1\}\([^"]*\)"\{0,1\}[[:space:]]*$/\1/p' "$SRC/service.yml" 2>/dev/null | head -1)"
  CMD="mvn $EXTRA -U -B -ntp clean verify"
fi
START=$(date +%s)
echo "\$ $CMD -Dmaven.repo.local=$LR $TAILARG" >> "$LOG"
(cd "$SRC" && eval "$CMD -Dmaven.repo.local=\"$LR\" $TAILARG") >> "$LOG" 2>&1
RC=$?
SECS=$(( $(date +%s) - START ))

# The consumer's own parent chain as Maven resolved it (groupId:artifactId:pom:version).
LINEAGE="" EXPR=project.parent
for _ in 1 2 3 4 5 6; do
  id="$(cd "$SRC" && mvn -B -q -N org.apache.maven.plugins:maven-help-plugin:3.5.1:evaluate \
    -Dexpression="$EXPR.id" -DforceStdout -Dmaven.repo.local="$LR" $TAILARG 2>>"$LOG")"
  echo "$id" | grep -qE '^[A-Za-z0-9_.-]+:[A-Za-z0-9_.-]+:[A-Za-z0-9_-]+:[^ :]+$' || break
  LINEAGE="$LINEAGE $id"; EXPR="$EXPR.parent"
done

# Which versions of our artifacts did the build actually resolve? A version other
# than V in the local repository was downloaded from a remote. It counts against
# injection if it holds more than a POM (a jar, zip or tarball was used) or is in
# the consumer's own parent chain. Other POM-only versions parent a dependency's
# POM; they are listed but don't decide this build's versions.
# What was used at V comes from the consumer's own dependency list (same local
# repository, so the same resolution), since install-file put every V artifact there.
OTHER="" POMONLY=""
for a in $CHECK; do
  [ -d "$LR/$GROUP_PATH/$a" ] || continue
  for vd in "$LR/$GROUP_PATH/$a"/*/; do
    v="$(basename "$vd")"
    [ "$v" = "$VERSION" ] && continue
    files="$(ls "$vd" 2>/dev/null | grep -vE '\.(sha1|sha256|sha512|md5|asc|lastUpdated)$|^_remote\.repositories$|^resolver-status\.properties$')"
    [ -n "$files" ] || continue
    if echo "$files" | grep -qvE '\.pom$' || echo "$LINEAGE" | tr ' ' '\n' | grep -qx "$GROUP:$a:pom:$v"; then
      OTHER="$OTHER $a:$v"
    else
      POMONLY="$POMONLY $a:$v"
    fi
  done
done
DEPS="$WORK/results/$NAME-$LABEL.deps"; rm -f "$DEPS"
(cd "$SRC" && mvn -B -q org.apache.maven.plugins:maven-dependency-plugin:3.8.1:list -DoutputFile="$DEPS" \
  -DappendOutput=true -Dmaven.repo.local="$LR" $TAILARG) >> "$LOG" 2>&1
USED="$(grep -hoE "$GROUP:[^:]+:[^ ]*$VERSION" "$DEPS" 2>/dev/null | cut -d: -f2 | sort -u | tr '\n' ' ')"
# Dependencies of ours at another version (also catches ones read from MAVEN_TAIL,
# which never land in $LR; parent POMs are only caught without a tail, as in CI).
for line in $(grep -hoE "$GROUP:[^ ]+" "$DEPS" 2>/dev/null | sort -u); do
  a="$(echo "$line" | cut -d: -f2)"; v="$(echo "$line" | awk -F: '{print $(NF-1)}')"
  echo "$CHECK" | grep -qx "$a" && [ "$v" != "$VERSION" ] && OTHER="$OTHER $a:$v"
done
OTHER="$(echo "$OTHER" | tr ' ' '\n' | grep -v '^$' | sort -u | tr '\n' ' ')"
{
  echo "resolved at $VERSION:${USED:- none}"
  echo "resolved at other versions:${OTHER:- none}"
  echo "own parent chain:${LINEAGE:- none}"
  echo "POMs of dependencies' parents at other versions (not counted):${POMONLY:- none}"
  echo "failing test classes:"
  find "$SRC" \( -path '*/target/surefire-reports/TEST-*.xml' -o -path '*/target/failsafe-reports/TEST-*.xml' \) \
    -exec grep -lE '<(failure|error)[ >]' {} + 2>/dev/null \
    | sed 's|.*/TEST-||; s|\.xml$||' | sort -u | sed 's/^/  /'
} > "$WORK/results/$NAME-$LABEL.summary"

if [ "$MODE" = bazel ] && [ -n "$OTHER" ]; then STATUS=NOT-INJECTED
elif [ "$MODE" = bazel ] && [ -z "$USED" ] && [ "$RC" -eq 0 ]; then STATUS=NOT-CONSUMED
elif [ "$RC" -eq 0 ]; then STATUS=PASS
else STATUS=FAIL; fi
echo "RESULT: $STATUS $CONSUMER $LABEL rc=$RC ${SECS}s" | tee -a "$WORK/results/$NAME-$LABEL.summary"
[ "$STATUS" = PASS ]
