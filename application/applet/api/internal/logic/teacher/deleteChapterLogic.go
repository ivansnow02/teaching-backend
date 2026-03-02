// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"
	"teaching-backend/pkg/xcode"

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
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	_, err = l.svcCtx.CourseRPC.DeleteChapter(l.ctx, &course.DeleteChapterReq{
		Id:         req.Id,
		OperatorId: userId,
	})
	if err != nil {
		l.Errorf("删除章节失败: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
