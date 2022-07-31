package exams

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"time"
	"zmed_exam_manager/app_errors"
	"zmed_exam_manager/interface_input"
	"zmed_exam_manager/model"
)

func (s *service) RegisterExam(ctx context.Context, dto interface_input.RegisterRequestDTO) (
	*model.Exam, app_errors.AppError) {

	patient, appError := s.patientProvider.GetPatient(dto.Document)
	if appError != nil {
		return nil, appError
	}

	if patient.Id == "" || patient.Status != model.StatusActive {
		return nil, app_errors.NewPatientError("Patient not eligible", errors.New("id or status error"))
	}

	data := model.Exam{
		Id:        uuid.New().String(),
		PatientId: patient.Id,
		Status:    "Register",
		ExamType:  *dto.ExamType,
		CreatedAt: time.Now().String(),
		UpdatedAt: time.Now().String(),
		IsRevoked: false,
	}
	exam, appError := s.examsProvider.Persist(ctx, &data)
	if appError != nil {
		return nil, appError
	}
	return exam, nil
}
