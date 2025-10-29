package methods

import "github.com/alagunto/tb/telegram"

// EditMessageTextRequest represents the request for editMessageText method.
type EditMessageTextRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of the message to edit
	MessageID string `json:"message_id,omitempty"`

	// Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// New text of the message, 1-4096 characters after entities parsing
	Text string `json:"text"`

	// Mode for parsing entities in the message text
	ParseMode string `json:"parse_mode,omitempty"`

	// List of special entities that appear in message text
	Entities interface{} `json:"entities,omitempty"`

	// Link preview generation options for the message
	LinkPreviewOptions interface{} `json:"link_preview_options,omitempty"`

	// Additional interface options
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`

	// Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
}

// EditMessageTextResponse represents the response for editMessageText method.
type EditMessageTextResponse = telegram.Message
