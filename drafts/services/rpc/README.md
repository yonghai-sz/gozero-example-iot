
`goctl rpc template -o transform.proto`

`goctl rpc protoc transform.proto --go_out=./pb --go-grpc_out=./pb --zrpc_out=.`



要修改的地方：  
* etc/example.yaml
* internal/config/config.go
* internal/svc/servicecontext.go
* internal/logic/*.go


