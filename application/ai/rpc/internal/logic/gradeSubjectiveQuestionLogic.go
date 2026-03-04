package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeSubjectiveQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGradeSubjectiveQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeSubjectiveQuestionLogic {
	return &GradeSubjectiveQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GradeSubjectiveQuestionLogic) GradeSubjectiveQuestion(in *pb.GradeSubjectiveQuestionReq) (*pb.GradeSubjectiveQuestionRes, error) {
	if in.Input == nil {
		return nil, code.AiGradingFailed
	}
	res, err := l.svcCtx.AgentClient.Svc().GradeSubjective(l.ctx, &agentpb.GradeSubjectiveReq{
		RecordId:        in.Input.RecordId,
		QuestionId:      in.Input.QuestionId,
		QuestionContent: in.Input.QuestionContent,
		StandardAnswer:  in.Input.StandardAnswer,
		KnowledgePoints: in.Input.KnowledgePoints,
		UserAnswer:      in.Input.UserAnswer,
		MaxScore:        in.Input.MaxScore,
	})
	if err != nil {
		l.Errorf("GradeSubjectiveQuestion question_id=%d error: %v", in.Input.QuestionId, err)
		return nil, code.AiGradingFailed
	}
	return &pb.GradeSubjectiveQuestionRes{
		QuestionId: res.QuestionId,
		Score:      res.Score,
		IsCorrect:  res.IsCorrect,
		AiComment:  res.AiComment,
	}, nil
}
