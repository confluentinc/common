#!/usr/bin/env bash
# parity_check.sh — local publish-parity harness for one artifact.
#
# Compares the Bazel-produced artifact against the Maven-produced one, so a human
# reviewer has evidence for the Phase-5 / DP-19123 go/no-go. Copy into the repo
# (e.g. tools-bazel/parity_check.sh) and adapt the TODOs.
#
# Handles two artifact shapes:
#   - JAR artifacts: compares normalized class/resource content (sha256, via
#     normalize_jar.sh) and the generated POM coordinates.
#   - pom-only artifacts (packaging=pom, e.g. a parent POM / BOM): compares the
#     POM coordinates + <dependencies> only (no jar).
# What it does NOT check: byte-identical jars are NOT expected (manifest,
# timestamps, entry order differ — normalize_jar.sh strips/sorts those away).
# --check-signatures is a CI-only stub: it no-ops locally (the session cannot
# sign); real signature/provenance parity is verified in CI.
#
# Prints exactly one final line: `RESULT: PASS|FAIL|BASELINE-UNAVAILABLE <coords>`
# (exit 0 for PASS/BASELINE-UNAVAILABLE, non-zero for FAIL). Exception: a CLI
# usage error (unknown/missing argument) exits 2 with NO RESULT line — it fails
# before any parity work, so there are no <coords> to report; callers must treat
# exit 2 as "bad invocation", distinct from a parity FAIL.
#
# Usage:
#   parity_check.sh --coordinates io.confluent:rest-utils:8.5.0-0 \
#                   --maven-module core \
#                   --bazel-target //core:rest-utils \
#                   [--check-signatures]
set -euo pipefail

COORDS="" MAVEN_MODULE="" BAZEL_TARGET="" CHECK_SIGS=0
while [ $# -gt 0 ]; do
  case "$1" in
    --coordinates)  COORDS="$2"; shift 2;;
    --maven-module) MAVEN_MODULE="$2"; shift 2;;
    --bazel-target) BAZEL_TARGET="$2"; shift 2;;
    --check-signatures) CHECK_SIGS=1; shift;;
    *) echo "unknown arg: $1" >&2; exit 2;;
  esac
done
[ -n "$COORDS" ] && [ -n "$MAVEN_MODULE" ] && [ -n "$BAZEL_TARGET" ] || {
  echo "usage: parity_check.sh --coordinates G:A:V --maven-module DIR --bazel-target //label [--check-signatures]" >&2; exit 2; }

ARTIFACT_ID="$(echo "$COORDS" | cut -d: -f2)"
WORK="$(mktemp -d)"; trap 'rm -rf "$WORK"' EXIT
echo "== parity check for $COORDS =="

# 1) Build the Maven artifact (adjust flags to the repo; -DskipTests keeps it fast).
echo "-- building Maven artifact ($MAVEN_MODULE) --"
# common: skip checkstyle/spotbugs/enforcer — they gate the build but do not
# affect jar contents or the POM, which is what parity compares.
# Capture output so a build failure can be classified: a resolver/auth failure
# (no CodeArtifact creds locally) is an environment limitation and reported as
# BASELINE-UNAVAILABLE (exit 0); any other failure is a genuine defect and must
# surface as RESULT: FAIL — not be waved through as an environment limit.
MVN_BUILD_LOG="$WORK/mvn-build.log"
if ! mvn -q -pl "$MAVEN_MODULE" -am package -DskipTests \
  -Dcheckstyle.skip=true -Dspotbugs.skip=true -Denforcer.skip=true -Dcyclonedx.skip=true \
  > "$MVN_BUILD_LOG" 2>&1; then
  cat "$MVN_BUILD_LOG" >&2
  MVN_RESOLVER_FAILURE_RE='(status code:? *401|Unauthorized|Could not transfer artifact|Could not resolve dependencies|codeartifact)'
  if grep -qiE "$MVN_RESOLVER_FAILURE_RE" "$MVN_BUILD_LOG"; then
    echo "!! Maven baseline unavailable — dependency resolution/auth failure (CodeArtifact creds not available locally)" >&2
    echo "RESULT: BASELINE-UNAVAILABLE $COORDS"; exit 0
  else
    echo "!! Maven baseline build genuinely failed for $MAVEN_MODULE (not a resolver/auth pattern — treat as a real defect)" >&2
    echo "RESULT: FAIL $COORDS"; exit 1
  fi
fi
MVN_JAR="$(ls -1 "$MAVEN_MODULE"/target/"$ARTIFACT_ID"-*.jar 2>/dev/null | grep -vE 'sources|javadoc|tests' | head -n1 || true)"

