// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package progress

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudyProgressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新学习进度
func NewUpdateStudyProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudyProgressLogic {
	return &UpdateStudyProgressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStudyProgressLogic) UpdateStudyProgress(req *types.UpdateStudyProgressReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
