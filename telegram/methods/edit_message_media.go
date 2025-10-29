package methods

import "github.com/alagunto/tb/telegram"

// EditMessageMediaRequest represents the request for editMessageMedia method
type EditMessageMediaRequest struct {
	ChatID          string                `json:"chat_id,omitempty"`
	MessageID       string                `json:"message_id,omitempty"`
	InlineMessageID string                `json:"inline_message_id,omitempty"`
	Media           string                `json:"media"` // JSON-serialized InputMedia
	ReplyMarkup     *telegram.ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageMediaResponse represents the response for editMessageMedia method
type EditMessageMediaResponse = telegram.Message
