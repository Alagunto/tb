package tb

import (
	"fmt"
	"strings"

	"github.com/alagunto/tb/telegram"
)

// ReplyMarkup controls two convenient options for bot-user communications
// such as reply keyboard and inline "keyboard" (a grid of buttons as a part
// of the message).
type ReplyMarkup struct {
	// InlineKeyboard is a grid of InlineButtons displayed in the message.
	//
	// Note: DO NOT confuse with ReplyKeyboard and other keyboard properties!
	InlineKeyboard [][]telegram.InlineButton `json:"inline_keyboard,omitempty"`

	// ReplyKeyboard is a grid, consisting of keyboard buttons.
	//
	// Note: you don't need to set HideCustomKeyboard field to show custom keyboard.
	ReplyKeyboard [][]telegram.ReplyButton `json:"keyboard,omitempty"`

	// ForceReply forces Telegram clients to display
	// a reply interface to the user (act as if the user
	// has selected the bot‘s message and tapped "Reply").
	ForceReply bool `json:"force_reply,omitempty"`

	// Requests clients to resize the keyboard vertically for optimal fit
	// (e.g. make the keyboard smaller if there are just two rows of buttons).
	//
	// Defaults to false, in which case the custom keyboard is always of the
	// same height as the app's standard keyboard.
	ResizeKeyboard bool `json:"resize_keyboard,omitempty"`

	// Requests clients to hide the reply keyboard as soon as it's been used.
	//
	// Defaults to false.
	OneTimeKeyboard bool `json:"one_time_keyboard,omitempty"`

	// Requests clients to remove the reply keyboard.
	//
	// Defaults to false.
	RemoveKeyboard bool `json:"remove_keyboard,omitempty"`

	// Use this param if you want to force reply from
	// specific users only.
	//
	// Targets:
	// 1) Users that are @mentioned in the text of the Message object;
	// 2) If the bot's message is a reply (has SendOptions.ReplyTo),
	//       sender of the original message.
	Selective bool `json:"selective,omitempty"`

	// Placeholder will be shown in the input field when the reply is active.
	Placeholder string `json:"input_field_placeholder,omitempty"`

	// IsPersistent allows to control when the keyboard is shown.
	IsPersistent bool `json:"is_persistent,omitempty"`
}

func (r *ReplyMarkup) copy() *ReplyMarkup {
	cp := *r

	if len(r.ReplyKeyboard) > 0 {
		cp.ReplyKeyboard = make([][]telegram.ReplyButton, len(r.ReplyKeyboard))
		for i, row := range r.ReplyKeyboard {
			cp.ReplyKeyboard[i] = make([]telegram.ReplyButton, len(row))
			copy(cp.ReplyKeyboard[i], row)
		}
	}

	if len(r.InlineKeyboard) > 0 {
		cp.InlineKeyboard = make([][]telegram.InlineButton, len(r.InlineKeyboard))
		for i, row := range r.InlineKeyboard {
			cp.InlineKeyboard[i] = make([]telegram.InlineButton, len(row))
			copy(cp.InlineKeyboard[i], row)
		}
	}

	return &cp
}

// Btn is a constructor button, which will later become either a reply, or an inline button.
type Btn struct {
	Unique          string                   `json:"unique,omitempty"`
	Text            string                   `json:"text,omitempty"`
	URL             string                   `json:"url,omitempty"`
	Data            string                   `json:"callback_data,omitempty"`
	InlineQuery     string                   `json:"switch_inline_query,omitempty"`
	InlineQueryChat string                   `json:"switch_inline_query_current_chat,omitempty"`
	Login           *telegram.Login          `json:"login_url,omitempty"`
	WebApp          *telegram.WebApp         `json:"web_app,omitempty"`
	Contact         bool                     `json:"request_contact,omitempty"`
	Location        bool                     `json:"request_location,omitempty"`
	Poll            telegram.PollType        `json:"request_poll,omitempty"`
	User            *telegram.ReplyRecipient `json:"request_user,omitempty"`
	Chat            *telegram.ReplyRecipient `json:"request_chat,omitempty"`
}

