package config

import (
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/core/stores/cache"
	"github.com/zeromicro/go-queue/kq"
)

type Config struct {
	service.ServiceConf
	MaterialKq      kq.KqConf
	StudyProgressKq kq.KqConf
	DataSource      string
	CacheRedis      cache.CacheConf
}
