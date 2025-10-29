package methods

import "github.com/alagunto/tb/telegram"

// GetChatMemberRequest represents the request for getChatMember method.
type GetChatMemberRequest struct {
	// Unique identifier for the target chat or username of the target channel
	ChatID string `json:"chat_id"`

	// Unique identifier of the target user
	UserID string `json:"user_id"`
}

// GetChatMemberResponse represents the response for getChatMember method.
type GetChatMemberResponse = telegram.ChatMember
