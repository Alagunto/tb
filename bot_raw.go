package tb

import (
	"context"
	"encoding/json"
	"time"

	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/telegram"
)

// Raw lets you call any method of Bot API manually.
// This method is for external use only - bot.go methods should use ApiRequester directly.
func (b *Bot[RequestType]) Raw(method string, payload interface{}) ([]byte, error) {
	// Convert payload to map[string]any if it isn't already
	var params map[string]any
	switch p := payload.(type) {
	case map[string]any:
		params = p
	case map[string]string:
		// Convert map[string]string to map[string]any
		params = make(map[string]any, len(p))
		for k, v := range p {
			params[k] = v
		}
	case nil:
		params = make(map[string]any)
	default:
		// Try to marshal and unmarshal to convert to map
		data, err := json.Marshal(payload)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		if err := json.Unmarshal(data, &params); err != nil {
			return nil, errors.Wrap(err)
		}
	}

	r := NewApiRequester[map[string]any, json.RawMessage](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), method, params)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return *result, nil
}

// RawBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) RawBackground(method string, payload any) ([]byte, error) {
	return b.Raw(method, payload)
}

// Internal helper methods that use ApiRequester with proper types

func (b *Bot[RequestType]) GetMe() (*telegram.User, error) {
	r := NewApiRequester[map[string]any, telegram.User](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getMe", make(map[string]any))
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return result, nil
}

// GetUpdates fetches updates from the Telegram API.
// Do not use this method directly by default, instead use Start() to start the Poller to fetch updates automatically.
// Use it only if you need to fetch updates manually, without starting the bot as usual.
func (b *Bot[RequestType]) GetUpdates(offset, limit int, timeout time.Duration, allowed []string) ([]telegram.Update, error) {
	params := map[string]any{
		"offset":  offset,
		"timeout": int(timeout / time.Second),
	}

	if len(allowed) > 0 {
		params["allowed_updates"] = allowed
	}

	if limit != 0 {
		params["limit"] = limit
	}

	r := NewApiRequester[map[string]any, []telegram.Update](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getUpdates", params)
	if err != nil {
		return nil, errors.Wrap(err)
	}
	return *result, nil
}

// GetUpdatesBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) GetUpdatesBackground(offset, limit int, timeout time.Duration, allowed []string) ([]telegram.Update, error) {
	return b.GetUpdates(offset, limit, timeout, allowed)
}
