#!/usr/bin/env bash
set -euo pipefail

docker run \
  --rm \
  -t \
  -v "$(pwd):/src" \
  -e GOPROXY=https://goproxy.cn,direct \
  -w /src \
  golang:1.24 \
  bash -lc 'go test ./...'
