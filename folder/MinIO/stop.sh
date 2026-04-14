#!/bin/bash

set -e


docker container stop my-minio-container
docker container rm my-minio-container

