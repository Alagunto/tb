package request

import (
	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/telegram"
)

// Context wraps an update and represents the context of current event/request to the bot.
type Interface interface {
	bot.API

	// Bot returns the bot instance.
	Bot() interface{}

	// Update returns the original update object.
	Update() *telegram.Update

	// Message returns message that was sent by the user.
	Message() *telegram.Message

	// Callback returns stored callback if such presented.
	CallbackQuery() *telegram.CallbackQuery

	// CallbackMessage returns the message that was used to trigger the callback query if such presented.
	CallbackMessage() *telegram.Message

	// InlineQuery returns stored inline query if such presented.
	InlineQuery() *telegram.InlineQuery

	// InlineResult returns stored inline result if such presented.
	InlineResult() *telegram.InlineResult

	// ShippingQuery returns stored shipping query if such presented.
	ShippingQuery() *telegram.ShippingQuery

	// PreCheckoutQuery returns stored pre checkout query if such presented.
	PreCheckoutQuery() *telegram.PreCheckoutQuery

	// Payment returns payment instance.
	Payment() *telegram.Payment

	// Poll returns stored poll if such presented.
	Poll() *telegram.Poll

	// PollAnswer returns stored poll answer if such presented.
	PollAnswer() *telegram.PollAnswer

	// ChatMember returns chat member changes.
	ChatMember() *telegram.ChatMember

	// ChatJoinRequest returns the chat join request.
	ChatJoinRequest() *telegram.ChatJoinRequest

	// Migration returns both migration from and to chat IDs.
	Migration() (int64, int64)

	// Thread returns the thread changes.
	Thread() *telegram.Thread

	// Boost returns the boost instance.
	Boost() *telegram.BoostUpdated

	// BoostRemoved returns the boost removed from a chat instance.
	BoostRemoved() *telegram.BoostRemoved

	// Sender returns the current recipient, depending on the context type.
	// Returns nil if user is not presented.
	Sender() *telegram.User

	// Chat returns the current chat, depending on the context type.
	// Returns nil if chat is not presented.
	Chat() *telegram.Chat

	// Recipient combines both Sender and Chat functions. If there is no user
	// the chat will be returned. The native context cannot be without sender,
	// but it is useful in the case when the context created intentionally
	// by the NewContext constructor and have only Chat field inside.
	Recipient() bot.Recipient

	// Text returns the message text, depending on the context type.
	// In the case when no related data presented, returns an empty string.
	Text() string

	// ThreadID returns the current message thread ID.
	ThreadID() int

	// Data returns the current callback query data, depending on the context type.
	Data() string

	// Args returns a raw slice of command or callback arguments as strings.
	// The message arguments split by space, while the callback's ones by a "|" symbol.
	Args() []string
}
