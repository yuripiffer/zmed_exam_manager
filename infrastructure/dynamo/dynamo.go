package dynamo

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
	app_errors2 "zmed_exam_manager/pkg/app_errors"
	"zmed_exam_manager/pkg/model/zmed_model"
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

func (r *repository) Persist(ctx context.Context, data *zmed_model.Exam) (*zmed_model.Exam, app_errors2.AppError) {
	dataMap, err := attributevalue.MarshalMap(data)
	if err != nil {
		return nil, app_errors2.NewInternalServerError("Error in persist Dynamodb", err)
	}

	params := &dynamodb.PutItemInput{
		TableName: aws.String(r.table),
		Item:      dataMap,
	}

	result, err := r.client.PutItem(ctx, params)
	if err != nil {
		return nil, app_errors2.NewInternalServerError("Error in persist Dynamodb", err)
	}

	exam := zmed_model.Exam{}
	err = attributevalue.UnmarshalMap(result.Attributes, &exam)

	return &exam, nil
}

func (r *repository) FindById(id string) (*zmed_model.Exam, app_errors2.AppError) {
	out, err := r.client.GetItem(context.TODO(), &dynamodb.GetItemInput{
		TableName: aws.String(r.table),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})
	if err != nil {
		return nil, app_errors2.NewInternalServerError("Error in persist Dynamodb", err)
	}
	var exam *zmed_model.Exam
	err = attributevalue.UnmarshalMap(out.Item, exam)
	if err != nil {
		return nil, app_errors2.NewInternalServerError("Error in Dynamodb unmarshal", err)
	}
	return exam, nil
}

func (r *repository) FindExamsByPatientId(ctx context.Context, patientId string) ([]*zmed_model.Exam, app_errors2.AppError) {
	keyConditions := map[string]types.Condition{
		"patient_id": {
			ComparisonOperator: "EQ",
			AttributeValueList: []types.AttributeValue{
				&types.AttributeValueMemberS{Value: patientId},
			},
		},
	}

	return r.runQuery(ctx, keyConditions)
}

func (r *repository) runQuery(ctx context.Context, keyConditions map[string]types.Condition) ([]*zmed_model.Exam, app_errors2.AppError) {
	var queryInput = dynamodb.QueryInput{
		TableName:     aws.String(r.table),
		KeyConditions: keyConditions,
	}

	response, err := r.client.Query(ctx, &queryInput)
	if err != nil {
		return nil, app_errors2.NewInternalServerError("Error in Dynamodb", err)
	}
	var queryResult []*zmed_model.Exam

	err = attributevalue.UnmarshalListOfMaps(response.Items, &queryResult)
	if err != nil {
		return nil, app_errors2.NewInternalServerError("Error in Dynamodb", err)
	}
	return queryResult, nil
}
