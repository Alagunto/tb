package methods

// DeleteMessageRequest represents the request for deleteMessage method.
type DeleteMessageRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID int64 `json:"chat_id"`

	// Identifier of the message to delete
	MessageID string `json:"message_id"`
}

// DeleteMessageResponse represents the response for deleteMessage method.
type DeleteMessageResponse bool
