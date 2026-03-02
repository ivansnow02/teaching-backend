// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新课程(教师)
func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateCourseLogic) UpdateCourse(req *types.UpdateCourseReq) (resp *types.Empty, err error) {
	_, err = l.svcCtx.CourseRPC.UpdateCourse(l.ctx, &course.UpdateCourseReq{
		Id:          req.Id,
		Title:       req.Title,
		Cover:       req.Cover,
		Description: req.Description,
		Status:      int32(req.Status),
	})
	if err != nil {
		l.Errorf("更新课程失败: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
