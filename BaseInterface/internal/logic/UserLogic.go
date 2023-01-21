package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/sql"
	tools "SimpleTikTok/tools/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLogic {
	return &UserLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserLogic) User(req *types.UserHandlerRequest) (resp *types.UserHandlerResponse, err error) {
	// todo: add your logic here and delete this line
	ok, id := tools.CheckToke(req.Token)
	if !ok {
		return &types.UserHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "请登录！",
			User:       types.User{},
		}, err
	}
	ui ,ok:=  sql.CheckUserInf(int(req.UserID),id)
	if !ok {
		return &types.UserHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "查询的用户不存在！",
			User:       types.User{},
		}, err
	}
	return &types.UserHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功！",
		User:       types.User{
						UserId:         ui.User.UserID, 
						Name:           ui.User.UserName, 
						FollowCount:    ui.User.FollowCount, 
						FollowerCount:  ui.User.FollowerCount, 
						IsFollow:       ui.IsFollow,
					},
	}, err
}
