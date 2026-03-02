package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type MaterialListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMaterialListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MaterialListLogic {
	return &MaterialListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 课件列表
func (l *MaterialListLogic) MaterialList(in *pb.MaterialListReq) (*pb.MaterialListRes, error) {
	materials, err := l.svcCtx.CourseMaterialModel.FindListByChapterId(l.ctx, uint64(in.ChapterId))
	if err != nil {
		l.Errorf("查询章节课件列表失败: %v", err)
		return nil, xcode.ServerErr
	}

	var list []*pb.MaterialItem
	for _, m := range materials {
		list = append(list, &pb.MaterialItem{
			Id:        int64(m.Id),
			ChapterId: int64(m.ChapterId),
			Title:     m.Title,
			Type:      int32(m.Type),
			Url:       m.Url,
			FileHash:  m.FileHash,
			FileSize:  m.FileSize,
			AiStatus:  int32(m.AiStatus),
			Sort:      int32(m.Sort),
		})
	}

	return &pb.MaterialListRes{
		List: list,
	}, nil
}
