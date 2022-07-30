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

type provider struct {
	client *dynamodb.Client
	table  string
}

func NewProvider(awsConfig aws.Config, table string) *provider {
	client := dynamodb.NewFromConfig(awsConfig)
	fmt.Println(client)
	return &provider{
		client: client,
		table:  table,
	}
}

func (p *provider) Persist(ctx context.Context, data *model.Exam) app_errors.AppError {
	dataMap, err := attributevalue.MarshalMap(data)
	if err != nil {
		return app_errors.NewInternalServerError("Error in persist Dynamodb", err)
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String(p.table),
		Item:      dataMap,
	}

	_, err = p.client.PutItem(ctx, params)
	if err != nil {
		return app_errors.NewInternalServerError("Error in persist Dynamodb", err)
	}
	return nil
}

func (p *provider) FindById(id string) (*model.Exam, app_errors.AppError) {
	out, err := p.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(p.table),
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
