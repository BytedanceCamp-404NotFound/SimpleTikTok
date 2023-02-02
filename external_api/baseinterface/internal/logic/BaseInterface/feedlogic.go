package BaseInterface

import (
	"context"
	"encoding/base64"
	"errors"
	"strings"
	"time"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/oprations/commonerror"
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
	// userIdTTT := fmt.Sprintf("%v", l.ctx.Value("userIdentity"))
	// // 这里的key和生成jwt token时传入的key一致
	// userId, err := strconv.Atoi(userIdTTT)
	// logx.Infof("userId: %v",userId)

	ok, userId, err := tools.CheckToke(req.Token)
	if err != nil {
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
			VideoList:  []types.VideoTest{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}
	if !ok {
		logx.Infof("[pkg]logic [func]Feed [msg]feedUserInfo.Name is nuil ")
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARAMETER_FAILED),
			StatusMsg:  "登录过期，请重新登陆",
			VideoList:  []types.VideoTest{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}

	feedUserInfo, err := mysqlconnect.GetFeedUserInfo(userId)
	if err != nil {
		logx.Errorf("[pkg]logic [func]Feed [msg]gorm GetFeedUserInfo [err]%v", err)
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "获取用户信息失败",
			VideoList:  []types.VideoTest{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}
	if feedUserInfo.UserNickName == "" {
		logx.Infof("[pkg]logic [func]Feed [msg]feedUserInfo.Name is nuil ")
		return &types.FeedHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARAMETER_FAILED),
			StatusMsg:  "用户信息为空",
			VideoList:  []types.VideoTest{},
			NextTime:   time.Now().Unix(), // 暂时返回当前时间
		}, nil
	}
	var respFeedUserInfo types.User
	respFeedUserInfo.UserId = feedUserInfo.UserID
	respFeedUserInfo.Name = feedUserInfo.UserNickName
	respFeedUserInfo.FollowCount = feedUserInfo.FollowCount
	respFeedUserInfo.FollowerCount = feedUserInfo.FollowerCount
	respFeedUserInfo.IsFollow = feedUserInfo.IsFollow

	var feedVideLists []mysqlconnect.VideoInfo
	feedVideLists, err = mysqlconnect.GetFeedVideoList(userId)
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

	var respFeedVideoList = make([]types.VideoTest, len(feedVideLists))
	for index, val := range feedVideLists {
		// respFeedVideoList[index].Id = val.VideoID
		respFeedVideoList[index].Id = 1
		respFeedVideoList[index].Author = respFeedUserInfo
		// respFeedVideoList[index].PlayUrl = "175.178.93.55:9001/test-minio/vidoeFile/790ae4b3-cce7-43fe-bfd6-8ee82a23ca74-video_test6.mp4"
		// respFeedVideoList[index].PlayUrl = "http://175.178.93.55:9001/test-minio/vidoeFile/94e010c6-4c6e-4c2c-8d0b-b9afea9760d6-video_test12.mp4"
		keyval ,_ := DecodeFileKey(val.PlayUrl)
		respFeedVideoList[index].PlayUrl = keyval.Key
		// keyval, _ := DecodeFileKey(val.CoverUrl)
		// respFeedVideoList[index].CoverUrl = val.CoverUrl
		respFeedVideoList[index].CoverUrl = "http://www.baidu.com/img/PCtm_d9c8750bed0b3c7d089fa7d55720d6cf.png"
		respFeedVideoList[index].FavoriteCount = val.FavoriteCount
		respFeedVideoList[index].CommentCount = val.CommentCount
		respFeedVideoList[index].IsFavotite = val.IsFavotite
		// respFeedVideoList[index].VideoTitle = val.VideoTitle
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
