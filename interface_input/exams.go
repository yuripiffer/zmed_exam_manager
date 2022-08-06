package interface_input

import (
	"context"
	"zmed_exam_manager/pkg/app_errors"
	"zmed_exam_manager/pkg/model/zmed_model"
)

type ExamsUseCase interface {
	RegisterExam(ctx context.Context, dto RegisterRequestDTO) (*zmed_model.Exam, app_errors.AppError)
	FindExams(ctx context.Context, dto FindRequestDTO) ([]*zmed_model.Exam, app_errors.AppError)
	StartExam(ctx context.Context, dto StartRequestDTO) (string, app_errors.AppError)
	RevokeExam() app_errors.AppError
	CommunicatePatient() app_errors.AppError
}

type RegisterRequestDTO struct {
	Document string `json:"document"`
	ExamType *int   `json:"exam_type"`
}

type FindRequestDTO struct {
	Document string `json:"document"`
}

type StartRequestDTO struct {
	Document *string `json:"document"`
	ExamId   *string `json:"exam_id"`
	ExamType *int    `json:"exam_type"`
}

type CommunicateRequestDTO struct {
}
