// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package examRecord

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 提交答卷
func NewSubmitExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitExamLogic {
	return &SubmitExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitExamLogic) SubmitExam(req *types.SubmitExamReq) (resp *types.SubmitExamRes, err error) {
	// todo: add your logic here and delete this line

	return
}
