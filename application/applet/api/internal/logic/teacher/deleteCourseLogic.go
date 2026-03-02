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

type DeleteCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除课程(教师)
func NewDeleteCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCourseLogic {
	return &DeleteCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteCourseLogic) DeleteCourse(req *types.DeleteCourseReq) (resp *types.Empty, err error) {
	_, err = l.svcCtx.CourseRPC.DeleteCourse(l.ctx, &course.DeleteCourseReq{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("删除课程失败: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
