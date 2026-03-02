// Code scaffolded by goctl. Safe to edit.
// goctl 1.9.2

package middleware

import (
	"encoding/json"
	"net/http"

	"teaching-backend/application/applet/api/internal/code"

	"github.com/zeromicro/go-zero/rest/httpx"
)

type CheckTeacherRoleMiddleware struct {
}

func NewCheckTeacherRoleMiddleware() *CheckTeacherRoleMiddleware {
	return &CheckTeacherRoleMiddleware{}
}

func (m *CheckTeacherRoleMiddleware) Handle(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		roleValue := r.Context().Value("role")
		if roleValue == nil {
			httpx.Error(w, code.NoPermission)
			return
		}

		roleNumber, ok := roleValue.(json.Number)
		if !ok {
			// 兜底处理：有些情况下可能是 float64 或 int64
			httpx.Error(w, code.NoPermission)
			return
		}

		role, err := roleNumber.Int64()
		if err != nil || role == 1 { // 1:学生 2:教师 3:管理员
			httpx.Error(w, code.NoPermission)
			return
		}

		next(w, r)
	}
}
