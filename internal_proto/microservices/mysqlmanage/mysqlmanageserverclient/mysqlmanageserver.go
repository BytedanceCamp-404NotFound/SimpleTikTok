// Code generated by goctl. DO NOT EDIT.
// Source: mysqlmanage.proto

package mysqlmanageserverclient

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	AddVideoFavoriteRequest              = mysqlmanageserver.AddVideoFavoriteRequest
	AddVideoFavoriteResponse             = mysqlmanageserver.AddVideoFavoriteResponse
	CheckIsFollowRequest                 = mysqlmanageserver.CheckIsFollowRequest
	CheckIsFollowResponse                = mysqlmanageserver.CheckIsFollowResponse
	CheckUserInfRequest                  = mysqlmanageserver.CheckUserInfRequest
	CheckUserInfResponse                 = mysqlmanageserver.CheckUserInfResponse
	CreatePublishActionViedeInfoRequest  = mysqlmanageserver.CreatePublishActionViedeInfoRequest
	CreatePublishActionViedeInfoResponse = mysqlmanageserver.CreatePublishActionViedeInfoResponse
	FavoriteVideoNumRequest              = mysqlmanageserver.FavoriteVideoNumRequest
	FavoriteVideoNumResponse             = mysqlmanageserver.FavoriteVideoNumResponse
	FeedUserInfo                         = mysqlmanageserver.FeedUserInfo
	GetFavoriteVideoListRequest          = mysqlmanageserver.GetFavoriteVideoListRequest
	GetFavoriteVideoListResponse         = mysqlmanageserver.GetFavoriteVideoListResponse
	GetFeedUserInfoRequest               = mysqlmanageserver.GetFeedUserInfoRequest
	GetFeedUserInfoResponse              = mysqlmanageserver.GetFeedUserInfoResponse
	GetFeedVideoListRequest              = mysqlmanageserver.GetFeedVideoListRequest
	GetFeedVideoListResponse             = mysqlmanageserver.GetFeedVideoListResponse
	GetVideoListRequest                  = mysqlmanageserver.GetVideoListRequest
	GetVideoListResponse                 = mysqlmanageserver.GetVideoListResponse
	IsFavotiteRequest                    = mysqlmanageserver.IsFavotiteRequest
	IsFavotiteResponse                   = mysqlmanageserver.IsFavotiteResponse
	PublishActionVideoInfo               = mysqlmanageserver.PublishActionVideoInfo
	RelationActionRequest                = mysqlmanageserver.RelationActionRequest
	RelationActionResponse               = mysqlmanageserver.RelationActionResponse
	RelationFollowListRequest            = mysqlmanageserver.RelationFollowListRequest
	RelationFollowListResponse           = mysqlmanageserver.RelationFollowListResponse
	RelationFollowerListRequest          = mysqlmanageserver.RelationFollowerListRequest
	RelationFollowerListResponse         = mysqlmanageserver.RelationFollowerListResponse
	RelationFriendListRequest            = mysqlmanageserver.RelationFriendListRequest
	RelationFriendListResponse           = mysqlmanageserver.RelationFriendListResponse
	RelationUser                         = mysqlmanageserver.RelationUser
	SubVideoFavoriteRequest              = mysqlmanageserver.SubVideoFavoriteRequest
	SubVideoFavoriteResponse             = mysqlmanageserver.SubVideoFavoriteResponse
	UserInf                              = mysqlmanageserver.UserInf
	UserLoginRequest                     = mysqlmanageserver.UserLoginRequest
	UserLoginResponse                    = mysqlmanageserver.UserLoginResponse
	UserRegisterRequest                  = mysqlmanageserver.UserRegisterRequest
	UserRegisterResponse                 = mysqlmanageserver.UserRegisterResponse
	Users                                = mysqlmanageserver.Users
	VideoInfo                            = mysqlmanageserver.VideoInfo
	VideoNumRequest                      = mysqlmanageserver.VideoNumRequest
	VideoNumResponse                     = mysqlmanageserver.VideoNumResponse

	MySQLManageServer interface {
		// ??????????????????????????????
		GetFeedUserInfo(ctx context.Context, in *GetFeedUserInfoRequest, opts ...grpc.CallOption) (*GetFeedUserInfoResponse, error)
		// ??????????????????????????????
		GetFeedVideoList(ctx context.Context, in *GetFeedVideoListRequest, opts ...grpc.CallOption) (*GetFeedVideoListResponse, error)
		// ??????????????????
		UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
		// ????????????
		UserRigster(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
		// ????????????????????????????????????
		CreatePublishActionViedeInfo(ctx context.Context, in *CreatePublishActionViedeInfoRequest, opts ...grpc.CallOption) (*CreatePublishActionViedeInfoResponse, error)
		// ??????????????????
		CheckUserInf(ctx context.Context, in *CheckUserInfRequest, opts ...grpc.CallOption) (*CheckUserInfResponse, error)
		// ????????????
		CheckIsFollow(ctx context.Context, in *CheckIsFollowRequest, opts ...grpc.CallOption) (*CheckIsFollowResponse, error)
		// ????????????
		IsFavotite(ctx context.Context, in *IsFavotiteRequest, opts ...grpc.CallOption) (*IsFavotiteResponse, error)
		// ????????????????????????
		GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...grpc.CallOption) (*GetVideoListResponse, error)
		// ????????????????????????
		VideoNum(ctx context.Context, in *VideoNumRequest, opts ...grpc.CallOption) (*VideoNumResponse, error)
		// ?????????????????????
		RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error)
		// ????????????
		RelationFollowerList(ctx context.Context, in *RelationFollowerListRequest, opts ...grpc.CallOption) (*RelationFollowerListResponse, error)
		// ????????????
		RelationFollowList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListResponse, error)
		// ??????????????????
		RelationFriendList(ctx context.Context, in *RelationFriendListRequest, opts ...grpc.CallOption) (*RelationFriendListResponse, error)
		// ??????????????????
		FavoriteVideoNum(ctx context.Context, in *FavoriteVideoNumRequest, opts ...grpc.CallOption) (*FavoriteVideoNumResponse, error)
		// ??????????????????
		GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListResponse, error)
		// ??????
		AddVideoFavorite(ctx context.Context, in *AddVideoFavoriteRequest, opts ...grpc.CallOption) (*AddVideoFavoriteResponse, error)
		// ????????????
		SubVideoFavorite(ctx context.Context, in *SubVideoFavoriteRequest, opts ...grpc.CallOption) (*SubVideoFavoriteResponse, error)
	}

	defaultMySQLManageServer struct {
		cli zrpc.Client
	}
)

