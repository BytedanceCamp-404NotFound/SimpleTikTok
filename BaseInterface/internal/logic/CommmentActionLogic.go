package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mongodb"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type CommmentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommmentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommmentActionLogic {
	return &CommmentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommmentActionLogic) CommmentAction(req *types.CommmentActionHandlerRequest) (resp *types.CommmentActionHandlerResponse, err error) {
	// todo: add your logic here and delete this line
	//parse token
	resp = new(types.CommmentActionHandlerResponse)
	resp.StatusCode = 400
	flag, userId, err := tools.CheckToke(req.Token)
	if !flag {
		logx.Errorf("parse token failed, err:%v", err)
		resp.StatusMsg = "parse token failed"
		return resp, err
	}
	//get collection from mongodb
	collection := mongodb.MongoDBCollection
	actionType := req.ActionType
	videoId := req.VideoId
	if actionType == 2 {
		//delete comment
		commentId := req.CommentId
		filter := bson.D{{
			Key:   "_id",
			Value: commentId,
		},
			{
				Key:   "video_id",
				Value: videoId,
			}}
		_, err = collection.DeleteOne(context.Background(), filter)
		if err != nil {
			logx.Errorf("delete comment failed, err:%v", err)
			resp.StatusMsg = "delete comment failed"
			return resp, err
		}
		resp.StatusCode = 200
		resp.StatusMsg = "delete success"
	} else {
		//insert comment
		db := mysqlconnect.GormDB

		var user types.User
		err = db.Table("user_info").Where("user_id=?", userId).First(&user).Error
		if err != nil {
			logx.Errorf("search user_info failed, err:%v", err)
			resp.StatusMsg = "search user_info failed"
			return resp, err
		}
		content := req.CommentText
		date := time.Now()
		createDate := fmt.Sprintf("%d-%v", date.Month(), date.Day())
		id, err := mongodb.GetId(collection)
		if err != nil {
			logx.Errorf("get id failed, err:%v", err)
			resp.StatusMsg = "get id failed"
			return resp, err
		}
		comment := types.Comment{
			Id:         id,
			VideoId:    videoId,
			User:       user,
			Content:    content,
			CreateDate: createDate,
		}
		_, err = collection.InsertOne(context.Background(), comment)
		if err != nil {
			logx.Errorf("insert comment failed, err:%v", err)
			resp.StatusMsg = "insert comment failed"
			return resp, err
		}
		resp.StatusCode = 200
		resp.StatusMsg = "insert success"
		resp.Comment = comment
	}
	return
}
