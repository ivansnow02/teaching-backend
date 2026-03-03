package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/model"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/application/course/rpc/internal/code"

	"github.com/zeromicro/go-zero/core/logx"
)

type EnrollCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewEnrollCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *EnrollCourseLogic {
	return &EnrollCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 选课
func (l *EnrollCourseLogic) EnrollCourse(in *pb.EnrollCourseReq) (*pb.EnrollCourseRes, error) {
	existing, err := l.svcCtx.CourseEnrollmentModel.FindOneByUserIdCourseId(l.ctx, uint64(in.UserId), uint64(in.CourseId))
	if err != nil && err != model.ErrNotFound {
		l.Errorf("查询选课记录失败: %v", err)
		return nil, code.EnrollFailed
	}

	if err == nil {
		if existing.Status == 1 {
			return &pb.EnrollCourseRes{}, nil
		}
		// 曾退课，重新激活
		existing.Status = 1
		if err = l.svcCtx.CourseEnrollmentModel.Update(l.ctx, existing); err != nil {
			l.Errorf("重新激活选课记录失败: %v", err)
			return nil, code.EnrollFailed
		}
		return &pb.EnrollCourseRes{}, nil
	}

	// 插入新选课记录
	_, err = l.svcCtx.CourseEnrollmentModel.Insert(l.ctx, &model.CourseEnrollment{
		UserId:   uint64(in.UserId),
		CourseId: uint64(in.CourseId),
		Status:   1,
	})
	if err != nil {
		l.Errorf("插入选课记录失败: %v", err)
		return nil, code.EnrollFailed
	}

	return &pb.EnrollCourseRes{}, nil
}
