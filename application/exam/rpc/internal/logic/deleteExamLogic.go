package logic

import (
	"context"
	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteExamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteExamLogic {
	return &DeleteExamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除试卷
func (l *DeleteExamLogic) DeleteExam(in *pb.DeleteExamReq) (*pb.DeleteExamRes, error) {
	_, err := l.svcCtx.ExamModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("ExamModel.FindOne error: %v", err)
		return nil, code.ExamNotFound
	}

	err = l.svcCtx.ExamModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("ExamModel.Delete error: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.DeleteExamRes{}, nil
}
