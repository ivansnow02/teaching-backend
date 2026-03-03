// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新试卷(教师)
func NewUpdateExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExamLogic {
	return &UpdateExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateExamLogic) UpdateExam(req *types.UpdateExamReq) (resp *types.Empty, err error) {
	_, err = l.svcCtx.ExamRPC.UpdateExam(l.ctx, &exam.UpdateExamReq{
		Id:         req.Id,
		Title:      req.Title,
		TotalScore: req.TotalScore,
		PassScore:  req.PassScore,
		Duration:   int32(req.Duration),
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		Status:     int32(req.Status),
	})
	if err != nil {
		l.Errorf("UpdateExam error: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
