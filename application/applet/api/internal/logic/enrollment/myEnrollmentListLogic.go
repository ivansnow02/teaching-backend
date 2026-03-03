// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package enrollment

import (
	"context"
	"encoding/json"

	"teaching-backend/application/applet/api/internal/code"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type MyEnrollmentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 我的选课列表
func NewMyEnrollmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MyEnrollmentListLogic {
	return &MyEnrollmentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MyEnrollmentListLogic) MyEnrollmentList(req *types.MyEnrollmentListReq) (resp *types.MyEnrollmentListRes, err error) {
	uid, ok := l.ctx.Value("userId").(json.Number)
	if !ok {
		return nil, xcode.AccessDenied
	}
	userId, err := uid.Int64()
	if err != nil || userId <= 0 {
		return nil, xcode.AccessDenied
	}

	enrollResp, err := l.svcCtx.CourseRPC.GetEnrollmentList(l.ctx, &course.GetEnrollmentListReq{
		UserId: userId,
		Page:   int64(req.Page),
		Size:   int64(req.Size),
	})
	if err != nil {
		l.Errorf("获取选课列表失败: %v", err)
		return nil, code.GetEnrollmentFailed
	}

	if len(enrollResp.List) == 0 {
		return &types.MyEnrollmentListRes{List: []types.EnrollmentCourseItem{}, Total: enrollResp.Total}, nil
	}

	// 批量查课程详情（RPC 层已有缓存）
	list := make([]types.EnrollmentCourseItem, 0, len(enrollResp.List))
	for _, item := range enrollResp.List {
		detail, detailErr := l.svcCtx.CourseRPC.CourseDetail(l.ctx, &course.CourseDetailReq{
			Id: item.CourseId,
		})
		if detailErr != nil {
			l.Errorf("查询课程详情失败 courseId=%d: %v", item.CourseId, detailErr)
			return nil, code.GetEnrollmentFailed
		}
		c := detail.Course
		list = append(list, types.EnrollmentCourseItem{
			Id:          c.Id,
			Title:       c.Title,
			Cover:       c.Cover,
			Description: c.Description,
			TeacherId:   c.TeacherId,
			Status:      int(c.Status),
			EnrollTime:  item.EnrollTime,
		})
	}

	return &types.MyEnrollmentListRes{
		List:  list,
		Total: enrollResp.Total,
	}, nil
}
