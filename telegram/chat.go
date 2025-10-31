package telegram

import "strconv"

// Chat object represents a Telegram user, bot, group or a channel.
type Chat struct {
	ID int64 `json:"id"`

	// See ChatType and consts.
	Type ChatType `json:"type"`

	// Won't be there for ChatPrivate.
	Title string `json:"title"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// Recipient returns chat id as string for sending messages.
// If chat has a username, returns @username instead.
func (c *Chat) Recipient() string {
	if c.Username != "" {
		return "@" + c.Username
	}
	return strconv.FormatInt(c.ID, 10)
}

// ChatType represents one of the possible chat types.
type ChatType string

const (
	ChatPrivate    ChatType = "private"
	ChatGroup      ChatType = "group"
	ChatSuperGroup ChatType = "supergroup"
	ChatChannel    ChatType = "channel"
)
