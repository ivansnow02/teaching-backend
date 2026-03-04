// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package svc

import (
	"teaching-backend/application/ai/rpc/aibridge"
	"teaching-backend/application/applet/api/internal/config"
	"teaching-backend/application/applet/api/internal/middleware"
	"teaching-backend/application/course/rpc/course"
	"teaching-backend/application/exam/rpc/exam"
	"teaching-backend/application/user/rpc/user"
	"teaching-backend/pkg/interceptors"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	CheckTeacherRole    rest.Middleware
	UserRPC             user.User
	CourseRPC           course.Course
	ExamRPC             exam.Exam
	AiRPC               aibridge.AiBridge
	BizRedis            *redis.Redis
	StudyProgressPusher *kq.Pusher
}

func NewServiceContext(c config.Config) *ServiceContext {

	userRPC := zrpc.MustNewClient(c.UserRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	courseRPC := zrpc.MustNewClient(c.CourseRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	examRPC := zrpc.MustNewClient(c.ExamRPC, zrpc.WithUnaryClientInterceptor(interceptors.ClientErrorInterceptor()))
	aiRPC := zrpc.MustNewClient(c.AiRPC)
	return &ServiceContext{
		Config:              c,
		CheckTeacherRole:    middleware.NewCheckTeacherRoleMiddleware().Handle,
		UserRPC:             user.NewUser(userRPC),
		CourseRPC:           course.NewCourse(courseRPC),
		ExamRPC:             exam.NewExam(examRPC),
		AiRPC:               aibridge.NewAiBridge(aiRPC),
		BizRedis:            redis.New(c.BizRedis.Host, redis.WithPass(c.BizRedis.Pass)),
		StudyProgressPusher: kq.NewPusher(c.StudyProgressKafka.Brokers, c.StudyProgressKafka.Topic),
	}
}
