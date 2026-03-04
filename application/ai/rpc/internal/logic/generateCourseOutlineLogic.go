package logic

import (
	"context"

	"teaching-backend/application/ai/rpc/agentpb"
	"teaching-backend/application/ai/rpc/internal/code"
	"teaching-backend/application/ai/rpc/internal/svc"
	"teaching-backend/application/ai/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCourseOutlineLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGenerateCourseOutlineLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCourseOutlineLogic {
	return &GenerateCourseOutlineLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GenerateCourseOutlineLogic) GenerateCourseOutline(in *pb.GenerateCourseOutlineReq) (*pb.GenerateTaskRes, error) {
	res, err := l.svcCtx.AgentClient.Svc().GenerateCourseOutline(l.ctx, &agentpb.GenerateCourseOutlineReq{
		CourseId:     in.CourseId,
		Topic:        in.Topic,
		Requirements: in.Requirements,
	})
	if err != nil {
		l.Errorf("GenerateCourseOutline course_id=%d error: %v", in.CourseId, err)
		return nil, code.AiGenerateFailed
	}
	return &pb.GenerateTaskRes{TaskId: res.TaskId}, nil
}
