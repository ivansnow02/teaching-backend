package logic

import (
	"context"

	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/model"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type StartExamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewStartExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *StartExamLogic {
	return &StartExamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ========== 考试作答 ==========
func (l *StartExamLogic) StartExam(in *pb.StartExamReq) (*pb.StartExamRes, error) {
	// 检查试卷是否存在
	_, err := l.svcCtx.ExamModel.FindOne(l.ctx, uint64(in.ExamId))
	if err != nil {
		l.Errorf("ExamModel.FindOne error: %v", err)
		return nil, code.ExamNotFound
	}

	// 检查是否已有记录
	record, err := l.svcCtx.UserExamRecordModel.FindOneByExamIdUserId(l.ctx, uint64(in.ExamId), uint64(in.UserId))
	if err == nil {
		// 如果已有且是进行中，直接返回
		if record.Status == 0 {
			return &pb.StartExamRes{RecordId: int64(record.Id)}, nil
		}
		// 如果已提交，目前逻辑返回错误（不可重考）
		return nil, code.ExamAlreadySubmitted
	}

	// 创建新记录
	newRecord := &model.UserExamRecord{
		ExamId: uint64(in.ExamId),
		UserId: uint64(in.UserId),
		Status: 0, // 答题中
	}

	res, err := l.svcCtx.UserExamRecordModel.Insert(l.ctx, newRecord)
	if err != nil {
		l.Errorf("UserExamRecordModel.Insert error: %v", err)
		return nil, xcode.ServerErr
	}
	id, _ := res.LastInsertId()

	return &pb.StartExamRes{
		RecordId: id,
	}, nil
}
