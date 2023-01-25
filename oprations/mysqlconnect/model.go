package mysqlconnect

import (
	"time"

	"gorm.io/gorm"
)

// 默认情况下，名为 `ID` 的字段会作为表的主键
type VideoInfo struct {
	gorm.Model
	VideID        int64 `gorm:"index;default: 0;unsigned"`
	AuthorID      int64
	PlayUrl       string // 视频播放地址
	CoverUrl      string // 视频封面地址
	FavoriteCount int64  // 视频的点赞总数
	CommentCount  int64  // 视频的评论总数
	IsFavotite    bool   // true-已点赞
	VideoTitle    string // 视频标题
}

type UserInfo struct {
	gorm.Model
	UserID        int64  // 用户ID
	Name          string // 用户名称
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	IsFollow      bool   // true-已关注
}


type Favorite_list struct {
	Favorite_video_id int64 
	Favorite_user_id int
	Record_time time.Time
}
