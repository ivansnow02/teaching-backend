// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package question

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 题目详情
func NewQuestionDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionDetailLogic {
	return &QuestionDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuestionDetailLogic) QuestionDetail(req *types.QuestionDetailReq) (resp *types.QuestionDetailRes, err error) {
	rpcResp, err := l.svcCtx.ExamRPC.QuestionDetail(l.ctx, &exam.QuestionDetailReq{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("QuestionDetail error: %v", err)
		return nil, err
	}

	return &types.QuestionDetailRes{
		Question: types.QuestionItem{
			Id:              rpcResp.Question.Id,
			CourseId:        rpcResp.Question.CourseId,
			Type:            int(rpcResp.Question.Type),
			Content:         rpcResp.Question.Content,
			Answer:          rpcResp.Question.Answer,
			Analysis:        rpcResp.Question.Analysis,
			KnowledgePoints: rpcResp.Question.KnowledgePoints,
			Score:           rpcResp.Question.Score,
			Difficulty:      int(rpcResp.Question.Difficulty),
			CreateTime:      rpcResp.Question.CreateTime,
		},
	}, nil
}
