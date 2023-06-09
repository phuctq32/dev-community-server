package common

import (
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"net/http"
)

var ErrNotFound = mongo.ErrNoDocuments

type AppError struct {
	StatusCode      int               `json:"status_code"`
	Key             string            `json:"error_key"`
	Message         string            `json:"message"`
	RootErr         error             `json:"-"`
	Log             string            `json:"-"`
	ValidationError []ValidationError `json:"validation_errors,omitempty"`
}

type AppErrorLog struct {
	Key             string            `json:"error_key"`
	RootErr         error             `json:"root_error"`
	Log             string            `json:"log"`
	ValidationError []ValidationError `json:"validation_errors,omitempty"`
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
	errLog := AppErrorLog{
		Key:     e.Key,
		Log:     e.Log,
		RootErr: e.RootErr,
	}

	errJson, _ := json.MarshalIndent(&errLog, "", "\t")
	log.Println("\n", string(errJson))
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

func NewCustomBadRequestError(msg string) *AppError {
	return NewBadRequestError(msg, errors.New(msg))
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

func NewExistingError(entity string) *AppError {
	return NewErrorResponse(
		http.StatusUnprocessableEntity,
		"EXISTING",
		fmt.Sprintf("%v already exist", entity),
		fmt.Sprintf("%v already exist", entity),
		nil,
	)
}

func NewValidationError(err []ValidationError) *AppError {
	return &AppError{
		StatusCode:      http.StatusUnprocessableEntity,
		Key:             "VALIDATION_ERROR",
		Message:         fmt.Sprintf("%v: %v", err[0].Field, err[0].Message),
		ValidationError: err,
		RootErr:         &err[0],
	}
}
