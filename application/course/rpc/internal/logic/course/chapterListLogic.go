package courselogic

import (
	"context"

	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type ChapterListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewChapterListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ChapterListLogic {
	return &ChapterListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 章节列表
func (l *ChapterListLogic) ChapterList(in *pb.ChapterListReq) (*pb.ChapterListRes, error) {
	chapters, err := l.svcCtx.CourseChapterModel.FindListByCourseId(l.ctx, uint64(in.CourseId))
	if err != nil {
		l.Errorf("查询课程章节列表失败: %v", err)
		return nil, xcode.ServerErr
	}

	var list []*pb.ChapterItem
	for _, c := range chapters {
		materials, err := l.svcCtx.CourseMaterialModel.FindListByChapterId(l.ctx, c.Id)
		if err != nil {
			l.Errorf("查询章节课件列表失败 (ChapterId: %d): %v", c.Id, err)
			return nil, xcode.ServerErr
		}

		var materialItems []*pb.MaterialItem
		for _, m := range materials {
			materialItems = append(materialItems, &pb.MaterialItem{
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

		list = append(list, &pb.ChapterItem{
			Id:        int64(c.Id),
			CourseId:  int64(c.CourseId),
			Title:     c.Title,
			Sort:      int32(c.Sort),
			Materials: materialItems,
		})
	}

	return &pb.ChapterListRes{
		List: list,
	}, nil
}
