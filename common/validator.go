package common

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

type Validator interface {
	Validate(obj interface{}) []*ValidationError
}

type ValidationError struct {
	Field   string      `json:"field"`
	Message string      `json:"message"`
	Value   interface{} `json:"value"`
}

func (e *ValidationError) Error() string {
	return fmt.Sprintf("Validation Error: field %q: %v", e.Field, e.Message)
}

type myValidator struct {
	validator *validator.Validate
}

func NewValidator() *myValidator {
	var v myValidator
	v.validator = validator.New()
	v.validator.RegisterTagNameFunc(func(fld reflect.StructField) string {
		name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return &v
}

func (v *myValidator) Validate(obj interface{}) []*ValidationError {
	if err := v.validator.Struct(obj); err != nil {
		return validationErrorsConverter(err)
	}

	return nil
}

func formatValidationError(err validator.FieldError) *ValidationError {
	res := &ValidationError{Field: err.Field(), Value: err.Value()}
	switch err.ActualTag() {
	case "required":
		res.Message = "Must be not empty"
	case "email":
		res.Message = "Invalid email"
	case "gte":
		res.Message = fmt.Sprintf("Must be greater than or equal %v", err.Param())
	case "eqfield":
		res.Message = fmt.Sprintf("Does not match to %v", strings.ToLower(err.Param()))
	case "nefield":
		res.Message = fmt.Sprintf("Must be not match to %v", strings.ToLower(err.Param()))
	case "min":
		if _, ok := res.Value.(int); ok {
			res.Message = fmt.Sprintf("Must be greater than or equal %v", err.Param())
		} else if _, ok := res.Value.(string); ok {
			res.Message = fmt.Sprintf("Min length of string is %v", err.Param())
		} else {
			res.Message = fmt.Sprintf("Min length of array is %v", err.Param())
		}
	case "mongodb":
		res.Message = "Invalid ObjectId"
	}

	return res
}

func validationErrorsConverter(errs error) []*ValidationError {
	var res []*ValidationError
	for _, err := range errs.(validator.ValidationErrors) {
		res = append(res, formatValidationError(err))
	}

	return res
}
