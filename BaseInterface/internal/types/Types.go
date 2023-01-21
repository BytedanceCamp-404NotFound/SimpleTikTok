// Code generated by goctl. DO NOT EDIT.
package types

type FeedHandlerRequest struct {
	LastestTime int64  `form:"lastest_time"`
	Token       string `form:"token"`
}

type FeedHandlerResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	VideoList  Video  `json:"video_list"`
	NextTime   int64  `json:"next_time"`
}

type Video struct {
	Id            int64  `form:"id"`
	Author        User   `form:"author"`
	PlayUrl       string `form:"play_url"`
	CoverUrl      string `form:"cover_url"`
	FavoriteCount int64  `form:"favorite_count"`
	CommentCount  int64  `form:"comment_count"`
	IsFavotite    bool   `form:"is_favorite"`
	Title         string `form:"title"`
}

type User struct {
	UserId        int64  `gorm:"user_id" form:"user_id" bson:"user_id"`
	Name          string `gorm:"user_nick_name" form:"name" bson:"name"`
	FollowCount   int64  `gorm:"follow_count" form:"follow_count" bson:"follow_count"`
	FollowerCount int64  `gorm:"follower_count" form:"follower_count" bson:"follower_count"`
	IsFollow      bool   `form:"is_follow" bson:"is_follow"`
}

type UserHandlerRequest struct {
	UserID int64  `form:"user_id"`
	Token  string `form:"token"`
}

type UserHandlerResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	User       User   `json:"user"`
}

type UserRegisterHandlerRequest struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type UserRegisterHandlerResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type UserloginHandlerRequest struct {
	UserName string `form:"username"`
	PassWord string `form:"password"`
}

type UserloginHandlerResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type PublishActionHandlerRequest struct {
	Token string `form:"token"`
	Title string `form:"title"`
}

type PublishActionHandlerResponse struct {
	StatusCode int32  `form:"status_code,default=400"`
	StatusMsg  string `form:"status_msg,optional"`
}

type PublishListHandlerRequest struct {
	StatusCode int32 `form:"status_code"`
}

type PublishListHandlerResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type Comment struct {
	VideoId    int64  `form:"video_id" bson:"video_id"` //视频id
	User       User   `form:"user" bson:"user"`
	Content    string `form:"content" bson:"content"`
	CreateDate string `form:"create_date" bson:"create_date"`
}

type CommmentActionHandlerRequest struct {
	Token       string `form:"token"`
	VideoId     int64  `form:"video_id"`
	ActionType  int32  `form:"action_type"`
	CommentText string `form:"comment_text"`
	CommentId   int64  `form:"comment_id,optional"`
}

type CommmentActionHandlerResponse struct {
	StatusCode int32   `json:"status_code"`
	StatusMsg  string  `json:"status_msg"`
	Comment    Comment `json:"comment"`
}

type CommmentListHandlerRequest struct {
	Token   string `form:"token"`
	VideoId int64  `form:"video_id"`
}

type CommmentListHandlerResponse struct {
	StatusCode  int32   `json:"status_code"`
	StatusMsg   string  `json:"status_msg"`
	CommentList Comment `json:"comment_list"`
}

type RelationActionHandlerRequest struct {
	Token       string `form:"token"`
	To_user_id  int32  `form:"to_user_id"`
	Sction_type int32  `form:"action_type"`
}

type RelationActionHandlerResponse struct {
	StatusCode int32  `from:"status_code"`
	StatusMsg  string `from:"status_msg"`
}

type RelationFollowListHandlerRequest struct {
	Token  string `form:"token"`
	UserId int32  `form:"user_id"`
}

type RelationFollowListHandlerResponse struct {
	StatusCode int32          `from:"status_code"`
	StatusMsg  string         `from:"status_msg"`
	UserList   []RelationUser `from:"user_list"`
}

type RelationFollowerListHandlerRequest struct {
	Token  string `form:"token"`
	UserId int32  `form:"user_id"`
}

type RelationFollowerListHandlerResponse struct {
	StatusCode int32          `from:"status_code"`
	StatusMsg  string         `from:"status_msg"`
	UserList   []RelationUser `from:"user_list"`
}

type RelationUser struct {
	Id            int64  `from:"id" gorm:"column:user_id"`
	Name          string `from:"name" gorm:"column:user_nick_name"`
	FollowCount   int32  `form:"follow_count" gorm:"column:follow_count"`
	FollowerCount int32  `from:"follower_count" gorm:"column:follower_count"`
	IsFollow      bool   `from:"is_follow"`
}
