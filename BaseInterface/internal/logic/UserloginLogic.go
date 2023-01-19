package logic

import (
	"context"
	"fmt"

	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/sql"
	tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserloginLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserloginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserloginLogic {
	return &UserloginLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserloginLogic) Userlogin(req *types.UserloginHandlerRequest) (resp *types.UserloginHandlerResponse, err error) {
	fmt.Println(req)
	uid := sql.CheckUser(req.UserName, req.PassWord)
	fmt.Println(uid)
	if uid == -1 {
		return &types.UserloginHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "用户名或密码错误，请重试",
			UserID:     -1,
			Token:      "",
		}, err
	}
	TokenString := tools.CreateToken(uid)
	return &types.UserloginHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserID:     int64(uid),
		Token:      TokenString,
	}, err
}
