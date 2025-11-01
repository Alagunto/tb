package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// SendPoll sends a native poll. A native poll can't be sent to a private chat.
func (b *Bot[RequestType]) SendPoll(ctx context.Context, to bot.Recipient, poll *telegram.SendPollParams, opts ...params.SendOptions) (*telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrTelebot, "recipient", nil)
	}
	if poll == nil {
		return nil, errors.WithInvalidParam(errors.ErrTelebot, "poll", nil)
	}

	sendOpts := params.Merge(opts...)

	p := params.New().
		Add("chat_id", to.Recipient()).
		Add("question", poll.Question).
		Add("options", poll.Options).
		AddBool("is_anonymous", poll.IsAnonymous).
		Add("type", poll.Type).
		AddBool("allows_multiple_answers", poll.AllowsMultipleAnswers).
		Add("explanation", poll.Explanation).
		Add("explanation_parse_mode", poll.ExplanationParseMode).
		Add("explanation_entities", poll.ExplanationEntities)

	if poll.CorrectOptionID != nil {
		p.AddInt("correct_option_id", *poll.CorrectOptionID)
	}
	if poll.OpenPeriod != nil {
		p.AddInt("open_period", *poll.OpenPeriod)
	}
	if poll.CloseDate != nil {
		p.AddInt64("close_date", *poll.CloseDate)
	}

	p.AddBool("is_closed", poll.IsClosed).
		With(sendOpts)

	r := NewApiRequester[map[string]any, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(ctx, "sendPoll", p.Build())
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SendPollBackground sends a native poll using context.Background(). A native poll can't be sent to a private chat.
func (b *Bot[RequestType]) SendPollBackground(to bot.Recipient, poll *telegram.SendPollParams, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.SendPoll(context.Background(), to, poll, opts...)
}
