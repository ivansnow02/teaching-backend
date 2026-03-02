// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package course

import (
	"context"

	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/client/course"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 课程列表
func NewCourseListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseListLogic {
	return &CourseListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CourseListLogic) CourseList(req *types.CourseListReq) (resp *types.CourseListRes, err error) {
	rpcResp, err := l.svcCtx.CourseRPC.CourseList(l.ctx, &course.CourseListReq{
		Page: int64(req.Page),
		Size: int64(req.Size),
	})
	if err != nil {
		l.Errorf("查询课程列表失败: %v", err)
		return nil, err
	}

	list := make([]types.CourseItem, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		list = append(list, types.CourseItem{
			Id:          item.Id,
			Title:       item.Title,
			Cover:       item.Cover,
			Description: item.Description,
			TeacherId:   item.TeacherId,
			Status:      int(item.Status),
			CreateTime:  item.CreateTime,
		})
	}

	return &types.CourseListRes{
		List:  list,
		Total: rpcResp.Total,
	}, nil
}
