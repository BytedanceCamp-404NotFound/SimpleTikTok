package sql

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
 *	函数功能：连接数据库
 *	输入参数 user:数据库账号
 *		     pwd:数据库密码
 *	返回值 *gorm.DB为链接上的数据库
 *  Tips:如果数据库参数不一致可以依照参数表注释修改
 */
func SqlConnect(user string, pwd string) *gorm.DB {
	//配置MySQL连接参数
	username := user         //账号
	password := pwd          //密码
	host := "127.0.0.1"      //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "SimpleTikTok" //数据库名
	timeout := "10s"         //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return db
}
