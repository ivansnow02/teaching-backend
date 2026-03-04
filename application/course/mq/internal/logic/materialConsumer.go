package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"strconv"

	"teaching-backend/application/ai/rpc/aibridge"
	"teaching-backend/application/course/mq/internal/svc"
	"teaching-backend/application/course/rpc/course"
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

func (c *MaterialConsumer) processMaterial(ctx context.Context, row map[string]any) error {
	idStr, _ := row["id"].(string)
	id, _ := strconv.ParseInt(idStr, 10, 64)
	url, _ := row["url"].(string)
	courseIdStr, _ := row["course_id"].(string)
	courseId, _ := strconv.ParseInt(courseIdStr, 10, 64)
	title, _ := row["title"].(string)
	typeStr, _ := row["type"].(string)
	materialType, _ := strconv.Atoi(typeStr)

	if id == 0 || url == "" {
		return nil
	}

	logx.Infof("触发课件向量化: id=%d, url=%s, type=%d, title=%s, courseId=%d", id, url, materialType, title, courseId)

	// 调用 AI RPC 进行向量化处理
	_, err := c.svcCtx.AiRPC.EmbedMaterial(ctx, &aibridge.EmbedMaterialReq{
		MaterialId: id,
		CourseId:   courseId,
		Title:      title,
		Url:        url,
		Type:       int32(materialType),
	})
	if err != nil {
		logx.Errorf("调用 AI RPC 失败 material_id: %d err: %v", id, err)
		// 修改数据库状态为失败
		_, dbErr := c.svcCtx.CourseRPC.UpdateMaterialAiStatus(ctx, &course.UpdateMaterialAiStatusReq{
			MaterialId: id,
			AiStatus:   3,
		})
		if dbErr != nil {
			logx.Errorf("更新课件 ai_status 失败 material_id: %d err: %v", id, dbErr)
		}
		// 消费异常通常不阻塞进度，所以只打印日志
		return nil
	}

	// 更新课件的 ai_status（1=处理中）
	_, err = c.svcCtx.CourseRPC.UpdateMaterialAiStatus(ctx, &course.UpdateMaterialAiStatusReq{
		MaterialId: id,
		AiStatus:   1,
	})
	if err != nil {
		return fmt.Errorf("更新课件 ai_status 失败: %w", err)
	}

	logx.Infof("课件 id=%d 已提交向量化并标记为处理中", id)
	return nil
}
