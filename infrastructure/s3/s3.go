package s3

import (
	"github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/aws/aws-sdk-go-v2/service/s3"
	"zmed_exam_manager/infrastructure/config"
)

type provider struct {
	client       *s3.Client
	bucket       string
	completedKey string
	processedKey string
	stuckKey     string
	deniedKey    string
}

func NewProvider(awsConfig aws.Config) *provider {
	return &provider{
		client:       s3.NewFromConfig(awsConfig),
		bucket:       config.ENV.S3Bucket,
		completedKey: config.ENV.S3CompletedKey,
		processedKey: config.ENV.S3ProcessedKey,
		stuckKey:     config.ENV.S3StuckKey,
		deniedKey:    config.ENV.S3DeniedKey,
	}
}
