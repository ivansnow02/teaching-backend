// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package material

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaterialListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 课件列表
func NewMaterialListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialListLogic {
	return &MaterialListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MaterialListLogic) MaterialList(req *types.MaterialListReq) (resp *types.MaterialListRes, err error) {
	rpcResp, err := l.svcCtx.CourseRPC.MaterialList(l.ctx, &course.MaterialListReq{
		ChapterId: req.ChapterId,
	})
	if err != nil {
		l.Errorf("查询课件列表失败: %v", err)
		return nil, err
	}

	list := make([]types.MaterialItem, 0, len(rpcResp.List))
	for _, m := range rpcResp.List {
		list = append(list, types.MaterialItem{
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

	return &types.MaterialListRes{
		List: list,
	}, nil
}
