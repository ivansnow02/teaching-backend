package code

import "teaching-backend/pkg/xcode"

var (
	RegisterNameEmpty     = xcode.New(10001, "注册名字不能为空")
	RegisterPasswordEmpty = xcode.New(10002, "注册密码不能为空")
	RegisterCodeEmpty     = xcode.New(10003, "注册验证码不能为空")
	RegisterRoleEmpty     = xcode.New(10004, "注册角色不能为空")
	EmailEmpty            = xcode.New(10005, "邮箱不能为空")
	EmailInvalid          = xcode.New(10006, "邮箱格式不正确")
	VerificationCodeError = xcode.New(10007, "验证码错误")
	UserAlreadyRegister   = xcode.New(10008, "该用户已注册")
	UserNotRegister       = xcode.New(10009, "该用户未注册")
	PasswordIncorrect     = xcode.New(10010, "密码错误")
)
