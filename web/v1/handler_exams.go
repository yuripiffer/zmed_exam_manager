package v1

import (
	"net/http"
	"zmed_exam_manager/interface_input"
)

type ExamsV1Handler struct {
	UseCase interface_input.ExamsUseCase
}

func (h *ExamsV1Handler) NewExam(w http.ResponseWriter, r *http.Request) {
	_ = h.UseCase.RegisterExam()
}

func (h *ExamsV1Handler) FindExamsByPatientId(w http.ResponseWriter, r *http.Request) {
	_ = h.UseCase.FindExams()
}

func (h *ExamsV1Handler) StartExam(w http.ResponseWriter, r *http.Request) {
	_ = h.UseCase.StartExam()
}

func (h *ExamsV1Handler) FinishAndCommunicateExam(w http.ResponseWriter, r *http.Request) {
	_ = h.UseCase.CommunicatePatient()
}
