package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"
	"context"
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowerListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowerListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowerListLogic {
	return &RelationFollowerListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowerListLogic) RelationFollowerList(req *types.RelationFollowerListHandlerRequest) (resp *types.RelationFollowerListHandlerResponse, err error) {

	testUserList_ := make([]types.RelationUser, 0)
	tempUserList := types.RelationUser{}
	resultJson := types.RelationFollowerListHandlerResponse{StatusCode: 501, StatusMsg: "token失效，请重新登录"}
	result, TokenToUserID, err := tools.CheckToke(req.Token)
	if !result {
		return &resultJson, err
	}

	db := mysqlconnect.GormDB
	//#查找某个账号的粉丝列表   user_id：账号
	userInfoList, err := db.Raw(fmt.Sprintf("SELECT * FROM user_info where user_id in(SELECT follower_id FROM follow_and_follower_list where user_id=%d)", req.UserId)).Rows()

	if err != nil {
		resultJson.StatusCode = 500
		resultJson.StatusMsg = err.Error()
	} else {
		for userInfoList.Next() {
			userInfoList.Scan(&tempUserList.Id, &tempUserList.Name, &tempUserList.FollowCount, &tempUserList.FollowerCount, &tempUserList.IsFollow) //！！err :查询出来的列数不同、数据格式不同时会报错，不影响格式正确的变量
			testUserList_ = append(testUserList_, tempUserList)
		}
		//查询一遍上面查出来的id，是否已被当前登录的账号关注
		for i := 0; i < len(testUserList_); i++ {
			//是否已关注  follower_id：哪个账号想查询  user_id:哪个账号想被查询是否已被关注
			sqlString := fmt.Sprintf("SELECT * FROM follow_and_follower_list where follower_id=%d and user_id=%d", TokenToUserID, testUserList_[i].Id)
			if temtpp, _ := db.Raw(sqlString).Rows(); temtpp.Next() {
				testUserList_[i].IsFollow = true
			}
		}
		resultJson.StatusCode = 200
		resultJson.StatusMsg = "查询成功"
	}

	resultJson.UserList = testUserList_
	return &resultJson, err

}
