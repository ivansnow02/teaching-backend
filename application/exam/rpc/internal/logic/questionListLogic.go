package logic

import (
	"context"
	"fmt"
	"teaching-backend/application/exam/rpc/internal/svc"
	"teaching-backend/application/exam/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type QuestionListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewQuestionListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *QuestionListLogic {
	return &QuestionListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 题目列表
func (l *QuestionListLogic) QuestionList(in *pb.QuestionListReq) (*pb.QuestionListRes, error) {
	list, total, err := l.svcCtx.QuestionModel.FindList(l.ctx, uint64(in.CourseId), 0, uint64(in.Type), in.Page, in.Size)
	if err != nil {
		l.Errorf("QuestionModel.FindList error: %v", err)
		return nil, xcode.ServerErr
	}

	var resList []*pb.QuestionItem
	for _, item := range list {
		resList = append(resList, &pb.QuestionItem{
			Id:              int64(item.Id),
			CourseId:        int64(item.CourseId),
			TeacherId:       int64(item.TeacherId),
			Type:            int32(item.Type),
			Content:         item.Content,
			Answer:          item.Answer,
			Analysis:        item.Analysis.String,
			KnowledgePoints: item.KnowledgePoints,
			Score:           fmt.Sprintf("%.1f", item.Score),
			Difficulty:      int32(item.Difficulty),
			CreateTime:      item.CreateTime.Unix(),
		})
	}

	return &pb.QuestionListRes{
		List:  resList,
		Total: total,
	}, nil
}