// Row represents an array of buttons, a row.
type Row []Btn

// Row creates a row of buttons.
func (r *ReplyMarkup) Row(many ...Btn) Row {
	return many
}

// Split splits the keyboard into the rows with N maximum number of buttons.
// For example, if you pass six buttons and 3 as the max, you get two rows with
// three buttons in each.
//
// `Split(3, []Btn{six buttons...}) -> [[1, 2, 3], [4, 5, 6]]`
// `Split(2, []Btn{six buttons...}) -> [[1, 2],[3, 4],[5, 6]]`
func (r *ReplyMarkup) Split(max int, btns []Btn) []Row {
	rows := make([]Row, (max-1+len(btns))/max)
	for i, b := range btns {
		i /= max
		rows[i] = append(rows[i], b)
	}
	return rows
}

func (r *ReplyMarkup) Inline(rows ...Row) {
	inlineKeys := make([][]telegram.InlineButton, 0, len(rows))
	for i, row := range rows {
		keys := make([]telegram.InlineButton, 0, len(row))
		for j, btn := range row {
			btn := btn.Inline()
			if btn == nil {
				panic(fmt.Sprintf(
					"telebot: button row %d column %d is not an inline button",
					i, j))
			}
			keys = append(keys, *btn)
		}
		inlineKeys = append(inlineKeys, keys)
	}

	r.InlineKeyboard = inlineKeys
}

func (r *ReplyMarkup) Reply(rows ...Row) {
	replyKeys := make([][]telegram.ReplyButton, 0, len(rows))
	for i, row := range rows {
		keys := make([]telegram.ReplyButton, 0, len(row))
		for j, btn := range row {
			btn := btn.Reply()
			if btn == nil {
				panic(fmt.Sprintf(
					"telebot: button row %d column %d is not a reply button",
					i, j))
			}
			keys = append(keys, *btn)
		}
		replyKeys = append(replyKeys, keys)
	}

	r.ReplyKeyboard = replyKeys
}

func (r *ReplyMarkup) Text(text string) Btn {
	return Btn{Text: text}
}

func (r *ReplyMarkup) Data(text, unique string, data ...string) Btn {
	return Btn{
		Unique: unique,
		Text:   text,
		Data:   strings.Join(data, "|"),
	}
}

func (r *ReplyMarkup) URL(text, url string) Btn {
	return Btn{Text: text, URL: url}
}

func (r *ReplyMarkup) Query(text, query string) Btn {
	return Btn{Text: text, InlineQuery: query}
}

func (r *ReplyMarkup) QueryChat(text, query string) Btn {
	return Btn{Text: text, InlineQueryChat: query}
}

func (r *ReplyMarkup) Contact(text string) Btn {
	return Btn{Contact: true, Text: text}
}

func (r *ReplyMarkup) Location(text string) Btn {
	return Btn{Location: true, Text: text}
}

func (r *ReplyMarkup) Poll(text string, poll telegram.PollType) Btn {
	return Btn{Poll: poll, Text: text}
}

func (r *ReplyMarkup) User(text string, user *telegram.ReplyRecipient) Btn {
	return Btn{Text: text, User: user}
}

func (r *ReplyMarkup) Chat(text string, chat *telegram.ReplyRecipient) Btn {
	return Btn{Text: text, Chat: chat}
}

func (r *ReplyMarkup) Login(text string, login *telegram.Login) Btn {
	return Btn{Login: login, Text: text}
}

func (r *ReplyMarkup) WebApp(text string, app *telegram.WebApp) Btn {
	return Btn{Text: text, WebApp: app}
}

// RecipientShared combines both UserShared and ChatShared objects.
type RecipientShared struct {
	ID       int32           `json:"request_id"` // chat, users
	ChatID   int64           `json:"chat_id"`    // chat only
	Title    string          `json:"title"`      // chat only
	Username string          `json:"username"`   // chat only
	Photo    *telegram.Photo `json:"photo"`      // chat only

	Users []struct {
		UserID    int64           `json:"user_id"`
		FirstName string          `json:"first_name"`
		LastName  string          `json:"last_name"`
		Username  string          `json:"username"`
		Photo     *telegram.Photo `json:"photo"`
	} `json:"users"` // users only

}

