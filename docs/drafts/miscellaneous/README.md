

When working on your **temporary** branch, you can:  
`git push --force-with-lease`  



`ssh -i /Users/chenyonghai/Desktop/SingaporeTest.pem root@47.84.204.211`    
`uname -m`    




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




# 系统架构图或数据流图
# Mermaid 图表
# 接口类型
* 无需校验
* 登录
* jwt 接口
* jwt + RBAC 接口
