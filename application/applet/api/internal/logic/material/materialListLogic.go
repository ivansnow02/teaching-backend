// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package material

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaterialListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 课件列表
func NewMaterialListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialListLogic {
	return &MaterialListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialListLogic) MaterialList(req *types.MaterialListReq) (resp *types.MaterialListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
