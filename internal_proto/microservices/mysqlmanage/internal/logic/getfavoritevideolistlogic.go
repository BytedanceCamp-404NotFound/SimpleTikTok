package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFavoriteVideoListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFavoriteVideoListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFavoriteVideoListLogic {
	return &GetFavoriteVideoListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取喜欢列表
func (l *GetFavoriteVideoListLogic) GetFavoriteVideoList(in *mysqlmanageserver.GetFavoriteVideoListRequest) (*mysqlmanageserver.GetFavoriteVideoListResponse, error) {
	// todo: add your logic here and delete this line

	//var list []VideoInfo
	var lsa FavoriteVideoNumLogic
	var lsb IsFavotiteLogic
	insa := mysqlmanageserver.FavoriteVideoNumRequest{UserID:in.UserID}
	n, _ := lsa.FavoriteVideoNum(&insa)
	if n.N == 0 {
		logx.Infof("[pkg]logic [func]GetVideoList [msg]FavoriteVideoNum is 0")
		return &mysqlmanageserver.GetFavoriteVideoListResponse{}, nil
	}

	var VideoIdList []int
	err := svc.DB.Table("favorite_list").Select("favorite_video_id").Where("favorite_user_id = ?", in.UserID).Find(&VideoIdList).Error //查所有喜欢的视频的id存到VideoIdList里面
	if err != nil {
		logx.Errorf("[pkg]logic [func]GetVideoList [msg]gorm VideoID.Find  [err]%v", err)
		return &mysqlmanageserver.GetFavoriteVideoListResponse{}, err
	}


	var getfavoritevideoResponse mysqlmanageserver.GetFavoriteVideoListResponse	


	for i := 0; i < len(VideoIdList); i++ {
		var vl VideoInfo
		err = svc.DB.Table("video_info").Where("video_id = ?", VideoIdList[i]).Take(&vl).Error
		if err != nil {
			logx.Errorf("[pkg]logic [func]GetVideoList [msg]gorm VideoInfo.Find [err]%v", err)
			return &mysqlmanageserver.GetFavoriteVideoListResponse{}, err
		}
		//vl.IsFavotite, err = lsb.IsFavotite(&insb)
		// 一定得查完视频信息后才能查询是否点赞，不然没有视频ID
		inss := mysqlmanageserver.IsFavotiteRequest{UserId: in.UserID, VideoId: vl.VideoID}
		outss, err := lsb.IsFavotite(&inss)
		vl.IsFavotite = outss.Ok
		if err != nil {
			logx.Errorf("[pkg]logic [func]GetVideoList [msg]func IsFavotite [err]%v", err)
			return &mysqlmanageserver.GetFavoriteVideoListResponse{}, nil
		}

		getfavoritevideoResponse.VideoInfo = append(getfavoritevideoResponse.VideoInfo, &mysqlmanageserver.VideoInfo{
			VideoId:       vl.VideoID,
			VideoTitle:    vl.VideoTitle,
			AuthorId:      vl.AuthorID,
			CoverUrl:      vl.CoverUrl,
			PlayUrl:       vl.PlayUrl,
			FavoriteCount: vl.FavoriteCount,
			CommentCount:  vl.CommentCount,
			IsFavotite: vl.IsFavotite,
		})
	}

	return &getfavoritevideoResponse, nil
}