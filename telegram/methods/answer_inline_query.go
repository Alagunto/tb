package methods

import "github.com/alagunto/tb/telegram"

// AnswerInlineQueryRequest represents the request for answerInlineQuery method.
type AnswerInlineQueryRequest struct {
	// Unique identifier for the answered query
	InlineQueryID string `json:"inline_query_id"`

	// Array of results for the inline query
	Results []telegram.Result `json:"results,omitempty"`

	// The maximum amount of time in seconds that the result of the inline query may be cached on the server
	CacheTime int `json:"cache_time,omitempty"`

	// Pass True if results may be cached on the server side only for the user that sent the query
	IsPersonal bool `json:"is_personal,omitempty"`

	// Pass the offset that a client should send in the next query with the same text to receive more results
	NextOffset string `json:"next_offset,omitempty"`

	// A button to be shown above inline query results
	Button interface{} `json:"button,omitempty"`
}

// AnswerInlineQueryResponse represents the response for answerInlineQuery method.
type AnswerInlineQueryResponse bool
