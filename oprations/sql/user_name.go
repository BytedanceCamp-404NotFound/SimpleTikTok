package sql

import (
	"fmt"
	"time"
)

type User_name struct {
	UserID       int
	UserName     string
	UserPwd      string
	RegisterDate time.Time
}

/*
 *	函数功能 向user_name表中插入新值
 *	输入参数 UserName:用户名
 *           password:密码
 *	返回值 id:新用户的uid
 *		若id=-1,则说明数据库添加不成功
 */
func CreateUser(UserName string, password string) int {
	id := -1
	db := SqlConnect("root", "Yu20020601*0518")
	user := User_name{UserName: UserName, UserPwd: password, RegisterDate: time.Now()}
	db.Table("user_name").Create(&user)
	db.Table("user_name").Select("user_id").Where("user_name = ? and user_pwd = ?", UserName, password).Find(&id)

	temp, err := db.DB()
	if err != nil {
		panic(err)
	}
	temp.Close()

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
	db := SqlConnect("root", "Yu20020601*0518")
	db.Table("user_name").Select("user_id").Where("user_name = ? and user_pwd = ?", UserName, password).Find(&id)
	fmt.Println(id)

	temp, err := db.DB()
	if err != nil {
		panic(err)
	}
	temp.Close()

	return id
}
