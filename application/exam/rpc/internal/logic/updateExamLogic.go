package logic

import (
	"context"
	"database/sql"
	"strconv"

	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/util"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateExamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateExamLogic {
	return &UpdateExamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新试卷
func (l *UpdateExamLogic) UpdateExam(in *pb.UpdateExamReq) (*pb.UpdateExamRes, error) {
	exam, err := l.svcCtx.ExamModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("ExamModel.FindOne error: %v", err)
		return nil, code.ExamNotFound
	}

	totalScore, _ := strconv.ParseFloat(in.TotalScore, 64)
	passScore, _ := strconv.ParseFloat(in.PassScore, 64)

	var startTime, endTime sql.NullTime
	if in.StartTime > 0 {
		st := util.AdaptiveTime(in.StartTime)
		startTime = sql.NullTime{Time: st, Valid: !st.IsZero()}
	}
	if in.EndTime > 0 {
		et := util.AdaptiveTime(in.EndTime)
		endTime = sql.NullTime{Time: et, Valid: !et.IsZero()}
	}

	exam.Title = in.Title
	exam.TotalScore = totalScore
	exam.PassScore = passScore
	exam.Duration = int64(in.Duration)
	exam.StartTime = startTime
	exam.EndTime = endTime
	exam.Status = int64(in.Status)
	exam.ExamType = int64(in.ExamType)

	err = l.svcCtx.ExamModel.Update(l.ctx, exam)
	if err != nil {
		l.Errorf("ExamModel.Update error: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.UpdateExamRes{}, nil
}
