package BaseInterface

import (
	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/oprations/commonerror"
	"SimpleTikTok/oprations/minioconnect"
	tools "SimpleTikTok/tools/token"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type PublishListLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishListLogic {
	return &PublishListLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *PublishListLogic) PublishList(req *types.PublishListHandlerRequest) (resp *types.PublishListHandlerResponse, err error) {

	if req.Token == "" { //APP没有登录账号就请求了该API，为了不让前端报错，直接返回空数据
		return &types.PublishListHandlerResponse{
			StatusCode: 0,
			VideoList:  []types.Video{},
		}, nil
	}

	ok, id, err := tools.CheckToke(req.Token)
	if !ok {
		logx.Infof("[pkg]logic [func]PublishList [msg]req.Token is wrong ")
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_PARSE_TOKEN_ERROR),
			StatusMsg:  "登录过期，请重新登陆",
			VideoList:  []types.Video{},
		}, nil
	}
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishListr [msg]func CheckToken [err]%v", err)
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "Token校验出错",
			VideoList:  []types.Video{},
		}, nil
	}

	user, err := l.svcCtx.MySQLManageRpc.CheckUserInf(l.ctx, &mysqlmanageserver.CheckUserInfRequest{UserId: req.UserID, FollowerId: int64(id)})
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishList [msg]rpc CheckUserInf %v", err)
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
			StatusMsg:  "获取用户信息失败",
			VideoList:  []types.Video{},
		}, nil
	}

	videoNumResponse, err := l.svcCtx.MySQLManageRpc.VideoNum(l.ctx, &mysqlmanageserver.VideoNumRequest{AuthorId: req.UserID})
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishList [msg]rpc VideoNum [err]%v", err)
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_DB_ERROR),
			StatusMsg:  "操作失败",
			VideoList:  []types.Video{},
		}, nil
	}

	getVideoListResponse, err := l.svcCtx.MySQLManageRpc.GetVideoList(l.ctx, &mysqlmanageserver.GetVideoListRequest{
		AuthorId: req.UserID,
		UserId:   int64(id),
	})
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishList [msg]rpc GetVideoList [err]%v", err)
		return &types.PublishListHandlerResponse{
			StatusCode: int32(commonerror.CommonErr_DB_ERROR),
			StatusMsg:  "操作失败",
			VideoList:  []types.Video{},
		}, nil
	}

	videolist := make([]types.Video, videoNumResponse.Num)
	for i, vi := range getVideoListResponse.VideoInfo {
		checkIsFollow, err := l.svcCtx.MySQLManageRpc.CheckIsFollow(l.ctx, &mysqlmanageserver.CheckIsFollowRequest{UserId: req.UserID, FollowerId: int64(id)})
		if err != nil {
			logx.Errorf("[pkg]logic [func]PublishList [msg]rpc checkIsFollow [err]%v", err)
			return &types.PublishListHandlerResponse{
				StatusCode: int32(commonerror.CommonErr_DB_ERROR),
				StatusMsg:  "操作失败",
				VideoList:  []types.Video{},
			}, nil
		}
		realPlayUrl, _ := minioconnect.GetPlayUrl(vi.PlayUrl)
		realCoverUrl, _ := minioconnect.GetPlayUrl(vi.CoverUrl)
		videolist[i] = types.Video{
			Id: vi.VideoId,
			Author: types.User{
				UserId:        user.User.Users.UserId,
				Name:          user.User.Users.UserNickName,
				FollowCount:   user.User.Users.FollowCount,
				FollowerCount: user.User.Users.FollowerCount,
				IsFollow:      checkIsFollow.Ok,
			},
			PlayUrl:       realPlayUrl,
			CoverUrl:      realCoverUrl,
			FavoriteCount: vi.FavoriteCount,
			CommentCount:  vi.CommentCount,
			IsFavotite:    vi.IsFavotite,
			VideoTitle:    vi.VideoTitle,
		}
	}

	return &types.PublishListHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "查询发布列表成功",
		VideoList:  videolist,
	}, err
}
