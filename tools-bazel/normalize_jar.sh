#!/usr/bin/env bash
# normalize_jar.sh — emit a content sha256 for one jar, independent of build
# envelope. Class/resource BYTES are compared verbatim; only envelope is
# normalized away:
#   - META-INF/maven/**            dropped (Maven-injected build metadata)
#   - MANIFEST.MF volatile lines   stripped, remaining lines sorted
#   - zip timestamps / entry order irrelevant (we hash extracted file contents)
# Usage: normalize_jar.sh <path-to-jar>  ->  prints 64-char hex sha256
set -euo pipefail
JAR="$1"
WORK="$(mktemp -d)"; trap 'rm -rf "$WORK"' EXIT

# Portable sha256: GNU coreutils `sha256sum` (Linux/CI) or `shasum -a 256`
# (a stock macOS ships shasum, not sha256sum). Both print "<hash>  <name>".
_sha256() {
  if command -v sha256sum >/dev/null 2>&1; then
    sha256sum "$@"
  else
    shasum -a 256 "$@"
  fi
}
unzip -q -o "$JAR" -d "$WORK"

# Drop Maven-injected metadata.
rm -rf "$WORK/META-INF/maven"

# Normalize the manifest: strip volatile lines, sort the remainder.
MF="$WORK/META-INF/MANIFEST.MF"
if [ -f "$MF" ]; then
  grep -avE '^(Build-Jdk|Build-Jdk-Spec|Created-By|Built-By|Bnd-LastModified|Tool|Archiver-Version|Commit-ID):' "$MF" \
    | sed '/^[[:space:]]*$/d' | LC_ALL=C sort > "$MF.norm"
  mv "$MF.norm" "$MF"
fi

# Digest: "relpath  sha256(content)" for every file, sorted, then hash the list.
( cd "$WORK" && find . -type f | LC_ALL=C sort | while read -r f; do
    printf '%s  %s\n' "${f#./}" "$(_sha256 "$f" | cut -d' ' -f1)"
  done ) | _sha256 | cut -d' ' -f1
