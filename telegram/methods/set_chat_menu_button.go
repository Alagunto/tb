package methods

import "github.com/alagunto/tb/telegram"

// SetChatMenuButtonRequest represents the request for setChatMenuButton method.
type SetChatMenuButtonRequest struct {
	// Unique identifier for the target private chat. If not specified, default bot's menu button will be changed
	ChatID string `json:"chat_id,omitempty"`

	// An object for the bot's new menu button
	MenuButton *telegram.MenuButton `json:"menu_button,omitempty"`
}

// SetChatMenuButtonResponse represents the response for setChatMenuButton method.
type SetChatMenuButtonResponse bool
