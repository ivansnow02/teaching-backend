// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/client/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateChapterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新章节
func NewUpdateChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChapterLogic {
	return &UpdateChapterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateChapterLogic) UpdateChapter(req *types.UpdateChapterReq) (resp *types.Empty, err error) {
	_, err = l.svcCtx.CourseRPC.UpdateChapter(l.ctx, &course.UpdateChapterReq{
		Id:    req.Id,
		Title: req.Title,
		Sort:  int32(req.Sort),
	})
	if err != nil {
		l.Errorf("更新章节失败: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
