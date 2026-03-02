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
)
