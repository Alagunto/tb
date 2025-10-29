package telegram

import "encoding/json"

// ReplyMarkup controls two convenient options for bot-user communications
// such as reply keyboard and inline "keyboard" (a grid of buttons as a part
// of the message).
type ReplyMarkup struct {
	// InlineKeyboard is a grid of InlineButtons displayed in the message.
	//
	// Note: DO NOT confuse with ReplyKeyboard and other keyboard properties!
	InlineKeyboard [][]InlineButton `json:"inline_keyboard,omitempty"`

	// ReplyKeyboard is a grid, consisting of keyboard buttons.
	//
	// Note: you don't need to set HideCustomKeyboard field to show custom keyboard.
	ReplyKeyboard [][]ReplyButton `json:"keyboard,omitempty"`

	// ForceReply forces Telegram clients to display
	// a reply interface to the user (act as if the user
	// has selected the bot's message and tapped "Reply").
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

// ReplyButton represents a button displayed in reply-keyboard.
//
// Set either Contact or Location to true in order to request
// sensitive info, such as user's phone number or current location.
type ReplyButton struct {
	Text string `json:"text"`

	Contact  bool            `json:"request_contact,omitempty"`
	Location bool            `json:"request_location,omitempty"`
	Poll     PollType        `json:"request_poll,omitempty"`
	User     *ReplyRecipient `json:"request_users,omitempty"`
	Chat     *ReplyRecipient `json:"request_chat,omitempty"`
	WebApp   *WebApp         `json:"web_app,omitempty"`
}

// InlineButton represents a button displayed in the message.
type InlineButton struct {
	// Unique slagish name for this kind of button,
	// try to be as specific as possible.
	//
	// It will be used as a callback endpoint.
	Unique string `json:"unique,omitempty"`

	Text                  string             `json:"text"`
	URL                   string             `json:"url,omitempty"`
	Data                  string             `json:"callback_data,omitempty"`
	InlineQuery           string             `json:"switch_inline_query,omitempty"`
	InlineQueryChat       string             `json:"switch_inline_query_current_chat"`
	InlineQueryChosenChat *SwitchInlineQuery `json:"switch_inline_query_chosen_chat,omitempty"`
	Login                 *Login             `json:"login_url,omitempty"`
	WebApp                *WebApp            `json:"web_app,omitempty"`
	CallbackGame          *CallbackGame      `json:"callback_game,omitempty"`
	Pay                   bool               `json:"pay,omitempty"`
}

// ReplyRecipient combines both KeyboardButtonRequestUser
// and KeyboardButtonRequestChat objects. Use inside ReplyButton
// to request the user or chat sharing with respective settings.
//
// To pass the pointers to bool use a special tele.Flag function,
// that way you will be able to reflect the three-state bool (nil, false, true).
type ReplyRecipient struct {
	ID int32 `json:"request_id"`

	Bot      *bool `json:"user_is_bot,omitempty"`     // user only, optional
	Premium  *bool `json:"user_is_premium,omitempty"` // user only, optional
	Quantity int   `json:"max_quantity,omitempty"`    // user only, optional

	Channel         bool             `json:"chat_is_channel,omitempty"`           // chat only, required
	Forum           *bool            `json:"chat_is_forum,omitempty"`             // chat only, optional
	WithUsername    *bool            `json:"chat_has_username,omitempty"`         // chat only, optional
	Created         *bool            `json:"chat_is_created,omitempty"`           // chat only, optional
	UserRights      *ChatPermissions `json:"user_administrator_rights,omitempty"` // chat only, optional
	BotRights       *ChatPermissions `json:"bot_administrator_rights,omitempty"`  // chat only, optional
	BotMember       *bool            `json:"bot_is_member,omitempty"`             // chat only, optional
	RequestTitle    *bool            `json:"request_title,omitempty"`             // chat only, optional
	RequestName     *bool            `json:"request_name,omitempty"`              // user only, optional
	RequestUsername *bool            `json:"request_username,omitempty"`          // user only, optional
	RequestPhoto    *bool            `json:"request_photo,omitempty"`             // user only, optional
}

// Login represents a parameter of the inline keyboard button
// used to automatically authorize a user. Serves as a great replacement
// for the Telegram Login Widget when the user is coming from Telegram.
type Login struct {
	URL         string `json:"url"`
	Text        string `json:"forward_text,omitempty"`
	Username    string `json:"bot_username,omitempty"`
	WriteAccess bool   `json:"request_write_access,omitempty"`
}

// CallbackGame represents a placeholder for a game to be sent.
type CallbackGame struct {
}

// CallbackUnique returns ReplyButton.Text.
func (t *ReplyButton) CallbackUnique() string {
	return t.Text
}

// CallbackUnique returns InlineButton.Unique.
func (t *InlineButton) CallbackUnique() string {
	return "\f" + t.Unique
}

// MarshalJSON implements json.Marshaler interface.
// It needed to avoid InlineQueryChat and Login or WebApp fields conflict.
// If you have Login or WebApp field in your button, InlineQueryChat must be skipped.
func (t *InlineButton) MarshalJSON() ([]byte, error) {
	type IB InlineButton

	if t.Login != nil || t.WebApp != nil {
		return json.Marshal(struct {
			IB
			InlineQueryChat string `json:"switch_inline_query_current_chat,omitempty"`
		}{
			IB: IB(*t),
		})
	}
	return json.Marshal(IB(*t))
}

// With returns a copy of the button with data.
func (t *InlineButton) With(data string) *InlineButton {
	return &InlineButton{
		Unique:          t.Unique,
		Text:            t.Text,
		URL:             t.URL,
		InlineQuery:     t.InlineQuery,
		InlineQueryChat: t.InlineQueryChat,
		Login:           t.Login,
		Data:            data,
	}
}
