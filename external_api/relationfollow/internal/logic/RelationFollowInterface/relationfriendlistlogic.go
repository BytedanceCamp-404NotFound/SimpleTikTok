package RelationFollowInterface

import (
	"context"

	"SimpleTikTok/external_api/relationfollow/internal/svc"
	"SimpleTikTok/external_api/relationfollow/internal/types"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/commonerror"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationFriendListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationFriendListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationFriendListLogic {
	return &RelationFriendListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RelationFriendListLogic) RelationFriendList(req *types.RelationFriendListHandlerRequest) (resp *types.RelationFriendListHandlerResponse, err error) {
	// todo: add your logic here and delete this line

	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.RelationFriendListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.RelationFriendListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
		}, nil
	}

	resultJson := &types.RelationFriendListHandlerResponse{}
	rflhr, err := l.svcCtx.MySQLManageRpc.RelationFriendList(l.ctx, &mysqlmanageserver.RelationFriendListRequest{
		LoginUserID: int64(id),
		UserID:      req.UserId,
	})
	if err != nil {
		resultJson.StatusCode = int32(commonerror.CommonErr_TIMEOUT)
		resultJson.StatusMsg = err.Error()
		return resultJson, err
	}
	Userlist := make([]types.FriendUser, 0)
	for i := 0; i < len(rflhr.RelationUser); i++ {
		Userlist = append(Userlist, types.FriendUser{
			Id:            rflhr.RelationUser[i].Id,
			Name:          rflhr.RelationUser[i].Name,
			FollowCount:   rflhr.RelationUser[i].FollowCount,
			FollowerCount: rflhr.RelationUser[i].FollowerCount,
			IsFollow:      rflhr.RelationUser[i].IsFollow,
			Avatar:        "http://www.bailinzhe.com/image/2019-11-08/0ab5979f578d4f8bf15b20e6e51f0f2a.jpg",
			Message:       "",
			MsgType:       1,
		})
		resultJson.StatusCode = 0
		resultJson.StatusMsg = "success"
	}
	resultJson.Userlist = Userlist
	return resultJson, err

}
