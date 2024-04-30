package handlers

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/go-playground/validator/v10"
)

var Validate *validator.Validate

// NewValidator create new validator
func NewValidator() *validator.Validate {

	Validate = validator.New(validator.WithRequiredStructEnabled())
	Validate.RegisterTagNameFunc(func(field reflect.StructField) string {
		name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]

		if name == "-" {
			return ""
		}

		return name
	})

	return Validate
}

type ApiError struct {
	Type    string  `json:"type"`
	Message string  `json:"message"`
	Code    int     `json:"code"`
	TraceID *string `json:"trace_id"`
	Field   *string `json:"field"`
}

func msgForTag(fe validator.FieldError) string {
	switch fe.Tag() {
	case "required":
		return "The " + fe.Field() + " field is required"
	case "email":
		return "The " + fe.Field() + " field must be a valid email address"
	case "min":
		return "The " + fe.Field() + " field must be at least " + fe.Param() + " characters"
	case "max":
		return "The " + fe.Field() + " field must be at most " + fe.Param() + " characters"
	case "eqfield":
		return "The " + fe.Field() + " field must be equal to the " + fe.Param() + " field"
	case "gte":
		return "The " + fe.Field() + " field must be greater than or equal to " + fe.Param()
	case "lte":
		return "The " + fe.Field() + " field must be less than or equal to " + fe.Param()
	case "unique":
		return "The " + fe.Field() + " field must be unique"
	default:
		return "The " + fe.Field() + " field is invalid"
	}

}

func ErrorValidationHandler(errs validator.ValidationErrors) []ApiError {
	output := make([]ApiError, len(errs))

	for index, err := range errs {
		field := err.Field()

		splitNamespace := strings.Split(err.Namespace(), ".")
		if len(splitNamespace) > 0 && len(splitNamespace) > 2 {
			fmt.Println("Split Namespace: ", splitNamespace)
			// field only containts the second until the last element of splitNamespace make it with slice and joins
			field = strings.Join(splitNamespace[1:], ".")
			// field = splitNamespace[]
		}

		output[index] = ApiError{
			Type:    "validation_error",
			Message: msgForTag(err),
			Code:    400,
			TraceID: nil,
			Field:   &field,
		}
	}

	return output

}

func ErrorValidationHandlerMap(errs validator.ValidationErrors) map[string]ApiError {
	output := make(map[string]ApiError, len(errs))
	fmt.Printf("%+v\n", output)

	for _, err := range errs {
		field := err.Field()

		splitNamespace := strings.Split(err.Namespace(), ".")
		if len(splitNamespace) > 0 && len(splitNamespace) > 2 {
			// field only containts the second until the last element of splitNamespace
			field = strings.Join(splitNamespace[1:], ".")

		}

		output[field] = ApiError{
			Type:    "validation_error",
			Message: msgForTag(err),
			Code:    400,
			TraceID: nil,
		}
	}

	return output

}
