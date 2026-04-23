# gozero-example-iot



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






## Development

### Quick start

Bring everything up:

```bash
bash scripts/dev/docker-compose/up.sh
```

### Remote debugging (Delve)

Start the stack in debug mode (same app ports + extra Delve ports):

```bash
bash scripts/dev/docker-compose/debug.sh
```

Then attach your Go debugger to:

- `shorturl-api`: `localhost:2345`
- `transform-rpc`: `localhost:2346`

#### Note for Apple Silicon (arm64) Macs

If you build the debug images as `amd64` on an arm64 Mac, the Go binaries may run under Rosetta emulation (you'll see `/run/rosetta/rosetta` in `ps`). In that case Delve can attach, but **breakpoints may not stop** reliably.

Fix: build/run the debug images as **native arm64** (do not force `GOARCH=amd64`).

### Logs / stop

Logs:

```bash
bash scripts/dev/docker-compose/logs.sh
```

Stop and remove volumes:

```bash
bash scripts/dev/docker-compose/down.sh
```

### tests

Run unit tests in a clean Go container:

```bash
bash scripts/dev/test.sh
```



