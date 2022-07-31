package util

import (
	"github.com/SermoDigital/jose/crypto"
	"github.com/SermoDigital/jose/jws"
	"zmed_exam_manager/app_errors"
	"zmed_exam_manager/infrastructure/config"
)

func GenerateExamToken(patientId string, examId string, examType int) (string, app_errors.AppError) {
	claims := jws.Claims{"patient_id": patientId, "exam_id": examId, "exam_type": examType}
	token, err := jws.NewJWT(claims, crypto.SigningMethodHS256).Serialize([]byte(config.ENV.JWTKey))
	if err != nil {
		return "", app_errors.NewInternalServerError("jwt token", err)
	}
	return string(token), nil
}

func GetPatientId(serializedToken string) (string, app_errors.AppError) {
	token, err := jws.ParseJWT([]byte(serializedToken))
	if err != nil {
		return "", app_errors.NewInternalServerError("jwt token", err)
	}
	return token.Claims().Get("patient_id").(string), nil
}

func GetExamId(serializedToken string) (string, app_errors.AppError) {
	token, err := jws.ParseJWT([]byte(serializedToken))
	if err != nil {
		return "", app_errors.NewInternalServerError("jwt token", err)
	}
	return token.Claims().Get("exam_id").(string), nil
}
