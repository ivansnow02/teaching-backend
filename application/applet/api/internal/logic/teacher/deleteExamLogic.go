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

type DeleteExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 删除试卷(教师)
func NewDeleteExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExamLogic {
	return &DeleteExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *DeleteExamLogic) DeleteExam(req *types.DeleteExamReq) (resp *types.Empty, err error) {
	_, err = l.svcCtx.ExamRPC.DeleteExam(l.ctx, &exam.DeleteExamReq{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("DeleteExam error: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
