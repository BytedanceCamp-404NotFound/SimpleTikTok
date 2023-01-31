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
	CommentGetUserByUserIdRequest  = mysqlmanageserver.CommentGetUserByUserIdRequest
	CommentGetUserByUserIdResponse = mysqlmanageserver.CommentGetUserByUserIdResponse
	FavoriteVideoNumRequest        = mysqlmanageserver.FavoriteVideoNumRequest
	FavoriteVideoNumResponse       = mysqlmanageserver.FavoriteVideoNumResponse
	UserLoginRequest               = mysqlmanageserver.UserLoginRequest
	UserLoginResponse              = mysqlmanageserver.UserLoginResponse
	UserRegisterRequest            = mysqlmanageserver.UserRegisterRequest
	UserRegisterResponse           = mysqlmanageserver.UserRegisterResponse

	MySQLManageServer interface {
		// 1
		CommentGetUserByUserId(ctx context.Context, in *CommentGetUserByUserIdRequest, opts ...grpc.CallOption) (*CommentGetUserByUserIdResponse, error)
		// 2
		FavoriteVideoNum(ctx context.Context, in *FavoriteVideoNumRequest, opts ...grpc.CallOption) (*FavoriteVideoNumResponse, error)
		// 用户登陆校验
		UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error)
		// 用户注册
		UserRigster(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error)
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

// 1
func (m *defaultMySQLManageServer) CommentGetUserByUserId(ctx context.Context, in *CommentGetUserByUserIdRequest, opts ...grpc.CallOption) (*CommentGetUserByUserIdResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.CommentGetUserByUserId(ctx, in, opts...)
}

// 2
func (m *defaultMySQLManageServer) FavoriteVideoNum(ctx context.Context, in *FavoriteVideoNumRequest, opts ...grpc.CallOption) (*FavoriteVideoNumResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.FavoriteVideoNum(ctx, in, opts...)
}

// 用户登陆校验
func (m *defaultMySQLManageServer) UserLogin(ctx context.Context, in *UserLoginRequest, opts ...grpc.CallOption) (*UserLoginResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.UserLogin(ctx, in, opts...)
}

// 用户注册
func (m *defaultMySQLManageServer) UserRigster(ctx context.Context, in *UserRegisterRequest, opts ...grpc.CallOption) (*UserRegisterResponse, error) {
	client := mysqlmanageserver.NewMySQLManageServerClient(m.cli.Conn())
	return client.UserRigster(ctx, in, opts...)
}