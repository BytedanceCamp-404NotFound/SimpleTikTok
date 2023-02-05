package BaseInterface

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/mysqlconnect"

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
	// ok, userId, err := tools.CheckToke(req.Token)
	// if err != nil {
	// 	return &types.FeedHandlerResponse{
	// 		StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
	// 		StatusMsg:  "Token校验出错",
	// 		VideoList:  []types.VideoTest{},
	// 		NextTime:   time.Now().Unix(), // 暂时返回当前时间
	// 	}, nil
	// }
	// if !ok {
	// 	logx.Infof("[pkg]logic [func]Feed [msg]feedUserInfo.Name is nuil ")
	// 	return &types.FeedHandlerResponse{
	// 		StatusCode: int32(commonerror.CommonErr_PARAMETER_FAILED),
	// 		StatusMsg:  "登录过期，请重新登陆",
	// 		VideoList:  []types.VideoTest{},
	// 		NextTime:   time.Now().Unix(), // 暂时返回当前时间
	// 	}, nil
	// }

	// var feedVideLists []mysqlconnect.VideoInfo
	feedVideLists, err := l.svcCtx.MySQLManageRpc.GetFeedVideoList(l.ctx, &mysqlmanageserver.GetFeedVideoListRequest{})
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
		if val.PlayUrl == ""{
			continue
		}
		tmpAuthor, _ := getUserInfo(int(val.AuthorId))
		if tmpAuthor.Name == ""{
			continue
		}
		// realPlayUrl, _ := getPlayUrl(val.PlayUrl)
		respFeedVideoList = append(respFeedVideoList, types.VideoTest{
			Id:            val.VideoId,
			Author:        tmpAuthor,
			PlayUrl:       "http://175.178.93.55:9001/test-minio/vidoeFile/94e010c6-4c6e-4c2c-8d0b-b9afea9760d6-video_test12.mp4",
			CoverUrl:      "http://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png",
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

const (
	sourcetype = "minio"
	separator  = "_"
)

type MinioKeyVal struct {
	SourceType string
	Bucket     string
	Key        string
}

func DecodeFileKey(key string) (*MinioKeyVal, error) {
	keyval := &MinioKeyVal{}
	if !strings.Contains(key, separator) {
		return nil, errors.New("invalid filekey fail")
	}
	keyparts := strings.Split(key, separator)
	if len(keyparts) != 2 {
		return nil, errors.New("cant Split")
	}
	keyval.SourceType = keyparts[0]
	keyData, err := base64.StdEncoding.DecodeString(keyparts[1])
	if err != nil {
		logx.Errorf("decode base64 error:", err.Error())
		return nil, err
	}

	decodeString := string(keyData)
	index := strings.Index(decodeString, separator)
	if index <= 0 {
		return nil, errors.New("cant find separator")
	}

	keyval.Bucket = decodeString[:index]
	keyval.Key = decodeString[index+len(separator):]
	return keyval, nil
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

func getPlayUrl(playUrl string) (string, error) {
	if playUrl == "" {
		logx.Infof("[pkg]BaseInterface [func]getPlayUrl [msg]playUrl is nil")
		return "", nil
	}
	decodeKey, err := DecodeFileKey(playUrl)
	if err != nil {
		logx.Errorf("decode base64 error:%v", err)
		return "", err
	}
	minioUrl := "http://175.178.93.55:9001"

	resPlayUrl := fmt.Sprintf("%s/%s/%s", minioUrl, decodeKey.Bucket, decodeKey.Key)
	return resPlayUrl, nil
}
