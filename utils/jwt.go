package utils

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"zmed_exam_manager/infrastructure/config"
	app_errors2 "zmed_exam_manager/pkg/app_errors"
)

func GenerateExamToken(patientId string, examId string, examType int) (string, app_errors2.AppError) {
	claims := jws.Claims{"patient_id": patientId, "exam_id": examId, "exam_type": examType}
	token, err := jws.NewJWT(claims, crypto.SigningMethodHS256).Serialize([]byte(config.ENV.JWTKey))
	if err != nil {
		return "", app_errors2.NewInternalServerError("jwt token", err)
	}
	return string(token), nil
}

func GetPatientId(serializedToken string) (string, app_errors2.AppError) {
	token, err := jws.ParseJWT([]byte(serializedToken))
	if err != nil {
		return "", app_errors2.NewInternalServerError("jwt token", err)
	}
	return token.Claims().Get("patient_id").(string), nil
}

func GetExamId(serializedToken string) (string, app_errors2.AppError) {
	token, err := jws.ParseJWT([]byte(serializedToken))
	if err != nil {
		return "", app_errors2.NewInternalServerError("jwt token", err)
	}
	return token.Claims().Get("exam_id").(string), nil
}
