package logic

import (
	"context"
	"database/sql"
	"strconv"
	"teaching-backend/application/exam/rpc/model"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateQuestionLogic {
	return &CreateQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ========== 题库 management ==========
func (l *CreateQuestionLogic) CreateQuestion(in *pb.CreateQuestionReq) (*pb.CreateQuestionRes, error) {
	score, _ := strconv.ParseFloat(in.Score, 64)

	question := &model.Question{
		CourseId:        uint64(in.CourseId),
		TeacherId:       uint64(in.TeacherId),
		Type:            int64(in.Type),
		Content:         in.Content,
		Answer:          in.Answer,
		Analysis:        sql.NullString{String: in.Analysis, Valid: in.Analysis != ""},
		KnowledgePoints: in.KnowledgePoints,
		Score:           score,
		Difficulty:      int64(in.Difficulty),
	}

	res, err := l.svcCtx.QuestionModel.Insert(l.ctx, question)
	if err != nil {
		l.Errorf("QuestionModel.Insert error: %v", err)
		return nil, xcode.ServerErr
	}
	id, _ := res.LastInsertId()

	return &pb.CreateQuestionRes{
		Id: int64(id),
	}, nil
}
