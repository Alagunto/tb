package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// SendPoll sends a native poll. A native poll can't be sent to a private chat.
func (b *Bot[RequestType]) SendPoll(to bot.Recipient, poll *telegram.SendPollParams, opts ...params.SendOptions) (*telegram.Message, error) {
	sendOpts := params.Merge(opts...)
	paramsMap := sendOpts.ToMap()
	paramsMap["chat_id"] = to.Recipient()

	// Copy poll parameters
	paramsMap["question"] = poll.Question
	paramsMap["options"] = poll.Options

	if poll.IsAnonymous {
		paramsMap["is_anonymous"] = poll.IsAnonymous
	}
	if poll.Type != "" {
		paramsMap["type"] = poll.Type
	}
	if poll.AllowsMultipleAnswers {
		paramsMap["allows_multiple_answers"] = poll.AllowsMultipleAnswers
	}
	if poll.CorrectOptionID != nil {
		paramsMap["correct_option_id"] = *poll.CorrectOptionID
	}
	if poll.Explanation != "" {
		paramsMap["explanation"] = poll.Explanation
	}
	if poll.ExplanationParseMode != "" {
		paramsMap["explanation_parse_mode"] = poll.ExplanationParseMode
	}
	if poll.ExplanationEntities != nil {
		paramsMap["explanation_entities"] = poll.ExplanationEntities
	}
	if poll.OpenPeriod != nil {
		paramsMap["open_period"] = *poll.OpenPeriod
	}
	if poll.CloseDate != nil {
		paramsMap["close_date"] = *poll.CloseDate
	}
	if poll.IsClosed {
		paramsMap["is_closed"] = poll.IsClosed
	}

	r := NewApiRequester[map[string]any, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "sendPoll", paramsMap)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StopPoll stops a poll which was sent by the bot and returns the stopped poll.
func (b *Bot[RequestType]) StopPoll(msg bot.Editable, opts ...params.SendOptions) (*telegram.Poll, error) {
	msgID, chatID := msg.MessageSig()

	sendOpts := params.Merge(opts...)
	paramsMap := sendOpts.ToMap()
	paramsMap["chat_id"] = chatID
	paramsMap["message_id"] = msgID

	r := NewApiRequester[map[string]any, telegram.Poll](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "stopPoll", paramsMap)
	if err != nil {
		return nil, err
	}
	return result, nil
}
