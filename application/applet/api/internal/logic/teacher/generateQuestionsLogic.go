// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/ai/rpc/pb"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateQuestionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AI 智能出题 - 异步
func NewGenerateQuestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateQuestionsLogic {
	return &GenerateQuestionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateQuestionsLogic) GenerateQuestions(req *types.GenerateQuestionsReq) (resp *types.GenerateTaskRes, err error) {
	rpcRes, err := l.svcCtx.AiRPC.GenerateQuestions(l.ctx, &pb.GenerateQuestionsReq{
		CourseId:        req.CourseId,
		KnowledgePoints: req.KnowledgePoints,
		Count:           int32(req.Count),
		Type:            int32(req.Type),
		Difficulty:      int32(req.Difficulty),
	})
	if err != nil {
		l.Errorf("GenerateQuestions rpc error: %v", err)
		return nil, err
	}

	return &types.GenerateTaskRes{
		TaskId: rpcRes.TaskId,
	}, nil
}
