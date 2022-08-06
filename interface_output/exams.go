package interface_output

import (
	"context"
	"zmed_exam_manager/pkg/app_errors"
	"zmed_exam_manager/pkg/model/zmed_model"
)

type PatientProvider interface {
	GetPatient(document string) (*zmed_model.Patient, app_errors.AppError)
}

type ExamsProvider interface {
	Persist(ctx context.Context, data *zmed_model.Exam) (*zmed_model.Exam, app_errors.AppError)
	FindById(id string) (*zmed_model.Exam, app_errors.AppError)
	FindExamsByPatientId(ctx context.Context, patientId string) ([]*zmed_model.Exam, app_errors.AppError)
}

type RegisterResponseDTO struct {
	Document string `json:"document"`
}

type findResponseDTO struct {
	Document string `json:"document"`
}

type startResponseDTO struct {
	Document string `json:"document"`
	ExamId   string `json:"exam_id"`
}

type communicateResponseDTO struct {
}
