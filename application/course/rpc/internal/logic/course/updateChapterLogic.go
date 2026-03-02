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

type UpdateChapterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateChapterLogic {
	return &UpdateChapterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新章节
func (l *UpdateChapterLogic) UpdateChapter(in *pb.UpdateChapterReq) (*pb.UpdateChapterRes, error) {
	chapter, err := l.svcCtx.CourseChapterModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.ChapterNotFound
		}
		l.Errorf("查询章节详情失败: %v", err)
		return nil, xcode.ServerErr
	}

	chapter.Title = in.Title
	chapter.Sort = int64(in.Sort)

	err = l.svcCtx.CourseChapterModel.Update(l.ctx, chapter)
	if err != nil {
		l.Errorf("更新章节记录失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.UpdateChapterRes{}, nil
}
