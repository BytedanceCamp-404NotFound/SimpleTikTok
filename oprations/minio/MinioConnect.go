package minio

import (
	"encoding/base64"
	"path/filepath"

	"github.com/google/uuid"
	"github.com/minio/minio-go/v6"
	"github.com/sirupsen/logrus"
	"github.com/zeromicro/go-zero/core/logx"
)

func MinioConnect() (*minio.Client, error) {
	Endpoint := "39.106.72.165:9001" // 用的API端口，不是Console
	AccessKeyID := "minio"
	SecretAccessKey := "minio123"
	UseSSL := false

	minioClient, err := minio.New(Endpoint, AccessKeyID, SecretAccessKey, UseSSL)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	logx.Infof("%v", minioClient) // minioClient初使化成功
	return minioClient, nil
}

func MinioMakeBucket() error {
	// 创建一个叫mymusic的存储桶。
	// bucketName := "test-minio"
	// location := "yzx_bucket" // 存储桶被创建的region

	// err := minioClient.MakeBucket(bucketName, location)
	// if err != nil {
	// 	// 检查存储桶是否已经存在。
	// 	exists, err := minioClient.BucketExists(bucketName)
	// 	if err == nil && exists {
	// 		logx.Debug("We already own %s\n", bucketName)
	// 	} else {
	// 		logx.Error(err)
	// 		return err
	// 	}
	// }
	// logx.Debug("Successfully created %s\n", bucketName)
	return nil
}
func MinioFileUploader(minioClient *minio.Client, bucketName string, objectPre string, filePath string) (string, error) {
	newfilePath := filepath.Base(filePath)
	newfilePath = uuid.New().String() + "-" + newfilePath
	logx.Infof("%s", newfilePath)

	objectName := objectPre + newfilePath // 要上传的文件的名字
	logx.Infof("MinioFileUploader, objectName:%s, newfilePath:%s", objectName, newfilePath)
	contentType := "video/mp4" // 类型

	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logx.Error(err)
		return "", err
	}
	logx.Infof("Successfully uploaded %s of size %d\n", objectName, n)
	str := bucketName + "_" + objectName
	strbytes := []byte(str)
	bucket_filepath := "minio_" + base64.StdEncoding.EncodeToString(strbytes)
	return bucket_filepath, nil
}
