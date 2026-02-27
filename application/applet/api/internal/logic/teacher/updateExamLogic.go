// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新试卷(教师)
func NewUpdateExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExamLogic {
	return &UpdateExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExamLogic) UpdateExam(req *types.UpdateExamReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
