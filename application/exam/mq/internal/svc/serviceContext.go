package svc

import (
	"teaching-backend/application/exam/mq/internal/config"
	"teaching-backend/application/exam/rpc/exam"
	"teaching-backend/application/exam/rpc/model"

	"github.com/zeromicro/go-zero/core/stores/sqlx"
	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config              config.Config
	Conn                sqlx.SqlConn
	ExamRPC             exam.Exam
	QuestionModel       model.QuestionModel
	ExamModel           model.ExamModel
	ExamQuestionModel   model.ExamQuestionModel
	UserExamRecordModel model.UserExamRecordModel
	UserAnswerModel     model.UserAnswerModel
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := sqlx.NewMysql(c.DataSource)
	return &ServiceContext{
		Config:              c,
		Conn:                conn,
		ExamRPC:             exam.NewExam(zrpc.MustNewClient(c.ExamRPC)),
		QuestionModel:       model.NewQuestionModel(conn, c.CacheRedis),
		ExamModel:           model.NewExamModel(conn, c.CacheRedis),
		ExamQuestionModel:   model.NewExamQuestionModel(conn, c.CacheRedis),
		UserExamRecordModel: model.NewUserExamRecordModel(conn, c.CacheRedis),
		UserAnswerModel:     model.NewUserAnswerModel(conn, c.CacheRedis),
	}
}
