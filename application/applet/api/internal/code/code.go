package code

import "teaching-backend/pkg/xcode"

var (
	RegisterEmailEmpty      = xcode.New(100001, "注册邮箱不能为空")
	VerificationCodeEmpty   = xcode.New(100002, "验证码不能为空")
	EmailHasRegistered      = xcode.New(100003, "邮箱已经注册")
	LoginEmailEmpty         = xcode.New(100004, "邮箱不能为空")
	RegisterPasswdEmpty      = xcode.New(100005, "密码不能为空")
	VerificationCodeError   = xcode.New(100006, "验证码错误")
	VerificationCodeExpired = xcode.New(100007, "验证码已过期")
	VerificationCodeLimitPerDay = xcode.New(100008, "验证码发送次数已达上限")
)
