package RelationFollowInterface

import (
	"net/http"

	"SimpleTikTok/external_api/relationfollow/internal/logic/RelationFollowInterface"
	"SimpleTikTok/external_api/relationfollow/internal/svc"
	"SimpleTikTok/external_api/relationfollow/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func MessageActionHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.MessageActionHandlerRequest
		if err := httpx.Parse(r, &req); err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
			return
		}

		l := RelationFollowInterface.NewMessageActionLogic(r.Context(), svcCtx)
		resp, err := l.MessageAction(&req)
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
