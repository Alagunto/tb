package tb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// Me returns the bot's user information.
func (b *Bot[RequestType]) Me() *telegram.User {
	return b.me
}

// ChatByID fetches a chat by its ID or username.
func (b *Bot[RequestType]) ChatByID(chatID string) (*telegram.Chat, error) {
	// Try to parse as int64 first
	var chatIDValue any
	if id, err := strconv.ParseInt(chatID, 10, 64); err == nil {
		chatIDValue = id
	} else {
		// Treat as username
		chatIDValue = chatID
	}

	p := params.New().
		Add("chat_id", chatIDValue).
		Build()

	r := NewApiRequester[map[string]any, telegram.Chat](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getChat", p)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat: %w", err)
	}
	return result, nil
}

// ChatByUsername fetches a chat by its username.
func (b *Bot[RequestType]) ChatByUsername(username string) (*telegram.Chat, error) {
	return b.ChatByID(username)
}

// GetWebhookInfo returns current webhook status.
func (b *Bot[RequestType]) GetWebhookInfo() (*telegram.WebhookInfo, error) {
	r := NewApiRequester[map[string]any, telegram.WebhookInfo](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getWebhookInfo", make(map[string]any))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteWebhook removes webhook integration.
func (b *Bot[RequestType]) DeleteWebhook(dropPendingUpdates bool) error {
	p := params.New().
		AddBool("drop_pending_updates", dropPendingUpdates).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "deleteWebhook", p)
	return err
}
