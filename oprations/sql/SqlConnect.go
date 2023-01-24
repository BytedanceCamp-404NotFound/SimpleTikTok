package sql

import (
	"fmt"
	"log"
	"os"
	"time"
	"SimpleTikTok/oprations/viperconfigread"

	"gorm.io/gorm/schema"
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
	mysqlConfig, err := viperconfigread.ConfigReadToMySQL()
	if err != nil {
		logx.Errorf("SqlConnect error:%v", err)
		return nil, err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s",
		mysqlConfig.UserName, mysqlConfig.PassWord, mysqlConfig.Host,
		mysqlConfig.Port, mysqlConfig.DBname, mysqlConfig.TimeOut)
	db, err := gormInit(dsn)
	if err != nil {
		logx.Errorf("gorm init fail, error:%v", err.Error())
		return nil, err
	}
	// err = gormTableInit(db)
	// if err != nil {
	// 	logx.Errorf("init tables fail, error:%v", err.Error())
	// 	return nil, err
	// }

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
		PrepareStmt:            false, 
		AllowGlobalUpdate:      true, // 在没有任何条件的情况下执行批量删除，GORM 不会执行该操作
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true, //使用单数表名，启用该选项.
		},
	})
	if err != nil {
		return nil, err
	}

	return db, nil
}

// // 通过gorm 创建数据库表
// func gormTableInit(db *gorm.DB) error {
// 	if err := db.Set("gorm:table_options", "ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;").AutoMigrate(&VideoInfo{}, &UserInfo{}); err != nil {
// 		logx.Error("opendb fialed", err)
// 		return err
// 	}
// 	return nil
// }
