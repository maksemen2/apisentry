package coreerrors

var (
	ErrInvalidTotalEndpoints = NewValidationError("total_endpoints",
		CodeOutOfRange,
		"Total endpoints should be greater than zero",
	)

	ErrInvalidUptimePercentage = NewValidationError("uptime_percentage",
		CodeOutOfRange,
		"Uptime percentage should be greater than zero",
	)

	ErrInvalidUptime24Hours = NewValidationError("uptime_24_hours",
		CodeOutOfRange,
		"Uptime 24 hours should be greater than zero",
	)

	ErrInvalidAvgResponseTime = NewValidationError("avg_response_time",
		CodeOutOfRange,
		"Avg response time should be greater than zero",
	)
)
