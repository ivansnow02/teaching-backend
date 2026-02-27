// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建试卷(教师)
func NewCreateExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExamLogic {
	return &CreateExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateExamLogic) CreateExam(req *types.CreateExamReq) (resp *types.CreateExamRes, err error) {
	// todo: add your logic here and delete this line

	return
}
