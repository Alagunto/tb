package methods

import "github.com/alagunto/tb/telegram"

// StopPollRequest represents the request for stopPoll method
type StopPollRequest struct {
	telegram.HasReplyMarkup

	ChatID    int64  `json:"chat_id"`
	MessageID string `json:"message_id"`
}

// StopPollResponse represents the response for stopPoll method
type StopPollResponse = telegram.Poll
