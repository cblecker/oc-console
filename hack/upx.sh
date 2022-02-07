#!/usr/bin/env bash
set -o errexit
set -o nounset
set -o pipefail

if ! command -v upx; then
  echo "upx not available. skipping.."
  exit 0
fi

upx dist/oc-console*/oc-console*
