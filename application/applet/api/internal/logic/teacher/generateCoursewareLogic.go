// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"
	"encoding/json"

	"teaching-backend/application/ai/rpc/pb"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCoursewareLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AI 生成课件内容 - 异步
func NewGenerateCoursewareLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCoursewareLogic {
	return &GenerateCoursewareLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateCoursewareLogic) GenerateCourseware(req *types.GenerateCoursewareReq) (resp *types.GenerateTaskRes, err error) {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	rpcRes, err := l.svcCtx.AiRPC.GenerateCourseware(l.ctx, &pb.GenerateCoursewareReq{
		ChapterId:    req.ChapterId,
		Requirements: req.Requirements,
		UserId:       userId,
		SessionId:    req.SessionId,
	})
	if err != nil {
		l.Errorf("GenerateCourseware rpc error: %v", err)
		return nil, err
	}

	return &types.GenerateTaskRes{
		TaskId: rpcRes.TaskId,
	}, nil
}
