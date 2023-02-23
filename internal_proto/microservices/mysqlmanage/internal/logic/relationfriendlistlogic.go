package logic

import (
	"context"
	"fmt"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFriendListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 读取好友列表
func (l *RelationFriendListLogic) RelationFriendList(in *mysqlmanageserver.RelationFriendListRequest) (*mysqlmanageserver.RelationFriendListResponse, error) {
	// todo: add your logic here and delete this line
	testUserList_ := make([]*mysqlmanageserver.RelationUser, 0)

	db := svc.DB
	//#查询某个帐号的互关列表   follower_id：账号
	userInfoList, err := db.Raw(fmt.Sprintf("SELECT * FROM follow_and_follower_list where follower_id in(SELECT user_id FROM follow_and_follower_list where follower_id = %d) and user_id = %d", in.UserID, in.UserID)).Rows()
	if err != nil {
		return &mysqlmanageserver.RelationFriendListResponse{}, err
	} else {
		for userInfoList.Next() {
			tempUserList := mysqlmanageserver.RelationUser{}
			userInfoList.Scan(&tempUserList.Id, &tempUserList.Name, &tempUserList.FollowCount, &tempUserList.FollowerCount, &tempUserList.IsFollow) //！！err :查询出来的列数不同、数据格式不同时会报错，不影响格式正确的变量
			testUserList_ = append(testUserList_, &tempUserList)
		}
		//查询一遍上面查出来的id，是否已被当前登录的账号关注
		for i := 0; i < len(testUserList_); i++ {
			var ls CheckIsFollowLogic
			isFollow := mysqlmanageserver.CheckIsFollowRequest{UserId: int64(testUserList_[i].Id), FollowerId: int64(in.LoginUserID)}
			result, _ := ls.CheckIsFollow(&isFollow)
			if result.Ok {
				testUserList_[i].IsFollow = true
			}
		}
	}
	return &mysqlmanageserver.RelationFriendListResponse{RelationUser: testUserList_}, nil
}
