// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"teaching-backend/application/applet/api/internal/config"
	"teaching-backend/application/applet/api/internal/middleware"
	"teaching-backend/application/user/rpc/user"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config           config.Config
	CheckTeacherRole rest.Middleware
	UserRPC          user.User
	BizRedis         *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {

	userRPC := zrpc.MustNewClient(c.UserRPC)
	return &ServiceContext{
		Config:           c,
		CheckTeacherRole: middleware.NewCheckTeacherRoleMiddleware().Handle,
		UserRPC:          user.NewUser(userRPC),
		BizRedis:         redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
	}
}
