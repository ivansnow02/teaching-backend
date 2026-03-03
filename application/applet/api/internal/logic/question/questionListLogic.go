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

type QuestionListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 题目列表
func NewQuestionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionListLogic {
	return &QuestionListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *QuestionListLogic) QuestionList(req *types.QuestionListReq) (resp *types.QuestionListRes, err error) {
	rpcResp, err := l.svcCtx.ExamRPC.QuestionList(l.ctx, &exam.QuestionListReq{
		CourseId:   req.CourseId,
		Type:       int32(req.Type),
		Difficulty: int32(req.Difficulty),
		Page:       int64(req.Page),
		Size:       int64(req.Size),
	})
	if err != nil {
		l.Errorf("QuestionList error: %v", err)
		return nil, err
	}

	list := make([]types.QuestionItem, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		list = append(list, types.QuestionItem{
			Id:              item.Id,
			CourseId:        item.CourseId,
			Type:            int(item.Type),
			Content:         item.Content,
			Answer:          item.Answer,
			Analysis:        item.Analysis,
			KnowledgePoints: item.KnowledgePoints,
			Score:           item.Score,
			Difficulty:      int(item.Difficulty),
			CreateTime:      item.CreateTime,
		})
	}

	return &types.QuestionListRes{
		List:  list,
		Total: rpcResp.Total,
	}, nil
}
