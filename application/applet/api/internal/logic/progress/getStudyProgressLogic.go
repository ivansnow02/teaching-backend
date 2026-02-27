// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package progress

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudyProgressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取学习进度
func NewGetStudyProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudyProgressLogic {
	return &GetStudyProgressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudyProgressLogic) GetStudyProgress(req *types.GetStudyProgressReq) (resp *types.GetStudyProgressRes, err error) {
	// todo: add your logic here and delete this line

	return
}
