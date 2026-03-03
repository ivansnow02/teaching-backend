package logic

import (
	"context"
	"encoding/json"

	"teaching-backend/application/course/mq/internal/svc"
	"teaching-backend/pkg/kafkatypes"

	"github.com/zeromicro/go-zero/core/logx"
)

// StudyProgressConsumer 学习进度消费者（API -> Kafka -> 此处消费落盘）
type StudyProgressConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewStudyProgressConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *StudyProgressConsumer {
	return &StudyProgressConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// Consume 实现 kq.ConsumeHandler 接口
func (c *StudyProgressConsumer) Consume(ctx context.Context, key, value string) error {
	logx.Infof("【Debug】StudyProgressConsumer 进入 Consume: key=%s, value=%s", key, value)

	var msg kafkatypes.StudyProgressMsg
	if err := json.Unmarshal([]byte(value), &msg); err != nil {
		logx.Errorf("解析学习进度消息失败: %v, value=%s", err, value)
		return nil
	}

	if msg.UserId <= 0 || msg.CourseId <= 0 {
		logx.Errorf("学习进度消息参数无效: %+v", msg)
		return nil
	}

	// 使用 INSERT ... ON DUPLICATE KEY UPDATE 实现 Upsert
	query := "insert into `study_progress` (`user_id`, `course_id`, `chapter_id`, `material_id`, `progress`) " +
		"values (?, ?, ?, ?, ?) " +
		"on duplicate key update `progress` = values(`progress`), `chapter_id` = values(`chapter_id`)"

	_, err := c.svcCtx.Conn.ExecCtx(ctx, query,
		msg.UserId, msg.CourseId, msg.ChapterId, msg.MaterialId, msg.Progress,
	)
	if err != nil {
		logx.Errorf("持久化学习进度失败: %v, msg=%+v", err, msg)
		return err // 返回 err 触发重试
	}

	logx.Infof("学习进度已落盘: userId=%d, courseId=%d, materialId=%d, progress=%d",
		msg.UserId, msg.CourseId, msg.MaterialId, msg.Progress)
	return nil
}
