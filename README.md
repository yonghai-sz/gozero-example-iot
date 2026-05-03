# gozero-example-iot

Multi-tenant platform (To B)

English | [简体中文](README_zh.md)

## Clone

This repository uses [Git submodules](https://git-scm.com/book/en/v2/Git-Tools-Submodules) for shared Cursor rules and AI context (`.cursorrules`, `.github/ai-context`, `.ai-context/zero-skills`).

```bash
git clone --recurse-submodules <repository-url>
```

If you already cloned without submodules:

```bash
git submodule update --init --recursive
```

## Development

### IDE

**Cursor** or **Visual Studio Code**

### Unit Testing

Run unit tests in a clean Go container:

```bash
bash scripts/test/test.sh
```

### Quick start

Bring everything up:

```bash
bash scripts/dev/up.sh
```

Viewing Logs:

```bash
bash scripts/dev/logs.sh
```

Stop and remove volumes:

```bash
bash scripts/dev/down.sh
```

### Remote debugging

Start the stack in debug mode (same app ports + extra Delve ports):

```bash
bash scripts/dev/debug.sh
```

Then attach your Go debugger to:

- `shorturl-api`: `localhost:2345`
- `transform-rpc`: `localhost:2346`

## Deploy

For small to medium-sized projects, Docker Compose is usually sufficient. However, for large-scale systems, Kubernetes is the standard deployment solution due to its high availability, auto-scaling, and powerful orchestration capabilities. Fortunately, go-zero has great built-in support for cloud-native setups.

### Option 1: Docker Compose

### Option 2: Kubernetes



