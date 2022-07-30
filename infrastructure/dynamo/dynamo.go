package dynamo

import (
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
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
