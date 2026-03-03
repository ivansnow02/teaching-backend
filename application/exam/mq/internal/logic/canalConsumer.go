package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"strconv"
	"time"

	"teaching-backend/application/exam/mq/internal/svc"
	"teaching-backend/pkg/kafkatypes"

	"github.com/zeromicro/go-zero/core/logx"
)

type CanalConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCanalConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *CanalConsumer {
	return &CanalConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *CanalConsumer) Consume(ctx context.Context, key, value string) error {
	var msg kafkatypes.CanalMsg
	if err := json.Unmarshal([]byte(value), &msg); err != nil {
		logx.Errorf("解析 Canal 消息失败: %v, value=%s", err, value)
		return nil
	}

	logx.Infof("CanalConsumer 收到消息: DB=%s, Table=%s, Type=%s", msg.Database, msg.Table, msg.Type)

	// 过滤需要的表和类型
	if msg.Database != "teaching_exam" {
		return nil
	}

	switch msg.Table {
	case "user_answer":
		return c.handleUserAnswer(ctx, msg)
	}

	return nil
}

func (c *CanalConsumer) handleUserAnswer(ctx context.Context, msg kafkatypes.CanalMsg) error {
	if msg.Type != "INSERT" && msg.Type != "UPDATE" {
		return nil
	}

	for _, row := range msg.Data {
		// 转换 ID
		idVal, ok := row["id"]
		if !ok {
			continue
		}
		var id uint64
		switch v := idVal.(type) {
		case float64:
			id = uint64(v)
		case json.Number:
			idUint, _ := strconv.ParseUint(v.String(), 10, 64)
			id = idUint
		case string:
			idUint, _ := strconv.ParseUint(v, 10, 64)
			id = idUint
		default:
			logx.Errorf("CanalConsumer: 无法解析 id 类型: %T", idVal)
			continue
		}

		aiStatusVal, ok := row["ai_status"]
		if !ok {
			continue
		}
		var aiStatus int64
		switch v := aiStatusVal.(type) {
		case float64:
			aiStatus = int64(v)
		case json.Number:
			aiStatus, _ = v.Int64()
		case string:
			aiStatus, _ = strconv.ParseInt(v, 10, 64)
		}

		// 只有 ai_status 为 0 (待处理) 时才触发
		if aiStatus != 0 {
			continue
		}

		// 查询该回答对应的题目类型
		ua, err := c.svcCtx.UserAnswerModel.FindOne(ctx, id)
		if err != nil {
			logx.Errorf("CanalConsumer: FindOne(ua_id=%d) error: %v", id, err)
			continue
		}

		q, err := c.svcCtx.QuestionModel.FindOne(ctx, ua.QuestionId)
		if err != nil {
			logx.Errorf("CanalConsumer: QuestionModel.FindOne(qid=%d) error: %v", ua.QuestionId, err)
			continue
		}

		// 假设类型 > 3 为主观题 (如 4:简答, 5:编程)
		if q.Type <= 3 {
			continue
		}

		logx.Infof("CanalConsumer: 检测到主观题(id=%d, type=%d), 准备触发异步 AI 批改...", id, q.Type)

		// TODO: 调用 AI LLM 接口
		// 此处模拟 AI 异步处理
		go func(uaId uint64) {
			// 延迟模拟
			time.Sleep(2 * time.Second)

			// 模拟回写 AI 评分和状态
			uaUpdate, err := c.svcCtx.UserAnswerModel.FindOne(context.Background(), uaId)
			if err != nil {
				return
			}
			uaUpdate.AiStatus = 2 // 已完成
			uaUpdate.AiComment = sql.NullString{String: "AI 自动评分：逻辑清晰，回答完整。", Valid: true}
			uaUpdate.Score = 10.0 // 模拟给 10 分
			_ = c.svcCtx.UserAnswerModel.Update(context.Background(), uaUpdate)

			logx.Infof("CanalConsumer: AI 批改记录(id=%d)完成", uaId)

			// 检查是否全卷批改完成，并更新总分
			c.checkAndFinalizeExam(context.Background(), uaUpdate.RecordId)
		}(id)
	}
	return nil
}

func (c *CanalConsumer) checkAndFinalizeExam(ctx context.Context, recordId uint64) {
	// 1. 获取该记录下所有的回答
	answers, err := c.svcCtx.UserAnswerModel.FindListByRecordId(ctx, recordId)
	if err != nil {
		logx.Errorf("checkAndFinalizeExam: FindListByRecordId(record_id=%d) error: %v", recordId, err)
		return
	}

	// 2. 检查是否还有题目处于待批改状态 (ai_status == 0)
	var totalScore float64
	allCompleted := true
	for _, ans := range answers {
		if ans.AiStatus == 0 {
			allCompleted = false
			break
		}
		totalScore += ans.Score
	}

	// 3. 如果全部完成，更新考试记录
	if allCompleted {
		record, err := c.svcCtx.UserExamRecordModel.FindOne(ctx, recordId)
		if err != nil {
			logx.Errorf("checkAndFinalizeExam: FindOne record error: %v", err)
			return
		}

		record.Score = totalScore
		record.Status = 2 // 已完成
		err = c.svcCtx.UserExamRecordModel.Update(ctx, record)
		if err != nil {
			logx.Errorf("checkAndFinalizeExam: Update record error: %v", err)
			return
		}
		logx.Infof("checkAndFinalizeExam: 考试记录(id=%d)已全卷批改完成，最终得分: %v", recordId, totalScore)
	}
}
