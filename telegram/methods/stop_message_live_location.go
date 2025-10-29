package methods

import "github.com/alagunto/tb/telegram"

// StopMessageLiveLocationRequest represents the request for stopMessageLiveLocation method
type StopMessageLiveLocationRequest struct {
	telegram.HasReplyMarkup

	ChatID          string `json:"chat_id,omitempty"`
	MessageID       string `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// StopMessageLiveLocationResponse represents the response for stopMessageLiveLocation method
type StopMessageLiveLocationResponse = telegram.Message
