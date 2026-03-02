package courselogic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/code"
	"teaching-backend/application/course/rpc/internal/model"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateMaterialLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateMaterialLogic {
	return &CreateMaterialLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 上传课件
func (l *CreateMaterialLogic) CreateMaterial(in *pb.CreateMaterialReq) (*pb.CreateMaterialRes, error) {
	if in.Title == "" {
		return nil, code.MaterialTitleEmpty
	}

	res, err := l.svcCtx.CourseMaterialModel.Insert(l.ctx, &model.CourseMaterial{
		ChapterId: uint64(in.ChapterId),
		Title:     in.Title,
		Type:      int64(in.Type),
		Url:       in.Url,
		FileHash:  in.FileHash,
		FileSize:  in.FileSize,
		AiStatus:  0, // 未处理
		Sort:      int64(in.Sort),
	})
	if err != nil {
		l.Errorf("插入课件记录失败: %v", err)
		return nil, xcode.ServerErr
	}

	id, err := res.LastInsertId()
	if err != nil {
		l.Errorf("获取课件 LastInsertId 失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.CreateMaterialRes{
		Id: id,
	}, nil
}
