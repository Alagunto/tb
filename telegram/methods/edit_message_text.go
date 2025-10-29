package methods

import "github.com/alagunto/tb/telegram"

// EditMessageTextRequest represents the request for editMessageText method.
type EditMessageTextRequest struct {
	telegram.HasReplyMarkup
	telegram.HasBusinessConnection

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
	Entities telegram.Entities `json:"entities,omitempty"`

	// Link preview generation options for the message
	LinkPreviewOptions *telegram.PreviewOptions `json:"link_preview_options,omitempty"`
}

// SetParseMode implements SetsParseMode interface.
func (r *EditMessageTextRequest) SetParseMode(mode telegram.ParseMode) {
	r.ParseMode = string(mode)
}

// SetEntities implements SetsEntities interface.
func (r *EditMessageTextRequest) SetEntities(entities telegram.Entities) {
	r.Entities = entities
}

// EditMessageTextResponse represents the response for editMessageText method.
type EditMessageTextResponse = telegram.Message
