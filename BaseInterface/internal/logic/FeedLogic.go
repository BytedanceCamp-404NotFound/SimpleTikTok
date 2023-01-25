package logic

import (
	"context"
	"time"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedHandlerRequest) (resp *types.FeedHandlerResponse, err error) {
	ok, userId, err := tools.CheckToke(req.Token)
	if !ok {
		return &types.FeedHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	db, err := mysqlconnect.SqlConnect() //连接数据库
	if err != nil {
		logx.Errorf("FeedLogic select feedUserInfo error:%v", err)
		return nil, err
	}

	var feedUserInfo mysqlconnect.UserInfo
	err = db.Model(&mysqlconnect.UserInfo{}).Where("user_id = ?", userId).First(&feedUserInfo).Error
	if err != nil {
		logx.Errorf("UserID: %v FeedLogic select feedUserInfo error:%v", userId, err)
		return nil, err
	}
	var respFeedUserInfo types.User
	respFeedUserInfo.UserId = feedUserInfo.UserID
	respFeedUserInfo.Name = feedUserInfo.Name
	respFeedUserInfo.FollowCount = feedUserInfo.FollowCount
	respFeedUserInfo.FollowerCount = feedUserInfo.FollowerCount
	respFeedUserInfo.IsFollow = feedUserInfo.IsFollow

	var feedVideLists []mysqlconnect.VideoInfo
	err = db.Model(&mysqlconnect.VideoInfo{}).Where("user_id = ?", userId).Scan(&feedVideLists).Limit(10).Error
	if err != nil {
		logx.Errorf("FeedLogic select VideoInfo error:%v", err)
		return nil, err
	}

	var respFeedVideoList = make([]types.Video, len(feedVideLists))
	for index, val := range feedVideLists {
		respFeedVideoList[index].Id = val.VideID
		respFeedVideoList[index].Author = respFeedUserInfo
		respFeedVideoList[index].PlayUrl = val.PlayUrl
		respFeedVideoList[index].CoverUrl = val.CoverUrl
		respFeedVideoList[index].FavoriteCount = val.FavoriteCount
		respFeedVideoList[index].CommentCount = val.CommentCount
		respFeedVideoList[index].IsFavotite = val.IsFavotite
		respFeedVideoList[index].VideoTitle = val.VideoTitle
	}

	return &types.FeedHandlerResponse{
		StatusCode: 200,
		StatusMsg:  "feed video success",
		VideoList:  respFeedVideoList,
		NextTime:   time.Now().Unix(), // 暂时返回当前时间
	}, nil
}

func tokenIsNull() error {
	return nil
}
