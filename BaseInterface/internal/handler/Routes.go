// Code generated by goctl. DO NOT EDIT.
package handler

import (
	"net/http"

	"SimpleTikTok/BaseInterface/internal/svc"

	"github.com/zeromicro/go-zero/rest"
)

func RegisterHandlers(server *rest.Server, serverCtx *svc.ServiceContext) {
	server.AddRoutes(
		[]rest.Route{
			{
				Method:  http.MethodGet,
				Path:    "/douyin/feed",
				Handler: FeedHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/register",
				Handler: UserRegisterHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/user/login",
				Handler: UserloginHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/user",
				Handler: UserHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/publish/action",
				Handler: PublishActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/publish/list",
				Handler: PublishListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/commment/action",
				Handler: CommmentActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/commment/list",
				Handler: CommmentListHandler(serverCtx),
			},
			{
				Method:  http.MethodPost,
				Path:    "/douyin/relation/action",
				Handler: RelationActionHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follow/list",
				Handler: RelationFollowListHandler(serverCtx),
			},
			{
				Method:  http.MethodGet,
				Path:    "/douyin/relation/follower/list",
				Handler: RelationFollowerListHandler(serverCtx),
			},
		},
	)
}
