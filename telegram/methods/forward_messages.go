package methods

import "github.com/alagunto/tb/telegram"

// ForwardMessagesRequest represents the request for forwardMessages method
type ForwardMessagesRequest struct {
	ChatID              string   `json:"chat_id"`
	FromChatID          int64    `json:"from_chat_id"`
	MessageIDs          []string `json:"message_ids"`
	MessageThreadID     int      `json:"message_thread_id,omitempty"`
	DisableNotification bool     `json:"disable_notification,omitempty"`
	ProtectContent      bool     `json:"protect_content,omitempty"`
}

// ForwardMessagesResponse represents the response for forwardMessages method
type ForwardMessagesResponse = []telegram.Message
