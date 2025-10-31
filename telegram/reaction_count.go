package telegram

// ReactionCount represents a reaction count.
type ReactionCount struct {
	Type          string `json:"type"`
	Emoji         string `json:"emoji,omitempty"`
	CustomEmojiID string `json:"custom_emoji_id,omitempty"`
	TotalCount    int    `json:"total_count"`
}