func NewMySQLManageServer(cli zrpc.Client) MySQLManageServer {
	return &defaultMySQLManageServer{
		cli: cli,
	}
}

// ??????????????????????????????
func (m *defaultMySQLManageServer) GetFeedUserInfo(ctx context.Context, in *GetFeedUserInfoRequest, opts ...grpc.CallOption) (*GetFeedUserInfoResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.GetFeedUserInfo(ctx, in, opts...)
}

// ??????????????????????????????
func (m *defaultMySQLManageServer) GetFeedVideoList(ctx context.Context, in *GetFeedVideoListRequest, opts ...grpc.CallOption) (*GetFeedVideoListResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.GetFeedVideoList(ctx, in, opts...)
}

// ??????????????????
func (m *defaultMySQLManageServer) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

// ????????????
func (m *defaultMySQLManageServer) UserRigster(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.UserRigster(ctx, in, opts...)
}

// ????????????????????????????????????
func (m *defaultMySQLManageServer) CreatePublishActionViedeInfo(ctx context.Context, in *CreatePublishActionViedeInfoRequest, opts ...grpc.CallOption) (*CreatePublishActionViedeInfoResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.CreatePublishActionViedeInfo(ctx, in, opts...)
}

// ??????????????????
func (m *defaultMySQLManageServer) CheckUserInf(ctx context.Context, in *CheckUserInfRequest, opts ...grpc.CallOption) (*CheckUserInfResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.CheckUserInf(ctx, in, opts...)
}

// ????????????
func (m *defaultMySQLManageServer) CheckIsFollow(ctx context.Context, in *CheckIsFollowRequest, opts ...grpc.CallOption) (*CheckIsFollowResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.CheckIsFollow(ctx, in, opts...)
}

// ????????????
func (m *defaultMySQLManageServer) IsFavotite(ctx context.Context, in *IsFavotiteRequest, opts ...grpc.CallOption) (*IsFavotiteResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.IsFavotite(ctx, in, opts...)
}

// ????????????????????????
func (m *defaultMySQLManageServer) GetVideoList(ctx context.Context, in *GetVideoListRequest, opts ...grpc.CallOption) (*GetVideoListResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.GetVideoList(ctx, in, opts...)
}

// ????????????????????????
func (m *defaultMySQLManageServer) VideoNum(ctx context.Context, in *VideoNumRequest, opts ...grpc.CallOption) (*VideoNumResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.VideoNum(ctx, in, opts...)
}

// ?????????????????????
func (m *defaultMySQLManageServer) RelationAction(ctx context.Context, in *RelationActionRequest, opts ...grpc.CallOption) (*RelationActionResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.RelationAction(ctx, in, opts...)
}

// ????????????
func (m *defaultMySQLManageServer) RelationFollowerList(ctx context.Context, in *RelationFollowerListRequest, opts ...grpc.CallOption) (*RelationFollowerListResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.RelationFollowerList(ctx, in, opts...)
}

// ????????????
func (m *defaultMySQLManageServer) RelationFollowList(ctx context.Context, in *RelationFollowListRequest, opts ...grpc.CallOption) (*RelationFollowListResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.RelationFollowList(ctx, in, opts...)
}

// ??????????????????
func (m *defaultMySQLManageServer) RelationFriendList(ctx context.Context, in *RelationFriendListRequest, opts ...grpc.CallOption) (*RelationFriendListResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.RelationFriendList(ctx, in, opts...)
}

// ??????????????????
func (m *defaultMySQLManageServer) FavoriteVideoNum(ctx context.Context, in *FavoriteVideoNumRequest, opts ...grpc.CallOption) (*FavoriteVideoNumResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.FavoriteVideoNum(ctx, in, opts...)
}

// ??????????????????
func (m *defaultMySQLManageServer) GetFavoriteVideoList(ctx context.Context, in *GetFavoriteVideoListRequest, opts ...grpc.CallOption) (*GetFavoriteVideoListResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.GetFavoriteVideoList(ctx, in, opts...)
}

// ??????
func (m *defaultMySQLManageServer) AddVideoFavorite(ctx context.Context, in *AddVideoFavoriteRequest, opts ...grpc.CallOption) (*AddVideoFavoriteResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.AddVideoFavorite(ctx, in, opts...)
}

// ????????????
func (m *defaultMySQLManageServer) SubVideoFavorite(ctx context.Context, in *SubVideoFavoriteRequest, opts ...grpc.CallOption) (*SubVideoFavoriteResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.SubVideoFavorite(ctx, in, opts...)
}
