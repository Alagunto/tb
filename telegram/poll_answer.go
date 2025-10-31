package telegram

// PollAnswer represents an answer of a user in a non-anonymous poll.
type PollAnswer struct {
	PollID  string `json:"poll_id"`
	Sender  *User  `json:"user"`
	Chat    *Chat  `json:"voter_chat"`
	Options []int  `json:"option_ids"`
}
