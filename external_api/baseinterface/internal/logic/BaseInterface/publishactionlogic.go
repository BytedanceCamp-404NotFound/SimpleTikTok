package BaseInterface

import (
	"bytes"
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"

	"github.com/disintegration/imaging"
	"github.com/google/uuid"
	ffmpeg "github.com/u2takey/ffmpeg-go"

	"SimpleTikTok/external_api/baseinterface/internal/svc"
	"SimpleTikTok/external_api/baseinterface/internal/types"

	// "SimpleTikTok/oprations/commonerror"
	minio "SimpleTikTok/oprations/minioconnect"
	"SimpleTikTok/oprations/mysqlconnect"
	// tools "SimpleTikTok/tools/token"

	"github.com/zeromicro/go-zero/core/logx"
)

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
func (l *PublishActionLogic) PublishAction(req *types.PublishActionHandlerRequest, r *http.Request) (resp *types.PublishActionHandlerResponse, err error) {
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

	go UploadData(r, userId, req.Title) //任务量太多，开协程传输

	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishAction [msg]CreatePublishActionViedeInfo is err [err]%v", err)
		return nil, err
	}
	return &types.PublishActionHandlerResponse{
		StatusCode: 0,
		StatusMsg:  "上传成功",
	}, err
}

func UploadData(r *http.Request, userId int, title string) {

	DownloadFile(r)

	// videoPath := "/yzx/src/SimpleTikTok/source/video/video_test2.mp4"
	// snapshotPath := "/yzx/src/SimpleTikTok/source/pic/test"
	// frameNum := 1
	// snapshotName, err := GetSnapshot(videoPath, snapshotPath, frameNum)
	// if err != nil {
	// 	logx.Errorf("生成缩略图失败：", err)
	// }
	// _ = snapshotName

	// _, _, err = minioUpDate(r) // Minio 上传文件
	minioVideoUrl, minioPictureUrl, err := minioUpDate(r) // Minio 上传文件
	if err != nil {
		logx.Errorf("[pkg]logic [func]PublishAction [msg]minioUpDate is fail [err]%v", err)
		// return &types.PublishActionHandlerResponse{
		// 	StatusCode: int32(commonerror.CommonErr_PAGE_NOT_EXIT),
		// 	StatusMsg:  "没有收到视频文件或者出现其他错误",
		// }, err
	}

	VideoInfo := &mysqlconnect.PublishActionVideoInfo{
		Video_id:       int32(uuid.New().ID()),
		Author_id:      int64(userId),
		Play_url:       minioVideoUrl,
		Cover_url:      minioPictureUrl,
		Favorite_count: 0,
		Comment_count:  0,
		Video_title:    title,
	}
	// gorm创建一条信息
	err = mysqlconnect.CreatePublishActionViedeInfo(VideoInfo)
}

// TODO: 将保存的文件名转化为传输的名字
// https://zhuanlan.zhihu.com/p/136410759
func DownloadFile(r *http.Request) error {
	fileType := r.PostFormValue("data")
	_ = fileType
	file, _, err := r.FormFile("data")
	if err != nil {
		logx.Errorf("err:%v", err)
		return err
	}
	defer file.Close()
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		logx.Errorf("err:%v", err)
		return err
	}
	filetypecheck := http.DetectContentType(fileBytes)
	_ = filetypecheck

	newPath := "/yzx/src/SimpleTikTok/source/video/tmp/1.mp4"
	newFile, err := os.Create(newPath)
	if err != nil {
		logx.Errorf("err:%v", err)
		return err
	}
	defer newFile.Close()
	_, err = newFile.Write(fileBytes)
	if err != nil {
		logx.Errorf("err:%v", err)
		return err
	}
	return nil
}

func GetSnapshot(videoPath, snapshotPath string, frameNum int) (snapshotName string, err error) {
	buf := bytes.NewBuffer(nil)
	err = ffmpeg.Input(videoPath).
		Filter("select", ffmpeg.Args{fmt.Sprintf("gte(n,%d)", frameNum)}).
		Output("pipe:", ffmpeg.KwArgs{"vframes": 1, "format": "image2", "vcodec": "mjpeg"}).
		WithOutput(buf, os.Stdout).
		Run()
	if err != nil {
		logx.Errorf("生成缩略图失败：", err)
		return "", err
	}

	img, err := imaging.Decode(buf)
	if err != nil {
		logx.Errorf("生成缩略图失败：", err)
		return "", err
	}

	err = imaging.Save(img, snapshotPath+".png")
	if err != nil {
		logx.Errorf("生成缩略图失败：", err)
		return "", err
	}

	names := strings.Split(snapshotPath, "\\")
	snapshotName = names[len(names)-1] + ".png"
	return
}

// 图片和视频上传到minio
func minioUpDate(r *http.Request) (string, string, error) {
	// TOTEMP
	// exePath, _ := os.Executable()
	// sourceFile := filepath.Dir(filepath.Dir(exePath))
	//vidoeFile := fmt.Sprintf("%s/source/video/video_test1.mp4", sourceFile)
	// pictureFile := fmt.Sprintf("%s/source/pic/pic_test2.jpg", sourceFile)
	// content, err := os.ReadFile(vidoeFile)

	bucketName := "testminio"
	minioClient, err := minio.MinioConnect()
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]MinioConnect fail [err]%v", err)
		return "", "", err
	}

	file, FileHeader, err := r.FormFile("data") //可以优化
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]r.FormFile fail [err]%v", err)
		return "", "", err
	}

	//minioVideoUrl, err := minio.MinioFileUploader(minioClient, bucketName, "vidoeFile/", vidoeFile) //上传本地文件
	minioVideoUrl, err := minio.MinioFileUploader_byte(minioClient, bucketName, "vidoeFile/", FileHeader.Filename, file, FileHeader.Size) // 上传字节类型文件
	if err != nil {
		logx.Errorf("[pkg]logic [func]minioUpDate [msg]minio video upload to fail [err]%v", err)
		return "", "", err // TODO: 上传失败暂时中断接口
	}

	minioPictureUrl := ""
	// minioPictureUrl, err := minio.MinioFileUploader(minioClient, bucketName, "pictureFile/", pictureFile)
	// if err != nil {
	// 	logx.Errorf("[pkg]logic [func]minioUpDate [msg]minio picture upload to fail [err]%v", err)
	// 	// return "", "", err // TODO: 上传图片失败不中断接口
	// }
	return minioVideoUrl, minioPictureUrl, err
}
