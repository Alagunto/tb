package bot

import (
	"io"
	"time"

	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// type Instance interface {
// API
// }

type Recipient interface {
	Recipient() string
}

type Editable interface {
	MessageSig() (messageID string, chatID int64)
}

// API specifies the methods that are available for the bot that can be called both from Request context and from the bot instance itself.
type API interface {
	Raw(method string, payload interface{}) ([]byte, error)
	GetUpdates(offset, limit int, timeout time.Duration, allowed []string) ([]telegram.Update, error)
	SendTo(to Recipient, what interface{}, opts ...params.SendOptions) (*telegram.Message, error)
	SendAlbumTo(to Recipient, a telegram.InputAlbum, opts ...params.SendOptions) ([]telegram.Message, error)
	ReplyTo(to *telegram.Message, what interface{}, opts ...params.SendOptions) (*telegram.Message, error)
	ForwardTo(to Recipient, msg Editable, opts ...params.SendOptions) (*telegram.Message, error)
	ForwardManyTo(to Recipient, msgs []Editable, opts ...params.SendOptions) ([]telegram.Message, error)
	CopyTo(to Recipient, msg Editable, opts ...params.SendOptions) (*telegram.Message, error)
	CopyManyTo(to Recipient, msgs []Editable, opts ...params.SendOptions) ([]telegram.Message, error)
	Edit(msg Editable, what interface{}, opts ...params.SendOptions) (*telegram.Message, error)
	EditReplyMarkup(msg Editable, markup *telegram.ReplyMarkup) (*telegram.Message, error)
	EditCaption(msg Editable, caption string, opts ...params.SendOptions) (*telegram.Message, error)
	EditMedia(msg Editable, media telegram.InputMedia, opts ...params.SendOptions) (*telegram.Message, error)
	Delete(msg Editable) error
	DeleteMany(msgs []Editable) error
	Notify(to Recipient, action telegram.ChatAction, opts ...params.SendOptions) error
	Ship(query *telegram.ShippingQuery, what ...interface{}) error
	Accept(query *telegram.PreCheckoutQuery, errorMessage ...string) error
	RespondToCallback(c *telegram.Callback, resp ...*telegram.CallbackResponse) error
	AnswerInlineQuery(query *telegram.InlineQuery, resp *telegram.InlineQueryResponse) error
	// AnswerWebAppQuery(query *telegram.WebApp, r telegram.Result) (*telegram.WebAppMessage, error)
	FileByID(fileID string) (files.FileReference, error)
	Download(file *files.FileReference, localFilename string) error
	File(file *files.FileReference) (io.ReadCloser, error)
	StopLiveLocation(msg Editable, opts ...params.SendOptions) (*telegram.Message, error)
	StopPoll(msg Editable, opts ...params.SendOptions) (*telegram.Poll, error)
	Leave(chat Recipient) error
	Pin(msg Editable) error
	Unpin(chat Recipient, messageID ...int) error
	UnpinAll(chat Recipient) error
	ChatByID(id int64) (*telegram.Chat, error)
	ChatByUsername(name string) (*telegram.Chat, error)
	ProfilePhotosOf(user *telegram.User) ([]telegram.PhotoSize, error)
	ChatMemberOf(chat, user Recipient) (*telegram.ChatMember, error)
	SetMenuButton(chat *telegram.User, mb interface{}) error
	Logout() (bool, error)
	SetMyName(name, language string) error
	SetMyDescription(desc, language string) error
	SetMyShortDescription(desc, language string) error
	GetMe() (*telegram.User, error)
}
