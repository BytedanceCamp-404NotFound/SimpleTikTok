package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFeedUserInfoLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFeedUserInfoLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFeedUserInfoLogic {
	return &GetFeedUserInfoLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 取流接口获取用户信息
func (l *GetFeedUserInfoLogic) GetFeedUserInfo(in *mysqlmanageserver.GetFeedUserInfoRequest) (*mysqlmanageserver.GetFeedUserInfoResponse, error) {
	// todo: add your logic here and delete this line

	return &mysqlmanageserver.GetFeedUserInfoResponse{}, nil
}
