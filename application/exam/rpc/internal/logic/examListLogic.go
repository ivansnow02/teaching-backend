package logic

import (
	"context"
	"fmt"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExamListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExamListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamListLogic {
	return &ExamListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 试卷列表
func (l *ExamListLogic) ExamList(in *pb.ExamListReq) (*pb.ExamListRes, error) {
	list, total, err := l.svcCtx.ExamModel.FindList(l.ctx, uint64(in.CourseId), in.Page, in.Size)
	if err != nil {
		l.Errorf("ExamModel.FindList error: %v", err)
		return nil, xcode.ServerErr
	}

	var resList []*pb.ExamItem
	for _, item := range list {
		resList = append(resList, &pb.ExamItem{
			Id:         int64(item.Id),
			CourseId:   int64(item.CourseId),
			Title:      item.Title,
			TotalScore: fmt.Sprintf("%.1f", item.TotalScore),
			PassScore:  fmt.Sprintf("%.1f", item.PassScore),
			Duration:   int32(item.Duration),
			StartTime:  item.StartTime.Time.Unix(),
			EndTime:    item.EndTime.Time.Unix(),
			Status:     int32(item.Status),
			ExamType:   int32(item.ExamType),
			CreateTime: item.CreateTime.Unix(),
		})
	}

	return &pb.ExamListRes{
		List:  resList,
		Total: total,
	}, nil
}
