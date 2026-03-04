package svc

import (
	"teaching-backend/application/ai/rpc/internal/agentclient"
	"teaching-backend/application/ai/rpc/internal/config"
)

type ServiceContext struct {
	Config      config.Config
	AgentClient *agentclient.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:      c,
		AgentClient: agentclient.NewClient(c.AgentRpc),
	}
}
