// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package examRecord

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAnswerSnapshotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 考试过程心跳与答案快照同步
func NewSaveAnswerSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAnswerSnapshotLogic {
	return &SaveAnswerSnapshotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveAnswerSnapshotLogic) SaveAnswerSnapshot(req *types.SaveAnswerSnapshotReq) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
