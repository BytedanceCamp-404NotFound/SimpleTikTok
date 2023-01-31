// Code generated by goctl. DO NOT EDIT.
// Source: mysqlmanage.proto

package server

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/logic"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
)

type MySQLManageServerServer struct {
	svcCtx *svc.ServiceContext
	mysqlmanageserver.UnimplementedMySQLManageServerServer
}

func NewMySQLManageServerServer(svcCtx *svc.ServiceContext) *MySQLManageServerServer {
	return &MySQLManageServerServer{
		svcCtx: svcCtx,
	}
}

// 1
func (s *MySQLManageServerServer) CommentGetUserByUserId(ctx context.Context, in *mysqlmanageserver.CommentGetUserByUserIdRequest) (*mysqlmanageserver.CommentGetUserByUserIdResponse, error) {
	l := logic.NewCommentGetUserByUserIdLogic(ctx, s.svcCtx)
	return l.CommentGetUserByUserId(in)
}

// 2
func (s *MySQLManageServerServer) FavoriteVideoNum(ctx context.Context, in *mysqlmanageserver.FavoriteVideoNumRequest) (*mysqlmanageserver.FavoriteVideoNumResponse, error) {
	l := logic.NewFavoriteVideoNumLogic(ctx, s.svcCtx)
	return l.FavoriteVideoNum(in)
}

// 用户登陆校验
func (s *MySQLManageServerServer) UserLogin(ctx context.Context, in *mysqlmanageserver.UserLoginRequest) (*mysqlmanageserver.UserLoginResponse, error) {
	l := logic.NewUserLoginLogic(ctx, s.svcCtx)
	return l.UserLogin(in)
}

// 用户注册
func (s *MySQLManageServerServer) UserRigster(ctx context.Context, in *mysqlmanageserver.UserRegisterRequest) (*mysqlmanageserver.UserRegisterResponse, error) {
	l := logic.NewUserRigsterLogic(ctx, s.svcCtx)
	return l.UserRigster(in)
}