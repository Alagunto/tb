package telegram

// InlineButton represents a button displayed in the message.
type InlineButton struct {
	Text                  string                `json:"text"`
	URL                   string                `json:"url,omitempty"`
	CallbackData          string                `json:"callback_data,omitempty"`
	WebApp                *WebAppInfo           `json:"web_app,omitempty"`
	LoginURL              *LoginURL             `json:"login_url,omitempty"`
	SwitchInlineQuery       string                `json:"switch_inline_query,omitempty"`
	SwitchInlineQueryCurrentChat string         `json:"switch_inline_query_current_chat,omitempty"`
	SwitchInlineQueryChosenChat *SwitchInlineQuery `json:"switch_inline_query_chosen_chat,omitempty"`
	CallbackGame          *CallbackGame         `json:"callback_game,omitempty"`
	Pay                   bool                  `json:"pay,omitempty"`
}
