// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"context"

	"teaching-backend/application/applet/api/internal/code"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
	"teaching-backend/application/course/rpc/course"
	"teaching-backend/application/user/rpc/user"
	"teaching-backend/pkg/encrypt"
	"teaching-backend/pkg/xcode"

	"github.com/zeromicro/go-zero/core/logx"
)

type CourseStudentsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

// 查看课程已选学生列表(教师)
func NewCourseStudentsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CourseStudentsLogic {
	return &CourseStudentsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CourseStudentsLogic) CourseStudents(req *types.CourseStudentsReq) (resp *types.CourseStudentsRes, err error) {

	if req.CourseId <= 0 {
		return nil, xcode.RequestErr
	}

	rpcResp, err := l.svcCtx.CourseRPC.GetCourseStudents(l.ctx, &course.GetCourseStudentsReq{
		CourseId: req.CourseId,
		Page:     int64(req.Page),
		Size:     int64(req.Size),
	})
	if err != nil {
		l.Errorf("获取课程已选学生失败: %v", err)
		return nil, code.GetStudentsFailed
	}

	if len(rpcResp.List) == 0 {
		return &types.CourseStudentsRes{List: []types.CourseStudentItem{}, Total: rpcResp.Total}, nil
	}

	list := make([]types.CourseStudentItem, 0, len(rpcResp.List))
	for _, item := range rpcResp.List {
		u, userErr := l.svcCtx.UserRPC.FindById(l.ctx, &user.FindByIdReq{
			UserId: item.UserId,
		})
		if userErr != nil {
			l.Errorf("查询学生信息失败 userId=%d: %v", item.UserId, userErr)
			return nil, code.GetStudentInfoFailed
		}

		email, decErr := encrypt.DecEmail(u.Email)
		if decErr != nil {
			l.Errorf("邮箱解密失败 userId=%d: %v", item.UserId, decErr)
			return nil, code.GetStudentInfoFailed
		}

		list = append(list, types.CourseStudentItem{
			UserId:     item.UserId,
			Nickname:   u.Nickname,
			Avatar:     u.Avatar,
			Email:      email,
			EnrollTime: item.EnrollTime,
		})
	}

	return &types.CourseStudentsRes{
		List:  list,
		Total: rpcResp.Total,
	}, nil
}
