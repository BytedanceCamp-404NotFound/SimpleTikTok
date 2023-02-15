package BaseInterface

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"

	// "SimpleTikTok/internal_proto/microservices/mysqlmanage/types/mysqlmanageserver"

	// "SimpleTikTok/oprations/commonerror"
	minio "SimpleTikTok/oprations/minioconnect"
	"SimpleTikTok/oprations/mysqlconnect"

	// tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

// var exePath string
// var videoName string
// var snapshotName string

type PublishActionLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewPublishActionLogic(ctx context.Context, svcCtx *svc.ServiceContext) *PublishActionLogic {
	return &PublishActionLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// yzx
// TODO 还需要处理传输过来的byte,暂时先用本地MP4和png代替
// TODO 兼容多种视频和图片格式？avi,mp4
func (l *PublishActionLogic) PublishAction(req *types.PublishActionHandlerRequest, httpReq *http.Request) (resp *types.PublishActionHandlerResponse, err error) {
	// ok, userId, err := tools.CheckToke(req.Token)
	// if err != nil {
	// 	return &types.PublishActionHandlerResponse{
	// 		StatusCode: int32(commonerror.CommonErr_INTERNAL_ERROR),
	// 		StatusMsg:  "Token校验出错",
	// 	}, nil
	// }
	// if !ok {
	// 	logx.Infof("[pkg]logic [func]PublishAction [msg]feedUserInfo.Name is nuil ")
	// 	return &types.PublishActionHandlerResponse{
	// 		StatusCode: int32(commonerror.CommonErr_PARAMETER_FAILED),
	// 		StatusMsg:  "登录过期，请重新登陆",
	// 	}, nil
	// }
	userId := 1

	go l.UploadData(httpReq, userId, req.Title) //任务量太多，开协程传输

	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishAction [msg]CreatePublishActionViedeInfo is err [err]%v", err)
		return nil, err
	}
	return &types.PublishActionHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "上传成功",
	}, err
}

// 上传视频的总协程
func (l *PublishActionLogic) UploadData(httpReq *http.Request, userId int, videoTitle string) {
	exePath, _ := os.Executable()
	exePath = filepath.Dir(exePath)
	videoName := fmt.Sprintf("%s/%s.mp4", exePath, videoTitle)
	snapshotName := fmt.Sprintf("%s/%s.png", exePath, videoTitle)
	frameNum := 1
	_ = DownloadVideo(httpReq, videoName, videoTitle) // 下载视频文件

	_ = GetSnapshot(videoName, snapshotName, frameNum)

	// _, _, err = minioUpDate(httpReq) // Minio 上传文件
	minioVideoUrl, minioPictureUrl, err := minioUpDate(httpReq, snapshotName) // Minio 上传文件
	defer os.Remove(videoName)
	defer os.Remove(snapshotName)
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishAction [msg]minioUpDate is fail [err]%v", err)
	}

	// tmpvideoInfo := &mysqlmanageserver.PublishActionVideoInfo{
	// 	VideoId:       int32(uuid.New().ID()),
	// 	AuthorId:      int64(userId),
	// 	PlayUrl:       minioVideoUrl,
	// 	CoverUrl:      minioPictureUrl,
	// 	FavoriteCount: 0,
	// 	CommentCount:  0,
	// 	VideoTitle:    videoTitle,
	// }
	// err = mysqlconnect.CreatePublishActionViedeInfo(VideoInfo)
	// resp, err := l.svcCtx.MySQLManageRpc.CreatePublishActionViedeInfo(l.ctx, &mysqlmanageserver.CreatePublishActionViedeInfoRequest{
	// 	VideoInfo: tmpvideoInfo,
	// })
	// _ = resp

	// gorm创建一条信息
	VideoInfo := &mysqlconnect.PublishActionVideoInfo{
		Video_id:       int32(uuid.New().ID()),
		Author_id:      int64(userId),
		Play_url:       minioVideoUrl,
		Cover_url:      minioPictureUrl,
		Favorite_count: 0,
		Comment_count:  0,
	}
	// gorm创建一条信息
	err = mysqlconnect.CreatePublishActionViedeInfo(VideoInfo)
}

// TODO: 将保存的文件名转化为传输的名字
// https://zhuanlan.zhihu.com/p/136410759
func DownloadVideo(httpReq *http.Request, videoName string, videoTitle string) error {
	fileType := httpReq.PostFormValue("data")
	_ = fileType
	file, _, err := httpReq.FormFile("data")
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	filetypecheck := http.DetectContentType(fileBytes)
	_ = filetypecheck

	videoFile, err := os.Create(videoName)

	defer videoFile.Close()
	_, err = videoFile.Write(fileBytes)
	if err != nil {
		logx.Errorf("err:%v", err)
		return err
	}
	return nil
}

func GetSnapshot(videoName, snapshotName string, frameNum int) error {
	buf := bytes.NewBuffer(nil)
	err := ffmpeg.Input(videoName).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()

	img, err := imaging.Decode(buf)

	err = imaging.Save(img, snapshotName)
	if err != nil {
		logx.Errorf("生成缩略图失败：", err)
		return err
	}
	return nil
}

// 图片和视频上传到minio
func minioUpDate(httpReq *http.Request, picPath string) (string, string, error) {
	bucketName := "testminio"
	minioClient, err := minio.MinioConnect()
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]MinioConnect fail [err]%v", err)
		return "", "", err
	}

	file, FileHeader, err := httpReq.FormFile("data") //可以优化
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]httpReq.FormFile fail [err]%v", err)
		return "", "", err
	}

	//minioVideoUrl, err := minio.MinioFileUploader(minioClient, bucketName, "vidoeFile/", vidoeFile) //上传本地文件
	minioVideoUrl, err := minio.MinioFileUploader_byte(minioClient, bucketName, "vidoeFile/", FileHeader.Filename, file, FileHeader.Size) // 上传字节类型文件
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]minio video upload to fail [err]%v", err)
		return "", "", err // TODO: 上传失败暂时中断接口
	}

	minioPictureUrl, err := minio.MinioFileUploader(minioClient, bucketName, "pictureFile/", picPath)
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]minio picture upload to fail [err]%v", err)
		// return "", "", err // TODO: 上传图片失败不中断接口
	}
	return minioVideoUrl, minioPictureUrl, err
}
