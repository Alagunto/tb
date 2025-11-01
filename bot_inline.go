package tb

import (
	"context"

	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// AnswerInlineQuery sends answers to an inline query.
// Results parameter is required and must be an array of inline query results.
func (b *Bot[RequestType]) AnswerInlineQuery(query *telegram.InlineQuery, response *telegram.InlineQueryResponse) error {
	if query == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "inline_query", nil)
	}
	if response == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "response", nil)
	}

	p := params.New().
		Add("inline_query_id", query.ID).
		Add("results", response.Results).
		AddInt("cache_time", response.CacheTime).
		AddBool("is_personal", response.IsPersonal).
		Add("next_offset", response.NextOffset).
		Add("button", response.Button).
		Add("switch_pm_text", response.SwitchPMText).
		Add("switch_pm_parameter", response.SwitchPMParameter).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "answerInlineQuery", p)
	return err
}

// AnswerInlineQueryBackground sends answers to an inline query using context.Background().
func (b *Bot[RequestType]) AnswerInlineQueryBackground(query *telegram.InlineQuery, response *telegram.InlineQueryResponse) error {
	return b.AnswerInlineQuery(query, response)
}

