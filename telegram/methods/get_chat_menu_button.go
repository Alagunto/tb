package methods

import "github.com/alagunto/tb/telegram"

// GetChatMenuButtonRequest represents the request for getChatMenuButton method.
type GetChatMenuButtonRequest struct {
	// Unique identifier for the target private chat. If not specified, default bot's menu button will be returned
	ChatID string `json:"chat_id,omitempty"`
}

// GetChatMenuButtonResponse represents the response for getChatMenuButton method.
type GetChatMenuButtonResponse = telegram.MenuButton
