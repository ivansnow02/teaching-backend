package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCoursewareLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateCoursewareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCoursewareLogic {
	return &GenerateCoursewareLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateCoursewareLogic) GenerateCourseware(in *pb.GenerateCoursewareReq) (*pb.GenerateTaskRes, error) {
	res, err := l.svcCtx.AgentClient.Svc().GenerateCourseware(l.ctx, &agentpb.GenerateCoursewareReq{
		ChapterId:    in.ChapterId,
		Requirements: in.Requirements,
	})
	if err != nil {
		l.Errorf("GenerateCourseware chapter_id=%d error: %v", in.ChapterId, err)
		return nil, code.AiGenerateFailed
	}
	return &pb.GenerateTaskRes{TaskId: res.TaskId}, nil
}
