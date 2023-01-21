package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/sql"
	tools "SimpleTikTok/tools/token"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {

	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

type Follow_and_follower_list struct {
	UserID     int
	Token      string
	RecordTime time.Time
}

func (l *RelationActionLogic) RelationAction(req *types.RelationActionHandlerRequest) (resp *types.RelationActionHandlerResponse, err error) {
	//	db.exec("sql语句")		//执行插入删除等操作使用
	//	db.raw("sql语句")		//执行查询操作时使用
	var sqlString string
	resultJson := types.RelationActionHandlerResponse{StatusCode: 501, StatusMsg: "token失效，请重新登录"}

	result, TokenToUserID := tools.CheckToke(req.Token)
	if !result {
		return &resultJson, err
	}

	if req.Sction_type == 1 { //关注
		//#关注账号 user_id：被关注的账号  follower_id：哪个账号要关注
		sqlString = fmt.Sprintf("INSERT into follow_and_follower_list(user_id,follower_id) VALUES(%d,%d)", req.To_user_id, TokenToUserID)
	} else if req.Sction_type == 2 { //取消关注
		//取消关注.  user_id：要被取消关注的账号   follower_id：哪个账号要取消关注
		sqlString = fmt.Sprintf("DELETE FROM follow_and_follower_list WHERE user_id=%d and follower_id=%d", req.To_user_id, TokenToUserID)
	}

	db, _ := sql.SqlConnect()
	b := db.Exec(sqlString)
	fmt.Println(b.RowsAffected)
	fmt.Printf("%+v\n\n\n", b)
	if b.RowsAffected <= 0 { //失败  ：传入的数据不正确
		resultJson.StatusCode = 500
		resultJson.StatusMsg = "请求失败，请稍后再试！"
	} else if req.Sction_type == 1 { //关注:以上是将关注的用户id绑定到数据库，以下是被关注的粉丝数+1
		sqlString = fmt.Sprintf("UPDATE user_info Set follower_count=follower_count+1 where user_id=%d", req.To_user_id)
		if db.Exec(sqlString).RowsAffected > 0 {
			resultJson.StatusCode = 200
			resultJson.StatusMsg = "关注成功"
		}
	} else if req.Sction_type == 2 { //取消关注  //关注:以上是将取消关注的用户id绑定信息从数据库删除，以下是被关注的粉丝数-1
		sqlString = fmt.Sprintf("UPDATE user_info Set follower_count=follower_count-1 where user_id=%d", req.To_user_id)
		if db.Exec(sqlString).RowsAffected > 0 {
			resultJson.StatusCode = 200
			resultJson.StatusMsg = "取消关注成功"
		}
	} else {
		resultJson.StatusCode = 500
		resultJson.StatusMsg = "服务异常"
	}
	return &resultJson, err
}
