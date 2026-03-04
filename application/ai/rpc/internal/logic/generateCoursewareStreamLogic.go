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

type GenerateCoursewareStreamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateCoursewareStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCoursewareStreamLogic {
	return &GenerateCoursewareStreamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// GenerateCoursewareStream 调用 Python Agent 的 gRPC 流式接口，逐帧转发 Markdown 片段
func (l *GenerateCoursewareStreamLogic) GenerateCoursewareStream(in *pb.GenerateCoursewareReq, stream pb.AiBridge_GenerateCoursewareStreamServer) error {
	agentStream, err := l.svcCtx.AgentClient.Svc().GenerateCoursewareStream(l.ctx, &agentpb.GenerateCoursewareReq{
		ChapterId:    in.ChapterId,
		Requirements: in.Requirements,
		UserId:       in.UserId,
		SessionId:    in.SessionId,
	})
	if err != nil {
		l.Errorf("GenerateCoursewareStream open stream error: %v", err)
		return code.AiServiceUnavailable
	}

	for {
		delta, err := agentStream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			l.Errorf("GenerateCoursewareStream recv error: %v", err)
			return code.AiStreamError
		}
		if err := stream.Send(&pb.GenerateCoursewareStreamRes{
			Delta:     delta.Delta,
			Finished:  delta.Finished,
			SessionId: delta.SessionId,
		}); err != nil {
			l.Errorf("GenerateCoursewareStream send error: %v", err)
			return code.AiStreamError
		}
	}
}
