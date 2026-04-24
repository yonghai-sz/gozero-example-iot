# gozero-example-iot

多租户平台（To B，面向企业）

## 开发

### IDE

**Visual Studio Code** or **Cursor**

### 单元测试

在纯净的 Go 容器中运行单元测试

```bash
bash scripts/test/test.sh
```

### 快速开始

启动所有服务：

```bash
bash scripts/dev-or-debug/up.sh
```

查看日志的方法：

```bash
bash scripts/dev-or-debug/logs.sh
```

停止并移除数据卷：

```bash
bash scripts/dev-or-debug/down.sh
```

### 远程调试

以调试模式启动服务栈（保留原有的应用端口 + 额外开放的 Delve 调试端口）：

```bash
bash scripts/dev-or-debug/debug.sh
```

接着，将你的 Go 调试器连接到以下地址：

- `shorturl-api`: `localhost:2345`
- `transform-rpc`: `localhost:2346`

## 部署

对于小项目，选 Docker Compose 通常就够了；但要是大规模系统，Kubernetes 是标准的部署方案（高可用、弹性伸缩和强大的服务编排能力）。go-zero 框架本身对云原生有非常好的支持。


### 选择一：Docker Compose

### 选择二：K8s













# 系统架构图或数据流图
# Mermaid 图表
# 接口类型
* 无需校验
* 登录
* jwt 接口
* jwt + RBAC 接口
