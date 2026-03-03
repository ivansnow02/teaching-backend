package logic

import (
	"context"
	"fmt"

	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAnswersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetUserAnswersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAnswersLogic {
	return &GetUserAnswersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取学生答题明细
func (l *GetUserAnswersLogic) GetUserAnswers(in *pb.GetUserAnswersReq) (*pb.GetUserAnswersRes, error) {
	_, err := l.svcCtx.UserExamRecordModel.FindOne(l.ctx, uint64(in.RecordId))
	if err != nil {
		l.Errorf("UserExamRecordModel.FindOne error: %v", err)
		return nil, code.ExamRecordNotFound
	}

	answers, err := l.svcCtx.UserAnswerModel.FindListByRecordId(l.ctx, uint64(in.RecordId))
	if err != nil {
		l.Errorf("UserAnswerModel.FindListByRecordId error: %v", err)
		return nil, xcode.ServerErr
	}

	var resList []*pb.UserAnswerItem
	for _, a := range answers {
		resList = append(resList, &pb.UserAnswerItem{
			QuestionId: int64(a.QuestionId),
			UserAnswer: a.UserAnswer.String,
			IsCorrect:  int32(a.IsCorrect),
			Score:      fmt.Sprintf("%.1f", a.Score),
			AiStatus:   int32(a.AiStatus),
			AiComment:  a.AiComment.String,
		})
	}

	return &pb.GetUserAnswersRes{
		List: resList,
	}, nil
}
