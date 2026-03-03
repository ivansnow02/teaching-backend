// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package exam

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"

	"github.com/zeromicro/go-zero/core/logx"
)

type ExamDetailLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 试卷详情
func NewExamDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamDetailLogic {
	return &ExamDetailLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ExamDetailLogic) ExamDetail(req *types.ExamDetailReq) (resp *types.ExamDetailRes, err error) {
	rpcResp, err := l.svcCtx.ExamRPC.ExamDetail(l.ctx, &exam.ExamDetailReq{
		Id: req.Id,
	})
	if err != nil {
		l.Errorf("ExamDetail error: %v", err)
		return nil, err
	}

	questions := make([]types.ExamQuestionItem, 0, len(rpcResp.Questions))
	for _, q := range rpcResp.Questions {
		questions = append(questions, types.ExamQuestionItem{
			QuestionId: q.QuestionId,
			Score:      q.Score,
			Sort:       int(q.Sort),
			Question: types.QuestionItem{
				Id:              q.Question.Id,
				CourseId:        q.Question.CourseId,
				Type:            int(q.Question.Type),
				Content:         q.Question.Content,
				Answer:          q.Question.Answer,
				Analysis:        q.Question.Analysis,
				KnowledgePoints: q.Question.KnowledgePoints,
				Score:           q.Question.Score,
				Difficulty:      int(q.Question.Difficulty),
				CreateTime:      q.Question.CreateTime,
			},
		})
	}

	return &types.ExamDetailRes{
		Exam: types.ExamItem{
			Id:         rpcResp.Exam.Id,
			CourseId:   rpcResp.Exam.CourseId,
			Title:      rpcResp.Exam.Title,
			TotalScore: rpcResp.Exam.TotalScore,
			PassScore:  rpcResp.Exam.PassScore,
			Duration:   int(rpcResp.Exam.Duration),
			StartTime:  rpcResp.Exam.StartTime,
			EndTime:    rpcResp.Exam.EndTime,
			Status:     int(rpcResp.Exam.Status),
			ExamType:   int(rpcResp.Exam.ExamType),
			CreateTime: rpcResp.Exam.CreateTime,
		},
		Questions: questions,
	}, nil
}
