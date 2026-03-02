package logic

import (
	"context"
	"database/sql"

	"teaching-backend/application/course/rpc/internal/code"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateCourseLogic {
	return &UpdateCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新课程
func (l *UpdateCourseLogic) UpdateCourse(in *pb.UpdateCourseReq) (*pb.UpdateCourseRes, error) {
	course, err := l.svcCtx.CourseModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.CourseNotFound
		}
		l.Errorf("查询课程详情失败: %v", err)
		return nil, xcode.ServerErr
	}

	if course.TeacherId != uint64(in.OperatorId) {
		return nil, code.NoPermission
	}

	course.Title = in.Title
	course.Cover = in.Cover
	course.Description = sql.NullString{String: in.Description, Valid: in.Description != ""}
	course.Status = int64(in.Status)

	err = l.svcCtx.CourseModel.Update(l.ctx, course)
	if err != nil {
		l.Errorf("更新课程记录失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.UpdateCourseRes{}, nil
}
