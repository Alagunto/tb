package telegram

// SwitchInlineQuery represents a switch to an inline query in a chosen chat.
type SwitchInlineQuery struct {
	Query string `json:"query"`
	AllowUserChats bool `json:"allow_user_chats,omitempty"`
	AllowBotChats bool `json:"allow_bot_chats,omitempty"`
	AllowGroupChats bool `json:"allow_group_chats,omitempty"`
	AllowChannelChats bool `json:"allow_channel_chats,omitempty"`
}
