package logic

import (
	"context"

	"go-zero-demo/BaseInterface/internal/svc"
	"go-zero-demo/BaseInterface/internal/types"

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
	// todo: add your logic here and delete this line

	return
}
