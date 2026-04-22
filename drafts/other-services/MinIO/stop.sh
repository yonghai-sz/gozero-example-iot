#!/usr/bin/env bash
set -euo pipefail


docker container stop my-minio-container
docker container rm my-minio-container

