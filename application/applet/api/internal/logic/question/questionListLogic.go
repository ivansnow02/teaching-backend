// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package question

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 题目列表
func NewQuestionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionListLogic {
	return &QuestionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuestionListLogic) QuestionList(req *types.QuestionListReq) (resp *types.QuestionListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
