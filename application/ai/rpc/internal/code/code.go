package code

import "teaching-backend/pkg/xcode"

var (
	AiServiceUnavailable = xcode.New(400001, "AI服务不可用")
	AiGradingFailed      = xcode.New(400002, "AI判卷失败")
	AiEmbedFailed        = xcode.New(400003, "课件向量化失败")
	AiTaskNotFound       = xcode.New(400004, "AI任务不存在")
	AiStreamError        = xcode.New(400005, "AI流式输出异常")
	AiGenerateFailed     = xcode.New(400006, "AI生成失败")
)
