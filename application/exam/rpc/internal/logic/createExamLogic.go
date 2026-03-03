package logic

import (
	"context"
	"database/sql"
	"strconv"

	"teaching-backend/application/exam/rpc/model"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/util"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExamLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExamLogic {
	return &CreateExamLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// ========== 试卷管理 ==========
func (l *CreateExamLogic) CreateExam(in *pb.CreateExamReq) (*pb.CreateExamRes, error) {
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

	exam := &model.Exam{
		CourseId:   uint64(in.CourseId),
		Title:      in.Title,
		TotalScore: totalScore,
		PassScore:  passScore,
		Duration:   int64(in.Duration),
		StartTime:  startTime,
		EndTime:    endTime,
		ExamType:   int64(in.ExamType),
		RuleJson:   sql.NullString{String: in.RuleJson, Valid: in.RuleJson != ""},
		Status:     0, // 默认未开始
	}

	res, err := l.svcCtx.ExamModel.Insert(l.ctx, exam)
	if err != nil {
		l.Errorf("ExamModel.Insert error: %v", err)
		return nil, xcode.ServerErr
	}
	id, _ := res.LastInsertId()

	return &pb.CreateExamRes{
		Id: int64(id),
	}, nil
}
