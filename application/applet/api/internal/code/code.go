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

	// Enrollment 2001xx
	CourseEnrollFailed    = xcode.New(200101, "选课失败")
	CourseDropFailed      = xcode.New(200102, "退课失败")
	GetEnrollmentFailed   = xcode.New(200103, "获取选课列表失败")
	CheckEnrollmentFailed = xcode.New(200104, "查询选课状态失败")
	GetStudentsFailed     = xcode.New(200105, "获取学生列表失败")
	GetStudentInfoFailed  = xcode.New(200106, "获取学生信息失败")

	// StudyProgress 2002xx
	UpdateProgressFailed = xcode.New(200201, "更新学习进度失败")
)
