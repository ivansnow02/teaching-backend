package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateQuestionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateQuestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateQuestionsLogic {
	return &GenerateQuestionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateQuestionsLogic) GenerateQuestions(in *pb.GenerateQuestionsReq) (*pb.GenerateTaskRes, error) {
	res, err := l.svcCtx.AgentClient.Svc().GenerateQuestions(l.ctx, &agentpb.GenerateQuestionsReq{
		CourseId:        in.CourseId,
		KnowledgePoints: in.KnowledgePoints,
		Count:           in.Count,
		Type:            in.Type,
		Difficulty:      in.Difficulty,
	})
	if err != nil {
		l.Errorf("GenerateQuestions course_id=%d error: %v", in.CourseId, err)
		return nil, code.AiGenerateFailed
	}
	return &pb.GenerateTaskRes{TaskId: res.TaskId}, nil
}
