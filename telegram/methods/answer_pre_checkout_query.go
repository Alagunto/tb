package methods

// AnswerPreCheckoutQueryRequest represents the request for answerPreCheckoutQuery method.
type AnswerPreCheckoutQueryRequest struct {
	// Unique identifier for the query to be answered
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`

	// Specify True if everything is alright (goods are available, etc.) and the bot is ready to proceed with the order. Use False if there are any problems.
	Ok bool `json:"ok"`

	// Required if ok is False. Error message in human readable form that explains the reason for failure to proceed with the checkout
	ErrorMessage string `json:"error_message,omitempty"`
}

// AnswerPreCheckoutQueryResponse represents the response for answerPreCheckoutQuery method.
type AnswerPreCheckoutQueryResponse bool
