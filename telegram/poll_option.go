package telegram

// PollOption contains information about one answer option in a poll.
type PollOption struct {
	Text       string          `json:"text"`
	VoterCount int             `json:"voter_count"`
	ParseMode  ParseMode       `json:"text_parse_mode,omitempty"`
	Entities   []MessageEntity `json:"text_entities,omitempty"`
}
