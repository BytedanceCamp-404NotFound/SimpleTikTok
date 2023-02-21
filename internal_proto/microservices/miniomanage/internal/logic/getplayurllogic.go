package logic

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"

	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/miniomanage/types/miniomanageserver"
	"SimpleTikTok/oprations/viperconfigread"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetPlayUrlLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetPlayUrlLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetPlayUrlLogic {
	return &GetPlayUrlLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

type MinioKeyVal struct {
	SourceType string
	Bucket     string
	Key        string
}

const (
	sourcetype = "minio"
	separator  = "_"
)

// 获取Minio视频播放的URL
func (l *GetPlayUrlLogic) GetPlayUrl(in *miniomanageserver.GetPlayUrlRequest) (*miniomanageserver.GetPlayUrlResponse, error) {
	if in.PlayUrl == "" {
		logx.Infof("[pkg]BaseInterface [func]GetPlayUrl [msg]playUrl is nil")
		return &miniomanageserver.GetPlayUrlResponse{}, nil
	}
	decodeKey, err := DecodeFileKey(in.PlayUrl)
	if err != nil {
		logx.Errorf("[pkg]BaseInterface [func]GetPlayUrl [msg]decode base64 error:%v", err)
		return &miniomanageserver.GetPlayUrlResponse{}, err
	}

	ConfigReadToMinio, err := viperconfigread.ConfigReadToMinio()
	if err != nil {
		logx.Errorf("[pkg]BaseInterface [func]GetPlayUrl [msg]SqlConnect error:%v", err)
		return &miniomanageserver.GetPlayUrlResponse{}, err
	}
	minioUrl := "http://" + ConfigReadToMinio.Endpoint
	resPlayUrl := fmt.Sprintf("%s/%s/%s", minioUrl, decodeKey.Bucket, decodeKey.Key)
	return &miniomanageserver.GetPlayUrlResponse{
		ResPlayUrl: resPlayUrl,
	}, nil
}

func DecodeFileKey(key string) (*MinioKeyVal, error) {
	keyval := &MinioKeyVal{}
	if !strings.Contains(key, separator) {
		return nil, errors.New("invalid filekey fail")
	}
	keyparts := strings.Split(key, separator)
	if len(keyparts) != 2 {
		return nil, errors.New("cant Split")
	}
	keyval.SourceType = keyparts[0]
	keyData, err := base64.StdEncoding.DecodeString(keyparts[1])
	if err != nil {
		logx.Errorf("decode base64 error:", err.Error())
		return nil, err
	}

	decodeString := string(keyData)
	index := strings.Index(decodeString, separator)
	if index <= 0 {
		return nil, errors.New("cant find separator")
	}

	keyval.Bucket = decodeString[:index]
	keyval.Key = decodeString[index+len(separator):]
	return keyval, nil
}
