package svc

import (
	"SimpleTikTok/external_api/relationfollow/internal/config"
	"SimpleTikTok/internal_proto/microservices/mongodbmanage/mongodbmanageserverclient"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/mysqlmanageserverclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	MySQLManageRpc mysqlmanageserverclient.MySQLManageServer
	MongoDBMangerRpc mongodbmanageserverclient.MongodbManageServer
}

func NewServiceContext(c config.Config) *ServiceContext {

	return &ServiceContext{
		Config:         c,
		MySQLManageRpc: mysqlmanageserverclient.NewMySQLManageServer(zrpc.MustNewClient(c.MySQLManageRpc)),
		MongoDBMangerRpc: mongodbmanageserverclient.NewMongodbManageServer(zrpc.MustNewClient(c.MongoDBManageRpc)),
	}
}
