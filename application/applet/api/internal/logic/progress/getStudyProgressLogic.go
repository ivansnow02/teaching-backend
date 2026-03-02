// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package progress

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudyProgressLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 获取学习进度
func NewGetStudyProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudyProgressLogic {
	return &GetStudyProgressLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *GetStudyProgressLogic) GetStudyProgress(req *types.GetStudyProgressReq) (resp *types.GetStudyProgressRes, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	rpcResp, err := l.svcCtx.CourseRPC.GetStudyProgress(l.ctx, &course.GetStudyProgressReq{
		UserId:   userId,
		CourseId: req.CourseId,
	})
	if err != nil {
		l.Errorf("查询学习进度失败: %v", err)
		return nil, err
	}

	list := make([]types.StudyProgressItem, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		list = append(list, types.StudyProgressItem{
			MaterialId: item.MaterialId,
			ChapterId:  item.ChapterId,
			Progress:   int(item.Progress),
		})
	}

	return &types.GetStudyProgressRes{
		List: list,
	}, nil
}
