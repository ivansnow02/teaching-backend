package courselogic

import (
	"context"
	"database/sql"

	"teaching-backend/application/course/rpc/internal/code"
	"teaching-backend/application/course/rpc/internal/model"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateCourseLogic {
	return &CreateCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建课程
func (l *CreateCourseLogic) CreateCourse(in *pb.CreateCourseReq) (*pb.CreateCourseRes, error) {
	if in.Title == "" {
		return nil, code.CourseTitleEmpty
	}

	res, err := l.svcCtx.CourseModel.Insert(l.ctx, &model.Course{
		Title:       in.Title,
		Cover:       in.Cover,
		Description: sql.NullString{String: in.Description, Valid: in.Description != ""},
		TeacherId:   uint64(in.TeacherId),
		Status:      0, // 默认未发布
	})
	if err != nil {
		l.Errorf("插入课程记录失败: %v", err)
		return nil, xcode.ServerErr
	}

	id, err := res.LastInsertId()
	if err != nil {
		l.Errorf("获取课程 LastInsertId 失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.CreateCourseRes{
		Id: id,
	}, nil
}
