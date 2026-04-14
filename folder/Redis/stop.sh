#!/bin/bash

set -e

docker container stop my-redis-container
docker container rm my-redis-container


