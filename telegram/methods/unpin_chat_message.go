package methods

// UnpinChatMessageRequest represents the request for unpinChatMessage method.
type UnpinChatMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`

	// Identifier of a message to unpin. If not specified, the most recent pinned message will be unpinned.
	MessageID int `json:"message_id,omitempty"`
}

// UnpinChatMessageResponse represents the response for unpinChatMessage method.
type UnpinChatMessageResponse bool
