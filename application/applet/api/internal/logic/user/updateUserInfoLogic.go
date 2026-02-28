// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/user/rpc/user"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateUserInfoLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 编辑个人信息
func NewUpdateUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateUserInfoLogic {
	return &UpdateUserInfoLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateUserInfoLogic) UpdateUserInfo(req *types.UpdateUserInfoReq) (resp *types.Empty, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	_, err = l.svcCtx.UserRPC.UpdateUser(l.ctx, &user.UpdateUserReq{
		UserId:   userId,
		Nickname: req.Nickname,
		Avatar:   req.Avatar,
	})
	if err != nil {
		l.Errorf("更新用户信息失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &types.Empty{}, nil
}
