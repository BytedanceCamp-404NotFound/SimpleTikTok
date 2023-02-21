package logic

import (
	"bytes"
	"context"
	"fmt"
	"os"

	"SimpleTikTok/internal_proto/microservices/utilserver/internal/svc"
	"SimpleTikTok/internal_proto/microservices/utilserver/types/utilserver"

	"github.com/disintegration/imaging"
	ffmpeg "github.com/u2takey/ffmpeg-go"
	"github.com/zeromicro/go-zero/core/logx"
)

type GetSnapshotLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetSnapshotLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetSnapshotLogic {
	return &GetSnapshotLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *GetSnapshotLogic) GetSnapshot(in *utilserver.GetSnapshotRequest) (*utilserver.GetSnapshotResponse, error) {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(in.VideoName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", in.FrameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	img, err := imaging.Decode(buf)

	err = imaging.Save(img, in.SnapshotName)
	if err != nil {
		logx.Errorf("生成缩略图失败：", err)
		return &utilserver.GetSnapshotResponse{}, nil
	}
	return &utilserver.GetSnapshotResponse{}, nil
}
