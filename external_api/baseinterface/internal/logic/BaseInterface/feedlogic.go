package BaseInterface

import (
	"context"
	"time"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/internal_proto/microservices/miniomanage/types/miniomanageserver"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"

	// tools "SimpleTikTok/tools/token"

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
	// 暂时不知道token什么作用
	_, userId, err := tools.CheckToke(req.Token)
	// if err != nil || !ok {   //表示Token已失效，或者当前APP没有登录账号

	// }

	// var feedVideLists []mysqlconnect.VideoInfo
	feedVideLists, err := l.svcCtx.MySQLManageRpc.GetFeedVideoList(l.ctx, &mysqlmanageserver.GetFeedVideoListRequest{UserId: int64(userId)})
	// feedVideLists, err = mysqlconnect.GetFeedVideoList()
	if err != nil {
		logx.Errorf("[pkg]logic [func]Feed [msg]gorm GetFeedVideoList [err]%v", err)
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "获取视频信息失败",
			VideoList:  []types.VideoTest{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}
	if feedVideLists == nil {
		logx.Infof("[pkg]logic [func]Feed [msg]feedVideLists is nil", err)
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "此用户没有视频信息",
			VideoList:  []types.VideoTest{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}

	// var respFeedVideoList = make([]types.VideoTest, len(feedVideLists.VideoInfo))
	var respFeedVideoList = make([]types.VideoTest, 0)
	for _, val := range feedVideLists.VideoInfo {
		if val.PlayUrl == "" {
			continue
		}
		tmpAuthor, _ := getUserInfo(int(val.AuthorId))
		if tmpAuthor.Name == "" {
			continue
		}

		// 通过rpc调用获取
		realPlayUrlresp, _ := l.svcCtx.MinioManageRpc.GetPlayUrl(l.ctx, &miniomanageserver.GetPlayUrlRequest{
			PlayUrl: val.PlayUrl,
		})
		realPlayUrl := realPlayUrlresp.ResPlayUrl
		realCoverUrlresp, _ := l.svcCtx.MinioManageRpc.GetPlayUrl(l.ctx, &miniomanageserver.GetPlayUrlRequest{
			PlayUrl: val.CoverUrl,
		})
		realCoverUrl := realCoverUrlresp.ResPlayUrl
		// realPlayUrl, _ := minioconnect.GetPlayUrl(val.PlayUrl)
		// realCoverUrl, _ := minioconnect.GetPlayUrl(val.CoverUrl)

		respFeedVideoList = append(respFeedVideoList, types.VideoTest{
			Id:       val.VideoId,
			Author:   tmpAuthor,
			PlayUrl:  realPlayUrl,
			CoverUrl: realCoverUrl,
			// PlayUrl:       "http://175.178.93.55:9001/test-minio/vidoeFile/94e010c6-4c6e-4c2c-8d0b-b9afea9760d6-video_test12.mp4",
			//	CoverUrl:      "http://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
			FavoriteCount: val.FavoriteCount,
			CommentCount:  val.CommentCount,
			IsFavotite:    val.IsFavotite,
		})
	}

	return &types.FeedHandlerResponse{
		StatusCode: 0, // 0是成功
		StatusMsg:  "feed video success",
		VideoList:  respFeedVideoList,
		NextTime:   time.Now().Unix(), // 暂时返回当前时间
	}, nil
}

// func (l *FeedLogic) getUserInfo(userID int) (types.User, error) {
// 	resp, err := l.svcCtx.MySQLManageRpc.GetFeedUserInfo(l.ctx, &mysqlmanageserver.GetFeedUserInfoRequest{
// 		UserID: int64(userID),
// 	})
// 	// feedUserInfo, err := mysqlconnect.GetFeedUserInfo(userID)
// 	if err != nil {
// 		logx.Errorf("[pkg]logic [func]Feed [msg]gorm GetFeedUserInfo [err]%v", err)
// 		return types.User{}, err
// 	}
// 	if resp.FeedUserInfo == nil {
// 		return types.User{}, err
// 	}

// 	var respFeedUserInfo types.User
// 	respFeedUserInfo.UserId = resp.FeedUserInfo.UserID
// 	respFeedUserInfo.Name = resp.FeedUserInfo.UserNickName
// 	respFeedUserInfo.FollowCount = resp.FeedUserInfo.FollowCount
// 	respFeedUserInfo.FollowerCount = resp.FeedUserInfo.FollowerCount
// 	respFeedUserInfo.IsFollow = resp.FeedUserInfo.IsFollow
// 	return respFeedUserInfo, nil
// }

func getUserInfo(userID int) (types.User, error) {
	feedUserInfo, err := mysqlconnect.GetFeedUserInfo(userID)
	if err != nil {
		logx.Errorf("[pkg]logic [func]Feed [msg]gorm GetFeedUserInfo [err]%v", err)
		return types.User{}, err
	}
	if feedUserInfo.UserNickName == "" {
		logx.Infof("[pkg]logic [func]Feed [msg]feedUserInfo.Name is nuil ")
		return types.User{}, nil
	}
	var respFeedUserInfo types.User
	respFeedUserInfo.UserId = feedUserInfo.UserID
	respFeedUserInfo.Name = feedUserInfo.UserNickName
	respFeedUserInfo.FollowCount = feedUserInfo.FollowCount
	respFeedUserInfo.FollowerCount = feedUserInfo.FollowerCount
	respFeedUserInfo.IsFollow = feedUserInfo.IsFollow
	return respFeedUserInfo, nil
}
