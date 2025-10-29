package methods

// DeleteMessagesRequest represents the request for deleteMessages method.
type DeleteMessagesRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID int64 `json:"chat_id"`

	// A list of 1-100 identifiers of messages to delete
	MessageIDs []string `json:"message_ids"`
}

// DeleteMessagesResponse represents the response for deleteMessages method.
type DeleteMessagesResponse bool
