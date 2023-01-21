package sql

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/spf13/viper"
	"github.com/zeromicro/go-zero/core/logx"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/*
 *	函数功能：连接数据库
 *	返回值 *gorm.DB为链接上的数据库
 *  Tips:如果数据库参数不一致可以依照参数表注释修改
 *  Tips:这里做了修改，将原本的配置信息放置在文件sqlConfig.yaml中
 */
func SqlConnect() (*gorm.DB, error) {
	// 获取当前可执行文件位置
	exePath, err := os.Executable()
	if err != nil {
		logx.Errorf("%v %v", exePath, err)
		return nil, err
	}
	//返回上级目录
	yamlFile := filepath.Dir(exePath)
	yamlFile = filepath.Dir(yamlFile)
	outputDir := fmt.Sprintf("%s/oprations/sql/sqlConfig.yaml", yamlFile)
	//配置MySQL连接参数
	viper.SetConfigFile(outputDir)
	content, err := ioutil.ReadFile(outputDir)
	if err != nil {
		logx.Errorf("SqlConnect ioutil获取配置文件失败！")
	}

	err = viper.ReadConfig(strings.NewReader(os.ExpandEnv(string(content))))
	if err != nil {
		logx.Errorf("viperhuoqu 配置文件失败！")
	}
	config := viper.Sub("database")

	username := config.Get("username") //账号
	password := config.Get("password") //密码
	host := config.Get("host")         //数据库地址，可以是Ip或者域名
	port := config.Get("port")         //数据库端口
	Dbname := config.Get("dbname")     //数据库名
	timeout := config.Get("timeout")   //连接超时，10秒

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)

	db, err := gormInit(dsn)
	if err != nil {
		logx.Errorf("gorm init fail, error:%v",err.Error())
		return nil, err
	}
	err = gormTableInit(db)
	if err != nil {
		logx.Errorf("init tables fail, error:%v",err.Error())
		return nil, err
	}	

	return db, nil
}

// gorm初始化
func gormInit(dsn string) (*gorm.DB, error) {
	// 日志的配置
	logLevel := logger.Warn
	if true {
		logx.Info("gorm with debug mode")
		logLevel = logger.Info
	}
	logger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logLevel,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)
	// 配置gorm
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger:                 logger,
		SkipDefaultTransaction: true, // 跳过默认开启事务模式
		PrepareStmt:            true, // 使用单数表名，启用该选项后，`User` 表将是`user`
		AllowGlobalUpdate:      true, // 在没有任何条件的情况下执行批量删除，GORM 不会执行该操作
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// 通过gorm 创建数据库表
func gormTableInit(db *gorm.DB) error {
	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;").AutoMigrate(&VideoInfo{},&UserInfo{}); err != nil {
		logx.Error("opendb fialed", err)
		return err
	}
	return nil
}

