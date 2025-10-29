package methods

import "github.com/alagunto/tb/telegram"

// ForwardMessageRequest represents the request for forwardMessage method.
type ForwardMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`

	// Unique identifier for the chat where the original message was sent
	FromChatID int64 `json:"from_chat_id"`

	// Message identifier in the chat specified in from_chat_id
	MessageID string `json:"message_id"`

	// Unique identifier for the target message thread
	MessageThreadID int `json:"message_thread_id,omitempty"`

	// Sends the message silently
	DisableNotification bool `json:"disable_notification,omitempty"`

	// Protects the contents of the forwarded message from forwarding and saving
	ProtectContent bool `json:"protect_content,omitempty"`
}

// ForwardMessageResponse represents the response for forwardMessage method.
type ForwardMessageResponse = telegram.Message
