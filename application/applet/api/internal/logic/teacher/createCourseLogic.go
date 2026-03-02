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

type CreateCourseLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建课程(教师)
func NewCreateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCourseLogic {
	return &CreateCourseLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateCourseLogic) CreateCourse(req *types.CreateCourseReq) (resp *types.CreateCourseRes, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	rpcResp, err := l.svcCtx.CourseRPC.CreateCourse(l.ctx, &course.CreateCourseReq{
		Title:       req.Title,
		Cover:       req.Cover,
		Description: req.Description,
		TeacherId:   userId,
	})
	if err != nil {
		l.Errorf("创建课程失败: %v", err)
		return nil, err
	}

	return &types.CreateCourseRes{
		Id: rpcResp.Id,
	}, nil
}
