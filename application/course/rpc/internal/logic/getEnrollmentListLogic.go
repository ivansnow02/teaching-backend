package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/code"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetEnrollmentListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetEnrollmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetEnrollmentListLogic {
	return &GetEnrollmentListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取学生已选课程列表
func (l *GetEnrollmentListLogic) GetEnrollmentList(in *pb.GetEnrollmentListReq) (*pb.GetEnrollmentListRes, error) {
	list, total, err := l.svcCtx.CourseEnrollmentModel.FindListByUserId(l.ctx, uint64(in.UserId), in.Page, in.Size)
	if err != nil {
		l.Errorf("查询学生选课列表失败: %v", err)
		return nil, code.GetEnrollmentFailed
	}

	var items []*pb.EnrollmentItem
	for _, item := range list {
		items = append(items, &pb.EnrollmentItem{
			CourseId:   int64(item.CourseId),
			EnrollTime: item.CreateTime.Unix(),
		})
	}

	return &pb.GetEnrollmentListRes{
		List:  items,
		Total: total,
	}, nil
}
