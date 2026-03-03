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

type AddExamQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 添加题目到试卷
func NewAddExamQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddExamQuestionLogic {
	return &AddExamQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AddExamQuestionLogic) AddExamQuestion(req *types.AddExamQuestionReq) (resp *types.Empty, err error) {
	_, err = l.svcCtx.ExamRPC.AddExamQuestion(l.ctx, &exam.AddExamQuestionReq{
		ExamId:     req.ExamId,
		QuestionId: req.QuestionId,
		Score:      req.Score,
		Sort:       int32(req.Sort),
	})
	if err != nil {
		l.Errorf("AddExamQuestion error: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
