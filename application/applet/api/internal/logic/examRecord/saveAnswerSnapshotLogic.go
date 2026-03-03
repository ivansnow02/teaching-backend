// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package examRecord

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type SaveAnswerSnapshotLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 考试过程心跳与答案快照同步
func NewSaveAnswerSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SaveAnswerSnapshotLogic {
	return &SaveAnswerSnapshotLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SaveAnswerSnapshotLogic) SaveAnswerSnapshot(req *types.SaveAnswerSnapshotReq) (resp *types.Empty, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	answers := make([]*pb.AnswerItem, 0, len(req.Answers))
	for _, a := range req.Answers {
		answers = append(answers, &pb.AnswerItem{
			QuestionId: a.QuestionId,
			Answer:     a.Answer,
		})
	}

	_, err = l.svcCtx.ExamRPC.SaveAnswerSnapshot(l.ctx, &exam.SaveAnswerSnapshotReq{
		RecordId: req.RecordId,
		UserId:   userId,
		Answers:  answers,
	})
	if err != nil {
		l.Errorf("SaveAnswerSnapshot error: %v", err)
		return nil, err
	}

	return &types.Empty{}, nil
}
