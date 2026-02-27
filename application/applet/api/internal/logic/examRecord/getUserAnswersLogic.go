// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package examRecord

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAnswersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取答题明细
func NewGetUserAnswersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAnswersLogic {
	return &GetUserAnswersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAnswersLogic) GetUserAnswers(req *types.GetUserAnswersReq) (resp *types.GetUserAnswersRes, err error) {
	// todo: add your logic here and delete this line

	return
}
