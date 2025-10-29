package methods

import "github.com/alagunto/tb/telegram"

// EditMessageReplyMarkupRequest represents the request for editMessageReplyMarkup method.
type EditMessageReplyMarkupRequest struct {
	telegram.HasReplyMarkup
	telegram.HasBusinessConnection

	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of the message to edit
	MessageID string `json:"message_id,omitempty"`

	// Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// EditMessageReplyMarkupResponse represents the response for editMessageReplyMarkup method.
type EditMessageReplyMarkupResponse = telegram.Message
