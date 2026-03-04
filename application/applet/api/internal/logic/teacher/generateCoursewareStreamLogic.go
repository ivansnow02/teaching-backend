// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"
	"encoding/json"
	"io"

	"teaching-backend/application/ai/rpc/pb"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GenerateCoursewareStreamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// AI 课件流式生成 SSE
func NewGenerateCoursewareStreamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GenerateCoursewareStreamLogic {
	return &GenerateCoursewareStreamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GenerateCoursewareStreamLogic) GenerateCoursewareStream(req *types.GenerateCoursewareReq, client chan<- *types.GenerateCoursewareStreamRes) error {
	userId, _ := l.ctx.Value("userId").(json.Number).Int64()

	stream, err := l.svcCtx.AiRPC.GenerateCoursewareStream(l.ctx, &pb.GenerateCoursewareReq{
		ChapterId:    req.ChapterId,
		Requirements: req.Requirements,
		UserId:       userId,
		SessionId:    req.SessionId,
	})
	if err != nil {
		l.Errorf("GenerateCoursewareStream rpc error: %v", err)
		return err
	}

	for {
		res, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			l.Errorf("GenerateCoursewareStream recv error: %v", err)
			return err
		}

		client <- &types.GenerateCoursewareStreamRes{
			Delta:     res.Delta,
			Finished:  res.Finished,
			SessionId: res.SessionId, // 首帧由 Agent 返回，客户端保存后续多轮传入
		}
	}
}
