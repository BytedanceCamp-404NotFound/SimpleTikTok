package logic

import (
	"context"
	"fmt"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"

	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
)

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

//yzx
func (l *PublishActionLogic) PublishAction(req *types.PublishActionHandlerRequest) (resp *types.PublishActionHandlerResponse, err error) {
	// todo: add your logic here and delete this line

	logrus.Debug("debug msg")
	fmt.Print("-------------------------")

	return &types.PublishActionHandlerResponse{
		StatusCode: 0,
		StatusMsg: "35425",
	},err
}
