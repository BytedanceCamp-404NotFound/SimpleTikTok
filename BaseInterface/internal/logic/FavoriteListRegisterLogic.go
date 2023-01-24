package logic

import (
	"context"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type FavoriteListRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewFavoriteListRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *FavoriteListRegisterLogic {
	return &FavoriteListRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *FavoriteListRegisterLogic) FavoriteListRegister(req *types.FavoriteListRegisterHandlerRequest) (resp *types.FavoriteListRegisterHandlerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
