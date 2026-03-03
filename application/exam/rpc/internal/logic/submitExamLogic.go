package logic

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"time"

	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/model"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type SubmitExamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSubmitExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitExamLogic {
	return &SubmitExamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 提交答卷
func (l *SubmitExamLogic) SubmitExam(in *pb.SubmitExamReq) (*pb.SubmitExamRes, error) {
	record, err := l.svcCtx.UserExamRecordModel.FindOne(l.ctx, uint64(in.RecordId))
	if err != nil {
		l.Errorf("UserExamRecordModel.FindOne error: %v", err)
		return nil, code.ExamRecordNotFound
	}
	if record.Status != 0 {
		return nil, code.ExamAlreadySubmitted
	}

	// 获取试卷题目列表
	examQuestions, err := l.svcCtx.ExamQuestionModel.FindListByExamId(l.ctx, record.ExamId)
	if err != nil {
		l.Errorf("ExamQuestionModel.FindListByExamId error: %v", err)
		return nil, xcode.ServerErr
	}

	// 合并提交答案与快照答案
	ansMap := make(map[uint64]string)
	// 先取快照
	key := fmt.Sprintf("exam:snapshot:%d", in.RecordId)
	snapshots, err := l.svcCtx.BizRedis.HgetallCtx(l.ctx, key)
	if err != nil {
		l.Errorf("SubmitExam HgetallCtx error: %v", err)
	}
	for k, v := range snapshots {
		qid, _ := strconv.ParseUint(k, 10, 64)
		ansMap[qid] = v
	}
	// 再由本次提交覆盖
	for _, a := range in.Answers {
		ansMap[uint64(a.QuestionId)] = a.Answer
	}

	var totalScore float64
	hasSubjective := false // 是否有简答题等需要人工/交互式批改的

	for _, eq := range examQuestions {
		userAnswerStr := ansMap[eq.QuestionId]

		q, err := l.svcCtx.QuestionModel.FindOne(l.ctx, eq.QuestionId)
		if err != nil {
			l.Errorf("QuestionModel.FindOne error: %v", err)
			continue
		}

		isCorrect := int64(0)
		score := 0.0

		// 简单自动判分逻辑
		switch q.Type {
		case 1, 3: // 单选(1)、判断(3)
			if userAnswerStr == q.Answer {
				isCorrect = 1
				score = eq.Score
			}
		case 2: // 多选匹配
			if userAnswerStr == q.Answer {
				isCorrect = 1
				score = eq.Score
			}
		default:
			hasSubjective = true
		}

		totalScore += score

		// 保存用户答案到DB (尝试 FindOneByRecordIdQuestionId 来支持重复提交时的幂等)
		ua, err := l.svcCtx.UserAnswerModel.FindOneByRecordIdQuestionId(l.ctx, record.Id, eq.QuestionId)
		if err == nil {
			ua.UserAnswer = sql.NullString{String: userAnswerStr, Valid: true}
			ua.IsCorrect = isCorrect
			ua.Score = score
			err = l.svcCtx.UserAnswerModel.Update(l.ctx, ua)
			if err != nil {
				l.Errorf("UserAnswerModel.Update error: %v", err)
			}
		} else {
			_, err = l.svcCtx.UserAnswerModel.Insert(l.ctx, &model.UserAnswer{
				RecordId:   record.Id,
				QuestionId: eq.QuestionId,
				UserAnswer: sql.NullString{String: userAnswerStr, Valid: true},
				IsCorrect:  isCorrect,
				Score:      score,
				AiStatus:   0,
			})
			if err != nil {
				l.Errorf("UserAnswerModel.Insert error: %v", err)
			}
		}
	}

	// 更新记录
	record.Score = totalScore
	if hasSubjective {
		record.Status = 1 // 待批改
	} else {
		record.Status = 2 // 已完成
	}
	record.SubmitTime = sql.NullTime{Time: time.Now(), Valid: true}

	err = l.svcCtx.UserExamRecordModel.Update(l.ctx, record)
	if err != nil {
		l.Errorf("UserExamRecordModel.Update error: %v", err)
		return nil, xcode.ServerErr
	}

	// 清理快照
	_, _ = l.svcCtx.BizRedis.DelCtx(l.ctx, key)

	return &pb.SubmitExamRes{
		Success: true,
	}, nil
}
