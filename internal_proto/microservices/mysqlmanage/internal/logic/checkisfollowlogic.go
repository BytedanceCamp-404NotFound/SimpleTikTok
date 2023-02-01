package logic

import (
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type CheckIsFollowLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCheckIsFollowLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CheckIsFollowLogic {
	return &CheckIsFollowLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 是否关注
func (l *CheckIsFollowLogic) CheckIsFollow(in *mysqlmanageserver.CheckIsFollowRequest) (*mysqlmanageserver.CheckIsFollowResponse, error) {
	// todo: add your logic here and delete this line
	var record int64 = 0
	fmt.Println(in.UserId, in.FollowerId)
	err := svc.DB.Table("follow_and_follower_list").Select("record_id").Where("user_id = ? and follower_id = ?",in.UserId, in.FollowerId).Take(&record).Error
	if err != nil {
		logx.Errorf("[pkg]logic [func]CheckIsFollow [msg]gorm follow_and_follower_list.Take %v",err)
		return &mysqlmanageserver.CheckIsFollowResponse{Ok: false},err
	}
	return &mysqlmanageserver.CheckIsFollowResponse{Ok: record != 0}, nil
}
