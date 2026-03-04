package svc

import (
	"teaching-backend/application/ai/rpc/aibridge"
	"teaching-backend/application/course/mq/internal/config"
	"teaching-backend/application/course/rpc/course"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config    config.Config
	Conn      sqlx.SqlConn
	AiRPC     aibridge.AiBridge
	CourseRPC course.Course
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:    c,
		Conn:      sqlx.NewMysql(c.DataSource),
		AiRPC:     aibridge.NewAiBridge(zrpc.MustNewClient(c.AiRPC)),
		CourseRPC: course.NewCourse(zrpc.MustNewClient(c.CourseRPC)),
	}
}
