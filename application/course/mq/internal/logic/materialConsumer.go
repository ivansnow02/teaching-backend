package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"teaching-backend/application/course/mq/internal/svc"
	"teaching-backend/pkg/kafkatypes"

	"github.com/zeromicro/go-zero/core/logx"
)

// MaterialConsumer 课件变更消费者（Canal Binlog -> Kafka -> 此处消费）
type MaterialConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMaterialConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialConsumer {
	return &MaterialConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Consume 实现 kq.ConsumeHandler 接口
func (c *MaterialConsumer) Consume(ctx context.Context, key, value string) error {
	logx.Infof("MaterialConsumer 收到消息: key=%s", key)

	var msg kafkatypes.CanalMsg
	if err := json.Unmarshal([]byte(value), &msg); err != nil {
		logx.Errorf("解析 Canal 消息失败: %v, value=%s", err, value)
		return nil // 格式错误不重试
	}

	// 只处理 course_material 表的 INSERT 和 UPDATE
	if msg.Table != "course_material" {
		return nil
	}
	if msg.Type != "INSERT" && msg.Type != "UPDATE" {
		return nil
	}

	for _, row := range msg.Data {
		if err := c.processMaterial(ctx, row); err != nil {
			logx.Errorf("处理课件向量化失败: %v", err)
		}
	}

	return nil
}

// processMaterial 处理单条课件数据，触发 AI 向量化
func (c *MaterialConsumer) processMaterial(ctx context.Context, row map[string]any) error {
	idStr, _ := row["id"].(string)
	id, _ := strconv.ParseInt(idStr, 10, 64)
	url, _ := row["url"].(string)
	fileHash, _ := row["file_hash"].(string)

	if id == 0 || url == "" {
		return nil
	}

	logx.Infof("触发课件向量化: id=%d, url=%s, fileHash=%s", id, url, fileHash)

	// TODO: 调用 AI RPC 进行向量化处理
	// aiResp, err := c.svcCtx.AiRPC.Vectorize(ctx, &ai.VectorizeReq{
	//     FileUrl:  url,
	//     FileHash: fileHash,
	// })
	// if err != nil {
	//     return fmt.Errorf("调用 AI RPC 失败: %w", err)
	// }

	// 更新课件的 ai_status（1=处理中, 2=已完成）
	_, err := c.svcCtx.Conn.ExecCtx(ctx, "update `course_material` set `ai_status` = ? where `id` = ?", 1, id)
	if err != nil {
		return fmt.Errorf("更新课件 ai_status 失败: %w", err)
	}

	logx.Infof("课件 id=%d 已标记为处理中", id)
	return nil
}
