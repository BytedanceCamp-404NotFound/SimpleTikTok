package logic

import (
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/types/mongodbmanageserver"
	"SimpleTikTok/oprations/mongodb"
	"SimpleTikTok/oprations/mysqlconnect"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type userForMongo struct {
	UserId        int64  `gorm:"column:user_id"        json:"id"               form:"user_id"        bson:"user_id"`
	Name          string `gorm:"column:user_nick_name" json:"name"             form:"name"           bson:"name"`
	FollowCount   int64  `gorm:"column:follow_count"   json:"follow_count"     form:"follow_count"   bson:"follow_count"`
	FollowerCount int64  `gorm:"column:follower_count" json:"follower_count"   form:"follower_count" bson:"follower_count"`
	IsFollow      bool   `json:"is_follow"             form:"is_follow"        bson:"is_follow"`
}

type commentForMongo struct {
	Id         int64        `json:"id"          form:"id"         bson:"_id"`
	VideoId    int64        `json:"video_id"    form:"video_id"    bson:"video_id"` //视频id
	User       userForMongo `json:"user"        form:"user"        bson:"user"`
	Content    string       `json:"content"     form:"content"     bson:"content"`
	CreateDate string       `json:"create_date" form:"create_date" bson:"create_date"`
}

type MakeCommentLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewMakeCommentLogic(ctx context.Context, svcCtx *svc.ServiceContext) *MakeCommentLogic {
	return &MakeCommentLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *MakeCommentLogic) MakeComment(in *mongodbmanageserver.CommentActionRequest) (*mongodbmanageserver.CommentActionResponse, error) {
	// todo: add your logic here and delete this line
	//get collection from mongodb
	collection := mongodb.MongoDatabase.Collection("comment")
	mongodb.InitAutoIncrement(collection)
	actionType := in.ActionType
	videoId := in.VideoId
	if actionType == 2 {
		//delete comment
		commentId := in.CommentId
		filter := bson.D{{
			Key:   "_id",
			Value: commentId,
		},
		{
			Key:   "video_id",
			Value: videoId,
		}}
		_, err := collection.DeleteOne(context.Background(), filter)
		if err != nil {
			logx.Errorf("[pkg]logic [func]CommentAction [msg]delete comment failed, [err]%v", err)
			return &mongodbmanageserver.CommentActionResponse{
				Comment: &mongodbmanageserver.Comment{},
			}, err
		}
		err = mysqlconnect.UpdateComment(videoId, actionType)
		if err != nil {
			logx.Errorf("[pkg]logic [func]CommentAction [msg]delete comment failed, [err]%v", err)
			return &mongodbmanageserver.CommentActionResponse{
				Comment: &mongodbmanageserver.Comment{},
			}, err
		}
		return &mongodbmanageserver.CommentActionResponse{
			Comment: &mongodbmanageserver.Comment{},
		}, nil
	} else {
		//insert comment
		comUser, err := mysqlconnect.CommentGetUserByUserId(int(in.UserId))
		if err != nil {
			return &mongodbmanageserver.CommentActionResponse{
				Comment: &mongodbmanageserver.Comment{},
			}, err
		}
		user := userForMongo{
			UserId:        comUser.UserId,
			Name:          comUser.Name,
			FollowCount:   comUser.FollowCount,
			FollowerCount: comUser.FollowerCount,
			IsFollow:      comUser.IsFollow,
		}
		content := in.CommentText
		date := time.Now()
		createDate := fmt.Sprintf("%d-%v", date.Month(), date.Day())
		id, err := mongodb.GetId(collection)
		if err != nil {
			return &mongodbmanageserver.CommentActionResponse{
				Comment: &mongodbmanageserver.Comment{},
			}, err
		}
		comment := commentForMongo{
			Id:         id,
			VideoId:    videoId,
			User:       user,
			Content:    *content,
			CreateDate: createDate,
		}
		_, err = collection.InsertOne(context.Background(), comment)
		if err != nil {
			return &mongodbmanageserver.CommentActionResponse{
				Comment: &mongodbmanageserver.Comment{},
			}, err
		}
		err = mysqlconnect.UpdateComment(videoId, actionType)
		if err != nil {
			logx.Errorf("[pkg]logic [func]CommentAction [msg]delete comment failed, [err]%v", err)
			return &mongodbmanageserver.CommentActionResponse{
				Comment: &mongodbmanageserver.Comment{},
			}, err
		}
		return &mongodbmanageserver.CommentActionResponse{
			Comment: &mongodbmanageserver.Comment{
				Id: id,
				User: &mongodbmanageserver.User{
					UserId:        comUser.UserId,
					UserNickName:  comUser.Name,
					FollowCount:   comUser.FollowCount,
					FollowerCount: comUser.FollowerCount,
					IsFollow:      comUser.IsFollow,
				},
				Content:    *in.CommentText,
				CreateDate: createDate,
			},
		}, nil
	}
}
