package logic

import (
	"context"
	"io"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type AskQuestionStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAskQuestionStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AskQuestionStreamLogic {
	return &AskQuestionStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// AskQuestionStream 调用 Python Agent 的 gRPC 流式接口，逐帧转发给上游客户端
func (l *AskQuestionStreamLogic) AskQuestionStream(in *pb.AskQuestionReq, stream pb.AiBridge_AskQuestionStreamServer) error {
	agentStream, err := l.svcCtx.AgentClient.Svc().AskQuestionStream(l.ctx, &agentpb.AskQuestionReq{
		CourseId:  in.CourseId,
		UserId:    in.UserId,
		Question:  in.Question,
		History:   in.History,
		SessionId: in.SessionId,
	})
	if err != nil {
		l.Errorf("AskQuestionStream open stream error: %v", err)
		return code.AiServiceUnavailable
	}

	for {
		delta, err := agentStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			l.Errorf("AskQuestionStream recv error: %v", err)
			return code.AiStreamError
		}
		if err := stream.Send(&pb.AskQuestionStreamRes{
			Delta:     delta.Delta,
			Sources:   delta.Sources,
			Finished:  delta.Finished,
			SessionId: delta.SessionId,
		}); err != nil {
			l.Errorf("AskQuestionStream send error: %v", err)
			return code.AiStreamError
		}
	}
}
