package logic

import (
	"context"
	"fmt"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/tools/encryption"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
)

type UserLoginLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUserLoginLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserLoginLogic {
	return &UserLoginLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户登陆校验
func (l *UserLoginLogic) UserLogin(in *mysqlmanageserver.UserLoginRequest) (*mysqlmanageserver.UserLoginResponse, error) {
	uid, err := CheckUser(l.svcCtx.GormDB, in.Username, in.Password)
	if err != nil {
		logx.Error("Check user rpc error: %v", err)
		return &mysqlmanageserver.UserLoginResponse{
			UserId: -1,
		}, err
	}

	return &mysqlmanageserver.UserLoginResponse{
		UserId: int32(uid),
	}, nil
}

// 函数功能 校验user_name表中的账户密码是否一致
// 返回id不为-1表示一致
// 返回id为-1表示不一致
func CheckUser(db *gorm.DB, UserName string, password string) (int32, error) {
	var id int32
	err := db.Table("user_login").Select("user_id").Where("user_name = ? and user_pwd = ?", UserName, encryption.HashEncode(password)).Find(&id).Error
	if err != nil {
		logx.Errorf("Check user fail, error:%v", err.Error())
		return -1, err
	}
	fmt.Println(id)

	return id, err
}
