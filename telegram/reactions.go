package telegram

// Reactions represents a list of reactions to set on a message.
type Reactions struct {
	Reactions []ReactionType `json:"reaction"`
	Big       bool           `json:"is_big,omitempty"`
}

