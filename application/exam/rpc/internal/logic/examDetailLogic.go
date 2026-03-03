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

type ExamDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewExamDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ExamDetailLogic {
	return &ExamDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 试卷详情(含题目)
func (l *ExamDetailLogic) ExamDetail(in *pb.ExamDetailReq) (*pb.ExamDetailRes, error) {
	exam, err := l.svcCtx.ExamModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("ExamModel.FindOne error: %v", err)
		return nil, code.ExamNotFound
	}

	questions, err := l.svcCtx.ExamQuestionModel.FindListByExamId(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("ExamQuestionModel.FindListByExamId error: %v", err)
		return nil, xcode.ServerErr
	}

	var resQuestions []*pb.ExamQuestionItem
	for _, eq := range questions {
		q, err := l.svcCtx.QuestionModel.FindOne(l.ctx, eq.QuestionId)
		if err != nil {
			continue // 稳健起见，跳过不存在的题目
		}
		resQuestions = append(resQuestions, &pb.ExamQuestionItem{
			Id:         int64(eq.Id),
			ExamId:     int64(eq.ExamId),
			QuestionId: int64(eq.QuestionId),
			Score:      fmt.Sprintf("%.1f", eq.Score),
			Sort:       int32(eq.Sort),
			Question: &pb.QuestionItem{
				Id:              int64(q.Id),
				CourseId:        int64(q.CourseId),
				TeacherId:       int64(q.TeacherId),
				Type:            int32(q.Type),
				Content:         q.Content,
				Answer:          q.Answer,
				Analysis:        q.Analysis.String,
				KnowledgePoints: q.KnowledgePoints,
				Score:           fmt.Sprintf("%.1f", q.Score),
				Difficulty:      int32(q.Difficulty),
				CreateTime:      q.CreateTime.Unix(),
			},
		})
	}

	return &pb.ExamDetailRes{
		Exam: &pb.ExamItem{
			Id:         int64(exam.Id),
			CourseId:   int64(exam.CourseId),
			Title:      exam.Title,
			TotalScore: fmt.Sprintf("%.1f", exam.TotalScore),
			PassScore:  fmt.Sprintf("%.1f", exam.PassScore),
			Duration:   int32(exam.Duration),
			StartTime:  exam.StartTime.Time.Unix(),
			EndTime:    exam.EndTime.Time.Unix(),
			Status:     int32(exam.Status),
			ExamType:   int32(exam.ExamType),
			CreateTime: exam.CreateTime.Unix(),
		},
		Questions: resQuestions,
	}, nil
}
