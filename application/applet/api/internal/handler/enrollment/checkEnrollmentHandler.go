// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package enrollment

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"teaching-backend/application/applet/api/internal/logic/enrollment"
	"teaching-backend/application/applet/api/internal/svc"
	"teaching-backend/application/applet/api/internal/types"
)

// 检查是否已选课
func CheckEnrollmentHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CheckEnrollmentReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := enrollment.NewCheckEnrollmentLogic(r.Context(), svcCtx)
		resp, err := l.CheckEnrollment(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
