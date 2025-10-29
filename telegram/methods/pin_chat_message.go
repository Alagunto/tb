package methods

// PinChatMessageRequest represents the request for pinChatMessage method.
type PinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID int64 `json:"chat_id"`

	// Identifier of a message to pin
	MessageID string `json:"message_id"`

	// Pass True if it is not necessary to send a notification to all chat members about the new pinned message
	DisableNotification bool `json:"disable_notification,omitempty"`
}

// PinChatMessageResponse represents the response for pinChatMessage method.
type PinChatMessageResponse bool
