// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"strings"

	"teaching-backend/application/applet/api/internal/code"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/user/rpc/user"
	"teaching-backend/pkg/encrypt"
	"teaching-backend/pkg/jwt"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
	"golang.org/x/crypto/bcrypt"
)

type LoginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户登录
func NewLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *LoginLogic {
	return &LoginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *LoginLogic) Login(req *types.LoginReq) (resp *types.LoginRes, err error) {
	// 1. 校验参数
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	if req.Email == "" {
		return nil, code.EmailEmpty
	}
	if req.Password == "" {
		return nil, code.PasswordEmpty
	}

	email, err := encrypt.EncEmail(req.Email)
	if err != nil {
		l.Errorf("邮箱加密失败: %v", err)
		return nil, xcode.ServerErr
	}

	// 检查用户是否已存在
	u, err := l.svcCtx.UserRPC.FindByEmail(l.ctx, &user.FindByEmailReq{
		Email: email,
	})
	if err != nil {
		l.Errorf("查询用户失败: %v", err)
		return nil, err
	}
	if u == nil || u.UserId <= 0 {
		return nil, code.UserNotFound
	}

	// 校验密码
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return nil, code.PasswordIncorrect
	}

	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": u.UserId,
			"role":   u.Role,
		},
	})
	if err != nil {
		l.Errorf("生成token失败: %v", err)
		return nil, err
	}

	return &types.LoginRes{
		UserId: u.UserId,
		Token:  token.AccessToken,
		Role:   int(u.Role),
	}, nil
}
