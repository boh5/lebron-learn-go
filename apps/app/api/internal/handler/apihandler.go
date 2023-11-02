package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"lebron-learn-go/apps/app/api/internal/logic"
	"lebron-learn-go/apps/app/api/internal/svc"
	"lebron-learn-go/apps/app/api/internal/types"
)

func ApiHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.Request
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewApiLogic(r.Context(), svcCtx)
		resp, err := l.Api(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
