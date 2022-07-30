package main

import (
	"fmt"
	"zmed_exam_manager/infrastructure/config"
	"zmed_exam_manager/infrastructure/dynamo"
)

func main() {
	config.PopulateEnv()

	awsConfig := config.InitAws()
	dynamoProvider := dynamo.NewProvider(awsConfig, config.ENV.DynamoExamManagerTableName)
	fmt.Println(dynamoProvider)
}
