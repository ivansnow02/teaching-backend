package logic

import (
	"context"
	"strconv"
	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/model"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddExamQuestionLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddExamQuestionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddExamQuestionLogic {
	return &AddExamQuestionLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 给试卷添加题目
func (l *AddExamQuestionLogic) AddExamQuestion(in *pb.AddExamQuestionReq) (*pb.AddExamQuestionRes, error) {
	_, err := l.svcCtx.ExamModel.FindOne(l.ctx, uint64(in.ExamId))
	if err != nil {
		l.Errorf("ExamModel.FindOne error: %v", err)
		return nil, code.ExamNotFound
	}

	_, err = l.svcCtx.QuestionModel.FindOne(l.ctx, uint64(in.QuestionId))
	if err != nil {
		l.Errorf("QuestionModel.FindOne error: %v", err)
		return nil, code.QuestionNotFound
	}

	score, _ := strconv.ParseFloat(in.Score, 64)

	eq := &model.ExamQuestion{
		ExamId:     uint64(in.ExamId),
		QuestionId: uint64(in.QuestionId),
		Score:      score,
		Sort:       int64(in.Sort),
	}

	_, err = l.svcCtx.ExamQuestionModel.Insert(l.ctx, eq)
	if err != nil {
		l.Errorf("ExamQuestionModel.Insert error: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.AddExamQuestionRes{}, nil
}
