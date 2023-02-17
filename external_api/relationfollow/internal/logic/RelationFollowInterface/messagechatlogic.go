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

type MessageChatLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewMessageChatLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MessageChatLogic {
	return &MessageChatLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *MessageChatLogic) MessageChat(req *types.MessageChatHandlerRequest) (resp *types.MessageChatHandlerResponse, err error) {
	ok, uid, err := tools.CheckToke(req.Token)
	if err!=nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]checktoken failed, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg: "checktoken failed",
			MessageList: []types.SingleMessage{},
		}, err
	}
	if !ok {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]请重新登录, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg: "请重新登录",
			MessageList: []types.SingleMessage{},
		}, err
	}

	messageList, err := l.svcCtx.MongoDBMangerRpc.GetMessage(l.ctx, &mongodbmanageserver.MessageChatRequest{
		ToUserId: req.ToUserId,
		FromUserId: int64(uid),
	})
	if err!=nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]get message list failed, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_RPC_RETUEN_ERROR),
			StatusMsg: "get message list failed",
			MessageList: []types.SingleMessage{},
		}, err
	}

	var res []types.SingleMessage
	for _,v:=range messageList.MessageList {
		message := types.SingleMessage{
			Id: v.Id,
			Content: v.Content,
			CreateTime: v.CreateTime,
		}
		res = append(res, message)
	}
	return &types.MessageChatHandlerResponse{
		StatusCode: 0,
		StatusMsg: "get message list success",
		MessageList: res,
	}, nil
}
