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

type GetAiTaskStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查询 AI 异步任务状态
func NewGetAiTaskStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetAiTaskStatusLogic {
	return &GetAiTaskStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetAiTaskStatusLogic) GetAiTaskStatus(req *types.GetAiTaskStatusReq) (resp *types.GetAiTaskStatusRes, err error) {
	rpcRes, err := l.svcCtx.AiRPC.GetAiTaskStatus(l.ctx, &pb.GetAiTaskStatusReq{
		TaskId: req.TaskId,
	})
	if err != nil {
		l.Errorf("GetAiTaskStatus rpc error: %v", err)
		return nil, err
	}

	return &types.GetAiTaskStatusRes{
		Status:  int(rpcRes.Status),
		Result:  rpcRes.Result,
		Message: rpcRes.Message,
	}, nil
}
