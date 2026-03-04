package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEmbedStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEmbedStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEmbedStatusLogic {
	return &GetEmbedStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetEmbedStatusLogic) GetEmbedStatus(in *pb.GetEmbedStatusReq) (*pb.GetEmbedStatusRes, error) {
	res, err := l.svcCtx.AgentClient.Svc().GetEmbedStatus(l.ctx, &agentpb.GetEmbedStatusReq{
		MaterialId: in.MaterialId,
	})
	if err != nil {
		l.Errorf("GetEmbedStatus material_id=%d error: %v", in.MaterialId, err)
		return nil, code.AiServiceUnavailable
	}
	return &pb.GetEmbedStatusRes{
		AiStatus: res.AiStatus,
		Message:  res.Message,
	}, nil
}
