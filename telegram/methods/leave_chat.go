package methods

// LeaveChatRequest represents the request for leaveChat method.
type LeaveChatRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`
}

// LeaveChatResponse represents the response for leaveChat method.
type LeaveChatResponse bool
