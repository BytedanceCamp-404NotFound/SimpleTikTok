package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListHandlerRequest) (resp *types.PublishListHandlerResponse, err error) {
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		return &types.PublishListHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "请登录",
			VideoList:  []types.Video{},
		}, err
	}
	user, ok := mysqlconnect.CheckUserInf(int(req.UserID), id)
	if !ok {
		return &types.PublishListHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "查询的用户不存在！",
			VideoList:  []types.Video{},
		}, err
	}

	n := mysqlconnect.VideoNum(int(req.UserID))
	v := mysqlconnect.GetVideoList(int(req.UserID))
	videolist := make([]types.Video, n)
	for i := 0; i < int(n); i++ {
		videolist[i] = types.Video{
			Id: v[i].VideID,
			Author: types.User{
				UserId:        user.User.UserID,
				Name:          user.User.UserNickName,
				FollowCount:   user.User.FollowCount,
				FollowerCount: user.User.FollowerCount,
				IsFollow:      user.IsFollow,
			},
			PlayUrl:       v[i].PlayUrl,
			CoverUrl:      v[i].CoverUrl,
			FavoriteCount: v[i].FavoriteCount,
			CommentCount:  v[i].CommentCount,
			IsFavotite:    v[i].IsFavotite,
			VideoTitle:    v[i].VideoTitle,
		}
	}
	return &types.PublishListHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "查询发布列表成功",
		VideoList:  videolist,
	}, err
}
