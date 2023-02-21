package logic

import (
	"context"
	"encoding/base64"
	"path/filepath"

	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/miniomanage/types/miniomanageserver"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v6"
	"github.com/zeromicro/go-zero/core/logx"
)

type PutFileUploaderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewPutFileUploaderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PutFileUploaderLogic {
	return &PutFileUploaderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 文件上传
func (l *PutFileUploaderLogic) PutFileUploader(in *miniomanageserver.PutFileUploaderRequest) (*miniomanageserver.PutFileUploaderResponse, error) {
	newfilePath := filepath.Base(in.FilePath)
	newfilePath = uuid.New().String() + "-" + newfilePath
	logx.Infof("%s", newfilePath)

	objectName := in.ObjectPre + newfilePath // 要上传的文件的名字
	logx.Infof("MinioFileUploader, objectName:%s, newfilePath:%s", objectName, newfilePath)
	contentType := ""
	n, err := svc.MinioClient.FPutObject(in.BucketName, objectName, in.FilePath, minio.PutObjectOptions{ContentType: contentType})

	if err != nil {
		logx.Error(err)
		return &miniomanageserver.PutFileUploaderResponse{}, err
	}
	logx.Infof("Successfully uploaded %s of size %d\n", objectName, n)
	str := in.BucketName + "_" + objectName
	strbytes := []byte(str)
	bucket_filepath := "minio_" + base64.StdEncoding.EncodeToString(strbytes)
	return &miniomanageserver.PutFileUploaderResponse{
		BucketFilepath: bucket_filepath,
	}, nil
}
