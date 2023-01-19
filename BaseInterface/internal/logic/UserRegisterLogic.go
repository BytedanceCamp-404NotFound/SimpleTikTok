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

type UserRegisterLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUserRegisterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRegisterLogic {
	return &UserRegisterLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UserRegisterLogic) UserRegister(req *types.UserRegisterHandlerRequest) (resp *types.UserRegisterHandlerResponse, err error) {
	fmt.Println(req)
	uid := sql.CreateUser(req.UserName, req.PassWord)
	fmt.Println(uid)
	if uid != -1 {
		return &types.UserRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "注册失败",
			UserID:     -1,
			Token:      "",
		}, err
	}
	TokenString := tools.CreateToken(uid)
	return &types.UserRegisterHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "注册成功",
		UserID:     int64(uid),
		Token:      TokenString,
	}, err
}
