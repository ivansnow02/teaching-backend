package logic

import (
	"context"
	"database/sql"

	"teaching-backend/application/course/rpc/internal/code"
	"teaching-backend/application/course/rpc/internal/svc"
	"teaching-backend/application/course/rpc/pb"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteChapterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteChapterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteChapterLogic {
	return &DeleteChapterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除章节
func (l *DeleteChapterLogic) DeleteChapter(in *pb.DeleteChapterReq) (*pb.DeleteChapterRes, error) {
	chapter, err := l.svcCtx.CourseChapterModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.ChapterNotFound
		}
		l.Errorf("查询待删除章节失败: %v", err)
		return nil, xcode.ServerErr
	}

	course, err := l.svcCtx.CourseModel.FindOne(l.ctx, chapter.CourseId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.CourseNotFound
		}
		l.Errorf("查询课程详情失败: %v", err)
		return nil, xcode.ServerErr
	}

	if course.TeacherId != uint64(in.OperatorId) {
		return nil, code.NoPermission
	}

	err = l.svcCtx.CourseChapterModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("执行删除章节失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.DeleteChapterRes{}, nil
}
