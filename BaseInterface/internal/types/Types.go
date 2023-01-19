// Code generated by goctl. DO NOT EDIT.
package types

type FeedHandlerRequest struct {
	LastestTime int64  `json:"lastest_time"`
	Token       string `json:"token"`
}

type FeedHandlerResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	VideoList  Video  `json:"video_list"`
	NextTime   int64  `json:"next_time"`
}

type Video struct {
	Id            int64  `json:"id"`
	Author        User   `json:"author"`
	PlayUrl       string `json:"play_url"`
	CoverUrl      string `json:"cover_url"`
	FavoriteCount int64  `json:"favorite_count"`
	CommentCount  int64  `json:"comment_count"`
	IsFavotite    bool   `json:"is_favorite"`
	Title         string `json:"title"`
}

type User struct {
	Id            int64  `json:"id"`
	Name          string `json:"name"`
	FollowCount   string `json:"follow_count"`
	FollowerCount string `json:"follower_count"`
	IsFollow      string `json:"is_follow"`
}

type UserRegisterHandlerRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserRegisterHandlerResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type UserloginHandlerRequest struct {
	UserName string `json:"username"`
	PassWord string `json:"password"`
}

type UserloginHandlerResponse struct {
	StatusCode int64  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	Token      string `json:"token"`
	UserID     int64  `json:"user_id"`
}

type UserHandlerRequest struct {
	UserID int64  `json:"user_id"`
	Token  string `json:"token"`
}

type UserHandlerResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	User       User   `json:"user"`
}

type PublishActionHandlerRequest struct {
	Token string `json:"token"`
	Data  int32  `json:"data"` // 存疑bytes
	Title string `json:"title"`
}

type PublishActionHandlerResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type PublishListHandlerRequest struct {
	StatusCode int32 `json:"status_code"`
}

type PublishListHandlerResponse struct {
	StatusCode string `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
}

type CommmentActionHandlerRequest struct {
	StatusCode int32 `json:"StatusCode"`
}

type CommmentActionHandlerResponse struct {
	StatusCode string `json:"UserName"`
	StatusMsg  string `json:"StatusMsg"`
}

type CommmentListHandlerRequest struct {
	StatusCode int32 `json:"StatusCode"`
}

type CommmentListHandlerResponse struct {
	StatusCode string `json:"UserName"`
	StatusMsg  string `json:"StatusMsg"`
}

type RelationActionHandlerRequest struct {
	StatusCode int32 `json:"StatusCode"`
}

type RelationActionHandlerResponse struct {
	StatusCode string `json:"UserName"`
	StatusMsg  string `json:"StatusMsg"`
}

type RelationFollowListHandlerRequest struct {
	StatusCode int32 `json:"StatusCode"`
}

type RelationFollowListHandlerResponse struct {
	StatusCode string `json:"UserName"`
	StatusMsg  string `json:"StatusMsg"`
}

type RelationFollowerListHandlerRequest struct {
	StatusCode int32 `json:"StatusCode"`
}

type RelationFollowerListHandlerResponse struct {
	StatusCode string `json:"UserName"`
	StatusMsg  string `json:"StatusMsg"`
}
