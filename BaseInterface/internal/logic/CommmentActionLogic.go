package logic

import (
	"context"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type CommmentActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewCommmentActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CommmentActionLogic {
	return &CommmentActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *CommmentActionLogic) CommmentAction(req *types.CommmentActionHandlerRequest) (resp *types.CommmentActionHandlerResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
