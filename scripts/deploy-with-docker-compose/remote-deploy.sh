#!/usr/bin/env bash
set -euo pipefail

DEPLOY_PATH="${DEPLOY_PATH:-/opt/gozero-example-iot}"
cd "$DEPLOY_PATH"

if ! command -v docker >/dev/null 2>&1; then
  echo "docker not found on server" >&2
  exit 1
fi

if ! docker compose version >/dev/null 2>&1; then
  echo "docker compose plugin not available on server" >&2
  exit 1
fi

COMPOSE_FILE="${COMPOSE_FILE:-needs-to-exist-on-the-server/docker-compose.prod.yml}"
if [[ ! -f "$COMPOSE_FILE" ]]; then
  echo "compose file not found: $DEPLOY_PATH/$COMPOSE_FILE" >&2
  exit 1
fi

TAG="${TAG:-latest}"

echo "Deploy path: $DEPLOY_PATH"
echo "Compose file: $COMPOSE_FILE"
echo "Image tag: $TAG"

export TAG

docker compose -f "$COMPOSE_FILE" pull
docker compose -f "$COMPOSE_FILE" up -d --remove-orphans
