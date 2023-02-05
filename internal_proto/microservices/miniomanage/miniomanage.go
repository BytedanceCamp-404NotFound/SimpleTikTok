package main

import (
	"flag"
	"fmt"

	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/config"
	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/server"
	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/miniomanage/types/miniomanageserver"
	"SimpleTikTok/oprations/viperconfigread"

	"github.com/minio/minio-go/v6"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/miniomanage.yaml", "the config file")
var MinioDB *minio.Client

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		miniomanageserver.RegisterMinioManageServerServer(grpcServer, server.NewMinioManageServerServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}

func init() {
	var err error
	MinioDB, err = minioConnect()
	if err != nil {
		logx.Errorf("get minio connect fail, err:%v", err)
	}
}

func minioConnect() (*minio.Client, error) {
	minioConfig, err := viperconfigread.ConfigReadToMinio()
	if err != nil {
		logx.Errorf("MinioConnect error:%v", err)
		return nil, err
	}
	minioClient, err := minio.New(minioConfig.Endpoint, minioConfig.AccessKeyID,
		minioConfig.SecretAccessKey, minioConfig.UseSSL)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logx.Infof("%v", minioClient) // minioClient初使化成功
	return minioClient, nil
}
