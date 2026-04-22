#!/usr/bin/env bash
set -euo pipefail



docker container stop my-mysql-container
docker container rm my-mysql-container



# docker volume prune
# docker volume rm $(docker volume ls -qf dangling=true)
