package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mongodbmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/types/mongodbmanageserver"
	"SimpleTikTok/oprations/mongodb"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type GetMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetMessageLogic {
	return &GetMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetMessageLogic) GetMessage(in *mongodbmanageserver.MessageChatRequest) (*mongodbmanageserver.MessageChatResponse, error) {
	// todo: add your logic here and delete this line
	toUserId := in.ToUserId
	fromUserId := in.FromUserId
	collection := mongodb.MongoDatabase.Collection("message")
	var messageList []*mongodbmanageserver.Message

	filter := bson.D{{
		Key:   "to_user_id",
		Value: toUserId,
	}, {
		Key:   "from_user_id",
		Value: fromUserId,
	}}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logx.Errorf("[pkg] logic [func]GetMessage [msg]get message list failed, [err]%v", err)
		return &mongodbmanageserver.MessageChatResponse{
			MessageList: nil,
		}, err
	}

	for cur.Next(context.Background()) {
		var message Message
		err = cur.Decode(&message)
		if err != nil {
			logx.Errorf("[pkg] logic [func]GetMessage [msg]decode single message failed, [err]%v", err)
			return &mongodbmanageserver.MessageChatResponse{
				MessageList: nil,
			}, err
		}
		temp := &mongodbmanageserver.Message{
			Id:         message.Id,
			ToUserId:   message.ToUserId,
			FromUserId: message.FromUserId,
			Content:    message.Content,
			CreateTime: message.CreateTime,
		}
		messageList = append(messageList, temp)
	}

	err = cur.Err()
	if err != nil {
		logx.Errorf("[pkg] logic [func]GetMessage [msg]cur has an error, [err]%v", err)
		return &mongodbmanageserver.MessageChatResponse{
			MessageList: nil,
		}, err
	}
	cur.Close(context.Background())
	return &mongodbmanageserver.MessageChatResponse{
		MessageList: messageList,
	}, nil
}
