// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateExamLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 创建试卷(教师)
func NewCreateExamLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateExamLogic {
	return &CreateExamLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateExamLogic) CreateExam(req *types.CreateExamReq) (resp *types.CreateExamRes, err error) {
	rpcResp, err := l.svcCtx.ExamRPC.CreateExam(l.ctx, &exam.CreateExamReq{
		CourseId:   req.CourseId,
		Title:      req.Title,
		TotalScore: req.TotalScore,
		PassScore:  req.PassScore,
		Duration:   int32(req.Duration),
		StartTime:  req.StartTime,
		EndTime:    req.EndTime,
		ExamType:   int32(req.ExamType),
		RuleJson:   req.RuleJson,
	})
	if err != nil {
		l.Errorf("CreateExam error: %v", err)
		return nil, err
	}

	return &types.CreateExamRes{
		Id: rpcResp.Id,
	}, nil
}
