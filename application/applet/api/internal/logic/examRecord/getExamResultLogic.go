// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package examRecord

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExamResultLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取考试结果
func NewGetExamResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExamResultLogic {
	return &GetExamResultLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetExamResultLogic) GetExamResult(req *types.GetExamResultReq) (resp *types.GetExamResultRes, err error) {
	rpcResp, err := l.svcCtx.ExamRPC.GetExamResult(l.ctx, &exam.GetExamResultReq{
		RecordId: req.RecordId,
	})
	if err != nil {
		l.Errorf("GetExamResult error: %v", err)
		return nil, err
	}

	return &types.GetExamResultRes{
		ExamId:     rpcResp.ExamId,
		Score:      rpcResp.Score,
		Status:     int(rpcResp.Status),
		SubmitTime: rpcResp.SubmitTime,
	}, nil
}
