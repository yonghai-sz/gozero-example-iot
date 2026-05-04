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

### Running Locally

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

### Debugging

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

#### GitHub

In GitHub: **Settings → Environments**  
- Create an environment named `staging`
- Create an environment named `production`

Add environment secrets   
- Environment secrets (per environment: `staging` / `production`):
  - `DEPLOY_HOST` (server hostname/IP)
  - `DEPLOY_PORT` (optional, defaults to 22)  
  - `DEPLOY_USER` (SSH username)
  - `DEPLOY_SSH_KEY` (private key for SSH)  
  - `DEPLOY_PATH` (working directory on server)
- Repository secrets:
  - `DOCKERHUB_USERNAME`
  - `DOCKERHUB_TOKEN`

In the `production` environment, enable **Required reviewers** so production deploys require manual approval.

Branch mapping:  
- Pushes to `develop` deploy to **staging**
- Pushes to `main` deploy to **production**

#### Server

Place `scripts/deploy/with-docker-compose/needs-to-exist-on-the-server` folder in the `$DEPLOY_PATH/` directory on the server.

In `needs-to-exist-on-the-server` directory, `mv .env.example .env`, and edit the values in `.env` file.

After changing `.env`, update `config/*.yaml` on the server so those YAML files stay aligned with the same hostnames, ports, credentials, and other settings reflected in `.env`.

### Option 2: Kubernetes



