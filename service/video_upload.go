package service

import (
	"dy/model"
	"errors"
	"fmt"
	"io"
	"sync"
)

type VideoUploadFlow struct {
}

var (
	videoUploadFlow     *VideoUploadFlow
	videoUploadFlowOnce sync.Once
)

func NewVideoUploadFlow() *VideoUploadFlow {
	videoUploadFlowOnce.Do(func() {
		videoUploadFlow = &VideoUploadFlow{}
	})
	return videoUploadFlow
}

func (f *VideoUploadFlow) Upload(title, filename string, authorId uint, videoReader io.Reader, size int64) error {
	videoDao := model.NewVideoDao()
	minioDao := model.NewMinioDao()

	info, err, _ := minioDao.UploadVideo(filename, videoReader, size)
	if err != nil {
		fmt.Println("minioDao.UploadVideo err: ", err)
		return errors.New("文件上传失败")
	}

	err = videoDao.CreateNewVideo(title, authorId, info.Key)
	if err != nil {
		fmt.Println("videoDao.UploadVideo err:", err)
		return errors.New("视频存储失败")
	}

	return nil
}
