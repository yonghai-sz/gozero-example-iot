# gozero-example-iot

多租户平台（To B，面向企业）

## Clone

本仓库使用 [Git 子模块](https://git-scm.com/book/zh/v2/Git-%E5%B7%A5%E5%85%B7-%E5%AD%90%E6%A8%A1%E5%9D%97) 管理共用的 Cursor 规则与 AI 上下文（`.cursorrules`、`.github/ai-context`、`.ai-context/zero-skills`）。

```bash
git clone --recurse-submodules <仓库地址>
```

若已克隆但未拉取子模块，在仓库根目录执行：

```bash
git submodule update --init --recursive
```

## 开发

### IDE

**Cursor** or **Visual Studio Code**

### 单元测试

在纯净的 Go 容器中运行单元测试

```bash
bash scripts/test/test.sh
```

### 快速开始

启动所有服务：

```bash
bash scripts/dev/up.sh
```

查看日志的方法：

```bash
bash scripts/dev/logs.sh
```

停止并移除数据卷：

```bash
bash scripts/dev/down.sh
```

### 远程调试

以调试模式启动服务栈（保留原有的应用端口 + 额外开放的 Delve 调试端口）：

```bash
bash scripts/dev/debug.sh
```

接着，将你的 Go 调试器连接到以下地址：

- `shorturl-api`: `localhost:2345`
- `transform-rpc`: `localhost:2346`

## 部署

对于中小型项目，选 Docker Compose 通常就够了；但要是大规模系统，Kubernetes 是标准的部署方案（高可用、弹性伸缩和强大的服务编排能力）。go-zero 框架本身对云原生有非常好的支持。

### 选择一：Docker Compose

### 选择二：Kubernetes