func (b Btn) Reply() *telegram.ReplyButton {
	if b.Unique != "" {
		return nil
	}

	return &telegram.ReplyButton{
		Text:     b.Text,
		Contact:  b.Contact,
		Location: b.Location,
		Poll:     b.Poll,
		User:     b.User,
		Chat:     b.Chat,
		WebApp:   b.WebApp,
	}
}

func (b Btn) Inline() *telegram.InlineButton {
	return &telegram.InlineButton{
		Unique:          b.Unique,
		Text:            b.Text,
		URL:             b.URL,
		Data:            b.Data,
		InlineQuery:     b.InlineQuery,
		InlineQueryChat: b.InlineQueryChat,
		Login:           b.Login,
		WebApp:          b.WebApp,
	}
}

// InlineKeyboardBuilder provides a fluent interface for building inline keyboards.
type InlineKeyboardBuilder struct {
	rows [][]telegram.InlineButton
}

// NewInlineKeyboard creates a new inline keyboard builder.
func NewInlineKeyboard() *InlineKeyboardBuilder {
	return &InlineKeyboardBuilder{
		rows: make([][]telegram.InlineButton, 0),
	}
}

// Row adds a new row of buttons to the keyboard.
func (kb *InlineKeyboardBuilder) Row(buttons ...telegram.InlineButton) *InlineKeyboardBuilder {
	if len(buttons) > 0 {
		kb.rows = append(kb.rows, buttons)
	}
	return kb
}

// Build returns the completed ReplyMarkup with the inline keyboard.
func (kb *InlineKeyboardBuilder) Build() *ReplyMarkup {
	return &ReplyMarkup{
		InlineKeyboard: kb.rows,
	}
}

// ReplyKeyboardBuilder provides a fluent interface for building reply keyboards.
type ReplyKeyboardBuilder struct {
	rows        [][]telegram.ReplyButton
	resize      bool
	oneTime     bool
	selective   bool
	placeholder string
	persistent  bool
}

// NewReplyKeyboard creates a new reply keyboard builder.
func NewReplyKeyboard() *ReplyKeyboardBuilder {
	return &ReplyKeyboardBuilder{
		rows: make([][]telegram.ReplyButton, 0),
	}
}

// Row adds a new row of buttons to the keyboard.
func (kb *ReplyKeyboardBuilder) Row(buttons ...telegram.ReplyButton) *ReplyKeyboardBuilder {
	if len(buttons) > 0 {
		kb.rows = append(kb.rows, buttons)
	}
	return kb
}

// Resize enables automatic keyboard resizing.
func (kb *ReplyKeyboardBuilder) Resize() *ReplyKeyboardBuilder {
	kb.resize = true
	return kb
}

// OneTime makes the keyboard hide after first use.
func (kb *ReplyKeyboardBuilder) OneTime() *ReplyKeyboardBuilder {
	kb.oneTime = true
	return kb
}

// Selective makes the keyboard visible only to specific users.
func (kb *ReplyKeyboardBuilder) Selective() *ReplyKeyboardBuilder {
	kb.selective = true
	return kb
}

// Placeholder sets the input field placeholder text.
func (kb *ReplyKeyboardBuilder) Placeholder(text string) *ReplyKeyboardBuilder {
	kb.placeholder = text
	return kb
}

// Persistent makes the keyboard persistent (shown even when bot messages scroll out of view).
func (kb *ReplyKeyboardBuilder) Persistent() *ReplyKeyboardBuilder {
	kb.persistent = true
	return kb
}

// Build returns the completed ReplyMarkup with the reply keyboard.
func (kb *ReplyKeyboardBuilder) Build() *ReplyMarkup {
	return &ReplyMarkup{
		ReplyKeyboard:   kb.rows,
		ResizeKeyboard:  kb.resize,
		OneTimeKeyboard: kb.oneTime,
		Selective:       kb.selective,
		Placeholder:     kb.placeholder,
		IsPersistent:    kb.persistent,
	}
}
