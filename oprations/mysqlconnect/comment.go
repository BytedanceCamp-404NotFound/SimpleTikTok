package mysqlconnect

import (
	"github.com/zeromicro/go-zero/core/logx"
)

type CommentUser struct {
	UserId        int64  `gorm:"column:user_id"        json:"id"               form:"user_id"        bson:"user_id"`
	Name          string `gorm:"column:user_nick_name" json:"name"             form:"name"           bson:"name"`
	FollowCount   int64  `gorm:"column:follow_count"   json:"follow_count"     form:"follow_count"   bson:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"   form:"follower_count" bson:"follower_count"`
	IsFollow      bool   `json:"is_follow"             form:"is_follow"        bson:"is_follow"`
}

type VideoMessage struct {
	VideoID       int64  `gorm:"column:video_id"`
	AuthorID      int64  `gorm:"column:author_id"`
	PlayUrl       string `gorm:"column:play_url"`
	CoverUrl      string `gorm:"column:cover_url"`
	FavoriteCount int64  `gorm:"column:favorite_count"`
	CommentCount  int64  `gorm:"column:comment_count"`
	UpdateTime    bool   `gorm:"column:update_time"`
	DeletedTime   string `gorm:"column:deleted_time"`
	VideoTitle    string `gorm:"column:video_title"`
}

func CommentGetUserByUserId(userId int) (CommentUser, error) {
	var user CommentUser
	db := GormDB
	err := db.Table("user_info").Where("user_id=?", userId).First(&user).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]CommentGetUserByUserId [msg]search user_info failed, [err]%v", err)
		return CommentUser{}, err
	}
	return user, err
}

func UpdateComment(videoId int64, actionType int32) error {
	var comm VideoMessage
	db := GormDB
	err := db.Table("video_info").Where("video_id=?", videoId).First(&comm).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]UpdateComment [msg]search comment_count failed, [err]%v", err)
		return err
	}
	if actionType == 2 {
		comm.CommentCount--
	} else {
		comm.CommentCount++
	}
	err = db.Table("video_info").Where("video_id=?", videoId).Updates(comm).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]UpdateCommentNumByVideoId [msg]update comment_count failed, [err]%v", err)
		return err
	}
	return nil
}
