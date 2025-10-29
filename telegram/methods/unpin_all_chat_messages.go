package methods

// UnpinAllChatMessagesRequest represents the request for unpinAllChatMessages method.
type UnpinAllChatMessagesRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`
}

// UnpinAllChatMessagesResponse represents the response for unpinAllChatMessages method.
type UnpinAllChatMessagesResponse bool
