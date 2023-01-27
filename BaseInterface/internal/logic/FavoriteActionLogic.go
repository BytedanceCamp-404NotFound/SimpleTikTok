package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/mysqlconnect"
	tools "SimpleTikTok/tools/token"
	"context"
	"fmt"
	"time"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteActionLogic {
	return &FavoriteActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteActionLogic) FavoriteAction(req *types.FavoriteActionHandlerRequest) (resp *types.FavoriteActionHandlerResponse, err error) {
	// todo: add your logic here and delete this line

	if req.Token == "" {
		return &types.FavoriteActionHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "未登录，请登录再点赞",
		}, nil
	}

	ok, userId, err := tools.CheckToke(req.Token)
	if !ok {
		return &types.FavoriteActionHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "登录过期，请重新登陆",
		}, nil
	}

	db := mysqlconnect.GormDB //连接数据库

	//直接让表中的数据自增或者自减，用原生的sql语句操作
	SqlStringAdd := fmt.Sprintf("UPDATE video_info SET favorite_count = favorite_count + 1 WHERE video_id = %d", req.VideoId)
	SqlStringSub := fmt.Sprintf("UPDATE video_info SET favorite_count = favorite_count - 1 WHERE video_id = %d", req.VideoId)

	if req.ActionType == 1 { //此时未点赞
		err1 := db.Create(mysqlconnect.Favorite_list{Favorite_video_id: req.VideoId, Favorite_user_id: userId, Record_time: time.Now()}).Error //插入点赞数据
		err2 := db.Exec(SqlStringAdd).Error                                                                                                    //将video的点赞总数+1
		if err1 != nil || err2 != nil {
			return &types.FavoriteActionHandlerResponse{
				StatusCode: -1,
				StatusMsg:  "点赞失败，请稍后重试",
			}, err

		} else {
			return &types.FavoriteActionHandlerResponse{
				StatusCode: 0,
				StatusMsg:  "点赞成功",
			}, nil
		}

	} else if req.ActionType == 2 {
		err1 := db.Unscoped().Where("Favorite_user_id = ?", userId).Delete(&mysqlconnect.Favorite_list{}).Error //根据用户操作硬删除整条数据
		err2 := db.Exec(SqlStringSub)                                                                           //将video的点赞总数-1
		if err1 != nil || err2 != nil {
			return &types.FavoriteActionHandlerResponse{
				StatusCode: -1,
				StatusMsg:  "取消点赞失败",
			}, err
		} else {
			return &types.FavoriteActionHandlerResponse{
				StatusCode: 0,
				StatusMsg:  "取消点赞成功",
			}, nil
		}
	}

	return &types.FavoriteActionHandlerResponse{
		StatusCode: 500,
		StatusMsg:  "操作状态异常",
	}, nil

}
