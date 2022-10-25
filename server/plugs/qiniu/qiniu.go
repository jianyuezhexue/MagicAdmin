package qiniu

import (
	"context"
	"io"

	"github.com/qiniu/api.v7/v7/auth/qbox"
	"github.com/qiniu/api.v7/v7/storage"
)

const (
	accessKey = "gxBccwgZgmSEeeKwZsx0BlrgLAds2JYD4vyf34PS"
	secretKey = "lVfG8etrs1iRk8IaTwO9m8g1g6cFG0uEJjiPb0K9"
	bucket    = "wjtest1"
)

func config() (cfg storage.Config) {

	cfg = storage.Config{}
	// 设置机房
	cfg.Zone = &storage.ZoneHuabei
	// 是否使用https
	cfg.UseHTTPS = false
	// 是否使用CDN上传加速
	cfg.UseCdnDomains = false

	return
}

func Upload(file io.Reader, filename string, size int64) (string, error) {

	putPolicy := storage.PutPolicy{
		Scope: bucket,
	}

	mac := qbox.NewMac(accessKey, secretKey)
	upToken := putPolicy.UploadToken(mac)

	cig := config()
	formUploader := storage.NewFormUploader(&cig)

	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}
	err := formUploader.Put(context.Background(), &ret, upToken, filename, file, size, &putExtra)

	if err != nil {
		return "", err
	}

	return ret.Key, nil
}
