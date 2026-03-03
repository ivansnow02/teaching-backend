package logic

import (
	"context"
	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteQuestionLogic {
	return &DeleteQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除题目
func (l *DeleteQuestionLogic) DeleteQuestion(in *pb.DeleteQuestionReq) (*pb.DeleteQuestionRes, error) {
	question, err := l.svcCtx.QuestionModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("QuestionModel.FindOne error: %v", err)
		return nil, code.QuestionNotFound
	}

	if question.TeacherId != uint64(in.OperatorId) {
		return nil, code.NoPermission
	}

	err = l.svcCtx.QuestionModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("QuestionModel.Delete error: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.DeleteQuestionRes{}, nil
}