# Detect pom-only artifacts (parent POM / BOM): no jar, packaging=pom.
POM_ONLY=0
if [ -z "$MVN_JAR" ]; then
  if grep -q "<packaging>pom</packaging>" "$MAVEN_MODULE/pom.xml" 2>/dev/null; then
    POM_ONLY=1
    echo "-- pom-only artifact (packaging=pom): comparing POMs only, no jar --"
  else
    echo "!! Maven baseline jar not found for $MAVEN_MODULE (local resolver limit)" >&2
    echo "RESULT: BASELINE-UNAVAILABLE $COORDS"; exit 0
  fi
fi

# 2) Build the Bazel artifact + its generated POM. Let Bazel print its own
#    errors (no 2>/dev/null) so failures are debuggable. Build the .publish
#    target only if it exists (plain java_library targets have none).
echo "-- building Bazel artifact ($BAZEL_TARGET) --"
# Bazel is expected to always succeed here (it's the migration target, not the
# baseline) — a failure is a genuine defect, so still emit the RESULT: line
# the contract requires before exiting non-zero.
bazel build "$BAZEL_TARGET" || { echo "RESULT: FAIL $COORDS"; exit 1; }
if bazel query "${BAZEL_TARGET}.publish" >/dev/null 2>&1; then
  bazel build "${BAZEL_TARGET}.publish" || { echo "RESULT: FAIL $COORDS"; exit 1; }   # materializes the generated POM
fi
# java_export / maven_export write <name>-pom.xml under bazel-bin.
BZL_POM="$(find -L bazel-bin -name "*$(echo "$BAZEL_TARGET" | sed 's#.*:##')*pom.xml" 2>/dev/null | head -n1 || true)"

# 3a) JAR: compare normalized content (sha256), not raw bytes or entry sets —
# manifest/timestamps/entry order differ even when class/resource bytes match,
# so byte- or listing-identical jars are NOT expected (see header).
if [ "$POM_ONLY" = 0 ]; then
  # No 2>/dev/null on the cquery (matches the lines 55-56 intent): a cquery
  # analysis error must be visible, not swallowed into an empty BZL_JAR.
  BZL_JAR="$(bazel cquery --output=files "$BAZEL_TARGET" | grep -E '\.jar$' | grep -vE 'sources|javadoc' | head -n1 || true)"
  [ -n "$BZL_JAR" ] || { echo "!! could not find Bazel jar for $BAZEL_TARGET" >&2; echo "RESULT: FAIL $COORDS"; exit 1; }
  echo "-- jar content parity (normalized sha256) --"
  NORM="$(dirname "$0")/normalize_jar.sh"
  MVN_SHA="$("$NORM" "$MVN_JAR")" || { echo "!! normalize_jar.sh failed on $MVN_JAR" >&2; echo "RESULT: FAIL $COORDS"; exit 1; }
  BZL_SHA="$("$NORM" "$BZL_JAR")" || { echo "!! normalize_jar.sh failed on $BZL_JAR" >&2; echo "RESULT: FAIL $COORDS"; exit 1; }
  echo "   maven  normalized sha256: $MVN_SHA"
  echo "   bazel  normalized sha256: $BZL_SHA"
  if [ "$MVN_SHA" != "$BZL_SHA" ]; then
    echo "!! jar content parity FAILED for $COORDS" >&2
    echo "RESULT: FAIL $COORDS"; exit 1
  fi
  echo "   OK: normalized jar content matches"
fi

# 3b) POM coordinates + dependency set (both shapes). Fail fast if the Bazel
#     POM wasn't produced — otherwise the comparison would be misleading.
[ -n "$BZL_POM" ] || { echo "!! could not locate the generated Bazel POM — check the <name>-pom / .publish target" >&2; echo "RESULT: FAIL $COORDS"; exit 1; }
MVN_POM="$(ls -1 "$MAVEN_MODULE"/target/*.pom 2>/dev/null | head -n1 || true)"; [ -n "$MVN_POM" ] || MVN_POM="$MAVEN_MODULE/pom.xml"

