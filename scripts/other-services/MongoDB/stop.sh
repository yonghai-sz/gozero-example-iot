#!/bin/bash

set -e

docker container stop my-mongo-container
docker container rm my-mongo-container
