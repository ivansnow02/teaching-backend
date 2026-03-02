package courselogic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetStudyProgressLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetStudyProgressLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetStudyProgressLogic {
	return &GetStudyProgressLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取学习进度
func (l *GetStudyProgressLogic) GetStudyProgress(in *pb.GetStudyProgressReq) (*pb.GetStudyProgressRes, error) {
	list, err := l.svcCtx.StudyProgressModel.FindListByUserIdCourseId(l.ctx, uint64(in.UserId), uint64(in.CourseId))
	if err != nil {
		l.Errorf("查询学习进度列表失败: %v", err)
		return nil, xcode.ServerErr
	}

	var res []*pb.StudyProgressItem
	for _, item := range list {
		res = append(res, &pb.StudyProgressItem{
			MaterialId: int64(item.MaterialId),
			ChapterId:  int64(item.ChapterId),
			Progress:   int32(item.Progress),
		})
	}

	return &pb.GetStudyProgressRes{
		List: res,
	}, nil
}
