package communications

import "github.com/alagunto/tb/telegram"

// Recipient is any possible endpoint you can send messages to: either user, group or a channel.
type Recipient interface {
	Recipient() string // must return legit Telegram chat_id or username
}

// Editable is an interface for all objects that provide "message signature",
// a pair of 32-bit message ID and 64-bit chat ID, both required for edit operations.
type Editable interface {
	// MessageSig is a "message signature".
	//
	// For inline messages, return chatID = 0.
	MessageSig() (messageID string, chatID int64)
}

// API is the minimal interface that represents interactions with Telegram Bot API.
type API interface {
	Raw(method string, payload interface{}) ([]byte, error)
}

// Sendable is any object that can send itself.
//
// This is pretty cool, since it lets bots implement
// custom Sendables for complex kind of media or
// chat objects spanning across multiple messages.
type Sendable interface {
	Send(API, Recipient, *SendOptions) (*telegram.Message, error)
}
