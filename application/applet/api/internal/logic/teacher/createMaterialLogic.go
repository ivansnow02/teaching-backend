// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传课件(绑定资源)
func NewCreateMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMaterialLogic {
	return &CreateMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMaterialLogic) CreateMaterial(req *types.CreateMaterialReq) (resp *types.CreateMaterialRes, err error) {
	// todo: add your logic here and delete this line

	return
}
