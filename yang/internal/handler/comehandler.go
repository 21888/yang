package handler

import (
	"net/http"

	"github.com/zeromicro/go-zero/rest/httpx"
	"service/yang/internal/logic"
	"service/yang/internal/svc"
	"service/yang/internal/types"
)

func ComeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.ComeReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.Error(w, err)
			return
		}

		l := logic.NewComeLogic(r.Context(), svcCtx)
		resp, err := l.Come(&req)
		if err != nil {
			httpx.Error(w, err)
		} else {
			httpx.OkJson(w, resp)
		}
	}
}
