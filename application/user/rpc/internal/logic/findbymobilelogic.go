package logic

import (
	"context"

	"teaching-backend/application/user/rpc/internal/svc"
	"teaching-backend/application/user/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByMobileLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByMobileLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByMobileLogic {
	return &FindByMobileLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据手机号获取用户信息
func (l *FindByMobileLogic) FindByMobile(in *pb.FindByMobileReq) (*pb.FindByMobileRes, error) {
	// todo: add your logic here and delete this line

	return &pb.FindByMobileRes{}, nil
}
