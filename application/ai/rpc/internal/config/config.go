package config

import (
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	zrpc.RpcServerConf
	AgentRpc zrpc.RpcClientConf // Python Agent gRPC 服务地址
}
