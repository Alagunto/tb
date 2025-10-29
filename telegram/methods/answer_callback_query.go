package methods

// AnswerCallbackQueryRequest represents the request for answerCallbackQuery method.
type AnswerCallbackQueryRequest struct {
	// Unique identifier for the query to be answered
	CallbackQueryID string `json:"callback_query_id"`

	// Text of the notification. If not specified, nothing will be shown to the user, 0-200 characters
	Text string `json:"text,omitempty"`

	// If True, an alert will be shown by the client instead of a notification at the top of the chat screen. Defaults to false.
	ShowAlert bool `json:"show_alert,omitempty"`

	// URL that will be opened by the user's client
	URL string `json:"url,omitempty"`

	// The maximum amount of time in seconds that the result of the callback query may be cached client-side
	CacheTime int `json:"cache_time,omitempty"`
}

// AnswerCallbackQueryResponse represents the response for answerCallbackQuery method.
type AnswerCallbackQueryResponse bool
