package dynamo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	"zmed_exam_manager/app_errors"
	"zmed_exam_manager/model"
)

type repository struct {
	client *dynamodb.Client
	table  string
}

func NewRepository(awsConfig aws.Config, table string) *repository {
	client := dynamodb.NewFromConfig(awsConfig)
	fmt.Println(client)
	return &repository{
		client: client,
		table:  table,
	}
}

func (r *repository) Persist(ctx context.Context, data *model.Exam) app_errors.AppError {
	dataMap, err := attributevalue.MarshalMap(data)
	if err != nil {
		return app_errors.NewInternalServerError("Error in persist Dynamodb", err)
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item:      dataMap,
	}

	_, err = r.client.PutItem(ctx, params)
	if err != nil {
		return app_errors.NewInternalServerError("Error in persist Dynamodb", err)
	}
	return nil
}

func (r *repository) FindById(id string) (*model.Exam, app_errors.AppError) {
	out, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, app_errors.NewInternalServerError("Error in persist Dynamodb", err)
	}
	var exam *model.Exam
	err = attributevalue.UnmarshalMap(out.Item, exam)
	if err != nil {
		return nil, app_errors.NewInternalServerError("Error in Dynamodb unmarshal", err)
	}
	return exam, nil
}
