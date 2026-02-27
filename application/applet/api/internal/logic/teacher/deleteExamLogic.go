// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除试卷(教师)
func NewDeleteExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExamLogic {
	return &DeleteExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteExamLogic) DeleteExam(req *types.DeleteExamReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
