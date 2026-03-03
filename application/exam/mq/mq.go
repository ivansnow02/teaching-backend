package main

import (
	"context"
	"flag"
	"fmt"

	"teaching-backend/application/exam/mq/internal/config"
	"teaching-backend/application/exam/mq/internal/logic"
	"teaching-backend/application/exam/mq/internal/svc"

	"github.com/zeromicro/go-queue/kq"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
)

var configFile = flag.String("f", "etc/mq.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	svcCtx := svc.NewServiceContext(c)
	ctx := context.Background()

	group := service.NewServiceGroup()
	defer group.Stop()

	// 1. Canal 消费者（处理 AI 批改）
	canalConsumer := logic.NewCanalConsumer(ctx, svcCtx)
	group.Add(kq.MustNewQueue(c.CanalKq, canalConsumer))

	// 2. 交卷事情消费者（处理异步落库/通知）
	submitConsumer := logic.NewSubmitConsumer(ctx, svcCtx)
	group.Add(kq.MustNewQueue(c.SubmitExamKq, submitConsumer))

	fmt.Printf("Starting exam-mq consumer: %s...\n", c.Name)
	group.Start()
}
