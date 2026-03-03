package logic

import (
	"context"
	"fmt"
	"strconv"

	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAnswerSnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSaveAnswerSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAnswerSnapshotLogic {
	return &SaveAnswerSnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 考试心跳与答案快照保存 (存Redis)
func (l *SaveAnswerSnapshotLogic) SaveAnswerSnapshot(in *pb.SaveAnswerSnapshotReq) (*pb.SaveAnswerSnapshotRes, error) {
	// 校验记录是否存在
	_, err := l.svcCtx.UserExamRecordModel.FindOne(l.ctx, uint64(in.RecordId))
	if err != nil {
		l.Errorf("UserExamRecordModel.FindOne error: %v", err)
		return nil, code.ExamRecordNotFound
	}

	// 存入Redis快照
	key := fmt.Sprintf("exam:snapshot:%d", in.RecordId)

	for _, ans := range in.Answers {
		err = l.svcCtx.BizRedis.HsetCtx(l.ctx, key, strconv.FormatInt(ans.QuestionId, 10), ans.Answer)
		if err != nil {
			l.Errorf("SaveAnswerSnapshot HsetCtx error: %v", err)
			return nil, xcode.ServerErr
		}
	}

	// 设置过期时间 (例如 24 小时)
	_ = l.svcCtx.BizRedis.ExpireCtx(l.ctx, key, 86400)

	return &pb.SaveAnswerSnapshotRes{}, nil
}
