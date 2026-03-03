package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/code"
	"teaching-backend/application/course/rpc/internal/model"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DropCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDropCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DropCourseLogic {
	return &DropCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 退课
func (l *DropCourseLogic) DropCourse(in *pb.DropCourseReq) (*pb.DropCourseRes, error) {
	existing, err := l.svcCtx.CourseEnrollmentModel.FindOneByUserIdCourseId(l.ctx, uint64(in.UserId), uint64(in.CourseId))
	if err != nil {
		if err == model.ErrNotFound {
			return &pb.DropCourseRes{}, nil
		}
		l.Errorf("查询选课记录失败: %v", err)
		return nil, code.DropFailed
	}

	if existing.Status == 2 {
		return &pb.DropCourseRes{}, nil
	}

	existing.Status = 2
	if err = l.svcCtx.CourseEnrollmentModel.Update(l.ctx, existing); err != nil {
		l.Errorf("退课失败: %v", err)
		return nil, code.DropFailed
	}

	return &pb.DropCourseRes{}, nil
}
