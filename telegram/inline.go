package telegram

// InlineQuery represents an incoming inline query. When the user sends an empty query,
// your bot could return some default or trending results.
type InlineQuery struct {
	ID     string    `json:"id"`
	Sender *User     `json:"from"`
	Query  string    `json:"query"`
	Offset string    `json:"offset"`
	Type   ChatType  `json:"chat_type,omitempty"`
	Loc    *Location `json:"location,omitempty"`
}

// ChosenInlineResult represents a result of an inline query that was chosen by a user and sent to their chat partner.
type ChosenInlineResult struct {
	ResultID        string    `json:"result_id"`
	Sender          *User     `json:"from"`
	Query           string    `json:"query"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Loc             *Location `json:"location,omitempty"`
}

// MessageSig satisfies Editable interface.
func (r *ChosenInlineResult) MessageSig() (string, int64) {
	return r.InlineMessageID, 0
}

// InlineQueryResponse builds a response to an inline query.
type InlineQueryResponse struct {
	Results           []InlineQueryResult     `json:"results"`
	CacheTime         int                     `json:"cache_time,omitempty"`
	IsPersonal        bool                    `json:"is_personal,omitempty"`
	NextOffset        string                  `json:"next_offset,omitempty"`
	Button            *InlineQueryResultsButton `json:"button,omitempty"`
	SwitchPMText      string                  `json:"switch_pm_text,omitempty"`
	SwitchPMParameter string                  `json:"switch_pm_parameter,omitempty"`
}

type InlineQueryResultsButton struct {
	Text           string      `json:"text"`
	WebApp         *WebAppInfo `json:"web_app,omitempty"`
	StartParameter string      `json:"start_parameter,omitempty"`
}
