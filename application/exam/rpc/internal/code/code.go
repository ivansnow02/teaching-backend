package code

import "teaching-backend/pkg/xcode"

var (
	QuestionNotFound     = xcode.New(300001, "题目不存在")
	ExamNotFound         = xcode.New(300002, "考试不存在")
	ExamRecordNotFound   = xcode.New(300003, "考试记录不存在")
	NoPermission         = xcode.New(300004, "无权操作该资源")
	ExamAlreadySubmitted = xcode.New(300005, "考试已提交，无法重复操作")
	QuestionInUse        = xcode.New(300006, "题目正在被试卷使用，无法删除")
	InvalidMemberType    = xcode.New(300007, "无效的成员身份")
)
