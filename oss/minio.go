package oss

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

const (
	endpoint        string = "106.14.141.61:30090"
	accessKeyID     string = "SxB3zoNzMouanGpCdMhe"
	secretAccessKey string = "3WZEulCl4A0UXJ2rekFNdtjfuG449Wr72jiCKOZB"
	useSSL          bool   = false
)

var (
	Client *minio.Client
	err    error
)

func init() {
	Client, err = minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKeyID, secretAccessKey, ""),
		Secure: useSSL})
	if err != nil {
		log.Fatalln("minio连接错误: ", err)
	}
}

// GetBucketList 查看桶的列表
func GetBucketList() ([]minio.BucketInfo, error) {
	buckets, err := Client.ListBuckets(context.Background())
	if err != nil {
		return nil, errors.New("not list buckets")
	}
	return buckets, nil
}

// UploadFile 文件上传
func UploadFile(ctx context.Context, bucket, object, filepath string) error {
	_, err := Client.FPutObject(ctx, bucket, object, filepath, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

// Download 文件下载
func Download(ctx context.Context, bucket, object, filepath string) error {
	err := Client.FGetObject(ctx, bucket, object, filepath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
