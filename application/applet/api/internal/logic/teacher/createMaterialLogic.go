// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMaterialLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 上传课件(绑定资源)
func NewCreateMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMaterialLogic {
	return &CreateMaterialLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CreateMaterialLogic) CreateMaterial(req *types.CreateMaterialReq) (resp *types.CreateMaterialRes, err error) {
	rpcResp, err := l.svcCtx.CourseRPC.CreateMaterial(l.ctx, &course.CreateMaterialReq{
		ChapterId: req.ChapterId,
		Title:     req.Title,
		Type:      int32(req.Type),
		Url:       req.Url,
		FileHash:  req.FileHash,
		FileSize:  req.FileSize,
		Sort:      int32(req.Sort),
	})
	if err != nil {
		l.Errorf("创建课件失败: %v", err)
		return nil, err
	}

	return &types.CreateMaterialRes{
		Id: rpcResp.Id,
	}, nil
}
