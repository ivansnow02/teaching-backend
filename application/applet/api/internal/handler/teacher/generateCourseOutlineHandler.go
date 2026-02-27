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

// AI 生成课件大纲 - 异步
func GenerateCourseOutlineHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.GenerateCourseOutlineReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := teacher.NewGenerateCourseOutlineLogic(r.Context(), svcCtx)
		resp, err := l.GenerateCourseOutline(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
