package logic

import (
	"context"
	"fmt"

	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetExamResultLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetExamResultLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetExamResultLogic {
	return &GetExamResultLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取考试结果
func (l *GetExamResultLogic) GetExamResult(in *pb.GetExamResultReq) (*pb.GetExamResultRes, error) {
	record, err := l.svcCtx.UserExamRecordModel.FindOne(l.ctx, uint64(in.RecordId))
	if err != nil {
		l.Errorf("UserExamRecordModel.FindOne error: %v", err)
		return nil, code.ExamRecordNotFound
	}

	return &pb.GetExamResultRes{
		ExamId:     int64(record.ExamId),
		UserId:     int64(record.UserId),
		Score:      fmt.Sprintf("%.1f", record.Score),
		Status:     int32(record.Status),
		SubmitTime: record.SubmitTime.Time.Unix(),
	}, nil
}
