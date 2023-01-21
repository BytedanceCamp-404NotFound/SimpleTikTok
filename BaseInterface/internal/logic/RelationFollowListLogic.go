package logic

import (
	"context"
	"fmt"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/sql"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFollowListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFollowListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFollowListLogic {
	return &RelationFollowListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFollowListLogic) RelationFollowList(req *types.RelationFollowListHandlerRequest) (resp *types.RelationFollowListHandlerResponse, err error) {

	testUserList_ := make([]types.RelationUser, 0)
	tempUserList := types.RelationUser{}
	resultJson := types.RelationFollowListHandlerResponse{StatusCode: 501, StatusMsg: "token失效，请重新登录"}
	result, TokenToUserID := tools.CheckToke(req.Token)
	if !result {
		return &resultJson, err
	}

	db, _ := sql.SqlConnect()
	//#查询某个账号关注的所有账号   follower_id：账号
	userInfoList, err := db.Raw(fmt.Sprintf("SELECT * FROM user_info where user_id in(SELECT user_id FROM follow_and_follower_list where follower_id = %d)", req.UserId)).Rows()

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
