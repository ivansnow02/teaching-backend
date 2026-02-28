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
	"teaching-backend/pkg/util"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"golang.org/x/crypto/bcrypt"
)

type RegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 用户注册
func NewRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RegisterLogic {
	return &RegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RegisterLogic) Register(req *types.RegisterReq) (resp *types.RegisterRes, err error) {
	// 1. 校验参数
	req.Email = strings.TrimSpace(req.Email)
	req.Password = strings.TrimSpace(req.Password)
	req.Code = strings.TrimSpace(req.Code)
	if req.Email == "" {
		return nil, code.RegisterEmailEmpty
	}
	if req.Password == "" {
		return nil, code.RegisterPasswdEmpty
	}
	if req.Code == "" {
		return nil, code.VerificationCodeEmpty
	}

	err = checkCode(l.svcCtx.BizRedis, req.Email, req.Code)
	if err != nil {
		l.Errorf("验证码错误: %v", err)
		return nil, code.VerificationCodeError
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
		return nil, xcode.ServerErr
	}
	if u != nil && u.UserId > 0 {
		return nil, code.EmailHasRegistered
	}
	// 密码加密
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		l.Errorf("密码加密失败: %v", err)
		return nil, xcode.ServerErr
	}

	// 生成默认昵称
	nickname := generateDefaultNickname(req.Email)

	regRet, err := l.svcCtx.UserRPC.Register(l.ctx, &user.RegisterReq{
		Email:    email,
		Password: string(hashed),
		Role:     int32(req.Role),
		Nickname: nickname,
	})
	if err != nil {
		l.Errorf("注册失败: %v", err)
		return nil, xcode.ServerErr
	}

	token, err := jwt.BuildTokens(jwt.TokenOptions{
		AccessSecret: l.svcCtx.Config.Auth.AccessSecret,
		AccessExpire: l.svcCtx.Config.Auth.AccessExpire,
		Fields: map[string]interface{}{
			"userId": regRet.UserId,
			"role":   req.Role,
		},
	})
	if err != nil {
		l.Errorf("生成token失败: %v", err)
		return nil, xcode.ServerErr
	}

	if err := delActivationCache(l.svcCtx.BizRedis, req.Email); err != nil {
		l.Errorf("删除激活缓存失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &types.RegisterRes{
		UserId: regRet.UserId,
		Token:  token.AccessToken,
	}, nil
}

func checkCode(rds *redis.Redis, email, reqCode string) error {
	cacheCode, err := getActivationCache(rds, email)
	if err != nil {
		return err
	}
	if cacheCode == "" {
		return code.VerificationCodeExpired
	}
	if cacheCode != reqCode {
		return code.VerificationCodeError
	}
	return nil
}

func generateDefaultNickname(email string) string {
	prefix := email
	if idx := strings.Index(email, "@"); idx > 0 {
		prefix = email[:idx]
	}
	if len(prefix) > 10 {
		prefix = prefix[:10]
	}
	return prefix + "_" + util.RandomNumeric(4)
}
