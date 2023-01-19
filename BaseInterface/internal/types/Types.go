// Code generated by goctl. DO NOT EDIT.
package types

type FeedHandlerRequest struct {
	LastestTime int64  `from:"LastestTime"`
	Token       string `from:"token"`
}

type FeedHandlerResponse struct {
	StatusCode int64  `from:"status_code"`
	StatusMsg  string `from:"status_msg"`
	VideoList  Video  `from:"video_list"`
	NextTime   int64  `from:"next_time"`
}

type Video struct {
	Id            int64  `from:"id"`
	Author        User   `from:"author"`
	PlayUrl       string `from:"play_url"`
	CoverUrl      string `from:"cover_url"`
	FavoriteCount int64  `from:"favorite_count"`
	CommentCount  int64  `from:"comment_count"`
	IsFavotite    bool   `from:"is_favorite"`
	Title         string `from:"title"`
}

type User struct {
	Id            int64  `from:"id"`
	Name          string `from:"name"`
	FollowCount   string `from:"follow_count"`
	FollowerCount string `from:"follower_count"`
	IsFollow      string `from:"is_follow"`
}

type UserRegisterHandlerRequest struct {
	UserName string `from:"username"`
	PassWord string `from:"password"`
}

type UserRegisterHandlerResponse struct {
	StatusCode int64  `from:"status_code"`
	StatusMsg  string `from:"status_msg"`
	Token      string `from:"token"`
	UserID     int64  `from:"user_id"`
}

type UserloginHandlerRequest struct {
	UserName string `from:"username"`
	PassWord string `from:"password"`
}

type UserloginHandlerResponse struct {
	StatusCode int64  `from:"status_code"`
	StatusMsg  string `from:"status_msg"`
	Token      string `from:"token"`
	UserID     int64  `from:"user_id"`
}

type UserHandlerRequest struct {
	UserID int64  `from:"user_id"`
	Token  string `from:"token"`
}

type UserHandlerResponse struct {
	StatusCode int32  `from:"status_code"`
	StatusMsg  string `from:"status_msg"`
	User       User   `from:"user"`
}

type PublishActionHandlerRequest struct {
	Token string `from:"token"`
	Data  int32  `from:"data"` // 存疑bytes
	Title string `from:"title"`
}

type PublishActionHandlerResponse struct {
	StatusCode int32  `from:"status_code"`
	StatusMsg  string `from:"status_msg"`
}

type PublishListHandlerRequest struct {
	StatusCode int32 `from:"status_code"`
}

type PublishListHandlerResponse struct {
	StatusCode string `from:"status_code"`
	StatusMsg  string `from:"status_msg"`
}

type CommmentActionHandlerRequest struct {
	StatusCode int32 `from:"StatusCode"`
}

type CommmentActionHandlerResponse struct {
	StatusCode string `from:"UserName"`
	StatusMsg  string `from:"StatusMsg"`
}

type CommmentListHandlerRequest struct {
	StatusCode int32 `from:"StatusCode"`
}

type CommmentListHandlerResponse struct {
	StatusCode string `from:"UserName"`
	StatusMsg  string `from:"StatusMsg"`
}

type RelationActionHandlerRequest struct {
	StatusCode int32 `from:"StatusCode"`
}

type RelationActionHandlerResponse struct {
	StatusCode string `from:"UserName"`
	StatusMsg  string `from:"StatusMsg"`
}

type RelationFollowListHandlerRequest struct {
	StatusCode int32 `from:"StatusCode"`
}

type RelationFollowListHandlerResponse struct {
	StatusCode string `from:"UserName"`
	StatusMsg  string `from:"StatusMsg"`
}

type RelationFollowerListHandlerRequest struct {
	StatusCode int32 `from:"StatusCode"`
}

type RelationFollowerListHandlerResponse struct {
	StatusCode string `from:"UserName"`
	StatusMsg  string `from:"StatusMsg"`
}
