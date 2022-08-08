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
	"zmed_exam_manager/infrastructure/config"
	app_errors2 "zmed_exam_manager/pkg/app_errors"
)

type Repository struct {
	client       *s3.Client
	bucket       string
	completedKey string
	processedKey string
	stuckKey     string
	deniedKey    string
	deletedKey   string
}

func NewRepository(awsConfig aws.Config) *Repository {
	return &Repository{
		client:       s3.NewFromConfig(awsConfig),
		bucket:       config.ENV.S3Bucket,
		completedKey: config.ENV.S3CompletedKey,
		processedKey: config.ENV.S3ProcessedKey,
		stuckKey:     config.ENV.S3StuckKey,
		deniedKey:    config.ENV.S3DeniedKey,
		deletedKey:   config.ENV.S3DeletedKey,
	}
}

func (r *Repository) PullS3CompletedExams(ctx context.Context) ([]types.Object, app_errors2.AppError) {
	params := &s3.ListObjectsInput{
		Bucket: aws.String(r.bucket),
		Prefix: aws.String(r.completedKey),
	}
	listObjectsOutput, err := r.client.ListObjects(ctx, params)
	if err != nil {
		return nil, app_errors2.NewInternalServerError("Error in S3", err)
	}
	if len(listObjectsOutput.Contents) > 0 {
		return listObjectsOutput.Contents, nil
	}
	return nil, nil
}

func (r *Repository) MoveExamToProcessedFolder(ctx context.Context, objectKey *string) app_errors2.AppError {
	srcKey := "/" + r.bucket + "/" + *objectKey
	destKey := fmt.Sprintf("/%s/%v_%v/%s", r.processedKey, time.Now().Year(), time.Now().Month(), path.Base(*objectKey))
	_, err := r.client.CopyObject(ctx,
		&s3.CopyObjectInput{
			Bucket:     aws.String(r.bucket),
			CopySource: aws.String(srcKey),
			Key:        aws.String(destKey),
		},
	)
	if err != nil {
		return app_errors2.NewInternalServerError("Error in S3", err)
	}
	return nil
}

func (r *Repository) MoveExamToDeletedFolder(ctx context.Context, objectKey *string) app_errors2.AppError {
	srcKey := "/" + r.bucket + "/" + *objectKey
	destKey := fmt.Sprintf("/%s/%v_%v/%s", r.deletedKey, time.Now().Year(), time.Now().Month(), path.Base(*objectKey))
	_, err := r.client.CopyObject(ctx,
		&s3.CopyObjectInput{
			Bucket:     aws.String(r.bucket),
			CopySource: aws.String(srcKey),
			Key:        aws.String(destKey),
		},
	)
	if err != nil {
		return app_errors2.NewInternalServerError("Error in S3", err)
	}
	return nil
}

func (r *Repository) MoveExamToStuckFolder(ctx context.Context, objectKey *string) app_errors2.AppError {
	srcKey := "/" + r.bucket + "/" + *objectKey
	destKey := fmt.Sprintf("/%s/%v_%v/%s", r.stuckKey, time.Now().Year(), time.Now().Month(), path.Base(*objectKey))
	_, err := r.client.CopyObject(ctx,
		&s3.CopyObjectInput{
			Bucket:     aws.String(r.bucket),
			CopySource: aws.String(srcKey),
			Key:        aws.String(destKey),
		},
	)
	if err != nil {
		return app_errors2.NewInternalServerError("Error in S3", err)
	}
	return nil
}
