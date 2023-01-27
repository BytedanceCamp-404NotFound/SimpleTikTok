package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mongodb"
	tools "SimpleTikTok/tools/token"
	"context"

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
	resp.StatusCode = 400
	var comments []types.Comment
	token := req.Token
	flag, _, err := tools.CheckToke(token)
	if !flag {
		logx.Errorf("parse token failed, err:%v", err)
		resp.StatusMsg = "parse token failed"
		return resp, err
	}
	videoId := req.VideoId
	collection := mongodb.MongoDBCollection
	filter := bson.D{{
		Key:   "video_id",
		Value: videoId,
	}}
	cur, err := collection.Find(context.Background(), filter)
	if err != nil {
		logx.Errorf("find comments failed, err:%v", err)
		resp.StatusMsg = "find comments failed"
		return resp, err
	}
	for cur.Next(context.Background()) {
		var comment types.Comment
		err = cur.Decode(&comment)
		if err != nil {
			logx.Errorf("decode comment failed, err:%v", err)
			resp.StatusMsg = "decode comment failed"
			return resp, err
		}
		comments = append(comments, comment)

	}
	err = cur.Err()
	if err != nil {
		logx.Errorf("cur has an error, err:%v", err)
		resp.StatusMsg = "cur has an error"
		return resp, err
	}
	cur.Close(context.Background())
	resp.StatusCode = 200
	resp.StatusMsg = "get commentList success"
	resp.CommentList = comments
	return resp, nil
}
