
## Who owns which env vars?

### GitHub (secrets)
Use GitHub Environments / repo secrets for **deploy connectivity** and **image publishing**. 

This repo’s deployment workflow uses **GitHub Environments** to separate staging vs production servers.

- Environment secrets (per environment: `staging` / `production`):
  - `DEPLOY_HOST` (server hostname/IP)
  - `DEPLOY_PORT` (optional, defaults to 22)  
  - `DEPLOY_USER` (SSH username)
  - `DEPLOY_SSH_KEY` (private key for SSH)  
  - `DEPLOY_PATH` (working directory on server)
- Repository secrets:
  - `DOCKERHUB_USERNAME`
  - `DOCKERHUB_TOKEN`

### Server (compose `.env`)
Use a **server-side** `.env` file for **runtime/config variables** used by `docker compose` (for example MySQL credentials and host ports).

- Location on server:
  - `$DEPLOY_PATH/needs-to-exist-on-the-server/.env`
- Create it from the template:
  - Copy `scripts/deploy/with-docker-compose/needs-to-exist-on-the-server/.env.example`
  - Edit the values on the server (do not commit real secrets)



## How to set up GitHub Environments / repo secrets

### 1) Create environments
In GitHub: **Settings → Environments**
- Create an environment named `staging`
- Create an environment named `production`

### 2) Add environment secrets
In each environment, add these secrets (same names in both environments, different values):
- `DEPLOY_SSH_KEY` (private key for SSH)
- `DEPLOY_HOST` (server hostname/IP)
- `DEPLOY_USER` (SSH username)
- `DEPLOY_PORT` (optional, defaults to 22)
- `DEPLOY_PATH` (path on server where compose files live)

The workflow also uses repository secrets:
- `DOCKERHUB_USERNAME`
- `DOCKERHUB_TOKEN`

### 3) Require approval for production
In the `production` environment, enable **Required reviewers** so production deploys require manual approval.

### 4) Branch mapping
- Pushes to `develop` deploy to **staging**
- Pushes to `main` deploy to **production**
