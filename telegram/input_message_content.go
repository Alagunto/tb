package telegram

// InputMessageContent represents the content of a message to be sent as a result of an inline query.
type InputMessageContent struct {
	// MessageText is the text of the message to be sent, 1-4096 characters.
	MessageText string `json:"message_text"`

	// ParseMode is the mode for parsing entities in the message text.
	ParseMode string `json:"parse_mode,omitempty"`

	// Entities is a list of special entities that appear in message text, which can be specified instead of parse_mode.
	Entities []MessageEntity `json:"entities,omitempty"`

	// DisableWebPagePreview disables link previews for links in this message.
	DisableWebPagePreview bool `json:"disable_web_page_preview,omitempty"`
}

