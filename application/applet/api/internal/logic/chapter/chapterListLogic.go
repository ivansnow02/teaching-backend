// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package chapter

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChapterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 章节列表
func NewChapterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterListLogic {
	return &ChapterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChapterListLogic) ChapterList(req *types.ChapterListReq) (resp *types.ChapterListRes, err error) {
	// todo: add your logic here and delete this line

	return
}
