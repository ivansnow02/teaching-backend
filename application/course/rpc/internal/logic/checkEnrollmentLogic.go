package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/model"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckEnrollmentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckEnrollmentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckEnrollmentLogic {
	return &CheckEnrollmentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 检查是否已选课
func (l *CheckEnrollmentLogic) CheckEnrollment(in *pb.CheckEnrollmentReq) (*pb.CheckEnrollmentRes, error) {
	enrollment, err := l.svcCtx.CourseEnrollmentModel.FindOneByUserIdCourseId(l.ctx, uint64(in.UserId), uint64(in.CourseId))
	if err != nil {
		if err == model.ErrNotFound {
			return &pb.CheckEnrollmentRes{IsEnrolled: false}, nil
		}
		l.Errorf("查询选课状态失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.CheckEnrollmentRes{IsEnrolled: enrollment.Status == 1}, nil
}
