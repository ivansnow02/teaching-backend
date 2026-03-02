package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/model"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateChapterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateChapterLogic {
	return &CreateChapterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建章节
func (l *CreateChapterLogic) CreateChapter(in *pb.CreateChapterReq) (*pb.CreateChapterRes, error) {
	res, err := l.svcCtx.CourseChapterModel.Insert(l.ctx, &model.CourseChapter{
		CourseId: uint64(in.CourseId),
		Title:    in.Title,
		Sort:     int64(in.Sort),
	})
	if err != nil {
		l.Errorf("插入章节记录失败: %v", err)
		return nil, xcode.ServerErr
	}

	id, err := res.LastInsertId()
	if err != nil {
		l.Errorf("获取章节 LastInsertId 失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.CreateChapterRes{
		Id: id,
	}, nil
}
