package model

import (
	"errors"
	"github.com/minio/minio-go/v7"
	"io"
	"math/rand"
	"sync"
)

type MinioDao struct {
}

var minioDao *MinioDao
var minioDaoOnce sync.Once

func NewMinioDao() *MinioDao {
	minioDaoOnce.Do(func() {
		minioDao = &MinioDao{}
	})
	return minioDao
}

func (m *MinioDao) UploadVideo(filename string, reader io.Reader, size int64) (uploadInfo *minio.UploadInfo, err error, objName string) {
	randomPre := generateRandomFilePre()
	count := 0
	for {
		// 使用 StatObject 检查对象是否存在
		_, err := MC.StatObject(MCctx, "videos", randomPre+filename, minio.StatObjectOptions{})
		if err != nil {
			break
		} else {
			randomPre = generateRandomFilePre()
		}

		count++
		if count > 20 {
			// 连续生成20次前缀被占用的可能性几乎为0，暂不处理。
			return &minio.UploadInfo{}, errors.New("生成文件名失效"), ""
		}
	}

	info, err := MC.PutObject(MCctx, "videos", randomPre+filename, reader, size, minio.PutObjectOptions{
		//ContentType: "application/octet-stream", // 这个好像是默认值
	})
	return &info, err, randomPre + filename
}

func generateRandomFilePre() string {
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	result := make([]byte, 16)
	for i := range result {
		result[i] = charset[rand.Intn(len(charset))]
	}
	return string(result) + "/"
}

func (m *MinioDao) GetSignedURL(file_url string) (string, error) {
	url, err := getPresignedObjUrl(MC, MCctx, "videos", file_url, 3600)
	// 以后可以在出错时设置一个默认Url
	return url, err
}
