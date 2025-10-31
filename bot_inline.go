package tb

import (
	"context"

	"github.com/alagunto/tb/telegram"
)

// AnswerInlineQuery sends answers to an inline query.
// Results parameter is required and must be an array of inline query results.
func (b *Bot[RequestType]) AnswerInlineQuery(query *telegram.InlineQuery, response *telegram.InlineQueryResponse) error {
	params := make(map[string]any)
	params["inline_query_id"] = query.ID
	params["results"] = response.Results

	if response.CacheTime > 0 {
		params["cache_time"] = response.CacheTime
	}
	if response.IsPersonal {
		params["is_personal"] = response.IsPersonal
	}
	if response.NextOffset != "" {
		params["next_offset"] = response.NextOffset
	}
	if response.Button != nil {
		params["button"] = response.Button
	}
	if response.SwitchPMText != "" {
		params["switch_pm_text"] = response.SwitchPMText
	}
	if response.SwitchPMParameter != "" {
		params["switch_pm_parameter"] = response.SwitchPMParameter
	}

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "answerInlineQuery", params)
	return err
}

