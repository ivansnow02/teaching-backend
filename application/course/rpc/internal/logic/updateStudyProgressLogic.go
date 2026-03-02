package logic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/model"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateStudyProgressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateStudyProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateStudyProgressLogic {
	return &UpdateStudyProgressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新学习进度
func (l *UpdateStudyProgressLogic) UpdateStudyProgress(in *pb.UpdateStudyProgressReq) (*pb.UpdateStudyProgressRes, error) {
	// 查找是否已存在记录
	progress, err := l.svcCtx.StudyProgressModel.FindOneByUserIdMaterialId(l.ctx, uint64(in.UserId), uint64(in.MaterialId))
	if err != nil && err != model.ErrNotFound {
		l.Errorf("查询原学习进度失败: %v", err)
		return nil, xcode.ServerErr
	}

	if err == model.ErrNotFound {
		// 插入新记录
		_, err = l.svcCtx.StudyProgressModel.Insert(l.ctx, &model.StudyProgress{
			UserId:     uint64(in.UserId),
			CourseId:   uint64(in.CourseId),
			ChapterId:  uint64(in.ChapterId),
			MaterialId: uint64(in.MaterialId),
			Progress:   int64(in.Progress),
		})
	} else {
		// 更新现有记录进度
		if int64(in.Progress) > progress.Progress {
			progress.Progress = int64(in.Progress)
			err = l.svcCtx.StudyProgressModel.Update(l.ctx, progress)
		}
	}

	if err != nil {
		l.Errorf("写入学习进度失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.UpdateStudyProgressRes{}, nil
}
