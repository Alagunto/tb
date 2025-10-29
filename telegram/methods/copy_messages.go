package methods

import "github.com/alagunto/tb/telegram"

// CopyMessagesRequest represents the request for copyMessages method
type CopyMessagesRequest struct {
	ChatID              string   `json:"chat_id"`
	FromChatID          int64    `json:"from_chat_id"`
	MessageIDs          []string `json:"message_ids"`
	MessageThreadID     int      `json:"message_thread_id,omitempty"`
	DisableNotification bool     `json:"disable_notification,omitempty"`
	ProtectContent      bool     `json:"protect_content,omitempty"`
	RemoveCaption       bool     `json:"remove_caption,omitempty"`
}

// CopyMessagesResponse represents the response for copyMessages method
type CopyMessagesResponse = []telegram.Message
