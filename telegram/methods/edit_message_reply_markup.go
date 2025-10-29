package methods

import "github.com/alagunto/tb/telegram"

// EditMessageReplyMarkupRequest represents the request for editMessageReplyMarkup method.
type EditMessageReplyMarkupRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of the message to edit
	MessageID string `json:"message_id,omitempty"`

	// Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// An object for an inline keyboard
	ReplyMarkup *telegram.ReplyMarkup `json:"reply_markup,omitempty"`

	// Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
}

// EditMessageReplyMarkupResponse represents the response for editMessageReplyMarkup method.
type EditMessageReplyMarkupResponse = telegram.Message
