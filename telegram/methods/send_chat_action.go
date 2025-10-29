package methods

// SendChatActionRequest represents the request for sendChatAction method.
type SendChatActionRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`

	// Type of action to broadcast
	Action string `json:"action"`

	// Unique identifier of the business connection
	BusinessConnectionID string `json:"business_connection_id,omitempty"`

	// Unique identifier for the target message thread
	MessageThreadID int `json:"message_thread_id,omitempty"`
}

// SendChatActionResponse represents the response for sendChatAction method.
type SendChatActionResponse bool
