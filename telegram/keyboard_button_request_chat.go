package telegram

// KeyboardButtonRequestChat represents the chat sharing settings.
type KeyboardButtonRequestChat struct {
	ID              int32            `json:"request_id"`
	Channel         bool             `json:"chat_is_channel"`
	Forum           *bool            `json:"chat_is_forum,omitempty"`
	WithUsername    *bool            `json:"chat_has_username,omitempty"`
	Created         *bool            `json:"chat_is_created,omitempty"`
	UserRights      *ChatPermissions `json:"user_administrator_rights,omitempty"`
	BotRights       *ChatPermissions `json:"bot_administrator_rights,omitempty"`
	BotMember       *bool            `json:"bot_is_member,omitempty"`
	RequestTitle    *bool            `json:"request_title,omitempty"`
	RequestUsername *bool            `json:"request_username,omitempty"`
	RequestPhoto    *bool            `json:"request_photo,omitempty"`
}
