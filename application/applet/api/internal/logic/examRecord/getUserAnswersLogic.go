// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package examRecord

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/exam/rpc/exam"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetUserAnswersLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取答题明细
func NewGetUserAnswersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetUserAnswersLogic {
	return &GetUserAnswersLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetUserAnswersLogic) GetUserAnswers(req *types.GetUserAnswersReq) (resp *types.GetUserAnswersRes, err error) {
	rpcResp, err := l.svcCtx.ExamRPC.GetUserAnswers(l.ctx, &exam.GetUserAnswersReq{
		RecordId: req.RecordId,
	})
	if err != nil {
		l.Errorf("GetUserAnswers error: %v", err)
		return nil, err
	}

	list := make([]types.UserAnswerItem, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		list = append(list, types.UserAnswerItem{
			QuestionId: item.QuestionId,
			UserAnswer: item.UserAnswer,
			IsCorrect:  int(item.IsCorrect),
			Score:      item.Score,
			AiStatus:   int(item.AiStatus),
			AiComment:  item.AiComment,
		})
	}

	return &types.GetUserAnswersRes{
		List: list,
	}, nil
}
