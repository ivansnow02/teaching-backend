// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChapterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除章节
func NewDeleteChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChapterLogic {
	return &DeleteChapterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteChapterLogic) DeleteChapter(req *types.DeleteChapterReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
