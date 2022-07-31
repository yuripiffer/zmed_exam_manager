package exams

import (
	"context"
	"errors"
	"time"
	"zmed_exam_manager/app_errors"
	"zmed_exam_manager/interface_input"
	"zmed_exam_manager/model"
	"zmed_exam_manager/util"
)

func (s *service) StartExam(ctx context.Context, dto interface_input.StartRequestDTO) (string, app_errors.AppError) {
	patient, appError := s.patientProvider.GetPatient(dto.Document)
	if appError != nil {
		return "", appError
	}

	if patient.Id == "" || patient.Status != model.StatusActive {
		return "", app_errors.NewPatientError("Patient not eligible", errors.New("id or status error"))
	}

	exam, appError := s.examsProvider.FindById(dto.ExamId)
	if appError != nil {
		return "", appError
	}

	if exam.PatientId != patient.Id {
		return "", app_errors.NewPatientError("Patient not eligible", errors.New("id or status error"))
	}

	data := model.Exam{
		Id:        exam.Id,
		Status:    "Started",
		UpdatedAt: time.Now().String(),
	}
	exam, appError = s.examsProvider.Persist(ctx, &data)
	if appError != nil {
		return "", appError
	}

	token, appError := util.GenerateExamToken(exam.PatientId, exam.Id, exam.ExamType)
	return token, nil
}
