// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 更新题目(教师)
func NewUpdateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateQuestionLogic {
	return &UpdateQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateQuestionLogic) UpdateQuestion(req *types.UpdateQuestionReq) (resp *types.Empty, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	_, err = l.svcCtx.ExamRPC.UpdateQuestion(l.ctx, &exam.UpdateQuestionReq{
		Id:              req.Id,
		OperatorId:      userId,
		Content:         req.Content,
		Answer:          req.Answer,
		Analysis:        req.Analysis,
		KnowledgePoints: req.KnowledgePoints,
		Score:           req.Score,
		Difficulty:      int32(req.Difficulty),
	})
	if err != nil {
		l.Errorf("UpdateQuestion error: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
