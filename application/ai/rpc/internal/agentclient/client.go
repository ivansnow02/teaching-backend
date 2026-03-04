// Package agentclient 封装对 Python gRPC Agent 的客户端调用
package agentclient

import (
	"teaching-backend/application/ai/rpc/agentpb"

	"github.com/zeromicro/go-zero/zrpc"
)

// Client 封装 gRPC 连接，提供 AgentServiceClient
type Client struct {
	inner zrpc.Client
}

func NewClient(conf zrpc.RpcClientConf) *Client {
	return &Client{inner: zrpc.MustNewClient(conf)}
}

// Svc 返回 protobuf 生成的 gRPC client，可直接调用所有 AgentService 方法
func (c *Client) Svc() agentpb.AgentServiceClient {
	return agentpb.NewAgentServiceClient(c.inner.Conn())
}
