// Code scaffolded by goctl. Safe to edit.
// goctl 1.10.1

package svc

import (
	"gozero-example-iot/services/api/internal/config"
	"gozero-example-iot/services/rpc-transform/transformer"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config      config.Config
	Transformer transformer.Transformer
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		Transformer: transformer.NewTransformer(zrpc.MustNewClient(c.Transform)),
	}
}
