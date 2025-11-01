package telegram

// MessageEntity object represents "special" parts of text messages,
// including hashtags, usernames, URLs, etc.
type MessageEntity struct {
	// Specifies entity type.
	Type EntityType `json:"type"`

	// Offset in UTF-16 code units to the start of the entity.
	Offset int `json:"offset"`

	// Length of the entity in UTF-16 code units.
	Length int `json:"length"`

	// (Optional) For EntityTextLink entity type only.
	//
	// URL will be opened after user taps on the text.
	URL string `json:"url,omitempty"`

	// (Optional) For EntityTMention entity type only.
	User *User `json:"user,omitempty"`

	// (Optional) For EntityCodeBlock entity type only.
	Language string `json:"language,omitempty"`

	// (Optional) For EntityCustomEmoji entity type only.
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
}

// Entities are used to set message's text entities as a send option.
type Entities []MessageEntity
