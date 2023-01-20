package logic

import (
	"context"

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
	logx.Infof("UserRegister UserName: %v PassWord: %v", req.UserName, req.PassWord == "")
	rsp := &types.UserRegisterHandlerResponse{StatusCode: -1}
	if req.PassWord == "" && req.UserName == "" {
		logx.Error("UserName and PassWord is nil")
		rsp.StatusCode = 400
		rsp.StatusMsg = "UserName and PassWord is null,register error"
		rsp.UserID = -1
		rsp.Token = ""
		return rsp, err
	}

	uid := sql.CreateUser(req.UserName, req.PassWord)
	logx.Infof("%d", uid)
	if uid == -1 {
		return &types.UserRegisterHandlerResponse{
			StatusCode: -1,
			StatusMsg:  "register error",
			UserID:     -1,
			Token:      "",
		}, err
	}
	TokenString := tools.CreateToken(uid)
	return &types.UserRegisterHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "register success",
		UserID:     int64(uid),
		Token:      TokenString,
	}, err
}
