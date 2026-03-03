// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package enrollment

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/code"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 学生退课
func NewDropCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropCourseLogic {
	return &DropCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DropCourseLogic) DropCourse(req *types.DropCourseReq) (resp *types.Empty, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}
	if req.CourseId <= 0 {
		return nil, xcode.RequestErr
	}

	_, err = l.svcCtx.CourseRPC.DropCourse(l.ctx, &course.DropCourseReq{
		UserId:   userId,
		CourseId: req.CourseId,
	})
	if err != nil {
		l.Errorf("退课失败: %v", err)
		return nil, code.CourseDropFailed
	}

	return &types.Empty{}, nil
}
