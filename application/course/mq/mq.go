package main

import (
	"context"
	"flag"
	"fmt"

	"teaching-backend/application/course/mq/internal/config"
	"teaching-backend/application/course/mq/internal/logic"
	"teaching-backend/application/course/mq/internal/svc"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-queue/kq"
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

	// 课件向量化消费者（Canal Binlog）
	materialConsumer := logic.NewMaterialConsumer(ctx, svcCtx)
	group.Add(kq.MustNewQueue(c.MaterialKq, materialConsumer))

	// 学习进度消费者（API 主动推送）
	studyProgressConsumer := logic.NewStudyProgressConsumer(ctx, svcCtx)
	group.Add(kq.MustNewQueue(c.StudyProgressKq, studyProgressConsumer))

	fmt.Println("Starting course-mq consumer...")
	group.Start()
}
