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

type DeleteCourseLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteCourseLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteCourseLogic {
	return &DeleteCourseLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除课程
func (l *DeleteCourseLogic) DeleteCourse(in *pb.DeleteCourseReq) (*pb.DeleteCourseRes, error) {
	_, err := l.svcCtx.CourseModel.FindOne(l.ctx, uint64(in.Id))
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, code.CourseNotFound
		}
		l.Errorf("查待删除课程失败: %v", err)
		return nil, xcode.ServerErr
	}

	err = l.svcCtx.CourseModel.Delete(l.ctx, uint64(in.Id))
	if err != nil {
		l.Errorf("执行删除课程失败: %v", err)
		return nil, xcode.ServerErr
	}

	return &pb.DeleteCourseRes{}, nil
}
