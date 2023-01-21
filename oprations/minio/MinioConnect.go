package minio

import (
	"log"

	"github.com/minio/minio-go/v6"
	"github.com/sirupsen/logrus"
)

func MinioConnect() (*minio.Client, error) {
	Endpoint := "39.106.72.165:9001"  // 用的API端口，不是Console
	AccessKeyID := "minio"
	SecretAccessKey := "minio123"
	UseSSL := false

	minioClient, err := minio.New(Endpoint, AccessKeyID, SecretAccessKey, UseSSL)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	log.Printf("%#v\n", minioClient) // minioClient初使化成功
	return minioClient, nil
}

func MinioFileUploader() error {
	minioClient, err := MinioConnect()
	// 创建一个叫mymusic的存储桶。
	bucketName := "test-minio"
	location := "yzx_bucket"

	err = minioClient.MakeBucket(bucketName, location)
	if err != nil {
		// 检查存储桶是否已经存在。
		exists, err := minioClient.BucketExists(bucketName)
		if err == nil && exists {
			logrus.Debug("We already own %s\n", bucketName)
		} else {
			logrus.Error(err)
			return err
		}
	}
	logrus.Debug("Successfully created %s\n", bucketName)

	objectName := "yzx_file1.zip"	// 要上传的文件的名字
	filePath := "./yzx_file1.zip"	// 本地文件的路径
	contentType := "application/zip"	// 类型

	n, err := minioClient.FPutObject(bucketName, objectName, filePath, minio.PutObjectOptions{ContentType: contentType})
	if err != nil {
		logrus.Error(err)
	}
	logrus.Error("Successfully uploaded %s of size %d\n", objectName, n)
	return nil
}
