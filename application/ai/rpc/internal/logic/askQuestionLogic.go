package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AskQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAskQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AskQuestionLogic {
	return &AskQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AskQuestionLogic) AskQuestion(in *pb.AskQuestionReq) (*pb.AskQuestionRes, error) {
	res, err := l.svcCtx.AgentClient.Svc().AskQuestion(l.ctx, &agentpb.AskQuestionReq{
		CourseId: in.CourseId,
		UserId:   in.UserId,
		Question: in.Question,
		History:  in.History,
	})
	if err != nil {
		l.Errorf("AskQuestion course_id=%d error: %v", in.CourseId, err)
		return nil, code.AiServiceUnavailable
	}
	return &pb.AskQuestionRes{
		Answer:  res.Answer,
		Sources: res.Sources,
	}, nil
}
