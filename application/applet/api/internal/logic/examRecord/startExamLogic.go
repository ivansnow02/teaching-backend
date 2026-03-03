// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package examRecord

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 开始考试
func NewStartExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartExamLogic {
	return &StartExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *StartExamLogic) StartExam(req *types.StartExamReq) (resp *types.StartExamRes, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	rpcResp, err := l.svcCtx.ExamRPC.StartExam(l.ctx, &exam.StartExamReq{
		ExamId: req.ExamId,
		UserId: userId,
	})
	if err != nil {
		l.Errorf("StartExam error: %v", err)
		return nil, err
	}

	return &types.StartExamRes{
		RecordId: rpcResp.RecordId,
	}, nil
}
