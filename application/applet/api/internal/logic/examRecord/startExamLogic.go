// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package examRecord

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 开始考试
func NewStartExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartExamLogic {
	return &StartExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartExamLogic) StartExam(req *types.StartExamReq) (resp *types.StartExamRes, err error) {
	// todo: add your logic here and delete this line

	return
}
