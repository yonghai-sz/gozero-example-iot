#!/usr/bin/env bash
set -euo pipefail

DIR="$( cd "$( dirname "${BASH_SOURCE[0]}" )" && pwd )"

cd "$DIR"

docker compose \
  -f docker-compose.yml \
  -f docker-compose.debug.yml \
  up \
  -d --build
