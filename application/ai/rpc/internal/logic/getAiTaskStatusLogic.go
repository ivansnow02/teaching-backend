package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetAiTaskStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetAiTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAiTaskStatusLogic {
	return &GetAiTaskStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetAiTaskStatusLogic) GetAiTaskStatus(in *pb.GetAiTaskStatusReq) (*pb.GetAiTaskStatusRes, error) {
	if in.TaskId == "" {
		return nil, code.AiTaskNotFound
	}
	res, err := l.svcCtx.AgentClient.Svc().GetTaskStatus(l.ctx, &agentpb.GetTaskStatusReq{
		TaskId: in.TaskId,
	})
	if err != nil {
		l.Errorf("GetAiTaskStatus task_id=%s error: %v", in.TaskId, err)
		return nil, code.AiServiceUnavailable
	}
	return &pb.GetAiTaskStatusRes{
		Status:  res.Status,
		Result:  res.Result,
		Message: res.Message,
	}, nil
}
