package handler

import (
	"net/http"

	"SimpleTikTok/BaseInterface/internal/logic"
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func FavoriteListRegisterHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.FavoriteListRegisterHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := logic.NewFavoriteListRegisterLogic(r.Context(), svcCtx)
		resp, err := l.FavoriteListRegister(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}