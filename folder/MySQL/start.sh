#!/bin/bash

set -e


# docker volume create --name mysql-data-vol

# -v /opt/docker_v/mysql-conf:/etc/mysql/conf.d

docker run \
  -d \
  --name my-mysql-container \
  --network dev-network \
  -p 3306:3306 \
  -e MYSQL_ROOT_PASSWORD=zsm9ccbE \
  -e MYSQL_DATABASE=example \
  -v mysql-data-vol:/var/lib/mysql \
  mysql:8.0



