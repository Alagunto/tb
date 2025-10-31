package telegram

// KeyboardButtonPollType represents a poll type that can be requested from the user.
type KeyboardButtonPollType struct {
	Type PollType `json:"type,omitempty"`
}
