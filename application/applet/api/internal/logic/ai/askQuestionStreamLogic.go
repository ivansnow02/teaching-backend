// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package ai

import (
	"context"
	"encoding/json"
	"io"

	"teaching-backend/application/ai/rpc/pb"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type AskQuestionStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 智能助教流式问答 SSE
func NewAskQuestionStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AskQuestionStreamLogic {
	return &AskQuestionStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *AskQuestionStreamLogic) AskQuestionStream(req *types.AskQuestionReq, client chan<- *types.AskQuestionStreamRes) error {
	// 从 JWT context 中获取 userId（go-zero JWT 中间件注入 json.Number）
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	stream, err := l.svcCtx.AiRPC.AskQuestionStream(l.ctx, &pb.AskQuestionReq{
		CourseId:  req.CourseId,
		UserId:    userId,
		Question:  req.Question,
		SessionId: req.SessionId,
	})
	if err != nil {
		l.Errorf("AskQuestionStream rpc error: %v", err)
		return err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			l.Errorf("AskQuestionStream recv error: %v", err)
			return err
		}

		client <- &types.AskQuestionStreamRes{
			Delta:     res.Delta,
			Sources:   res.Sources,
			Finished:  res.Finished,
			SessionId: res.SessionId, // 首帧由 Agent 返回，客户端保存后续多轮传入
		}
	}
}
