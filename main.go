package main

import (
	"fmt"
	"zmed_exam_manager/infrastructure/config"
	"zmed_exam_manager/infrastructure/dynamo"
	"zmed_exam_manager/infrastructure/s3"
)

func main() {
	config.PopulateEnv()

	awsConfig := config.InitAws()
	dynamoProvider := dynamo.NewProvider(awsConfig, config.ENV.DynamoExamManagerTableName)
	s3Provider := s3.NewProvider(awsConfig)
	fmt.Println(dynamoProvider, s3Provider)
}
