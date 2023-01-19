package minio

import(
	"github.com/minio/minio-go/v6"
	"log"
	"fmt"
)

func MinioConnect(){
	Endpoint := "http://39.106.72.165:9000"
	AccessKeyID := "minio"
	SecretAccessKey := "minio123"
	UseSSL := false


	fmt.Println("123")
	minioClient, err := minio.New(Endpoint,AccessKeyID,SecretAccessKey,UseSSL)
	if err != nil {
		log.Fatalln(err)
	}

    log.Printf("%#v\n", minioClient) // minioClient初使化成功
}