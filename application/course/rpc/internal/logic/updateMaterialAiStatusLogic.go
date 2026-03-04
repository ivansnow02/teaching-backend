package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateMaterialAiStatusLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateMaterialAiStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateMaterialAiStatusLogic {
	return &UpdateMaterialAiStatusLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新课件AI处理状态
func (l *UpdateMaterialAiStatusLogic) UpdateMaterialAiStatus(in *pb.UpdateMaterialAiStatusReq) (*pb.UpdateMaterialAiStatusRes, error) {
	err := l.svcCtx.CourseMaterialModel.UpdateAiStatus(l.ctx, in.MaterialId, int64(in.AiStatus))
	if err != nil {
		l.Errorf("更新课件AI状态失败 material_id: %d err: %v", in.MaterialId, err)
		return nil, xcode.ServerErr
	}

	return &pb.UpdateMaterialAiStatusRes{}, nil
}
