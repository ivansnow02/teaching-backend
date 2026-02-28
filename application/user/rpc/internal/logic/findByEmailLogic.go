package logic

import (
	"context"

	"teaching-backend/application/user/rpc/internal/svc"
	"teaching-backend/application/user/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type FindByEmailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewFindByEmailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FindByEmailLogic {
	return &FindByEmailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 根据邮箱获取用户信息
func (l *FindByEmailLogic) FindByEmail(in *pb.FindByEmailReq) (*pb.FindByEmailRes, error) {
	user, err := l.svcCtx.UserModel.FindByEmail(l.ctx, in.Email)
	if err != nil {
		l.Errorf("查询用户失败: %v", err)
		return nil, xcode.ServerErr
	}
	if user == nil {
		return &pb.FindByEmailRes{}, nil
	}

	return &pb.FindByEmailRes{
		UserId: int64(user.Id),
		Email: user.Email,
		Password: user.Password,
		Nickname: user.Nickname,
		Avatar: user.Avatar,
		Role: int32(user.Role),
	}, nil
}
