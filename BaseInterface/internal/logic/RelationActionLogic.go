package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"context"

	"github.com/zeromicro/go-zero/core/logx"
)

type RelationActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRelationActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RelationActionLogic {
	return &RelationActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// type FeedHandlerRequest struct {
// 	LastestTime int64  `json:"LastestTime"`
// 	Token       string `json:"Token"`
// }

// 时间戳做筛选，时间戳过久，只返回多久之前的视频
// 通过token判断给不同登录状态的用户不同的推送
func (l *RelationActionLogic) RelationAction(req *types.RelationActionHandlerRequest) (resp *types.RelationActionHandlerResponse, err error) {
	// todo: add your logic here and delete this line
	if req.StatusCode > 10000 {
		// fmt.Append("")


	}
	req.StatusCode = 1
	// if(req.Token == nil)
	return
}
