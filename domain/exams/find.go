package exams

import (
	"context"
	"errors"
	"zmed_exam_manager/interface_input"
	"zmed_exam_manager/pkg/app_errors"
	"zmed_exam_manager/pkg/model/zmed_model"
)

func (s *service) FindExams(ctx context.Context, dto interface_input.FindRequestDTO) ([]*zmed_model.Exam, app_errors.AppError) {
	patient, appError := s.patientProvider.GetPatient(dto.Document)
	if appError != nil {
		return nil, appError
	}

	if patient.Id == "" || patient.Status != zmed_model.StatusActive {
		return nil, app_errors.NewPatientError("Patient not eligible", errors.New("id or status error"))
	}

	return s.examsProvider.FindExamsByPatientId(ctx, patient.Id)
}
