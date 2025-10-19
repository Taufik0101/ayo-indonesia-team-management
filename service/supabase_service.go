package service

import (
	"context"
	"fmt"
	config2 "gin-ayo/config"
	storage "gin-ayo/pkg/supabase"
	"github.com/aws/aws-sdk-go-v2/aws"
	"mime/multipart"
	"path/filepath"
	"time"

	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"github.com/google/uuid"
)

type UploadService interface {
	UploadFile(file *multipart.FileHeader, bucket string) (string, error)
}

type uploadService struct{}

func NewUploadService() UploadService {
	return &uploadService{}
}

func (s *uploadService) UploadFile(file *multipart.FileHeader, bucket string) (string, error) {
	src, err := file.Open()
	if err != nil {
		return "", fmt.Errorf("failed to open file: %w", err)
	}
	defer src.Close()

	fileExt := filepath.Ext(file.Filename)
	fileName := fmt.Sprintf("%s%s", uuid.NewString(), fileExt)
	key := fmt.Sprintf("uploads/%d/%s", time.Now().Year(), fileName)

	_, err = storage.S3Client.PutObject(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   src,
		ACL:    types.ObjectCannedACLPublicRead, // make it public (for public bucket)
	})
	if err != nil {
		return "", fmt.Errorf("upload failed: %w", err)
	}

	// Construct public URL
	url := fmt.Sprintf("%s/storage/v1/object/public/%s/%s",
		config2.GetEnv("SUPABASE_URL", ""),
		bucket,
		key,
	)

	return url, nil
}
