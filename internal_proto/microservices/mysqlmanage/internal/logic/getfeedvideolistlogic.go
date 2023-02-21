package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedVideoListLogic {
	return &GetFeedVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取流接口获取视频列表
func (l *GetFeedVideoListLogic) GetFeedVideoList(in *mysqlmanageserver.GetFeedVideoListRequest) (*mysqlmanageserver.GetFeedVideoListResponse, error) {
	var tmpFeedVideoList []VideoInfo
	err := svc.DB.Table("video_info").Order("update_time DESC").Limit(30).Scan(&tmpFeedVideoList).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]GetFeedVideoList [msg]gorm [err]%v", err)
		return &mysqlmanageserver.GetFeedVideoListResponse{}, err
	}

	var getFeedVideoListResponse mysqlmanageserver.GetFeedVideoListResponse
	for _, val := range tmpFeedVideoList {
		inss := mysqlmanageserver.IsFavotiteRequest{UserId: in.UserId, VideoId: val.VideoID}
		var lss IsFavotiteLogic
		outss, err := lss.IsFavotite(&inss) // 获取是否点赞
		if err != nil {
			logx.Errorf("[pkg]logic [func]GetFeedVideoList [msg]rpc IsFavotite %v", err)
			return &mysqlmanageserver.GetFeedVideoListResponse{}, nil
		}
		getFeedVideoListResponse.VideoInfo = append(getFeedVideoListResponse.VideoInfo, &mysqlmanageserver.VideoInfo{
			VideoId:       val.VideoID,
			VideoTitle:    val.VideoTitle,
			AuthorId:      val.AuthorID,
			CoverUrl:      val.CoverUrl,
			PlayUrl:       val.PlayUrl,
			FavoriteCount: val.FavoriteCount,
			CommentCount:  val.CommentCount,
			IsFavotite:    outss.Ok,
		})
	}

	return &getFeedVideoListResponse, nil
}
