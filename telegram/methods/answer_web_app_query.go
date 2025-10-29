package methods

import "github.com/alagunto/tb/telegram"

// AnswerWebAppQueryRequest represents the request for answerWebAppQuery method.
type AnswerWebAppQueryRequest struct {
	// Unique identifier for the query to be answered
	WebAppQueryID string `json:"web_app_query_id"`

	// An object describing the message to be sent
	Result telegram.Result `json:"result"`
}

// AnswerWebAppQueryResponse represents the response for answerWebAppQuery method.
type AnswerWebAppQueryResponse = telegram.WebAppMessage
