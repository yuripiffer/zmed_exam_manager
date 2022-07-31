package web

import (
	"github.com/gorilla/mux"
	"zmed_exam_manager/interface_input"
	"zmed_exam_manager/web/v1"
)

func ConfigureExamsRoutes(useCase interface_input.ExamsUseCase, r *mux.Router) {
	examsHandler := v1.ExamsV1Handler{UseCase: useCase}

	r.HandleFunc("/exam/new", examsHandler.NewExam).Methods("POST")
	r.HandleFunc("/exams/info", examsHandler.FindExamsByPatientId).Methods("GET")
	r.HandleFunc("/exam/start", examsHandler.StartExam).Methods("POST")
	r.HandleFunc("/exams/communicate", examsHandler.FinishAndCommunicateExam).Methods("GET")
}
