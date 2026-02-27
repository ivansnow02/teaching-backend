// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveExamQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 从试卷移除题目
func NewRemoveExamQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveExamQuestionLogic {
	return &RemoveExamQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveExamQuestionLogic) RemoveExamQuestion(req *types.RemoveExamQuestionReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
