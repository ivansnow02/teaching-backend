package code

import "teaching-backend/pkg/xcode"

var (
	EmailEmpty                  = xcode.New(100001, "邮箱不能为空")
	PasswordEmpty               = xcode.New(100002, "密码不能为空")
	VerificationCodeEmpty       = xcode.New(100003, "验证码不能为空")
	EmailHasRegistered          = xcode.New(100004, "邮箱已经注册")
	UserNotFound                = xcode.New(100005, "用户不存在")
	PasswordIncorrect           = xcode.New(100006, "密码错误")
	VerificationCodeError       = xcode.New(100007, "验证码错误")
	VerificationCodeExpired     = xcode.New(100008, "验证码已过期")
	VerificationCodeLimitPerDay = xcode.New(100009, "验证码发送次数已达上限")

	// Course 2000xx
	CourseNotFound     = xcode.New(200001, "课程不存在")
	ChapterNotFound    = xcode.New(200002, "章节不存在")
	MaterialNotFound   = xcode.New(200003, "课件不存在")
	CourseTitleEmpty   = xcode.New(200004, "课程标题不能为空")
	NoPermission       = xcode.New(200005, "无权操作该资源")
	ChapterTitleEmpty  = xcode.New(200006, "章节标题不能为空")
	MaterialTitleEmpty = xcode.New(200007, "课件标题不能为空")
)
