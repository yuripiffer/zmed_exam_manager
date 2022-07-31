package app_response

import (
	"encoding/json"
	"fmt"
	"net/http"
	"zmed_exam_manager/app_errors"
)

type successResponse struct {
	Data interface{} `json:"data"`
}

type ErrorResponse struct {
	Title    string      `json:"title"`
	Status   int         `json:"status"`
	Detail   string      `json:"detail"`
	Err      string      `json:"err"`
	ErrorKey string      `json:"errorKey"`
	Data     interface{} `json:"data,omitempty"`
}

func ERROR(w http.ResponseWriter, statusCode int, err app_errors.AppError) {
	errorResponse := ErrorResponse{
		Title:    err.GetTitle(),
		Detail:   err.GetDetail(),
		ErrorKey: err.GetKey(),
		Err:      err.GetErr(),
		Status:   statusCode,
		Data:     err.GetData(),
	}

	w.WriteHeader(statusCode)

	build(w, statusCode, errorResponse)
}

func JSON(w http.ResponseWriter, statusCode int, data interface{}) {
	if data != nil {
		response := successResponse{
			Data: data,
		}
		build(w, statusCode, response)
		return
	}

	build(w, statusCode, nil)
}

func build(w http.ResponseWriter, statusCode int, data interface{}) {
	w.WriteHeader(statusCode)

	if data != nil {
		err := json.NewEncoder(w).Encode(data)
		if err != nil {
			fmt.Fprintf(w, "%s", err.Error())
		}
	}
}
