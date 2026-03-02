// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package chapter

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/client/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChapterListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 章节列表
func NewChapterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterListLogic {
	return &ChapterListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ChapterListLogic) ChapterList(req *types.ChapterListReq) (resp *types.ChapterListRes, err error) {
	rpcResp, err := l.svcCtx.CourseRPC.ChapterList(l.ctx, &course.ChapterListReq{
		CourseId: req.CourseId,
	})
	if err != nil {
		l.Errorf("查询章节列表失败: %v", err)
		return nil, err
	}

	list := make([]types.ChapterItem, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		materials := make([]types.MaterialItem, 0, len(item.Materials))
		for _, m := range item.Materials {
			materials = append(materials, types.MaterialItem{
				Id:        m.Id,
				ChapterId: m.ChapterId,
				Title:     m.Title,
				Type:      int(m.Type),
				Url:       m.Url,
				FileHash:  m.FileHash,
				FileSize:  m.FileSize,
				AiStatus:  int(m.AiStatus),
				Sort:      int(m.Sort),
			})
		}
		list = append(list, types.ChapterItem{
			Id:        item.Id,
			CourseId:  item.CourseId,
			Title:     item.Title,
			Sort:      int(item.Sort),
			Materials: materials,
		})
	}

	return &types.ChapterListRes{
		List: list,
	}, nil
}
