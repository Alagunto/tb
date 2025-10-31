package telegram

// MessageReactionCountUpdated represents reactions added to a message.
type MessageReactionCountUpdated struct {
	Chat      *Chat           `json:"chat"`
	MessageID int             `json:"message_id"`
	Date      int64           `json:"date"`
	Reactions []ReactionCount `json:"reactions"`
}
