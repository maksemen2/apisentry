package coreerrors

var (
	ErrInvalidResponseTime = NewValidationError("response_time",
		CodeOutOfRange,
		"Response time should be equal or greater than zero")
)
