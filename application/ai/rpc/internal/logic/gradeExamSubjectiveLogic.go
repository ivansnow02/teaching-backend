package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GradeExamSubjectiveLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGradeExamSubjectiveLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GradeExamSubjectiveLogic {
	return &GradeExamSubjectiveLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GradeExamSubjectiveLogic) GradeExamSubjective(in *pb.GradeExamSubjectiveReq) (*pb.GradeExamSubjectiveRes, error) {
	results := make([]*pb.GradeSubjectiveQuestionRes, 0, len(in.Questions))
	for _, q := range in.Questions {
		res, err := l.svcCtx.AgentClient.Svc().GradeSubjective(l.ctx, &agentpb.GradeSubjectiveReq{
			RecordId:        q.RecordId,
			QuestionId:      q.QuestionId,
			QuestionContent: q.QuestionContent,
			StandardAnswer:  q.StandardAnswer,
			KnowledgePoints: q.KnowledgePoints,
			UserAnswer:      q.UserAnswer,
			MaxScore:        q.MaxScore,
		})
		if err != nil {
			l.Errorf("GradeExamSubjective question_id=%d error: %v", q.QuestionId, err)
			return nil, code.AiGradingFailed
		}
		results = append(results, &pb.GradeSubjectiveQuestionRes{
			QuestionId: res.QuestionId,
			Score:      res.Score,
			IsCorrect:  res.IsCorrect,
			AiComment:  res.AiComment,
		})
	}
	return &pb.GradeExamSubjectiveRes{Results: results}, nil
}
