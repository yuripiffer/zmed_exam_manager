package app_errors

import "fmt"

func NewInternalServerError(detail string, err error) AppError {
	return &ErrorStruct{
		Title:    "Internal Server Error",
		Detail:   fmt.Sprintf("%s", detail),
		Err:      err.Error(),
		ErrorKey: KeyInternalServerError,
	}
}

func NewInputError(detail string, err error) AppError {
	return &ErrorStruct{
		Title:    "Input Error",
		Detail:   fmt.Sprintf("%s", detail),
		Err:      err.Error(),
		ErrorKey: keyInputError,
	}
}
