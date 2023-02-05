package logic

import (
	"context"
	"time"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/tools/encryption"

	"github.com/zeromicro/go-zero/core/logx"
)

type UserRigsterLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}
type User_login struct {
	UserID       int64     `gorm:"cloumn:user_id;primaryKey"`
	UserName     string    `gorm:"cloumn:user_id;"`
	UserPwd      string    `gorm:"cloumn:user_id;"`
	RegisterDate time.Time `gorm:"cloumn:register_date;"`
}
type User struct {
	UserID        int64  `gorm:"cloumn:user_id;primaryKey"`
	UserNickName  string `gorm:"cloumn:user_nick_name"`
	FollowCount   int64  `gorm:"cloumn:follow_count"`
	FollowerCount int64  `gorm:"cloumn:follower_count"`
}

func NewUserRigsterLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UserRigsterLogic {
	return &UserRigsterLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 用户注册
func (l *UserRigsterLogic) UserRigster(in *mysqlmanageserver.UserRegisterRequest) (*mysqlmanageserver.UserRegisterResponse, error) {
	// res, err := FindUserIsExist(l.svcCtx.GormDB, in.Username, in.Password)
	var count int64
	var uid int
	// db, _ := SqlConnect()
	err := l.svcCtx.GormDB.Table("user_login").Select("user_id").Where("user_name = ?", in.Username).Find(&uid).Count(&count).Error

	if err != nil {
		logx.Error("Rigster rpc search user_login error: %v", err)
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: -2,
		}, err
	}
	if count != 0 {
		logx.Error("Rigster rpc User already exist")
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: -3,
		}, nil
	}

	// uid, err := CreateUser(l.svcCtx.GormDB, in.Username, in.Password)
	user := User_login{UserName: in.Username, UserPwd: encryption.HashEncode(in.Password), RegisterDate: time.Now()}
	err = l.svcCtx.GormDB.Table("user_login").Create(&user).Error
	if err != nil {
		logx.Errorf("Rigster rpc Create user_login error: %v", err)
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: -1,
		}, err
	}
	l.svcCtx.GormDB.Table("user_login").Select("user_id").Where("user_name = ? and user_pwd = ?", in.Username, encryption.HashEncode(in.Password)).Find(&uid)

	// err = CreateInfo(db, UserName, int64(id))
	info := User{UserID: int64(uid), UserNickName: in.Username, FollowCount: 0, FollowerCount: 0}
	err = l.svcCtx.GormDB.Table("user_info").Create(&info).Error
	if err != nil {
		logx.Errorf("Rigster rpc Create user_info error: %v", err)
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: -1,
		}, err
	}

	return &mysqlmanageserver.UserRegisterResponse{
		UserId: int64(uid),
	}, nil
}
