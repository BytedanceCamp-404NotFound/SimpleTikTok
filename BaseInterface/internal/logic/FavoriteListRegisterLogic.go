package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListRegisterLogic {
	return &FavoriteListRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type VideoInfo struct {
	VideoTitle    string `gorm:"column:video_title" form:"video_title" bson:"video_title"`
	AuthorId      int64  `gorm:"column:author_id" form:"author_id" bson:"author_id"`
	CoverUrl      string `gorm:"column:cover_url" form:"cover_url" bson:"cover_url"`
	PlayUrl       string `gorm:"column:play_url" form:"play_url" bson:"play_url"`
	FavoriteCount int64  `gorm:"column:favorite_count" form:"favorite_count" bson:"favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count" form:"comment_count" bson:"comment_count"`
}

type UserInfo struct {
	UserNickName  string `gorm:"column:user_nick_name" form:"user_nick_name" bson:"user_nick_name"`
	FollowCount   int64  `gorm:"column:follow_count" form:"follow_count" bson:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" form:"follower_count" bson:"follower_count"`
}

type FollowAndFollowerList struct {
	UserId int64 `gorm:"column:user_id" form:"user_id" bson:"user_id"`
}

func (l *FavoriteListRegisterLogic) FavoriteListRegister(req *types.FavoriteListRegisterHandlerRequest) (resp *types.FavoriteListRegisterHandlerResponse, err error) {
	//todo: add your logic here and delete this line
	if req.Token == "" {
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "未登录，请登录再查看",
		}, nil
	}

	ok, userId, err := tools.CheckToke(req.Token)
	if !ok {
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "登录过期，请重新登陆",
		}, err
	}

	//
	db, _ := mysqlconnect.SqlConnect() //连接数据库
	//

	favoriteVideoId := make([]int, 100) //预分配足够的内存，提升性能

	err1 := db.Table("favorite_list").Where("favorite_user_id = ?", userId).Select("favorite_video_id").Find(&favoriteVideoId).Error //查询到的点赞的视频id存在favorite切片里面
	if err1 != nil {
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "操作失败",
		}, err1
	}
	videoInfo := make([]VideoInfo, len(favoriteVideoId))

	err2 := db.Table("video_info").Where("video_id = ?", favoriteVideoId).Select("video_title", "author_id", "cover_url", "play_url", "favorite_count", "comment_count").Find(&videoInfo).Error //查出数据在videoInfo中

	if err2 != nil {
		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "操作失败",
		}, err2
	}

	videoList := make([]types.Video, len(favoriteVideoId)) //定义一个类型为Video的切片，每一个元素都包含一个喜欢列表的视频信息，最后返回这个数据即可
	userid := make([]int64, len(favoriteVideoId))

	for i := 0; i < len(favoriteVideoId); i++ {

		videoList[i].IsFavotite = true //1已点赞

		videoList[i].Id = int64(favoriteVideoId[i]) //2视频id

		videoList[i].VideoTitle = videoInfo[i].VideoTitle //3视频标题

		videoList[i].Author.UserId = videoInfo[i].AuthorId //4作者id

		userid[i] = videoInfo[i].AuthorId //将作者id再存在一个切片中，为下面查表准备

		videoList[i].CoverUrl = videoInfo[i].CoverUrl //5视频标题

		videoList[i].PlayUrl = videoInfo[i].PlayUrl //6视频播放地址

		videoList[i].FavoriteCount = videoInfo[i].FavoriteCount //7视频点赞数

		videoList[i].CommentCount = videoInfo[i].CommentCount //8视频评论数

	}

	userInfo := make([]UserInfo, len(favoriteVideoId))

	err3 := db.Table("user_info").Where("user_id in (?)", userid).Select("user_nick_name", "follow_count", "follower_count").Find(&userInfo).Error //查出数据在videoInfo中

	if err3 != nil {

		return &types.FavoriteListRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "查找失败",
		}, nil

	}

	for i := 0; i < len(favoriteVideoId); i++ {

		videoList[i].Author.Name = userInfo[i].UserNickName
		videoList[i].Author.FollowCount = userInfo[i].FollowCount
		videoList[i].Author.FollowerCount = userInfo[i].FollowerCount

	}

	////db.Table("follow_and_follower_list").Where("follower_id = ?",userId).Select("user_id").Find(&followAndFollowerList)  //这里的userId是用户，通过用户id反向查找其关注的人

	for i := 0; i < len(favoriteVideoId); i++ {
		videoList[i].Author.IsFollow = mysqlconnect.CheckIsFollow(int(userid[i]), userId)
	}

	return &types.FavoriteListRegisterHandlerResponse{

		StatusCode: 0,
		StatusMsg:  "查找成功",
		VideoList:  videoList,
	}, nil

}
