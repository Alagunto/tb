package telegram

import "strconv"

// User object represents a Telegram user, bot.
type User struct {
	ID int64 `json:"id"`

	FirstName           string   `json:"first_name"`
	LastName            string   `json:"last_name"`
	IsForum             bool     `json:"is_forum"`
	Username            string   `json:"username"`
	LanguageCode        string   `json:"language_code"`
	IsBot               bool     `json:"is_bot"`
	IsPremium           bool     `json:"is_premium"`
	AddedToMenu         bool     `json:"added_to_attachment_menu"`
	Usernames           []string `json:"active_usernames"`
	StatusCustomEmojiID string   `json:"emoji_status_custom_emoji_id"`

	// Returns only in getMe
	CanJoinGroups        bool `json:"can_join_groups"`
	CanReadMessages      bool `json:"can_read_all_group_messages"`
	SupportsInline       bool `json:"supports_inline_queries"`
	CanConnectToBusiness bool `json:"can_connect_to_business"`
	HasMainWebApp        bool `json:"has_main_web_app"`
}

// Recipient returns user id as string for sending messages.
func (u *User) Recipient() string {
	return strconv.FormatInt(u.ID, 10)
}
