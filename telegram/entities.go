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
	CustomEmojiID string `json:"custom_emoji_id"`
}

// EntityType is a MessageEntity type.
type EntityType string

const (
	EntityMention       EntityType = "mention"
	EntityTMention      EntityType = "text_mention"
	EntityHashtag       EntityType = "hashtag"
	EntityCashtag       EntityType = "cashtag"
	EntityCommand       EntityType = "bot_command"
	EntityURL           EntityType = "url"
	EntityEmail         EntityType = "email"
	EntityPhone         EntityType = "phone_number"
	EntityBold          EntityType = "bold"
	EntityItalic        EntityType = "italic"
	EntityUnderline     EntityType = "underline"
	EntityStrikethrough EntityType = "strikethrough"
	EntityCode          EntityType = "code"
	EntityCodeBlock     EntityType = "pre"
	EntityTextLink      EntityType = "text_link"
	EntitySpoiler       EntityType = "spoiler"
	EntityCustomEmoji   EntityType = "custom_emoji"
	EntityBlockquote    EntityType = "blockquote"
	EntityEBlockquote   EntityType = "expandable_blockquote"
)

// Entities are used to set message's text entities as a send option.
type Entities []MessageEntity

// PreviewOptions used for link preview generation.
type PreviewOptions struct {
	// (Optional) True, if the link preview is disabled.
	Disabled bool `json:"is_disabled"`

	// (Optional) URL to use for the link preview. If empty, then the first URL found in the message text will be used.
	URL string `json:"url,omitempty"`

	// (Optional) True, if the media in the link preview is supposed to be shrunk; ignored if the URL isn't explicitly specified or media type is not supported.
	PreferSmallMedia bool `json:"prefer_small_media,omitempty"`

	// (Optional) True, if the media in the link preview is supposed to be enlarged; ignored if the URL isn't explicitly specified or media type is not supported.
	PreferLargeMedia bool `json:"prefer_large_media,omitempty"`

	// (Optional) True, if there is no need to show an animation in the link preview.
	ShowAboveText bool `json:"show_above_text,omitempty"`
}

// ParseMode defines how parse mode should be applied.
type ParseMode string

const (
	ModeDefault    ParseMode = ""
	ModeMarkdown   ParseMode = "Markdown"
	ModeMarkdownV2 ParseMode = "MarkdownV2"
	ModeHTML       ParseMode = "HTML"
)
