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

type CreateQuestionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建题目(教师)
func NewCreateQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateQuestionLogic {
	return &CreateQuestionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateQuestionLogic) CreateQuestion(req *types.CreateQuestionReq) (resp *types.CreateQuestionRes, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	rpcResp, err := l.svcCtx.ExamRPC.CreateQuestion(l.ctx, &exam.CreateQuestionReq{
		CourseId:        req.CourseId,
		TeacherId:       userId,
		Type:            int32(req.Type),
		Content:         req.Content,
		Answer:          req.Answer,
		Analysis:        req.Analysis,
		KnowledgePoints: req.KnowledgePoints,
		Score:           req.Score,
		Difficulty:      int32(req.Difficulty),
	})
	if err != nil {
		l.Errorf("CreateQuestion error: %v", err)
		return nil, err
	}

	return &types.CreateQuestionRes{
		Id: rpcResp.Id,
	}, nil
}
