package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreatePublishActionViedeInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreatePublishActionViedeInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreatePublishActionViedeInfoLogic {
	return &CreatePublishActionViedeInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 视频上传接口创建视频信息
func (l *CreatePublishActionViedeInfoLogic) CreatePublishActionViedeInfo(in *mysqlmanageserver.CreatePublishActionViedeInfoRequest) (*mysqlmanageserver.CreatePublishActionViedeInfoResponse, error) {
	err := svc.DB.Table("video_info").Create(&in.VideoInfo).Error
	if err != nil {
		logx.Errorf("[pkg]mysqlconnect [func]CreatePublishActionViedeInfo [msg]gorm [err]%v", err)
		return &mysqlmanageserver.CreatePublishActionViedeInfoResponse{}, err
	}
	return &mysqlmanageserver.CreatePublishActionViedeInfoResponse{}, nil
}
