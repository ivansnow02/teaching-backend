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

type SubmitExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 提交答卷
func NewSubmitExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SubmitExamLogic {
	return &SubmitExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SubmitExamLogic) SubmitExam(req *types.SubmitExamReq) (resp *types.SubmitExamRes, err error) {
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

	rpcResp, err := l.svcCtx.ExamRPC.SubmitExam(l.ctx, &exam.SubmitExamReq{
		RecordId: req.RecordId,
		ExamId:   req.ExamId,
		UserId:   userId,
		Answers:  answers,
	})
	if err != nil {
		l.Errorf("SubmitExam error: %v", err)
		return nil, err
	}

	return &types.SubmitExamRes{
		Success: rpcResp.Success,
	}, nil
}
