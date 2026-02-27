// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChapterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建章节
func NewCreateChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChapterLogic {
	return &CreateChapterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateChapterLogic) CreateChapter(req *types.CreateChapterReq) (resp *types.CreateChapterRes, err error) {
	// todo: add your logic here and delete this line

	return
}
