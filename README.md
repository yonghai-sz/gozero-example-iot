# gozero-example-iot

Multi-tenant platform (To B)

English | [简体中文](README_zh.md)

## Development

### IDE

**Visual Studio Code** or **Cursor**

### Unit Testing

Run unit tests in a clean Go container:

```bash
bash scripts/test/test.sh
```

### Quick start

Bring everything up:

```bash
bash scripts/dev-or-debug/up.sh
```

Viewing Logs:

```bash
bash scripts/dev-or-debug/logs.sh
```

Stop and remove volumes:

```bash
bash scripts/dev-or-debug/down.sh
```

### Remote debugging

Start the stack in debug mode (same app ports + extra Delve ports):

```bash
bash scripts/dev-or-debug/debug.sh
```

Then attach your Go debugger to:

- `shorturl-api`: `localhost:2345`
- `transform-rpc`: `localhost:2346`

## Deploy

For small projects, Option 1 (Docker Compose) is usually enough. But for large-scale systems, you should go with Option 2 (Kubernetes).

### Option 1: Docker Compose

### Option 2: Kubernetes















### Endpoints
- **Shorten**:
  - `GET http://localhost:8080/shorten?url=https://example.com`
- **Expand**:
  - `GET http://localhost:8080/expand?shorten=<key>`


test the API Gateway service
`curl -i "http://localhost:8080/shorten?url=http://www.example.cn"`
`curl -i "http://localhost:8080/expand?shorten=fb5cd9"`


### Useful ports
- **API**: `8080`
- **RPC**: `8081`
- **MySQL**: `3307` (root password: `zzz888`, database: `myexampledb`)
- **Redis**: `6379`
- **Etcd**: `2379`
