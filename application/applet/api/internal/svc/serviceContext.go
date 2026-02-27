// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"github.com/zeromicro/go-zero/rest"
	"teaching-backend/application/applet/api/internal/config"
	"teaching-backend/application/applet/api/internal/middleware"
)

type ServiceContext struct {
	Config           config.Config
	CheckTeacherRole rest.Middleware
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:           c,
		CheckTeacherRole: middleware.NewCheckTeacherRoleMiddleware().Handle,
	}
}
