#!/usr/bin/env bash
set -euo pipefail

docker container stop my-redis-container
docker container rm my-redis-container


