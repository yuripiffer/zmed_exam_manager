package interface_output

import (
	"context"
	"zmed_exam_manager/app_errors"
	"zmed_exam_manager/model"
)

type PatientProvider interface {
	GetPatient(document string) (*model.Patient, app_errors.AppError)
}

type ExamsProvider interface {
	Persist(ctx context.Context, data *model.Exam) (*model.Exam, app_errors.AppError)
	FindById(id string) (*model.Exam, app_errors.AppError)
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
