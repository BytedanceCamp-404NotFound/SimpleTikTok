package BaseInterface

import (
	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/commonerror"
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
	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		return &types.UserHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "请登录！",
			User:       types.User{},
		}, err
	}
	ui, err := l.svcCtx.MySQLManageRpc.CheckUserInf(l.ctx, &mysqlmanageserver.CheckUserInfRequest{
		UserId:     req.UserID,
		FollowerId: int64(id),
	})
	if err != nil {
		return &types.UserHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_DB_ERROR),
			StatusMsg: "该用户不存在！",
			User: types.User{},
		},nil
	}
	return &types.UserHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "查询成功！",
		User: types.User{
			UserId:        ui.User.Users.UserId,
			Name:          ui.User.Users.UserNickName,
			FollowCount:   ui.User.Users.FollowCount,
			FollowerCount: ui.User.Users.FollowerCount,
			IsFollow:      ui.User.IsFollow,
		},
	}, nil
}
