package telegram

// KeyboardButtonRequestUser represents the user sharing settings.
type KeyboardButtonRequestUser struct {
	ID      int32 `json:"request_id"`
	Bot     *bool `json:"user_is_bot,omitempty"`
	Premium *bool `json:"user_is_premium,omitempty"`
}
