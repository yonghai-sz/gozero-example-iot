#!/bin/bash

# docker volume create --name redis-data-vol

docker run \
  -d \
  --name my-redis-container \
  --network dev-network \
  -p 6379:6379 \
  -v redis-data-vol:/data \
  redis \
  redis-server --save 30 1 \
  --requirepass "zzsskja"


# save a snapshot of the DB every 30 seconds 
# if at least 1 write operation was performed.
# it will also lead to more logs, so the loglevel option may be desirable     
# --loglevel warning \


