package oss

import (
	"context"
	"errors"
	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
	"log"
)

// NewMinioClient 实例化链接
func NewMinioClient(endpoint, accessKey, secretKey string, useSSL bool) *minio.Client {
	client, err := minio.New(endpoint, &minio.Options{
		Creds:  credentials.NewStaticV4(accessKey, secretKey, ""),
		Secure: useSSL,
	})
	if err != nil {
		log.Fatal(err)

	}
	return client
}

// CreateBucket 创建桶
func CreateBucket(ctx context.Context, bucket, region string, client *minio.Client, locking bool) error {
	err := client.MakeBucket(ctx, bucket, minio.MakeBucketOptions{region, locking})
	if err != nil {
		return err
	}
	exists, _ := client.BucketExists(ctx, bucket)
	if exists {
		return errors.New("bucket already exists")
	}
	return nil
}

// GetBucketList 查看桶的列表
func GetBucketList(client *minio.Client) ([]minio.BucketInfo, error) {
	buckets, err := client.ListBuckets(context.Background())
	if err != nil {
		return nil, errors.New("not list buckets")
	}
	return buckets, nil
}

// UploadFile 文件上传
func UploadFile(ctx context.Context, bucket, object, filepath string, client *minio.Client) error {
	_, err := client.FPutObject(ctx, bucket, object, filepath, minio.PutObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}

// Download 文件下载
func Download(ctx context.Context, bucket, object, filepath string, client *minio.Client) error {
	err := client.FGetObject(ctx, bucket, object, filepath, minio.GetObjectOptions{})
	if err != nil {
		return err
	}
	return nil
}
