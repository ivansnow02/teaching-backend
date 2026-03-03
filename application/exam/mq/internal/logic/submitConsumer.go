package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"teaching-backend/application/exam/mq/internal/svc"
	"teaching-backend/application/exam/rpc/model"
	"teaching-backend/pkg/kafkatypes"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitConsumer struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSubmitConsumer(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitConsumer {
	return &SubmitConsumer{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (c *SubmitConsumer) Consume(ctx context.Context, key, value string) error {
	logx.Infof("SubmitConsumer 收到交卷通知: key=%s", key)

	var msg kafkatypes.SubmitExamMsg
	if err := json.Unmarshal([]byte(value), &msg); err != nil {
		logx.Errorf("解析交卷消息失败: %v, value=%s", err, value)
		return nil
	}

	// 1. 获取考试记录
	record, err := c.svcCtx.UserExamRecordModel.FindOne(ctx, uint64(msg.RecordId))
	if err != nil {
		logx.Errorf("SubmitConsumer: 未找到考试记录 id=%d", msg.RecordId)
		return nil
	}
	// 状态 1 表示处理中/待批改，异步交卷会将状态先置为 1
	if record.Status != 1 {
		logx.Infof("SubmitConsumer: 记录 id=%d 状态为 %d, 跳过处理", msg.RecordId, record.Status)
		return nil
	}

	// 2. 获取试卷题目关联信息（用于获取每道题在试卷中的分值）
	examQuestions, err := c.svcCtx.ExamQuestionModel.FindListByExamId(ctx, record.ExamId)
	if err != nil {
		logx.Errorf("ExamQuestionModel.FindListByExamId(examId=%d) error: %v", record.ExamId, err)
		return err
	}

	ansMap := make(map[int64]string)
	for _, a := range msg.Answers {
		ansMap[a.QuestionId] = a.Answer
	}

	var totalScore float64
	hasSubjective := false

	for _, eq := range examQuestions {
		userAnswerStr := ansMap[int64(eq.QuestionId)]

		// 获取题目详细内容（用于获取正确答案和类型）
		q, err := c.svcCtx.QuestionModel.FindOne(ctx, eq.QuestionId)
		if err != nil {
			logx.Errorf("QuestionModel.FindOne(qid=%d) error: %v", eq.QuestionId, err)
			continue
		}

		isCorrect := int64(0)
		score := 0.0

		// 异步判分核心逻辑
		switch q.Type {
		case 1, 3: // 单选、判断
			if userAnswerStr == q.Answer {
				isCorrect = 1
				score = eq.Score
			}
		case 2: // 多选
			if userAnswerStr == q.Answer {
				isCorrect = 1
				score = eq.Score
			}
		default:
			// 含有非客观题，标记为有主观部分
			hasSubjective = true
		}
		totalScore += score

		// 更新或插入用户答案表
		ua, err := c.svcCtx.UserAnswerModel.FindOneByRecordIdQuestionId(ctx, record.Id, eq.QuestionId)
		if err == nil {
			ua.UserAnswer = sql.NullString{String: userAnswerStr, Valid: true}
			ua.IsCorrect = isCorrect
			ua.Score = score
			_ = c.svcCtx.UserAnswerModel.Update(ctx, ua)
		} else {
			_, _ = c.svcCtx.UserAnswerModel.Insert(ctx, &model.UserAnswer{
				RecordId:   record.Id,
				QuestionId: eq.QuestionId,
				UserAnswer: sql.NullString{String: userAnswerStr, Valid: true},
				IsCorrect:  isCorrect,
				Score:      score,
				AiStatus:   0,
			})
		}
	}

	// 3. 更新考试记录终态分数与状态
	record.Score = totalScore
	if hasSubjective {
		record.Status = 1 // 维持 待批改 状态
	} else {
		record.Status = 2 // 全部客观题判分完成 -> 已完成
	}
	record.SubmitTime = sql.NullTime{Time: time.Now(), Valid: true}

	err = c.svcCtx.UserExamRecordModel.Update(ctx, record)
	if err != nil {
		logx.Errorf("SubmitConsumer: 更新记录 id=%d 最终状态失败: %v", record.Id, err)
		return err
	}

	logx.Infof("SubmitConsumer: 记录 id=%d 处理成功, 总分: %v, 是否含主观题: %v", record.Id, totalScore, hasSubjective)
	return nil
}
