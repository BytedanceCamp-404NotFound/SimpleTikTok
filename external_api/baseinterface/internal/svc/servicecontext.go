package svc

import (
	"SimpleTikTok/external_api/baseinterface/internal/config"
	"SimpleTikTok/internal_proto/microservices/miniomanage/miniomanageserverclient"
	"SimpleTikTok/internal_proto/microservices/mysqlmanage/mysqlmanageserverclient"
	"SimpleTikTok/internal_proto/microservices/utilserver/utilserverclient"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config         config.Config
	MinioManageRpc miniomanageserverclient.MinioManageServer
	MySQLManageRpc mysqlmanageserverclient.MySQLManageServer
	UtilServerRpc  utilserverclient.Utilserver
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:         c,
		MinioManageRpc: miniomanageserverclient.NewMinioManageServer(zrpc.MustNewClient(c.MinioManageRpc)),
		MySQLManageRpc: mysqlmanageserverclient.NewMySQLManageServer(zrpc.MustNewClient(c.MySQLManageRpc)),
		UtilServerRpc:  utilserverclient.NewUtilserver(zrpc.MustNewClient(c.UtilServerRpc)),
	}
}
