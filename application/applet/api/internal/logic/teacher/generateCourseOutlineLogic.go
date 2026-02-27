// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCourseOutlineLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AI 生成课件大纲 - 异步
func NewGenerateCourseOutlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCourseOutlineLogic {
	return &GenerateCourseOutlineLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateCourseOutlineLogic) GenerateCourseOutline(req *types.GenerateCourseOutlineReq) (resp *types.GenerateTaskRes, err error) {
	// todo: add your logic here and delete this line

	return
}
