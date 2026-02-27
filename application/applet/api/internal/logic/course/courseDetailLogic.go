// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package course

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 课程详情
func NewCourseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseDetailLogic {
	return &CourseDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CourseDetailLogic) CourseDetail(req *types.CourseDetailReq) (resp *types.CourseDetailRes, err error) {
	// todo: add your logic here and delete this line

	return
}
