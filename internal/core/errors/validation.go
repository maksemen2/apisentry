package coreerrors

import "encoding/json"

type ValidationError struct {
	Field   string
	Message string
	Code    string
}

func (e ValidationError) Error() string {
	return e.Message
}

// NewValidationError creates a new core validation error
func NewValidationError(field, code, message string) ValidationError {
	return ValidationError{
		Field:   field,
		Code:    code,
		Message: message,
	}
}

const (
	CodeInvalidValue = "invalid_value"
	CodeRequired     = "required"
	CodeOutOfRange   = "out_of_range"
)

type ValidationErrors struct {
	errors []ValidationError
}

func NewValidationErrors(errors []error) ValidationErrors {
	result := ValidationErrors{}
	for _, err := range errors {
		if validationErr, ok := err.(ValidationError); ok {
			result.errors = append(result.errors, validationErr)
		}
	}
	return result
}

func (e ValidationErrors) Error() string {
	result, err := json.Marshal(e.errors)
	if err != nil {
		return "failed to marshal validation errors"
	}
	return string(result)
}
