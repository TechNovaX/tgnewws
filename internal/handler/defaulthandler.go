package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"tgnewws/internal/logic"
	"tgnewws/internal/svc"
	"tgnewws/internal/types"
)

func defaultHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.DefaultReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewDefaultLogic(r.Context(), svcCtx)
		resp, err := l.Default(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
