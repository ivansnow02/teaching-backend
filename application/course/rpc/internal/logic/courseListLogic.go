package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCourseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseListLogic {
	return &CourseListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 课程列表
func (l *CourseListLogic) CourseList(in *pb.CourseListReq) (*pb.CourseListRes, error) {
	list, total, err := l.svcCtx.CourseModel.FindByTeacherId(l.ctx, in.TeacherId, in.Page, in.Size)
	if err != nil {
		l.Errorf("查询教师课程列表失败: %v", err)
		return nil, xcode.ServerErr
	}

	var res []*pb.CourseItem
	for _, item := range list {
		res = append(res, &pb.CourseItem{
			Id:          int64(item.Id),
			Title:       item.Title,
			Cover:       item.Cover,
			Description: item.Description.String,
			TeacherId:   int64(item.TeacherId),
			Status:      int32(item.Status),
			CreateTime:  item.CreateTime.Unix(),
		})
	}

	return &pb.CourseListRes{
		List:  res,
		Total: total,
	}, nil
}
