package telegram

// Callback object represents a query from a callback button in an
// inline keyboard.
type Callback struct {
	ID string `json:"id"`

	// For message sent to channels, Sender may be empty
	Sender *User `json:"from"`

	// Message will be set if the button that originated the query
	// was attached to a message sent by a bot.
	Message *Message `json:"message"`

	// MessageID will be set if the button was attached to a message
	// sent via the bot in inline mode.
	MessageID string `json:"inline_message_id"`

	// Data associated with the callback button. Be aware that
	// a bad client can send arbitrary data in this field.
	Data string `json:"data"`

	// ChatInstance is a global identifier, uniquely corresponding to
	// the chat to which the message with the callback button was sent.
	ChatInstance string `json:"chat_instance"`

	// GameShortName is a unique identifier of the game for which a URL
	// is requested from the bot when a user presses the Play button of
	// that game. GameShortName may be empty
	GameShortName string `json:"game_short_name"`
}

// CallbackResponse builds a response to a CallbackQuery query.
type CallbackResponse struct {
	Text      string `json:"text,omitempty"`
	ShowAlert bool   `json:"show_alert,omitempty"`
	URL       string `json:"url,omitempty"`
	CacheTime int    `json:"cache_time,omitempty"`
}

// MessageSig satisfies Editable interface.
func (c *Callback) MessageSig() (string, int64) {
	if c.IsInline() {
		return c.MessageID, 0
	}
	if c.Message != nil {
		// MessageSig needs to be handled at a higher level
		return c.Message.MessageSig()
	}
	return "", 0
}

// IsInline says whether message is an inline message.
func (c *Callback) IsInline() bool {
	return c.MessageID != ""
}
