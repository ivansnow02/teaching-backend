package logic

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"teaching-backend/application/exam/rpc/internal/code"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	examDetailCacheKeyPrefix = "exam:detail:"
	examDetailCacheTTL       = 30 * time.Minute
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

// ExamDetail 试卷详情(含题目)，内置 Cache-Aside 缓存（BizRedis，TTL=30min）
func (l *ExamDetailLogic) ExamDetail(in *pb.ExamDetailReq) (*pb.ExamDetailRes, error) {
	cacheKey := fmt.Sprintf("%s%d", examDetailCacheKeyPrefix, in.Id)

	// 1. 先读缓存
	if cached, err := l.svcCtx.BizRedis.GetCtx(l.ctx, cacheKey); err == nil && cached != "" {
		var res pb.ExamDetailRes
		if err = json.Unmarshal([]byte(cached), &res); err == nil {
			l.Infof("ExamDetail: cache hit, exam_id=%d", in.Id)
			return &res, nil
		}
	}

	// 2. 缓存 miss，从数据库构建
	l.Infof("ExamDetail: cache miss, exam_id=%d, querying DB", in.Id)
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
			continue
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

	res := &pb.ExamDetailRes{
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
	}

	// 3. 写入缓存
	if data, err := json.Marshal(res); err == nil {
		if err = l.svcCtx.BizRedis.SetexCtx(l.ctx, cacheKey, string(data), int(examDetailCacheTTL.Seconds())); err != nil {
			l.Errorf("ExamDetail: cache set error: %v", err)
		}
	}

	return res, nil
}

// InvalidateExamDetailCache 用于在试卷变更时主动删除缓存
func InvalidateExamDetailCache(ctx context.Context, svcCtx *svc.ServiceContext, examId int64) {
	cacheKey := fmt.Sprintf("%s%d", examDetailCacheKeyPrefix, examId)
	if _, err := svcCtx.BizRedis.DelCtx(ctx, cacheKey); err != nil {
		logx.Errorf("InvalidateExamDetailCache(exam_id=%d) error: %v", examId, err)
	}
}
