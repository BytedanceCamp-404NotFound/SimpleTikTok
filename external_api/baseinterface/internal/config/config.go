package config

import (
	"github.com/zeromicro/go-zero/rest"
	"github.com/zeromicro/go-zero/zrpc"
)

type Config struct {
	rest.RestConf
	// SecretKey string
	// Auth struct {
	// 	AccessSecret string
	// 	AccessExpire int64
	// }
	MinioManageRpc zrpc.RpcClientConf
	MySQLManageRpc zrpc.RpcClientConf
	UtilServerRpc  zrpc.RpcClientConf
}
