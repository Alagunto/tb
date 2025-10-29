package methods

import "github.com/alagunto/tb/telegram"

// SendMessageRequest represents the request for sendMessage method.
type SendMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`

	// Text of the message to be sent, 1-4096 characters
	Text string `json:"text"`

	// Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Unique identifier for the target message thread
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Mode for parsing entities in the message text
	ParseMode telegram.ParseMode `json:"parse_mode,omitempty"`

	// List of special entities that appear in message text
	Entities telegram.Entities `json:"entities,omitempty"`

	// Link preview generation options for the message
	LinkPreviewOptions interface{} `json:"link_preview_options,omitempty"`

	// Sends the message silently
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Unique identifier of the message effect to be added to the message
	MessageEffectID string `json:"message_effect_id,omitempty"`

	// Description of the message to reply to
	ReplyParameters *telegram.ReplyParams `json:"reply_parameters,omitempty"`

	// Additional interface options
	ReplyMarkup *telegram.ReplyMarkup `json:"reply_markup,omitempty"`
}

// SendMessageResponse represents the response for sendMessage method.
type SendMessageResponse = telegram.Message
