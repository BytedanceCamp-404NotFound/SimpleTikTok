package logic

import (
	"context"
	"time"

	"SimpleTikTok/internal_proto/microservices/mysqlmanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"
	"SimpleTikTok/tools/encryption"

	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/gorm"
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
	res, err := FindUserIsExist(l.svcCtx.GormDB, in.Username, in.Password)
	if err != nil {
		logx.Error("Rigster rpc error: %v", err)
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: -2,
		}, err
	}
	if res != 0 {
		logx.Error("User already exist")
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: -3,
		}, nil
	}

	uid, err := CreateUser(l.svcCtx.GormDB, in.Username, in.Password)

	if err != nil {
		logx.Error("Rigster rpc error: %v", err)
		return &mysqlmanageserver.UserRegisterResponse{
			UserId: int32(uid),
		}, err
	}

	return &mysqlmanageserver.UserRegisterResponse{
		UserId: int32(uid),
	}, nil
}

// 向user_login表中插入新值，同时更新user表
// 返回id不为-1表示注册成功
// 返回id为-1表示失败
func CreateUser(db *gorm.DB, UserName string, password string) (int, error) {
	id := -1
	// db, _ := SqlConnect()
	user := User_login{UserName: UserName, UserPwd: encryption.HashEncode(password), RegisterDate: time.Now()}
	err := db.Table("user_login").Create(&user).Error
	if err != nil {
		logx.Errorf("Create user_login error: %v", err)
		return id, err
	}
	db.Table("user_login").Select("user_id").Where("user_name = ? and user_pwd = ?", UserName, encryption.HashEncode(password)).Find(&id)

	err = CreateInfo(db, UserName, int64(id))
	if err != nil {
		logx.Errorf("Create user_info error: %v", err)
		return -1, err
	}
	return id, nil
}

// 检查注册用户是否存在,存在则修改密码,不存在才创建用户
// -1代表发生错误
// 0代表用户不存在
// 否则代表用户存在
func FindUserIsExist(db *gorm.DB, userName string, password string) (int, error) {
	var count int64
	var uid int
	// db, _ := SqlConnect()
	err := db.Table("user_login").Select("user_id").Where("user_name = ?", userName).Find(&uid).Count(&count).Error
	if err != nil {
		return -1, err
	}
	if count == 0 {
		return 0, nil
	}
	return uid, nil
}

// 在user表中同步创建注册用户的信息
// 返回err表示创建过程中有无错误
func CreateInfo(db *gorm.DB, UserName string, uid int64) error {
	info := User{UserID: int64(uid), UserNickName: UserName, FollowCount: 0, FollowerCount: 0}
	err := db.Table("user_info").Create(&info).Error

	return err
}
