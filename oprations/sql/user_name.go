package sql

import (
	"fmt"
	"time"
)

type User_name struct {
	UserID       int64     `gorm:"cloumn:user_id;primaryKey"`
	UserName     string    `gorm:"cloumn:user_id;"`
	UserPwd      string    `gorm:"cloumn:user_id;"`
	RegisterDate time.Time `gorm:"cloumn:register_date;"`
}

/*
 *	函数功能 向user_login表中插入新值，同时更新user表
 *	输入参数 UserName:用户名
 *           password:密码
 *	返回值 id:新用户的uid
 *		若id=-1,则说明数据库添加不成功
 */
func CreateUser(UserName string, password string) int {
	id := -1
	db, _ := SqlConnect()
	user := User_name{UserName: UserName, UserPwd: password, RegisterDate: time.Now()}
	db.Table("user_login").Create(&user)
	db.Table("user_login").Select("user_id").Where("user_name = ? and user_pwd = ?", UserName, password).Find(&id)

	CreateInfo(UserName, int64(id))
	return id
}

/*
 *	函数功能 校验user_name表中的账户密码是否一致
 *	输入参数 UserName:用户名
 *           password:密码
 *	返回值 id:判断插入是否成功
 *          id>0，查询成功
 *          id=-1 表中没有对应的项
 */
func CheckUser(UserName string, password string) int {
	id := -1
	db, _ := SqlConnect()
	db.Table("user_login").Select("user_id").Where("user_name = ? and user_pwd = ?", UserName, password).Find(&id)
	fmt.Println(id)

	return id
}

//检查注册用户是否存在,存在则修改密码,不存在才创建用户
//-1代表发生错误
//0代表用户不存在
//否则代表用户存在,修改密码
func FindUserIsExist(userName string, password string) (int, error) {
	var count int64
	var uid int64
	db, _ := SqlConnect()
	err := db.Table("user_login").Select("user_id").Where("user_name = ?", userName).Find(&uid).Count(&count).Error
	if err!=nil {
		return -1, err
	}
	if count==0 {
		return 0, nil
	}
	err = db.Table("user_login").Where("user_name = ?", userName).Update("user_pwd", password).Error
	if err!=nil {
		return -1, err
	}
	return int(uid), nil
}