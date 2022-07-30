package config

import (
	goenv "github.com/Netflix/go-env"
	"log"
)

type environment struct {
	DynamoExamManagerTableName string `env:"DYNAMO_EXAM_MANAGER_TABLE_NAME"`
}

var ENV environment

func PopulateEnv() {
	_, err := goenv.UnmarshalFromEnviron(&ENV)
	if err != nil {
		log.Fatal(err)
	}
}
