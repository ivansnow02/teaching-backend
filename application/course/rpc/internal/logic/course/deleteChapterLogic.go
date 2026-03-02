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
	_, err := l.svcCtx.CourseChapterModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.ChapterNotFound
		}
		l.Errorf("查询待删除章节失败: %v", err)
		return nil, xcode.ServerErr
	}

	err = l.svcCtx.CourseChapterModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("执行删除章节失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.DeleteChapterRes{}, nil
}
