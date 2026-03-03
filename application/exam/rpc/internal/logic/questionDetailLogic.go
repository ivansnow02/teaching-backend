package logic

import (
	"context"
	"fmt"
	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuestionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionDetailLogic {
	return &QuestionDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 题目详情
func (l *QuestionDetailLogic) QuestionDetail(in *pb.QuestionDetailReq) (*pb.QuestionDetailRes, error) {
	question, err := l.svcCtx.QuestionModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("QuestionModel.FindOne error: %v", err)
		return nil, code.QuestionNotFound
	}

	return &pb.QuestionDetailRes{
		Question: &pb.QuestionItem{
			Id:              int64(question.Id),
			CourseId:        int64(question.CourseId),
			TeacherId:       int64(question.TeacherId),
			Type:            int32(question.Type),
			Content:         question.Content,
			Answer:          question.Answer,
			Analysis:        question.Analysis.String,
			KnowledgePoints: question.KnowledgePoints,
			Score:           fmt.Sprintf("%.1f", question.Score),
			Difficulty:      int32(question.Difficulty),
			CreateTime:      question.CreateTime.Unix(),
		},
	}, nil
}
