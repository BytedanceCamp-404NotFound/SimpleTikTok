package RelationFollowInterface

import (
	"context"
	"sort"
	"time"

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
	if err != nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]checktoken failed, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode:  int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:   "checktoken failed",
			MessageList: []types.Message{},
		}, err
	}
	if !ok {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]请重新登录, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode:  int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:   "请重新登录",
			MessageList: []types.Message{},
		}, err
	}
	//查找我发过去的消息
	messageList, err := l.svcCtx.MongoDBMangerRpc.GetMessage(l.ctx, &mongodbmanageserver.MessageChatRequest{
		ToUserId:   req.ToUserId,
		FromUserId: int64(uid),
	})
	if err != nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]get message list failed, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode:  int32(commonerror.CommonErr_RPC_RETUEN_ERROR),
			StatusMsg:   "get message list failed",
			MessageList: []types.Message{},
		}, err
	}

	var res []types.Message
	for _, v := range messageList.MessageList {
		t, _ := time.Parse("2006-01-02 15:04:05", v.CreateTime)
		if t.Unix() > req.PreMsgTime {
			message := types.Message{
				Id:         v.Id,
				Content:    v.Content,
				CreateTime: t.Unix(),
				ToUserId:   v.ToUserId,
				FromUserId: v.FromUserId,
			}
			res = append(res, message)
		}
	}

	//查找对方发过来的消息
	messageList, err = l.svcCtx.MongoDBMangerRpc.GetMessage(l.ctx, &mongodbmanageserver.MessageChatRequest{
		ToUserId:   int64(uid),
		FromUserId: req.ToUserId,
	})
	if err != nil {
		logx.Errorf("[pkg]RelationFollowInterface [func]MessageChat [msg]get message list failed, [err]%v", err)
		return &types.MessageChatHandlerResponse{
			StatusCode:  int32(commonerror.CommonErr_RPC_RETUEN_ERROR),
			StatusMsg:   "get message list failed",
			MessageList: []types.Message{},
		}, err
	}

	for _, v := range messageList.MessageList {
		t, _ := time.Parse("2006-01-02 15:04:05", v.CreateTime)
		if t.Unix() > req.PreMsgTime {
			message := types.Message{
				Id:         v.Id,
				Content:    v.Content,
				CreateTime: t.Unix(),
				ToUserId:   v.ToUserId,
				FromUserId: v.FromUserId,
			}
			res = append(res, message)
		}
	}
	sort.Slice(res, func(i, j int) bool {
		if res[i].CreateTime < res[j].CreateTime {
			return true
		}
		return false
	})

	return &types.MessageChatHandlerResponse{
		StatusCode:  0,
		StatusMsg:   "get message list success",
		MessageList: res,
	}, nil
}
