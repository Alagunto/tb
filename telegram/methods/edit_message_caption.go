package methods

import "github.com/alagunto/tb/telegram"

// EditMessageCaptionRequest represents the request for editMessageCaption method.
type EditMessageCaptionRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id,omitempty"`

	// Identifier of the message to edit
	MessageID string `json:"message_id,omitempty"`

	// Identifier of the inline message
	InlineMessageID string `json:"inline_message_id,omitempty"`

	// New caption of the message, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the message caption
	ParseMode string `json:"parse_mode,omitempty"`

	// List of special entities that appear in the caption
	CaptionEntities interface{} `json:"caption_entities,omitempty"`

	// Pass True if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// Additional interface options
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`

	// Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
}

// EditMessageCaptionResponse represents the response for editMessageCaption method.
type EditMessageCaptionResponse = telegram.Message
