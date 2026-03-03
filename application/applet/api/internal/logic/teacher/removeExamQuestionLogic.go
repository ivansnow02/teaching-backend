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

type RemoveExamQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 从试卷移除题目
func NewRemoveExamQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveExamQuestionLogic {
	return &RemoveExamQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveExamQuestionLogic) RemoveExamQuestion(req *types.RemoveExamQuestionReq) (resp *types.Empty, err error) {
	_, err = l.svcCtx.ExamRPC.RemoveExamQuestion(l.ctx, &exam.RemoveExamQuestionReq{
		ExamId:     req.ExamId,
		QuestionId: req.QuestionId,
	})
	if err != nil {
		l.Errorf("RemoveExamQuestion error: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
