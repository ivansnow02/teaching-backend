package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/code"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetCourseStudentsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetCourseStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetCourseStudentsLogic {
	return &GetCourseStudentsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取课程已选学生列表（教师用）
func (l *GetCourseStudentsLogic) GetCourseStudents(in *pb.GetCourseStudentsReq) (*pb.GetCourseStudentsRes, error) {
	list, total, err := l.svcCtx.CourseEnrollmentModel.FindListByCourseId(l.ctx, uint64(in.CourseId), in.Page, in.Size)
	if err != nil {
		l.Errorf("查询课程已选学生失败: %v", err)
		return nil, code.GetStudentsFailed
	}

	var items []*pb.CourseStudentItem
	for _, item := range list {
		items = append(items, &pb.CourseStudentItem{
			UserId:     int64(item.UserId),
			EnrollTime: item.CreateTime.Unix(),
		})
	}

	return &pb.GetCourseStudentsRes{
		List:  items,
		Total: total,
	}, nil
}
