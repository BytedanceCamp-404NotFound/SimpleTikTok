package logic

import (
	"context"

	"SimpleTikTok/internal_proto/microservices/miniomanage/internal/svc"
	"SimpleTikTok/internal_proto/microservices/miniomanage/types/miniomanageserver"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetFileDownloaderLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetFileDownloaderLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetFileDownloaderLogic {
	return &GetFileDownloaderLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 文件下载
func (l *GetFileDownloaderLogic) GetFileDownloader(in *miniomanageserver.GetFileDownloaderRequest) (*miniomanageserver.GetFileDownloaderResponse, error) {
	// todo: add your logic here and delete this line

	return &miniomanageserver.GetFileDownloaderResponse{}, nil
}
