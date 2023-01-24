package logic

import (
	"context"
	"fmt"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mongodb"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
	"go.mongodb.org/mongo-driver/bson"
)

type CommmentListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommmentListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommmentListLogic {
	return &CommmentListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommmentListLogic) CommmentList(req *types.CommmentListHandlerRequest) (resp *types.CommmentListHandlerResponse, err error) {
	// todo: add your logic here and delete this line
	resp = new(types.CommmentListHandlerResponse)
	resp.StatusCode = -1
	var comments []types.Comment
	token := req.Token
	flag, _ := tools.CheckToke(token)
	if flag == false {
		resp.StatusMsg = "token invalid"
		return resp, nil
	}
	videoId := req.VideoId
	mongoUser := "admin"
	mongoPwd := "admin"
	mongoUrl := "192.168.31.132:27017"
	url := fmt.Sprintf("mongodb://%v:%v@%v", mongoUser, mongoPwd, mongoUrl)
	collection, err := mongodb.Connect("tiktok", "comment", url)
	if err != nil {
		resp.StatusMsg = "connect mongodb failed"
		return resp, nil
	}
	filter := bson.D{{
		Key:   "video_id",
		Value: videoId,
	}}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		resp.StatusMsg = "find comments failed"
		return resp, nil
	}
	for cur.Next(context.Background()) {
		var comment types.Comment
		err = cur.Decode(&comment)
		if err != nil {
			resp.StatusMsg = "decode comment failed"
			return resp, nil
		}
		comments = append(comments, comment)

	}
	err = cur.Err()
	if err != nil {
		return nil, err
	}
	cur.Close(context.Background())
	resp.StatusCode = 0
	resp.StatusMsg = "get commentList success"
	resp.CommentList = comments
	return resp, nil
}
