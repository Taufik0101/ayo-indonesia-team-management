package storage

import (
	"context"
	"fmt"
	config2 "gin-ayo/config"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/credentials"
	"github.com/aws/aws-sdk-go-v2/service/s3"
)

var S3Client *s3.Client

func InitSupabaseS3() {
	endpoint := fmt.Sprintf("%s/storage/v1/s3", config2.GetEnv("SUPABASE_URL", ""))
	accessKey := config2.GetEnv("SUPABASE_ACCESS_KEY", "")
	secretKey := config2.GetEnv("SUPABASE_SECRET_KEY", "")
	region := config2.GetEnv("SUPABASE_REGION", "")

	customResolver := aws.EndpointResolverWithOptionsFunc(func(service, region string, options ...interface{}) (aws.Endpoint, error) {
		return aws.Endpoint{
			URL:               endpoint,
			HostnameImmutable: true,
		}, nil
	})

	cfg, err := config.LoadDefaultConfig(context.TODO(),
		config.WithRegion(region),
		config.WithEndpointResolverWithOptions(customResolver),
		config.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(accessKey, secretKey, "")),
	)
	if err != nil {
		panic(fmt.Sprintf("failed to load AWS config: %v", err))
	}

	S3Client = s3.NewFromConfig(cfg)
}
