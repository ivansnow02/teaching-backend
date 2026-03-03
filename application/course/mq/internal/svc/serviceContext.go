package svc

import (
	"teaching-backend/application/course/mq/internal/config"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config config.Config
	Conn   sqlx.SqlConn
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		Conn:   sqlx.NewMysql(c.DataSource),
	}
}
