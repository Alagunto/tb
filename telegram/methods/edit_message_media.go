package methods

import "github.com/alagunto/tb/telegram"

// EditMessageMediaRequest represents the request for editMessageMedia method
type EditMessageMediaRequest struct {
	telegram.HasReplyMarkup

	ChatID          string `json:"chat_id,omitempty"`
	MessageID       string `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	Media           string `json:"media"` // JSON-serialized InputMedia
}

// EditMessageMediaResponse represents the response for editMessageMedia method
type EditMessageMediaResponse = telegram.Message
