// Code generated by goctl. DO NOT EDIT.
// Source: miniomanage.proto

package miniomanageserverclient

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/miniomanage/types/miniomanageserver"

	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
)

type (
	GetFileDownloaderRequest    = miniomanageserver.GetFileDownloaderRequest
	GetFileDownloaderResponse   = miniomanageserver.GetFileDownloaderResponse
	GetPlayUrlRequest           = miniomanageserver.GetPlayUrlRequest
	GetPlayUrlResponse          = miniomanageserver.GetPlayUrlResponse
	PutFileUploaderByteRequest  = miniomanageserver.PutFileUploaderByteRequest
	PutFileUploaderByteResponse = miniomanageserver.PutFileUploaderByteResponse
	PutFileUploaderRequest      = miniomanageserver.PutFileUploaderRequest
	PutFileUploaderResponse     = miniomanageserver.PutFileUploaderResponse

	MinioManageServer interface {
		// 文件上传
		PutFileUploader(ctx context.Context, in *PutFileUploaderRequest, opts ...grpc.CallOption) (*PutFileUploaderResponse, error)
		// byte形式文件上传
		PutFileUploaderByte(ctx context.Context, in *PutFileUploaderByteRequest, opts ...grpc.CallOption) (*PutFileUploaderByteResponse, error)
		// 文件下载
		GetFileDownloader(ctx context.Context, in *GetFileDownloaderRequest, opts ...grpc.CallOption) (*GetFileDownloaderResponse, error)
		// 获取Minio视频播放的URL
		GetPlayUrl(ctx context.Context, in *GetPlayUrlRequest, opts ...grpc.CallOption) (*GetPlayUrlResponse, error)
	}

	defaultMinioManageServer struct {
		cli zrpc.Client
	}
)

func NewMinioManageServer(cli zrpc.Client) MinioManageServer {
	return &defaultMinioManageServer{
		cli: cli,
	}
}

// 文件上传
func (m *defaultMinioManageServer) PutFileUploader(ctx context.Context, in *PutFileUploaderRequest, opts ...grpc.CallOption) (*PutFileUploaderResponse, error) {
	client := miniomanageserver.NewMinioManageServerClient(m.cli.Conn())
	return client.PutFileUploader(ctx, in, opts...)
}

// byte形式文件上传
func (m *defaultMinioManageServer) PutFileUploaderByte(ctx context.Context, in *PutFileUploaderByteRequest, opts ...grpc.CallOption) (*PutFileUploaderByteResponse, error) {
	client := miniomanageserver.NewMinioManageServerClient(m.cli.Conn())
	return client.PutFileUploaderByte(ctx, in, opts...)
}

// 文件下载
func (m *defaultMinioManageServer) GetFileDownloader(ctx context.Context, in *GetFileDownloaderRequest, opts ...grpc.CallOption) (*GetFileDownloaderResponse, error) {
	client := miniomanageserver.NewMinioManageServerClient(m.cli.Conn())
	return client.GetFileDownloader(ctx, in, opts...)
}

// 获取Minio视频播放的URL
func (m *defaultMinioManageServer) GetPlayUrl(ctx context.Context, in *GetPlayUrlRequest, opts ...grpc.CallOption) (*GetPlayUrlResponse, error) {
	client := miniomanageserver.NewMinioManageServerClient(m.cli.Conn())
	return client.GetPlayUrl(ctx, in, opts...)
}
