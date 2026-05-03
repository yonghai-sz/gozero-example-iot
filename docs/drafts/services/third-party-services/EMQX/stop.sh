#!/usr/bin/env bash
set -euo pipefail

docker container stop my-emqx-container
docker container rm my-emqx-container
