#!/usr/bin/env bash

set -o nounset
set -o pipefail

ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"

GITHUB_REF=${GITHUB_REF:-""}

if [[ -z "$GITHUB_REF" ]]; then
  echo "GITHUB_REF is empty"
  exit -1
fi

SCOOP_VERSION=${GITHUB_REF#refs/tags/}
SCOOP_VERSION=${SCOOP_VERSION#v}

SCOOP_HASH="$(sha256sum dist/go-musicfox_3.7.6_windows_amd64.zip)"

eval "cat <<EOF
$(< "$ROOT"/deploy/scoop/go-musicfox.json.tpl)
EOF
"  > "$ROOT"/deploy/scoop/go-musicfox.json