package request

import (
	"time"

	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/telegram"
)

// Context wraps an update and represents the context of current event/request to the bot.
type Interface interface {
	// Bot returns the bot instance.
	Bot() communications.API

	// Update returns the original update object.
	Update() *telegram.Update

	// Message returns stored message if such presented.
	Message() *telegram.Message

	// Callback returns stored callback if such presented.
	Callback() *telegram.CallbackQuery

	// Query returns stored query if such presented.
	Query() *telegram.InlineQuery

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

	// Topic returns the topic changes.
	Topic() *telegram.Topic

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
	Recipient() communications.Recipient

	// Text returns the message text, depending on the context type.
	// In the case when no related data presented, returns an empty string.
	Text() string

	// ThreadID returns the current message thread ID.
	ThreadID() int

	// Entities returns the message entities, whether it's media caption's or the text's.
	// In the case when no entities presented, returns a nil.
	Entities() telegram.Entities

	// Data returns the current data, depending on the context type.
	// If the context contains command, returns its arguments string.
	// If the context contains payment, returns its payload.
	// In the case when no related data presented, returns an empty string.
	Data() string

	// Args returns a raw slice of command or callback arguments as strings.
	// The message arguments split by space, while the callback's ones by a "|" symbol.
	Args() []string

	// Send sends a message to the current recipient.
	// See Send from bot.go.
	Send(what interface{}, opts ...interface{}) error

	// SendAlbum sends an album to the current recipient.
	// See SendAlbum from bot.go.
	SendAlbum(a telegram.Album, opts ...interface{}) error

	// Reply replies to the current message.
	// See Reply from bot.go.
	Reply(what interface{}, opts ...interface{}) error

	// Forward forwards the given message to the current recipient.
	// See Forward from bot.go.
	Forward(msg communications.Editable, opts ...interface{}) error

	// ForwardTo forwards the current message to the given recipient.
	// See Forward from bot.go
	ForwardTo(to communications.Recipient, opts ...interface{}) error

	// Edit edits the current message.
	// See Edit from bot.go.
	Edit(what interface{}, opts ...interface{}) error

	// EditCaption edits the caption of the current message.
	// See EditCaption from bot.go.
	EditCaption(caption string, opts ...interface{}) error

	// EditOrSend edits the current message if the update is callback,
	// otherwise the content is sent to the chat as a separate message.
	EditOrSend(what interface{}, opts ...interface{}) error

	// EditOrReply edits the current message if the update is callback,
	// otherwise the content is replied as a separate message.
	EditOrReply(what interface{}, opts ...interface{}) error

	// Delete removes the current message.
	// See Delete from bot.go.
	Delete() error

	// DeleteAfter waits for the duration to elapse and then removes the
	// message. It handles an error automatically using b.OnError callback.
	// It returns a Timer that can be used to cancel the call using its Stop method.
	DeleteAfter(d time.Duration) *time.Timer

	// Notify updates the chat action for the current recipient.
	// See Notify from bot.go.
	Notify(action telegram.ChatAction) error

	// Ship replies to the current shipping query.
	// See Ship from bot.go.
	Ship(what ...interface{}) error

	// Accept finalizes the current deal.
	// See Accept from bot.go.
	Accept(errorMessage ...string) error

	// Answer sends a response to the current inline query.
	// See Answer from bot.go.
	Answer(resp *telegram.QueryResponse) error

	// Respond sends a response for the current callback query.
	// See Respond from bot.go.
	Respond(resp ...*telegram.CallbackResponse) error

	// RespondText sends a popup response for the current callback query.
	RespondText(text string) error

	// RespondAlert sends an alert response for the current callback query.
	RespondAlert(text string) error

	// Get retrieves data from the context.
	Get(key string) interface{}

	// Set saves data in the context.
	Set(key string, val interface{})
}
