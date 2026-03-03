package svc

import (
	"teaching-backend/application/exam/rpc/internal/config"
	"teaching-backend/application/exam/rpc/internal/model"

	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/stores/sqlx"
)

type ServiceContext struct {
	Config              config.Config
	QuestionModel       model.QuestionModel
	ExamModel           model.ExamModel
	ExamQuestionModel   model.ExamQuestionModel
	UserExamRecordModel model.UserExamRecordModel
	UserAnswerModel     model.UserAnswerModel
	BizRedis            *redis.Redis
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:              c,
		QuestionModel:       model.NewQuestionModel(conn, c.CacheRedis),
		ExamModel:           model.NewExamModel(conn, c.CacheRedis),
		ExamQuestionModel:   model.NewExamQuestionModel(conn, c.CacheRedis),
		UserExamRecordModel: model.NewUserExamRecordModel(conn, c.CacheRedis),
		UserAnswerModel:     model.NewUserAnswerModel(conn, c.CacheRedis),
		BizRedis:            redis.MustNewRedis(c.BizRedis.RedisConf),
	}
}
