package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"zmed_exam_manager/domain/exams"
	"zmed_exam_manager/domain/script"
	"zmed_exam_manager/infrastructure/config"
	"zmed_exam_manager/infrastructure/dynamo"
	"zmed_exam_manager/infrastructure/patient_provider"
	"zmed_exam_manager/infrastructure/s3"
	"zmed_exam_manager/web"
)

func main() {
	config.PopulateEnv()

	awsConfig := config.InitAws()
	examsRepository := dynamo.NewRepository(awsConfig, config.ENV.DynamoExamManagerTableName)
	examResultProvider := s3.NewRepository(awsConfig)
	patientProvider := patient_provider.NewProvider()

	examsUseCase := exams.New(patientProvider, examsRepository)

	go script.HandleExamsResultProcessing(examResultProvider, examsRepository)

	r := mux.NewRouter()
	web.ConfigureExamsRoutes(examsUseCase, r)
	err := http.ListenAndServe(":85", r)
	if err != nil {
		panic(err)
	}
}
