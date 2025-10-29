package methods

import "github.com/alagunto/tb/telegram"

// AnswerShippingQueryRequest represents the request for answerShippingQuery method.
type AnswerShippingQueryRequest struct {
	// Unique identifier for the query to be answered
	ShippingQueryID string `json:"shipping_query_id"`

	// Specify True if delivery to the specified address is possible and False if there are any problems
	Ok bool `json:"ok"`

	// Required if ok is True. Array of available shipping options
	ShippingOptions []telegram.ShippingOption `json:"shipping_options,omitempty"`

	// Required if ok is False. Error message in human readable form
	ErrorMessage string `json:"error_message,omitempty"`
}

// AnswerShippingQueryResponse represents the response for answerShippingQuery method.
type AnswerShippingQueryResponse bool
