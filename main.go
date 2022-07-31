package main

import (
	"fmt"
	"zmed_exam_manager/infrastructure/config"
	"zmed_exam_manager/infrastructure/dynamo"
	"zmed_exam_manager/infrastructure/patient_provider"
	"zmed_exam_manager/infrastructure/s3"
)

func main() {
	config.PopulateEnv()

	awsConfig := config.InitAws()
	dynamoRepository := dynamo.NewRepository(awsConfig, config.ENV.DynamoExamManagerTableName)
	s3Provider := s3.NewRepository(awsConfig)
	patientProvider := patient_provider.NewProvider()

	fmt.Println(dynamoRepository, s3Provider, patientProvider)

	//r := mux.NewRouter()
	//web.ConfigureExamsRoutes()
}
