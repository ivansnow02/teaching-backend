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

type DeleteMaterialLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteMaterialLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteMaterialLogic {
	return &DeleteMaterialLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除课件
func (l *DeleteMaterialLogic) DeleteMaterial(in *pb.DeleteMaterialReq) (*pb.DeleteMaterialRes, error) {
	material, err := l.svcCtx.CourseMaterialModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.MaterialNotFound
		}
		l.Errorf("查询待删除课件失败: %v", err)
		return nil, xcode.ServerErr
	}

	chapter, err := l.svcCtx.CourseChapterModel.FindOne(l.ctx, material.ChapterId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.ChapterNotFound
		}
		l.Errorf("查询关联章节失败: %v", err)
		return nil, xcode.ServerErr
	}

	course, err := l.svcCtx.CourseModel.FindOne(l.ctx, chapter.CourseId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.CourseNotFound
		}
		l.Errorf("查询关联课程失败: %v", err)
		return nil, xcode.ServerErr
	}

	if course.TeacherId != uint64(in.OperatorId) {
		return nil, code.NoPermission
	}

	err = l.svcCtx.CourseMaterialModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("执行删除课件失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.DeleteMaterialRes{}, nil
}
