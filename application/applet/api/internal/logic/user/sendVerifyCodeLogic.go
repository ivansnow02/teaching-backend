// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package user

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"teaching-backend/application/applet/api/internal/code"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/pkg/util"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/stores/redis"
	"github.com/zeromicro/go-zero/core/threading"
)

const (
	prefixVerificationCount = "biz#verification#count#%s"
	verificationLimitPerDay = 10
	expireActivation        = 60 * 30
	prefixActivation        = "biz#activation#%s"
)

type SendVerifyCodeLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 发送验证码
func NewSendVerifyCodeLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendVerifyCodeLogic {
	return &SendVerifyCodeLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SendVerifyCodeLogic) SendVerifyCode(req *types.SendVerifyCodeReq) (resp *types.SendVerifyCodeRes, err error) {
	cnt, err := getVerificationCount(l.svcCtx.BizRedis, req.Email)
	if err != nil {
		l.Errorf("获取验证码次数失败: %v", err)
		return nil, xcode.ServerErr
	}
	if cnt >= verificationLimitPerDay {
		return nil, code.VerificationCodeLimitPerDay
	}

	c, err := getActivationCache(l.svcCtx.BizRedis, req.Email)
	if err != nil {
		l.Errorf("获取验证码失败: %v", err)
		return nil, xcode.ServerErr
	}
	if len(c) == 0 {
		c = util.RandomNumeric(6)
	}

	// 先存 Redis + 计数（同步，保证数据一致性）
	if err := saveActivationCache(l.svcCtx.BizRedis, req.Email, c); err != nil {
		l.Errorf("保存验证码失败: %v", err)
		return nil, xcode.ServerErr
	}

	if err := incrVerificationCount(l.svcCtx.BizRedis, req.Email); err != nil {
		l.Errorf("增加验证码次数失败: %v", err)
		return nil, xcode.ServerErr
	}

	// 异步发送邮件，不阻塞用户请求
	// 如果发送失败只记日志，用户可点击重新发送（复用缓存中的同一个验证码）
	threading.GoSafe(func() {
		if err := sendMail(req.Email, c); err != nil {
			logx.Errorf("异步发送验证码邮件失败, email: %s, err: %v", req.Email, err)
		}
	})

	return &types.SendVerifyCodeRes{}, nil
}

func sendMail(email, code string) error {
	// todo
	return nil
}

func getVerificationCount(rds *redis.Redis, email string) (int, error) {
	key := fmt.Sprintf(prefixVerificationCount, email)
	val, err := rds.Get(key)
	if err != nil {
		return 0, err
	}
	if len(val) == 0 {
		return 0, nil
	}
	return strconv.Atoi(val)
}

// Lua 脚本: INCR + EXPIREAT 原子操作
// KEYS[1] = key, ARGV[1] = 过期时间戳
// 只在第一次 INCR（结果为1）时设置过期时间
const luaIncrWithExpire = `
local cnt = redis.call("INCR", KEYS[1])
if cnt == 1 then
    redis.call("EXPIREAT", KEYS[1], ARGV[1])
end
return cnt
`

func incrVerificationCount(rds *redis.Redis, email string) error {
	key := fmt.Sprintf(prefixVerificationCount, email)
	expireAt := strconv.FormatInt(util.EndOfDay(time.Now()).Unix(), 10)
	_, err := rds.Eval(luaIncrWithExpire, []string{key}, []string{expireAt})
	return err
}

func getActivationCache(rds *redis.Redis, email string) (string, error) {
	key := fmt.Sprintf(prefixActivation, email)
	return rds.Get(key)
}

func saveActivationCache(rds *redis.Redis, email, code string) error {
	key := fmt.Sprintf(prefixActivation, email)
	logx.Infof("code is: %s", code)
	return rds.Setex(key, code, expireActivation)
}

func delActivationCache(rds *redis.Redis, email string) error {
	key := fmt.Sprintf(prefixActivation, email)
	_, err := rds.Del(key)
	return err
}
