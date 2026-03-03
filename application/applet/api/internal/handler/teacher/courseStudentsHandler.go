// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package teacher

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"teaching-backend/application/applet/api/internal/logic/teacher"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
)

// 查看课程已选学生列表(教师)
func CourseStudentsHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CourseStudentsReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := teacher.NewCourseStudentsLogic(r.Context(), svcCtx)
		resp, err := l.CourseStudents(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
