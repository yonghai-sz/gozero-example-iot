#!/usr/bin/env bash
set -euo pipefail

docker run \
  --rm \
  -t \
  -v "$(pwd):/src" \
  -w /src \
  golang:1.26 \
  bash -lc 'go test ./...'
