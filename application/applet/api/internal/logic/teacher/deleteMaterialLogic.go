// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除课件
func NewDeleteMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMaterialLogic {
	return &DeleteMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteMaterialLogic) DeleteMaterial(req *types.DeleteMaterialReq) (resp *types.Empty, err error) {
	_, err = l.svcCtx.CourseRPC.DeleteMaterial(l.ctx, &course.DeleteMaterialReq{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("删除课件失败: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
