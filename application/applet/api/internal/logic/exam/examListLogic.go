// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package exam

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExamListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 试卷列表
func NewExamListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamListLogic {
	return &ExamListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExamListLogic) ExamList(req *types.ExamListReq) (resp *types.ExamListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
