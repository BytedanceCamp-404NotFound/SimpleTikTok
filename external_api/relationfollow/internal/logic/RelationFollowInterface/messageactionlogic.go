package RelationFollowInterface

import (
	"context"

	"SimpleTikTok/external_api/relationfollow/internal/svc"
	"SimpleTikTok/external_api/relationfollow/internal/types"
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/types/mongodbmanageserver"
	"SimpleTikTok/oprations/commonerror"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type MessageActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageActionLogic {
	return &MessageActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageActionLogic) MessageAction(req *types.MessageActionHandlerRequest) (resp *types.MessageActionHandlerResponse, err error) {
	ok, uid, err := tools.CheckToke(req.Token)
	if err!=nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageAction [msg]checktoken failed, [err]%v", err)
		return &types.MessageActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg: "checktoken failed",
		}, err
	}
	if !ok {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageAction [msg]请重新登录, [err]%v", err)
		return &types.MessageActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg: "请重新登录",
		}, err
	}

	_, err = l.svcCtx.MongoDBMangerRpc.SendMessage(l.ctx, &mongodbmanageserver.MessageActionRequest{
		ToUserId: req.ToUserId,
		ActionType: req.ActionType,
		Content: req.Content,
		FromUserId: int64(uid),
	})
	if err!=nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageAction [msg]get message list failed, [err]%v", err)
		return &types.MessageActionHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_RPC_RETUEN_ERROR),
			StatusMsg: "send message failed",
		}, err
	}
	return &types.MessageActionHandlerResponse{
		StatusCode: int32(commonerror.CommonErr_STATUS_OK),
		StatusMsg: "send message success",
	}, nil
}
