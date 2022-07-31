package config

import (
	goenv "github.com/Netflix/go-env"
	"log"
)

type environment struct {
	DynamoExamManagerTableName string `env:"ZMED_DYNAMO_EXAM_MANAGER_TABLE_NAME"`
	S3Bucket                   string `env:"ZMED_S3_BUCKET"`
	S3CompletedKey             string `env:"ZMED_S3_COMPLETED_KEY"`
	S3ProcessedKey             string `env:"ZMED_S3_PROCESSED_KEY"`
	S3StuckKey                 string `env:"ZMED_S3_STUCK_KEY"`
	S3DeniedKey                string `env:"ZMED_S3_DENIED_KEY"`
	S3DeletedKey               string `env:"ZMED_S3_DELETED_KEY"`
	PatientManagerHost         string `env:"ZMED_PATIENT_MANAGER_HOST"`
	JWTKey                     string `env:"ZMED_JWT_KEY"`

	PatientManagerGetPath string
}

var ENV environment

func PopulateEnv() {
	_, err := goenv.UnmarshalFromEnviron(&ENV)
	if err != nil {
		log.Fatal(err)
	}
	ENV.PatientManagerGetPath = "document/find"
}