# Parse POMs with a real XML parser (namespace-agnostic). Reads the PROJECT-level
# groupId/artifactId/version, falling back to <parent> for inherited groupId /
# version, and only top-level <dependencies> (never dependency-nested groupIds).
extract_coords() {
  python3 - "$1" <<'PY'
import sys
# POM inputs are the engineer's own local repo files (trusted), but prefer the
# hardened parser when installed to avoid XXE / entity-expansion on odd inputs.
try:
    from defusedxml.ElementTree import parse as xml_parse
except Exception:
    from xml.etree.ElementTree import parse as xml_parse
def local(t): return t.rsplit('}', 1)[-1]
try:
    root = xml_parse(sys.argv[1]).getroot()
except Exception as e:
    # Fail LOUD (stderr + non-zero) — exiting 0 here makes extract_coords emit an
    # empty capture, and a diff of two empty outputs falsely reports "coordinates
    # match", masking the real failure (e.g. the Bazel POM was never built).
    print("parse-error:", e, file=sys.stderr); sys.exit(2)
def child(el, name):
    for c in list(el):
        if local(c.tag) == name: return c
    return None
def text(el, name):
    c = child(el, name); return c.text.strip() if (c is not None and c.text) else None
parent = child(root, "parent")
gid = text(root, "groupId") or (text(parent, "groupId") if parent is not None else None) or "?"
aid = text(root, "artifactId") or "?"
ver = text(root, "version") or (text(parent, "version") if parent is not None else None) or "?"
print("coords:", gid, aid, ver)
deps = child(root, "dependencies")
seen = set()
if deps is not None:
    for d in list(deps):
        if local(d.tag) != "dependency": continue
        g, a = text(d, "groupId"), text(d, "artifactId")
        if g and a: seen.add(g + ":" + a)
for c in sorted(seen): print("dep:", c)
PY
}
echo "-- POM coordinate + dependency diff (Maven < > Bazel) --"
extract_coords "$MVN_POM" > "$WORK/mvn.pom" || { echo "RESULT: FAIL $COORDS"; exit 1; }
extract_coords "$BZL_POM" > "$WORK/bzl.pom" || { echo "RESULT: FAIL $COORDS"; exit 1; }
echo "   expected coordinates: $COORDS"

# Assert BOTH POMs actually publish the requested G:A:V. Diffing the two POMs
# against each other is not enough: they could agree on the same WRONG artifactId
# or version and still "match". $COORDS is G:A:V; extract_coords emits a
# "coords: <groupId> <artifactId> <version>" line.
EXPECT_COORDS_LINE="coords: $(echo "$COORDS" | awk -F: '{print $1, $2, $3}')"
for side in mvn bzl; do
  if ! grep -Fxq "$EXPECT_COORDS_LINE" "$WORK/$side.pom"; then
    echo "!! $side POM coordinates do not match requested $COORDS" >&2
    echo "   expected: $EXPECT_COORDS_LINE" >&2
    echo "   got:      $(grep '^coords:' "$WORK/$side.pom" || echo '<none>')" >&2
    echo "RESULT: FAIL $COORDS"; exit 1
  fi
done

# Allowlisted, non-failing POM deltas (documented in the spec):
#  1) java_export emits <scope>runtime</scope> where Maven emits compile —
#     extract_coords never prints scope, so this is already invisible below;
#     kept so the filter still applies if scope diffing is added later.
#  2) common-logging's java_export declares 3 jackson artifacts directly —
#     jackson-annotations, jackson-core, jackson-databind (see
#     logging/BUILD.bazel) — because strict-deps needs the explicit edge for
#     direct ObjectMapper / @JsonProperty / @JsonRawValue usage. Maven gets
#     all three transitively via connect-json (none of the three appear in
#     logging/pom.xml's own <dependencies>); common's depMgmt does not pin
#     jackson, so both resolve the same 2.21.x line. Exactly these 3
#     coordinates are allowlisted — not a catch-all jackson pattern.
POM_ALLOWLIST_RE='(<scope>runtime</scope>|com\.fasterxml\.jackson\.core:jackson-(annotations|core|databind))'
grep -vE "$POM_ALLOWLIST_RE" "$WORK/mvn.pom" > "$WORK/mvn.pom.filtered" || true
grep -vE "$POM_ALLOWLIST_RE" "$WORK/bzl.pom" > "$WORK/bzl.pom.filtered" || true
if diff -u "$WORK/mvn.pom.filtered" "$WORK/bzl.pom.filtered"; then
  echo "   OK: POM coordinates + deps match (after allowlisted deltas)"
else
  echo "!! POM coordinate/dependency parity FAILED for $COORDS (review above)" >&2
  echo "RESULT: FAIL $COORDS"; exit 1
fi

if [ "$CHECK_SIGS" -eq 1 ]; then
  echo "   signature parity: SKIPPED (CI-only — local session cannot sign)"
fi

echo "== done. Signing/provenance parity is verified in CI, not here. =="
echo "RESULT: PASS $COORDS"
