package logic

import (
	"context"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FeedLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFeedLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FeedLogic {
	return &FeedLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FeedLogic) Feed(req *types.FeedHandlerRequest) (resp *types.FeedHandlerResponse, err error) {
	// todo: add your logic here and delete this line
	


	return &types.FeedHandlerResponse{
		StatusCode: 200,
		StatusMsg:  "feed video success",
		NextTime:   123,
	}, nil
}
