// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/user/rpc/user"
	"teaching-backend/pkg/encrypt"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取当前用户信息
func NewUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserInfoLogic {
	return &UserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserInfoLogic) UserInfo() (resp *types.UserInfoRes, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}
	u, err := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdReq{
		UserId: userId,
	})
	if err != nil {
		l.Errorf("查询用户失败: %v", err)
		return nil, xcode.ServerErr
	}
	if u == nil || u.UserId <= 0 {
		return nil, xcode.AccessDenied
	}
	email, err := encrypt.DecEmail(u.Email)
	if err != nil {
		l.Errorf("邮箱解密失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &types.UserInfoRes{
		UserId:   u.UserId,
		Email:    email,
		Nickname: u.Nickname,
		Avatar:   u.Avatar,
		Role:     int(u.Role),
	}, nil
}
