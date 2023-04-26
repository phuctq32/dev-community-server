package common

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type AppError struct {
	StatusCode int    `json:"status_code"`
	Key        string `json:"error_key"`
	Message    string `json:"message"`
	RootErr    error  `json:"-"`
	Log        string `json:"-"`
}

type AppErrorLog struct {
	StatusCode int    `json:"-"`
	Key        string `json:"error_key"`
	Message    string `json:"-"`
	RootErr    error  `json:"root_error"`
	Log        string `json:"log"`
}

func (err *AppError) Error() string {
	return err.RootError().Error()
}

func (err *AppError) RootError() error {
	if err, ok := err.RootErr.(*AppError); ok {
		return err.RootError()
	}

	return err.RootErr
}

func (e AppError) Logging() {
	e.RootErr = e.RootError()
	errLog := AppErrorLog(e)
	jsonAppErr, err := json.MarshalIndent(errLog, "", "  ")
	if err != nil {
		log.Println("fda")
	}

	log.Println(string(jsonAppErr))
}

func NewErrorResponse(statusCode int, key, msg, log string, rootErr error) *AppError {
	return &AppError{
		StatusCode: statusCode,
		Key:        key,
		Message:    msg,
		Log:        log,
		RootErr:    rootErr,
	}
}

func NewServerError(e error) *AppError {
	return &AppError{
		StatusCode: http.StatusInternalServerError,
		Key:        "INTERNAL_SERVER_ERROR",
		Message:    "Internal server error",
		Log:        e.Error(),
		RootErr:    e,
	}
}

func NewBadRequestError(msg string, e error) *AppError {
	return &AppError{
		StatusCode: http.StatusBadRequest,
		Key:        "BAD_REQUEST",
		Message:    msg,
		Log:        e.Error(),
		RootErr:    e,
	}
}

func NewNoPermissionError(e error) *AppError {
	return NewErrorResponse(
		http.StatusUnauthorized,
		"NO_PERMISSION",
		"You have no permission",
		e.Error(),
		e,
	)
}

func NewNotFoundError(entity string, e error) *AppError {
	return NewErrorResponse(
		http.StatusNotFound,
		"NOT_FOUND",
		fmt.Sprintf("%v not found", entity),
		e.Error(),
		e,
	)
}

func NewExistingError(entity string, e error) *AppError {
	return NewErrorResponse(
		http.StatusUnprocessableEntity,
		"EXISTING",
		fmt.Sprintf("%v already exist", entity),
		e.Error(),
		e,
	)
}
