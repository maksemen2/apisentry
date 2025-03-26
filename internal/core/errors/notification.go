package coreerrors

var (
	ErrInvalidNotificationType = NewValidationError("type",
		CodeInvalidValue,
		"Invalid notification type",
	)

	ErrEmailRequired = NewValidationError("email",
		CodeRequired,
		"Email is required",
	)

	ErrChatIDRequired = NewValidationError("chat_id",
		CodeRequired,
		"Chat ID is required",
	)

	ErrBotTokenRequired = NewValidationError("bot_token",
		CodeRequired,
		"Bot token is required",
	)

	ErrWebhookURLRequired = NewValidationError("url",
		CodeRequired,
		"URL is required",
	)
)
