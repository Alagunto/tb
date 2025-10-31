package telegram

// Poll contains information about a poll.
type Poll struct {
	ID                string          `json:"id"`
	Type              PollType        `json:"type"`
	Question          string          `json:"question"`
	Options           []PollOption    `json:"options"`
	VoterCount        int             `json:"total_voter_count"`
	IsClosed          bool            `json:"is_closed,omitempty"`
	CorrectOptionID   int             `json:"correct_option_id,omitempty"`
	AllowsMultipleAnswers bool        `json:"allows_multiple_answers,omitempty"`
	Explanation       string          `json:"explanation,omitempty"`
	ExplanationParseMode ParseMode    `json:"explanation_parse_mode,omitempty"`
	ExplanationEntities []MessageEntity `json:"explanation_entities,omitempty"`
	QuestionParseMode string          `json:"question_parse_mode,omitempty"`
	QuestionEntities  []MessageEntity `json:"question_entities,omitempty"`
	IsAnonymous       bool            `json:"is_anonymous"`
	OpenPeriod        int             `json:"open_period,omitempty"`
	CloseDate         int64           `json:"close_date,omitempty"`
}
