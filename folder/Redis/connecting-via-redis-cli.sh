#!/bin/bash

docker run \
  -it \
  --network prod_service-network \
  --rm \
  redis \
  redis-cli -h my-redis-container -p 26379 -a my9passwordz


