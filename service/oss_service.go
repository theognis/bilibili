package service

import (
	"bilibili/tool"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"mime/multipart"
)

type OssService struct {
}

//上传头像
func (f *OssService) UploadAvatar(file multipart.File, filename string) error {
	cfg := tool.GetCfg().Oss

	client, err := oss.New(cfg.EndPoint, cfg.AppKey, cfg.AppSecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(cfg.AvatarBucket)
	if err != nil {
		return err
	}

	err = bucket.PutObject(filename, file)
	return err
}

//上传到视频库
func (f *OssService) UploadVideoBucket(file multipart.File, filename string) error {
	cfg := tool.GetCfg().Oss

	client, err := oss.New(cfg.EndPoint, cfg.AppKey, cfg.AppSecret)
	if err != nil {
		return err
	}

	bucket, err := client.Bucket(cfg.VideosBucket)
	if err != nil {
		return err
	}

	err = bucket.PutObject(filename, file)
	return err
}

//上传封面
