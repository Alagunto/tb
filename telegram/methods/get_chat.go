package methods

import "github.com/alagunto/tb/telegram"

// GetChatRequest represents the request for getChat method.
type GetChatRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`
}

// GetChatResponse represents the response for getChat method.
type GetChatResponse = telegram.Chat
