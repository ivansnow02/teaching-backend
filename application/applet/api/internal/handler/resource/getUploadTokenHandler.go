// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package resource

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"teaching-backend/application/applet/api/internal/logic/resource"
	"teaching-backend/application/applet/api/internal/svc"
)

// 获取 OSS/MinIO 直传签名
func GetUploadTokenHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := resource.NewGetUploadTokenLogic(r.Context(), svcCtx)
		resp, err := l.GetUploadToken()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
