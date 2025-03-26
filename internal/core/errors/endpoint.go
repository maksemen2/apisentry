package coreerrors

var (
	ErrInvalidRetryStrategy = NewValidationError(
		"retry_strategy",
		CodeInvalidValue,
		"Invalid retry strategy",
	)

	ErrInvalidInterval = NewValidationError(
		"interval",
		CodeOutOfRange,
		"Interval should be greater than zero",
	)

	ErrInvalidStatus = NewValidationError(
		"status",
		CodeInvalidValue,
		"Invalid endpoint status",
	)

	ErrInvalidMaxRetries = NewValidationError(
		"max_retries",
		CodeOutOfRange,
		"Max retries should be greater than zero",
	)
)
