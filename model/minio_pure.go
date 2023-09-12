package model

// 注意，这个文件只放置纯函数（非严格定义，允许console输出提示）
// 即函数内所有的操作只与参数有关，无需操作函数外部的变量
// 只把计算结果通过返回值的传递出去，而本身不操作外部数据

import (
	"context"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

// createClient 生成 MinioClient
// 如果token不知道填什么可以先填空字符串
func createClient(endpoint, accessKey, secretKey, token string, useSSL bool) (*minio.Client, error) {
	minioClient, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, token),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatalln("Init Minio Client err", err)
		return nil, err
	}
	return minioClient, nil
}

// initBucket 初始化桶
func initBucket(mc *minio.Client, ctx context.Context, bucketName, location string) error {
	err := mc.MakeBucket(ctx, bucketName, minio.MakeBucketOptions{
		Region: location,
	})
	if err != nil {
		// Check to see if we already own this bucket (which happens if you run this twice)
		exists, errBucketExists := mc.BucketExists(ctx, bucketName)
		if errBucketExists == nil && exists {
			log.Printf("We already own %s bucket\n", bucketName)
			return nil
		} else {
			log.Fatalln("init Bucket unknown err:", err)
			return err
		}
	} else {
		log.Printf("Successfully created %s\n", bucketName)
		return nil
	}
}

// uploadFile 上传文件
// contentType MIME类型， 可以用 http.DetectContentType() 获取
func uploadFile(mc *minio.Client, ctx context.Context, bucketName, objectName, filePath, contentType string) (*minio.UploadInfo, error) {
	info, err := mc.FPutObject(ctx, bucketName, objectName, filePath, minio.PutObjectOptions{
		ContentType: contentType,
	})
	if err != nil {
		log.Fatalln("minio upload file error: ", err)
		return nil, err
	}
	return &info, nil
}
