package config

import (
	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	service.ServiceConf

	// Canal Binlog 消费 (用于异步 AI 批改)
	CanalKq kq.KqConf

	// 考试提交通知消费 (用于异步落库和计算)
	SubmitExamKq kq.KqConf

	DataSource string
	CacheRedis cache.CacheConf

	ExamRPC zrpc.RpcClientConf
	// AiRPC zrpc.RpcClientConf (预留)
}
