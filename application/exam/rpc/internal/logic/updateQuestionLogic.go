package logic

import (
	"context"
	"database/sql"
	"strconv"
	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateQuestionLogic {
	return &UpdateQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新题目
func (l *UpdateQuestionLogic) UpdateQuestion(in *pb.UpdateQuestionReq) (*pb.UpdateQuestionRes, error) {
	question, err := l.svcCtx.QuestionModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("QuestionModel.FindOne error: %v", err)
		return nil, code.QuestionNotFound
	}

	if question.TeacherId != uint64(in.OperatorId) {
		return nil, code.NoPermission
	}

	score, _ := strconv.ParseFloat(in.Score, 64)
	question.Content = in.Content
	question.Answer = in.Answer
	question.Analysis = sql.NullString{String: in.Analysis, Valid: in.Analysis != ""}
	question.KnowledgePoints = in.KnowledgePoints
	question.Score = score
	question.Difficulty = int64(in.Difficulty)

	err = l.svcCtx.QuestionModel.Update(l.ctx, question)
	if err != nil {
		l.Errorf("QuestionModel.Update error: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.UpdateQuestionRes{}, nil
}
