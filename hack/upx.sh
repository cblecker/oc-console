#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

if [[ -n ${GITHUB_ACTION:-} ]]; then
  apk --no-cache add upx
fi
if ! command -v upx; then
  echo "upx not available. skipping.."
  exit 0
fi

upx dist/oc-console*/oc-console*
