package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"zmed_exam_manager/domain/exams"
	"zmed_exam_manager/infrastructure/config"
	"zmed_exam_manager/infrastructure/dynamo"
	"zmed_exam_manager/infrastructure/patient_provider"
	"zmed_exam_manager/infrastructure/s3"
	"zmed_exam_manager/web"
)

func main() {
	config.PopulateEnv()

	awsConfig := config.InitAws()
	ExamsRepository := dynamo.NewRepository(awsConfig, config.ENV.DynamoExamManagerTableName)
	examResultProvider := s3.NewRepository(awsConfig)
	patientProvider := patient_provider.NewProvider()

	examsUseCase := exams.New(patientProvider, ExamsRepository)
	fmt.Println(examResultProvider)

	r := mux.NewRouter()
	web.ConfigureExamsRoutes(examsUseCase, r)
}
