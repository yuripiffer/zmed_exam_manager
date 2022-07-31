package interface_input

import "zmed_exam_manager/app_errors"

type ExamsUseCase interface {
	RegisterExam() app_errors.AppError
	FindExams() app_errors.AppError
	StartExam() app_errors.AppError
	RevokeExam() app_errors.AppError
	CommunicatePatient() app_errors.AppError
}
