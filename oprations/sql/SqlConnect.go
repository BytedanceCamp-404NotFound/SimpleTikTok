package sql

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

/*
 *	函数功能：连接数据库
 *	返回值 *gorm.DB为链接上的数据库
 *  Tips:如果数据库参数不一致可以依照参数表注释修改
 *  Tips:这里做了修改，将原本的配置信息放置在文件sqlConfig.yaml中
 */
func SqlConnect() *gorm.DB {
	//配置MySQL连接参数
	//oprations/sql/sqlConfig.yaml
	viper.SetConfigFile("/home/ss/Desktop/environment/gopath/src/github.com/SimpleTikTok/oprations/sql/sqlConfig.yaml")
	content, err := os.ReadFile("/home/ss/Desktop/environment/gopath/src/github.com/SimpleTikTok/oprations/sql/sqlConfig.yaml")
	if err != nil {
		fmt.Println("ioutil获取配置文件失败！")
		fmt.Println(err)
	}
	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		fmt.Println("viperhuoqu 配置文件失败！")
	}
	config := viper.Sub("database")

	username := config.Get("username") //账号
	password := config.Get("password") //密码
	host := config.Get("host")         //数据库地址，可以是Ip或者域名
	port := config.Get("port")         //数据库端口
	Dbname := config.Get("dbname")     //数据库名
	timeout := config.Get("timeout")   //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	return db
}
