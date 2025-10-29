package telegram

// Topic represents a forum topic.
type Thread struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
	ID                int    `json:"message_thread_id"`
}

// Game object represents a game.
type Game struct {
	Name string `json:"game_short_name"`

	Title       string `json:"title"`
	Description string `json:"description"`
	Photo       *Photo `json:"photo"`

	// (Optional)
	Text      string          `json:"text"`
	Entities  []MessageEntity `json:"text_entities"`
	Animation *Animation      `json:"animation"`
}

// ProximityAlert sent whenever a user in the chat triggers
// a proximity alert set by another user.
type ProximityAlert struct {
	Traveler *User `json:"traveler,omitempty"`
	Watcher  *User `json:"watcher,omitempty"`
	Distance int   `json:"distance"`
}

// AutoDeleteTimer represents a service message about a change in auto-delete timer settings.
type AutoDeleteTimer struct {
	Unixtime int `json:"message_auto_delete_time"`
}

// VideoChatStarted represents a service message about a video chat started in the chat.
type VideoChatStarted struct{}

// VideoChatEnded represents a service message about a video chat ended in the chat.
type VideoChatEnded struct {
	Duration int `json:"duration"`
}

// VideoChatParticipants represents a service message about new members invited to a video chat.
type VideoChatParticipants struct {
	Users []User `json:"users"`
}

// VideoChatScheduled represents a service message about a video chat scheduled in the chat.
type VideoChatScheduled struct {
	Unixtime int64 `json:"start_date"`
}

// RecipientShared represents information about a user or chat shared with the bot.
type RecipientShared struct {
	RequestID int   `json:"request_id"`
	UserID    int64 `json:"user_id,omitempty"`
	ChatID    int64 `json:"chat_id,omitempty"`
}

// ChatBackground represents a chat background.
type ChatBackground struct {
	Type string `json:"type"`
}

// ForumTopicCreated represents a service message about a new forum topic created in the chat.
type ForumTopicCreated struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicEdited represents a service message about an edited forum topic.
type ForumTopicEdited struct {
	Name              string `json:"name,omitempty"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id,omitempty"`
}

// ForumTopicClosed represents a service message about a forum topic closed in the chat.
type ForumTopicClosed struct{}

// ForumTopicReopened represents a service message about a forum topic reopened in the chat.
type ForumTopicReopened struct{}

// GeneralForumTopicHidden represents a service message about General forum topic hidden in the chat.
type GeneralForumTopicHidden struct{}

// GeneralForumTopicUnhidden represents a service message about General forum topic unhidden in the chat.
type GeneralForumTopicUnhidden struct{}
