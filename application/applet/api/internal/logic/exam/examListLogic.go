// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package exam

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExamListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 试卷列表
func NewExamListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamListLogic {
	return &ExamListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExamListLogic) ExamList(req *types.ExamListReq) (resp *types.ExamListRes, err error) {
	rpcResp, err := l.svcCtx.ExamRPC.ExamList(l.ctx, &exam.ExamListReq{
		CourseId: req.CourseId,
		Page:     int64(req.Page),
		Size:     int64(req.Size),
	})
	if err != nil {
		l.Errorf("ExamList error: %v", err)
		return nil, err
	}

	list := make([]types.ExamItem, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		list = append(list, types.ExamItem{
			Id:         item.Id,
			CourseId:   item.CourseId,
			Title:      item.Title,
			TotalScore: item.TotalScore,
			PassScore:  item.PassScore,
			Duration:   int(item.Duration),
			StartTime:  item.StartTime,
			EndTime:    item.EndTime,
			Status:     int(item.Status),
			ExamType:   int(item.ExamType),
			CreateTime: item.CreateTime,
		})
	}

	return &types.ExamListRes{
		List:  list,
		Total: rpcResp.Total,
	}, nil
}
