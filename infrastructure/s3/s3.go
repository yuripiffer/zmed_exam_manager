package s3

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	_ "github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	_ "github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/aws/aws-sdk-go-v2/service/s3/types"
	"path"
	"time"
	"zmed_exam_manager/app_errors"
	"zmed_exam_manager/infrastructure/config"
)

type provider struct {
	client       *s3.Client
	bucket       string
	completedKey string
	processedKey string
	stuckKey     string
	deniedKey    string
	deletedKey   string
}

func NewProvider(awsConfig aws.Config) *provider {
	return &provider{
		client:       s3.NewFromConfig(awsConfig),
		bucket:       config.ENV.S3Bucket,
		completedKey: config.ENV.S3CompletedKey,
		processedKey: config.ENV.S3ProcessedKey,
		stuckKey:     config.ENV.S3StuckKey,
		deniedKey:    config.ENV.S3DeniedKey,
		deletedKey:   config.ENV.S3DeletedKey,
	}
}

func (p *provider) PullS3CompletedExams(ctx context.Context) ([]types.Object, app_errors.AppError) {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(p.bucket),
		Prefix: aws.String(p.completedKey),
	}
	listObjectsOutput, err := p.client.ListObjects(ctx, params)
	if err != nil {
		return nil, app_errors.NewInternalServerError("Error in S3", err)
	}
	if len(listObjectsOutput.Contents) > 0 {
		return listObjectsOutput.Contents, nil
	}
	return nil, nil
}

func (p *provider) MoveExamToProcessedFolder(ctx context.Context, objectKey *string) app_errors.AppError {
	srcKey := "/" + p.bucket + "/" + *objectKey
	destKey := fmt.Sprintf("/%s/%v_%v/%s", p.processedKey, time.Now().Year(), time.Now().Month(), path.Base(*objectKey))
	_, err := p.client.CopyObject(ctx,
		&s3.CopyObjectInput{
			Bucket:     aws.String(p.bucket),
			CopySource: aws.String(srcKey),
			Key:        aws.String(destKey),
		},
	)
	if err != nil {
		return app_errors.NewInternalServerError("Error in S3", err)
	}
	return nil
}

func (p *provider) MoveExamToDeletedFolder(ctx context.Context, objectKey *string) app_errors.AppError {
	srcKey := "/" + p.bucket + "/" + *objectKey
	destKey := fmt.Sprintf("/%s/%v_%v/%s", p.deletedKey, time.Now().Year(), time.Now().Month(), path.Base(*objectKey))
	_, err := p.client.CopyObject(ctx,
		&s3.CopyObjectInput{
			Bucket:     aws.String(p.bucket),
			CopySource: aws.String(srcKey),
			Key:        aws.String(destKey),
		},
	)
	if err != nil {
		return app_errors.NewInternalServerError("Error in S3", err)
	}
	return nil
}

func (p *provider) MoveExamToStuckFolder(ctx context.Context, objectKey *string) app_errors.AppError {
	srcKey := "/" + p.bucket + "/" + *objectKey
	destKey := fmt.Sprintf("/%s/%v_%v/%s", p.stuckKey, time.Now().Year(), time.Now().Month(), path.Base(*objectKey))
	_, err := p.client.CopyObject(ctx,
		&s3.CopyObjectInput{
			Bucket:     aws.String(p.bucket),
			CopySource: aws.String(srcKey),
			Key:        aws.String(destKey),
		},
	)
	if err != nil {
		return app_errors.NewInternalServerError("Error in S3", err)
	}
	return nil
}
