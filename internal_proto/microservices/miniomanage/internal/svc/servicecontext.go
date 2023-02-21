package svc

import (
	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/config"
	"SimpleTikTok/oprations/viperconfigread"

	"github.com/minio/minio-go/v6"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
)

var MinioClient *minio.Client

type ServiceContext struct {
	Config  config.Config
	MinioDB *minio.Client
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		MinioDB: MinioClient,
	}
}

func init() {
	var err error
	MinioClient, err = MinioConnect()
	if err != nil {
		logx.Errorf("get minio connect fail, err:%v", err)
	}
}

func MinioConnect() (*minio.Client, error) {
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
