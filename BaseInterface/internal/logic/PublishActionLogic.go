package logic

import (
	"SimpleTikTok/BaseInterface/internal/svc"
	"SimpleTikTok/BaseInterface/internal/types"
	"SimpleTikTok/oprations/minio"
	"SimpleTikTok/oprations/sql"
	"context"
	"fmt"
	"os"
	"path/filepath"

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
// TODO
func (l *PublishActionLogic) PublishAction(req *types.PublishActionHandlerRequest) (resp *types.PublishActionHandlerResponse, err error) {
	// TODO 还需要处理传输过来的byte,暂时先用本地MP4和png代替
	// TODO 兼容多种视频和图片格式？avi,mp4
	exePath, _ := os.Executable()
	sourceFile := filepath.Dir(filepath.Dir(exePath))
	vidoeFile := fmt.Sprintf("%s/source/video/video_test1.mp4", sourceFile)
	pictureFile := fmt.Sprintf("%s/source/pic/pic_test1.png", sourceFile)
	content, err := os.ReadFile(vidoeFile)
	if err != nil {
		logx.Errorf("ioutil error:%v", err)
	}
	_ = content

	bucketName := "test-minio"
	minioClient, err := minio.MinioConnect()
	if err != nil {
		logx.Infof("minio connect is fail, error:%v\n", err)
		return nil, err
	}
	minioVideoUrl, err := minio.MinioFileUploader(minioClient, bucketName, "vidoeFile/", vidoeFile)
	if err != nil {
		logx.Infof("minio upload to fail, error:%v\n", err)
		return nil, err // TODO: 上传失败暂时中断接口
	}
	logx.Infof("vidoeFile:%s", minioVideoUrl)
	minioPictureUrl, err := minio.MinioFileUploader(minioClient, bucketName, "pictureFile/", pictureFile)
	if err != nil {
		logx.Infof("minio upload to fail, error:%v\n", err)
		return nil, err // TODO: 上传失败暂时中断接口
	}
	logx.Infof("pictureFile:%s", minioPictureUrl)
	// testUserList_ := make([]types.RelationUser, 0)
	// tempUserList := types.RelationUser{}
	// resultJson := types.PublishActionHandlerResponse{StatusCode: 501, StatusMsg: "token失效，请重新登录"}
	// result, TokenToUserID := tools.CheckToke(req.Token)
	// if !result {
	// 	return &resultJson, err
	// }
	// _ = TokenToUserID
	db, _ := sql.SqlConnect()
	// result := db.Table("user_info").Where("user_id = ?", UserID).Find(&u.User)
	// if result.RowsAffected == 0 {
	// 	return u, false
	// }

	videoInfo := &sql.VideoInfo{
		VideID:        0,
		AuthorID:      0,
		PlayUrl:       minioVideoUrl,
		CoverUrl:      minioPictureUrl,
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavotite:    false,
		VideoTitle:    req.Title,
	}

	// gorm创建一条信息
	err = db.Create(&videoInfo).Error
	if err != nil {
		logx.Errorf("PublishAction mysql error:%v", err)
		return nil, err
	}
	return &types.PublishActionHandlerResponse{
		StatusCode: 200,
		StatusMsg:  "up file success",
	}, err
}
