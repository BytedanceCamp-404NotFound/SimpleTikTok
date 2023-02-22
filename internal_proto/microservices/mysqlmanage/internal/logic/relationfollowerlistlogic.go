package logic

import (
	"context"
	"fmt"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/redisconnect"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	KVDBName = "FollowAndFollowerList:" //Redis-Set
)

type RelationUserInfo struct {
	UserID        int64  // 用户ID
	UserNickName  string // 用户名称
	FollowCount   int64  // 关注总数
	FollowerCount int64  // 粉丝总数
	IsFollow      bool   // true-已关注
}

// 粉丝列表
func (l *RelationFollowerListLogic) RelationFollowerList(in *mysqlmanageserver.RelationFollowerListRequest) (*mysqlmanageserver.RelationFollowerListResponse, error) {
	testUserList_ := make([]*mysqlmanageserver.RelationUser, 0)

	RedisDB, _ := redisconnect.RedisConnect()
	followerIDString := fmt.Sprintf("%s%s:%d", KVDBName, "follower_id", in.UserID)
	followerIDList, _ := RedisDB.SMembers(l.ctx, followerIDString).Result() // 找到粉丝的列表

	for _, val := range followerIDList {
		var userInfo RelationUserInfo
		_ = svc.DB.Table("user_info").Where("user_id = ?", val).Find(&userInfo).Error
		followerIDString := fmt.Sprintf("%s%s:%d", KVDBName, "follower_id", userInfo.UserID)
		userInfoFollowCountList, _ := RedisDB.SMembers(l.ctx, followerIDString).Result() // 找到关注的列表
		userIDString := fmt.Sprintf("%s%s:%d", KVDBName, "user_id", userInfo.UserID)
		userInfoFollowerCountList, _ := RedisDB.SMembers(l.ctx, userIDString).Result() // 找到粉丝的列表

		isFollowRedis := RedisDB.SIsMember(l.ctx, followerIDString, in.UserID).Val() // redis查看关注列表中是否有这个人
		testUserList_ = append(testUserList_, &mysqlmanageserver.RelationUser{
			Id:            userInfo.UserID,
			Name:          userInfo.UserNickName,
			FollowCount:   int64(len(userInfoFollowCountList)),
			FollowerCount: int64(len(userInfoFollowerCountList)),
			IsFollow:      isFollowRedis,
		})
	}

	return &mysqlmanageserver.RelationFollowerListResponse{RelationUser: testUserList_}, nil
}

// func (l *RelationFollowerListLogic) RelationFollowerList(in *mysqlmanageserver.RelationFollowerListRequest) (*mysqlmanageserver.RelationFollowerListResponse, error) {
// 	// todo: add your logic here and delete this line

// 	testUserList_ := make([]*mysqlmanageserver.RelationUser, 0)

// 	db := svc.DB
// 	//#查找某个账号的粉丝列表   user_id：账号
// 	userInfoList, err := db.Raw(fmt.Sprintf("SELECT * FROM user_info where user_id in(SELECT follower_id FROM follow_and_follower_list where user_id=%d)", in.UserID)).Rows()

// 	if err != nil {
// 		return &mysqlmanageserver.RelationFollowerListResponse{}, err
// 	} else {
// 		for userInfoList.Next() {
// 			tempUserList := mysqlmanageserver.RelationUser{}
// 			userInfoList.Scan(&tempUserList.Id, &tempUserList.Name, &tempUserList.FollowCount, &tempUserList.FollowerCount, &tempUserList.IsFollow) //！！err :查询出来的列数不同、数据格式不同时会报错，不影响格式正确的变量
// 			testUserList_ = append(testUserList_, &tempUserList)
// 		}
// 		//查询一遍上面查出来的id，是否已被当前登录的账号关注
// 		for i := 0; i < len(testUserList_); i++ {

// 			var ls CheckIsFollowLogic
// 			isFollow := mysqlmanageserver.CheckIsFollowRequest{UserId: int64(testUserList_[i].Id), FollowerId: int64(in.LoginUserID)}
// 			result, _ := ls.CheckIsFollow(&isFollow)
// 			if result.Ok {
// 				testUserList_[i].IsFollow = true
// 			}
// 		}
// 	}
// 	logx.Infof("\n\n%+v\n\n", &testUserList_)

// 	return &mysqlmanageserver.RelationFollowerListResponse{RelationUser: testUserList_}, nil
// }
