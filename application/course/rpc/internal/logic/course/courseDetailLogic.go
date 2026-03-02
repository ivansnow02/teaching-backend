package courselogic

import (
	"context"
	"database/sql"

	"teaching-backend/application/course/rpc/internal/code"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseDetailLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCourseDetailLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseDetailLogic {
	return &CourseDetailLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 课程详情
func (l *CourseDetailLogic) CourseDetail(in *pb.CourseDetailReq) (*pb.CourseDetailRes, error) {
	course, err := l.svcCtx.CourseModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.CourseNotFound
		}
		l.Errorf("查询课程详情失败: %v", err)
		return nil, xcode.ServerErr
	}

	// 查询章节
	chapters, err := l.svcCtx.CourseChapterModel.FindListByCourseId(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("查询课程章节列表失败: %v", err)
		return nil, xcode.ServerErr
	}

	var chapterItems []*pb.ChapterItem
	for _, c := range chapters {
		// 查询章节下的课件
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

		chapterItems = append(chapterItems, &pb.ChapterItem{
			Id:        int64(c.Id),
			CourseId:  int64(c.CourseId),
			Title:     c.Title,
			Sort:      int32(c.Sort),
			Materials: materialItems,
		})
	}

	return &pb.CourseDetailRes{
		Course: &pb.CourseItem{
			Id:          int64(course.Id),
			Title:       course.Title,
			Cover:       course.Cover,
			Description: course.Description.String,
			TeacherId:   int64(course.TeacherId),
			Status:      int32(course.Status),
			CreateTime:  course.CreateTime.Unix(),
		},
		Chapters: chapterItems,
	}, nil
}
