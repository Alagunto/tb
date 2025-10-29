package methods

import "github.com/alagunto/tb/telegram"

// CopyMessageRequest represents the request for copyMessage method.
type CopyMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`

	// Unique identifier for the chat where the original message was sent
	FromChatID int64 `json:"from_chat_id"`

	// Message identifier in the chat specified in from_chat_id
	MessageID string `json:"message_id"`

	// Unique identifier for the target message thread
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// New caption for media, 0-1024 characters after entities parsing
	Caption string `json:"caption,omitempty"`

	// Mode for parsing entities in the new caption
	ParseMode string `json:"parse_mode,omitempty"`

	// List of special entities that appear in the new caption
	CaptionEntities interface{} `json:"caption_entities,omitempty"`

	// Pass True if the caption must be shown above the message media
	ShowCaptionAboveMedia bool `json:"show_caption_above_media,omitempty"`

	// Sends the message silently
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the sent message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`

	// Description of the message to reply to
	ReplyParameters *telegram.ReplyParams `json:"reply_parameters,omitempty"`

	// Additional interface options
	ReplyMarkup interface{} `json:"reply_markup,omitempty"`
}

// CopyMessageResponse represents the response for copyMessage method.
type CopyMessageResponse = telegram.Message
