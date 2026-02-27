// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package resource

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUploadTokenLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取 OSS/MinIO 直传签名
func NewGetUploadTokenLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUploadTokenLogic {
	return &GetUploadTokenLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUploadTokenLogic) GetUploadToken() (resp *types.GetUploadTokenRes, err error) {
	// todo: add your logic here and delete this line

	return
}
