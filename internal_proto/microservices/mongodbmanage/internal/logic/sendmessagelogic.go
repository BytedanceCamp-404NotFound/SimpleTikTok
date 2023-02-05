package logic

import (
	"context"
	"errors"
	"time"

	"SimpleTikTok/internal_proto/microservices/mongodbmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/types/mongodbmanageserver"
	"SimpleTikTok/oprations/mongodb"

	"github.com/zeromicro/go-zero/core/logx"
)

type Message struct {
	Id         int64  `json:"id"                    form:"id"                    bson:"_id"`
	ToUserId   int64  `json:"to_user_id"            form:"to_user_id"            bson:"to_user_id"`
	FromUserId int64  `json:"from_user_id"          form:"from_user_id"          bson:"from_user_id"`
	Content    string `json:"content"               form:"content"               bson:"content"`
	CreateTime string `json:"create_time"           form:"create_time"           bson:"create_time"`
}

type SendMessageLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSendMessageLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SendMessageLogic {
	return &SendMessageLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *SendMessageLogic) SendMessage(in *mongodbmanageserver.MessageActionRequest) (*mongodbmanageserver.MessageActionResponse, error) {
	if in.ActionType!=1 {
		logx.Errorf("[pkg]logic [func]SendMessage [msg]wrong actionType, [err]%v", errors.New("ActioType错误"))
		return &mongodbmanageserver.MessageActionResponse{}, nil
	}
	collection := mongodb.MongoDatabase.Collection("message")
	mongodb.InitAutoIncrement(collection)
	id, _ := mongodb.GetId(collection)
	createTime := time.Now().Format("2006-01-02 15:04:05")
	insertMessage := Message{
		Id: id,
		ToUserId: in.ToUserId,
		FromUserId: in.FromUserId,
		Content: in.Content,
		CreateTime: createTime,
	}
	_, err := collection.InsertOne(context.Background(), insertMessage)
	if err!=nil {
		logx.Errorf("[pkg]logic [func]SendMessage [msg]insert message to mongodb error, [err]%v", err)
		return &mongodbmanageserver.MessageActionResponse{}, err
	}
	return &mongodbmanageserver.MessageActionResponse{}, nil
}
