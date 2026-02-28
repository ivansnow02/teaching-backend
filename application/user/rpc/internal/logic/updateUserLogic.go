package logic

import (
	"context"

	"teaching-backend/application/user/rpc/internal/svc"
	"teaching-backend/application/user/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserLogic {
	return &UpdateUserLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新用户信息
func (l *UpdateUserLogic) UpdateUser(in *pb.UpdateUserReq) (*pb.UpdateUserRes, error) {
	user, err := l.svcCtx.UserModel.FindOne(l.ctx, uint64(in.UserId))
	if err != nil {
		l.Errorf("查询用户失败: %v", err)
		return nil, xcode.ServerErr
	}

	if in.Nickname != "" {
		user.Nickname = in.Nickname
	}
	if in.Avatar != "" {
		user.Avatar = in.Avatar
	}

	if err := l.svcCtx.UserModel.Update(l.ctx, user); err != nil {
		l.Errorf("更新用户失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.UpdateUserRes{}, nil
}
