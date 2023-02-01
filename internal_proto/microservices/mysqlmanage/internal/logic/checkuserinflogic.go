package logic

import (
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)


type CheckUserInfLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckUserInfLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckUserInfLogic {
	return &CheckUserInfLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获得用户信息
func (l *CheckUserInfLogic) CheckUserInf(in *mysqlmanageserver.CheckUserInfRequest) (*mysqlmanageserver.CheckUserInfResponse, error) {
	// todo: add your logic here and delete this line
	var user User
	result := svc.DB.Table("user_info").Where("user_id = ?", in.UserId).Find(&user)
	err := result.Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]CheckUserInf [msg]gorm user_info.Find %v", err)
		return &mysqlmanageserver.CheckUserInfResponse{}, err
	}
	if result.RowsAffected == 0 {
		logx.Infof("[pkg]logic [func]CheckUserInf [msg]User does not exit")
		return &mysqlmanageserver.CheckUserInfResponse{}, nil
	}

	var ls CheckIsFollowLogic
	ins := mysqlmanageserver.CheckIsFollowRequest{UserId: in.UserId, FollowerId: in.FollowerId}
	out, err := ls.CheckIsFollow(&ins)
	if err != nil {
		logx.Errorf("[pkg]logic [func]CheckUserInf [msg]rpc CheckIsFollow %v", err)
		return &mysqlmanageserver.CheckUserInfResponse{}, err
	}
	return &mysqlmanageserver.CheckUserInfResponse{
		User: &mysqlmanageserver.UserInf{
			Users: &mysqlmanageserver.Users{
				UserId: user.UserID,
				UserNickName: user.UserNickName,
				FollowCount: user.FollowCount,
				FollowerCount: user.FollowerCount,
			},
			IsFollow: out.Ok,
		},
	}, nil
}
