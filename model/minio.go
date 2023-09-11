package model

import (
	"context"
	"dy/config"
	"log"

	"github.com/minio/minio-go/v7"
)

var MC *minio.Client
var MCctx context.Context

// MinioInit() minio初始化，如果初始化不成功，后续工作都无法进行
// 所以如果MinioInit()返回值不为nil应当调用panic中断程序
func MinioInit() error {
	endpoint, accessKey, secretAccessKey, useSSL := config.GetMinioConfig()
	ctx := context.Background()
	// Initialize minio client object.
	mc, err := createClient(endpoint, accessKey, secretAccessKey, "", useSSL)
	if err != nil {
		return err
	}
	MC = mc
	MCctx = ctx
	log.Printf("minio client start: %#v\n", mc) // minioClient is now set up

	err = initBucket(mc, ctx, "videos", "shenzhen")
	if err != nil {
		return err
	}
	return nil
}
