// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package progress

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/client/course"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudyProgressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新学习进度
func NewUpdateStudyProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudyProgressLogic {
	return &UpdateStudyProgressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateStudyProgressLogic) UpdateStudyProgress(req *types.UpdateStudyProgressReq) (resp *types.Empty, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	_, err = l.svcCtx.CourseRPC.UpdateStudyProgress(l.ctx, &course.UpdateStudyProgressReq{
		UserId:     userId,
		CourseId:   req.CourseId,
		ChapterId:  req.ChapterId,
		MaterialId: req.MaterialId,
		Progress:   int32(req.Progress),
	})
	if err != nil {
		l.Errorf("更新学习进度失败: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
