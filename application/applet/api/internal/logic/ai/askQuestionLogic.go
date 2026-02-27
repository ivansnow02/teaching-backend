// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ai

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AskQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 智能助教问答 (基于课件RAG)
func NewAskQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AskQuestionLogic {
	return &AskQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AskQuestionLogic) AskQuestion(req *types.AskQuestionReq) (resp *types.AskQuestionRes, err error) {
	// todo: add your logic here and delete this line

	return
}
