package exams

import (
	"zmed_exam_manager/interface_input"
	"zmed_exam_manager/interface_output"
)

type service struct {
	interface_input.ExamsUseCase
	patientProvider interface_output.PatientProvider
	examsProvider   interface_output.ExamsProvider
}

func New(patientProvider interface_output.PatientProvider, examsProvider interface_output.ExamsProvider) *service {
	return &service{
		patientProvider: patientProvider,
		examsProvider:   examsProvider,
	}
}
