package v1

import (
	"errors"
	"net/http"
	"zmed_exam_manager/app_errors"
	"zmed_exam_manager/app_response"
	"zmed_exam_manager/interface_input"
	"zmed_exam_manager/util"
)

type ExamsV1Handler struct {
	UseCase interface_input.ExamsUseCase
}

func (h *ExamsV1Handler) NewExam(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestDTO := interface_input.RegisterRequestDTO{}
	_, appError := util.UnmarshalDto(w, r, &requestDTO)
	if appError != nil {
		return
	}

	if requestDTO.Document == "" {
		app_response.ERROR(w, http.StatusBadRequest, app_errors.NewInputError("request field not found",
			errors.New("document")))
		return
	}
	if requestDTO.ExamType == nil {
		app_response.ERROR(w, http.StatusBadRequest, app_errors.NewInputError("request field not found",
			errors.New("exam_type")))
		return
	}
	response, appError := h.UseCase.RegisterExam(ctx, requestDTO)
	if appError != nil {
		app_response.ERROR(w, http.StatusInternalServerError, appError) //TODO
	}
	app_response.JSON(w, 200, response)

}

func (h *ExamsV1Handler) FindExamsByPatientId(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestDTO := interface_input.FindRequestDTO{}
	_, appError := util.UnmarshalDto(w, r, &requestDTO)
	if appError != nil {
		return
	}

	if requestDTO.Document == "" {
		app_response.ERROR(w, http.StatusBadRequest, app_errors.NewInputError("request field not found",
			errors.New("document")))
		return
	}
	response, appError := h.UseCase.FindExams(ctx, requestDTO)
	if appError != nil {
		app_response.ERROR(w, http.StatusInternalServerError, appError) //TODO
	}
	app_response.JSON(w, 200, response)
}

func (h *ExamsV1Handler) StartExam(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	requestDTO := interface_input.StartRequestDTO{}
	_, appError := util.UnmarshalDto(w, r, &requestDTO)
	if appError != nil {
		return
	}

	if requestDTO.ExamId == nil || requestDTO.ExamType == nil || requestDTO.Document == nil {
		app_response.ERROR(w, http.StatusBadRequest, app_errors.NewInputError("request field not found",
			errors.New("needs exam_id, patient_id and exam_type")))
	}

	response, appError := h.UseCase.StartExam(ctx, requestDTO)
	if appError != nil {
		app_response.ERROR(w, http.StatusInternalServerError, appError) //TODO
	}
	app_response.JSON(w, 200, response)
}

func (h *ExamsV1Handler) FinishAndCommunicateExam(w http.ResponseWriter, r *http.Request) {
	_ = h.UseCase.CommunicatePatient()
}
