package storage

import (
	"bytes"
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

type StorageManager struct {
	cli        *s3.Client
	bucketName string
}

func NewStorageManager(cfg aws.Config, bucketName string) *StorageManager {

	return &StorageManager{
		cli:        s3.NewFromConfig(cfg),
		bucketName: bucketName,
	}
}

func (sm *StorageManager) Upload(ctx context.Context, key string, data []byte) error {
	_, err := sm.cli.PutObject(ctx, &s3.PutObjectInput{
		Key:         aws.String(key),
		Bucket:      aws.String(sm.bucketName),
		Body:        bytes.NewReader(data),
		ContentType: aws.String("image/jpg"),
	})
	return err
}
