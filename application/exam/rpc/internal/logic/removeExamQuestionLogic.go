package logic

import (
	"context"
	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveExamQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRemoveExamQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveExamQuestionLogic {
	return &RemoveExamQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 移除试卷题目
func (l *RemoveExamQuestionLogic) RemoveExamQuestion(in *pb.RemoveExamQuestionReq) (*pb.RemoveExamQuestionRes, error) {
	eq, err := l.svcCtx.ExamQuestionModel.FindOneByExamIdQuestionId(l.ctx, uint64(in.ExamId), uint64(in.QuestionId))
	if err != nil {
		l.Errorf("ExamQuestionModel.FindOneByExamIdQuestionId error: %v", err)
		return nil, code.QuestionNotFound // 或者定义个更准确的比如 QuestionNotInExam
	}

	err = l.svcCtx.ExamQuestionModel.Delete(l.ctx, eq.Id)
	if err != nil {
		l.Errorf("ExamQuestionModel.Delete error: %v", err)
		return nil, xcode.ServerErr
	}

	// 试卷题目变更，清除 ExamDetail 缓存
	InvalidateExamDetailCache(l.ctx, l.svcCtx, in.ExamId)

	return &pb.RemoveExamQuestionRes{}, nil
}
