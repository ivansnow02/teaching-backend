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

type CheckEnrollmentLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 检查是否已选课
func NewCheckEnrollmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEnrollmentLogic {
	return &CheckEnrollmentLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CheckEnrollmentLogic) CheckEnrollment(req *types.CheckEnrollmentReq) (resp *types.CheckEnrollmentRes, err error) {
	if req.CourseId <= 0 {
		return nil, xcode.RequestErr
	}

	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	rpcResp, err := l.svcCtx.CourseRPC.CheckEnrollment(l.ctx, &course.CheckEnrollmentReq{
		UserId:   userId,
		CourseId: req.CourseId,
	})
	if err != nil {
		l.Errorf("检查选课状态失败: %v", err)
		return nil, code.CheckEnrollmentFailed
	}

	return &types.CheckEnrollmentRes{
		IsEnrolled: rpcResp.IsEnrolled,
	}, nil
}
