package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type EmbedMaterialLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEmbedMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EmbedMaterialLogic {
	return &EmbedMaterialLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *EmbedMaterialLogic) EmbedMaterial(in *pb.EmbedMaterialReq) (*pb.EmbedMaterialRes, error) {
	res, err := l.svcCtx.AgentClient.Svc().EmbedMaterial(l.ctx, &agentpb.EmbedMaterialReq{
		MaterialId: in.MaterialId,
		CourseId:   in.CourseId,
		Title:      in.Title,
		Url:        in.Url,
		Type:       in.Type,
	})
	if err != nil {
		l.Errorf("EmbedMaterial material_id=%d error: %v", in.MaterialId, err)
		return nil, code.AiEmbedFailed
	}
	return &pb.EmbedMaterialRes{Accepted: res.Accepted}, nil
}
