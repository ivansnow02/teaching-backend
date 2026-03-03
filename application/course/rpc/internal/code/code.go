package code

import "teaching-backend/pkg/xcode"

var (
	CourseNotFound     = xcode.New(200001, "课程不存在")
	ChapterNotFound    = xcode.New(200002, "章节不存在")
	MaterialNotFound   = xcode.New(200003, "课件不存在")
	CourseTitleEmpty   = xcode.New(200004, "课程标题不能为空")
	NoPermission       = xcode.New(200005, "无权操作该资源")
	ChapterTitleEmpty  = xcode.New(200006, "章节标题不能为空")
	MaterialTitleEmpty = xcode.New(200007, "课件标题不能为空")
	NotEnrolled        = xcode.New(200008, "未选该课程")
	AlreadyEnrolled    = xcode.New(200009, "已选该课程")

	// Enrollment 2001xx
	EnrollFailed        = xcode.New(200101, "选课失败")
	DropFailed          = xcode.New(200102, "退课失败")
	GetEnrollmentFailed = xcode.New(200103, "获取选课列表失败")
	GetStudentsFailed   = xcode.New(200104, "获取已选学生失败")
)
