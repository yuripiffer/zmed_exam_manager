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
	dynamoRepository := dynamo.NewRepository(awsConfig, config.ENV.DynamoExamManagerTableName)
	s3Provider := s3.NewRepository(awsConfig)
	fmt.Println(dynamoRepository, s3Provider)

	//r := mux.NewRouter()
	//web.ConfigureExamsRoutes()
}
