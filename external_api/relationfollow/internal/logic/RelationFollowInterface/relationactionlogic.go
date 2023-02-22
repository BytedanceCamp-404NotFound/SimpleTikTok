package RelationFollowInterface

import (
	"context"
	"fmt"
	"time"

	"SimpleTikTok/external_api/relationfollow/internal/svc"
	"SimpleTikTok/external_api/relationfollow/internal/types"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/redisconnect"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

const (
	KVDBName = "FollowAndFollowerList:" //Redis-Set
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
	ok, id, err := tools.CheckToke(req.Token)
	resultJson := types.RelationActionHandlerResponse{}
	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.RelationActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.RelationActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
		}, nil
	}

	RedisDB, _ := redisconnect.RedisConnect()

	userIDString := fmt.Sprintf("%s%s:%d", KVDBName, "user_id", id)
	followerIDString := fmt.Sprintf("%s%s:%d", KVDBName, "follower_id", req.To_user_id)

	// 由于使用Redis-Set记录，所以不需要在统计关注量，只需要统计set大小就可以了
	if req.Sction_type == 1 { //关注
		// user_id:1 10 11 每个用户的关注列表
		// 一个巨大的集合
		err = RedisDB.SAdd(l.ctx, userIDString, req.To_user_id).Err()
		// follower_id:10 1 每个用户的粉丝列表
		err = RedisDB.SAdd(l.ctx, followerIDString, id).Err()
	} else { // 取关
		// user_id:1 10 11 每个用户的关注列表
		// 从集合中删除元素
		err = RedisDB.SRem(l.ctx, userIDString, req.To_user_id).Err()
		// follower_id:10 1 每个用户的粉丝列表
		err = RedisDB.SRem(l.ctx, followerIDString, id).Err()
	}
	// res1, _ := RedisDB.SMembers(l.ctx, userIDString).Result()
	// res2, _ := RedisDB.SMembers(l.ctx, followerIDString).Result()
	// _, _ = res1, res2

	result, err := l.svcCtx.MySQLManageRpc.RelationAction(l.ctx, &mysqlmanageserver.RelationActionRequest{
		UserID:     int64(id),
		ToUserID:   req.To_user_id,
		ActionType: int32(req.Sction_type),
	})
	if err != nil {
		resultJson.StatusCode = int32(commonerror.CommonErr_TIMEOUT)
		resultJson.StatusMsg = err.Error()
		return &resultJson, err
	}
	if result.Ok {
		resultJson.StatusCode = 0
		resultJson.StatusMsg = "success"
	} else {
		resultJson.StatusCode = 500
		resultJson.StatusMsg = err.Error()
	}
	return &resultJson, err
}
