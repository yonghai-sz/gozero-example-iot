#!/usr/bin/env bash
set -euo pipefail

docker container stop my-mongo-container
docker container rm my-mongo-container
