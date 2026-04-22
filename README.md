# gozero-example-iot

## 说明

多租户平台（To B，面向企业）

# 部署

开发测试环境： Docker Compose
生产环境：Kubernetes



在开发、调试和测试阶段，使用 Docker Compose 是一种非常便捷的方式。它能让你通过一个简单的 YAML 文件来定义和运行所有服务（包括 go-zero 应用以及 MySQL、Redis 等依赖），实现一键启动整个环境。相比 K8s 集群，在单机上运行 Compose 对系统资源的要求更低。

当项目进入生产阶段，需要高可用、弹性伸缩和强大的服务编排能力时，Kubernetes 是标准的部署方案。go-zero 框架本身对云原生有非常好的支持。


在本地使用 Docker Compose 进行快速开发和功能验证，然后通过 CI/CD 流水线将应用打包成镜像，并最终部署到生产环境的 Kubernetes 集群中。



### Quick start

Bring everything up:

```bash
bash scripts/docker-compose/up.sh
```

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

### Remote debugging (Delve)

Start the stack in debug mode (same app ports + extra Delve ports):

```bash
bash scripts/docker-compose/debug.sh
```

Then attach your Go debugger to:

- `shorturl-api`: `localhost:2345`
- `transform-rpc`: `localhost:2346`

### Logs / stop / tests

Logs:

```bash
bash scripts/docker-compose/logs.sh
```

Stop and remove volumes:

```bash
bash scripts/docker-compose/down.sh
```

Run unit tests in a clean Go container:

```bash
bash scripts/docker-compose/test.sh
```



