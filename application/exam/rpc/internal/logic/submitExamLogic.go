package logic

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/kafkatypes"
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

	// 合并提交答案与快照答案
	key := fmt.Sprintf("exam:snapshot:%d", in.RecordId)
	snapshots, err := l.svcCtx.BizRedis.HgetallCtx(l.ctx, key)
	if err != nil {
		l.Errorf("SubmitExam HgetallCtx error: %v", err)
	}

	ansMap := make(map[uint64]string)
	for k, v := range snapshots {
		qid, _ := strconv.ParseUint(k, 10, 64)
		ansMap[qid] = v
	}
	// 再由本次提交覆盖
	for _, a := range in.Answers {
		ansMap[uint64(a.QuestionId)] = a.Answer
	}

	// 组装消息并推送到 Kafka
	msg := kafkatypes.SubmitExamMsg{
		RecordId: int64(record.Id),
		Answers:  make([]kafkatypes.SubmitAnswerItem, 0, len(ansMap)),
	}
	for qid, ans := range ansMap {
		msg.Answers = append(msg.Answers, kafkatypes.SubmitAnswerItem{
			QuestionId: int64(qid),
			Answer:     ans,
		})
	}

	body, err := json.Marshal(msg)
	if err != nil {
		l.Errorf("序列化交卷消息失败: %v", err)
		return nil, xcode.ServerErr
	}

	// 推送异步落库消息
	err = l.svcCtx.SubmitExamPusher.PushWithKey(l.ctx, strconv.FormatUint(record.Id, 10), string(body))
	if err != nil {
		l.Errorf("推送交卷消息到 Kafka 失败: %v", err)
		return nil, xcode.ServerErr
	}

	// 预更新记录状态为“处理中/批改中”
	record.Status = 1 // 1=待批改/处理中
	record.SubmitTime = sql.NullTime{Time: time.Now(), Valid: true}
	_ = l.svcCtx.UserExamRecordModel.Update(l.ctx, record)

	// 清理快照
	_, _ = l.svcCtx.BizRedis.DelCtx(l.ctx, key)

	return &pb.SubmitExamRes{
		Success: true,
	}, nil
}
