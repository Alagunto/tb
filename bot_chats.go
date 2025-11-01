package tb

import (
	"context"
	"fmt"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// Me returns the bot's user information.
func (b *Bot[RequestType]) Me() *telegram.User {
	return b.me
}

// ChatByID fetches a chat by its ID or username.
func (b *Bot[RequestType]) ChatByID(id int64) (*telegram.Chat, error) {
	if id == 0 {
		return nil, errors.WithInvalidParam(errors.ErrTelebot, "chat_id", nil)
	}

	p := params.New().
		Add("chat_id", id).
		Build()

	r := NewApiRequester[map[string]any, telegram.Chat](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getChat", p)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat: %w", err)
	}
	return result, nil
}

// ChatByIDBackground fetches a chat by its ID or username using background context.
func (b *Bot[RequestType]) ChatByIDBackground(id int64) (*telegram.Chat, error) {
	return b.ChatByID(id)
}

// ChatByUsername fetches a chat by its username.
func (b *Bot[RequestType]) ChatByUsername(username string) (*telegram.Chat, error) {
	if username == "" {
		return nil, errors.WithInvalidParam(errors.ErrTelebot, "username", nil)
	}

	p := params.New().
		Add("chat_id", username).
		Build()

	r := NewApiRequester[map[string]any, telegram.Chat](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getChat", p)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat: %w", err)
	}
	return result, nil
}

// ChatByUsernameBackground fetches a chat by its username using background context.
func (b *Bot[RequestType]) ChatByUsernameBackground(username string) (*telegram.Chat, error) {
	return b.ChatByUsername(username)
}

// GetWebhookInfo returns current webhook status.
func (b *Bot[RequestType]) GetWebhookInfo(ctx context.Context) (*telegram.WebhookInfo, error) {
	r := NewApiRequester[map[string]any, telegram.WebhookInfo](b.token, b.apiURL, b.client)
	result, err := r.Request(ctx, "getWebhookInfo", make(map[string]any))
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetWebhookInfoBackground returns current webhook status using background context.
func (b *Bot[RequestType]) GetWebhookInfoBackground() (*telegram.WebhookInfo, error) {
	return b.GetWebhookInfo(context.Background())
}

// DeleteWebhook removes webhook integration.
func (b *Bot[RequestType]) DeleteWebhook(ctx context.Context, dropPendingUpdates bool) error {
	p := params.New().
		AddBool("drop_pending_updates", dropPendingUpdates).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(ctx, "deleteWebhook", p)
	return err
}

// DeleteWebhookBackground removes webhook integration using background context.
func (b *Bot[RequestType]) DeleteWebhookBackground(dropPendingUpdates bool) error {
	return b.DeleteWebhook(context.Background(), dropPendingUpdates)
}

// ChatMemberOf returns information about a member of a chat.
func (b *Bot[RequestType]) ChatMemberOf(chat, user bot.Recipient) (*telegram.ChatMember, error) {
	if chat == nil || user == nil {
		return nil, errors.WithInvalidParam(errors.ErrTelebot, "chat or user", nil)
	}

	p := params.New().
		Add("chat_id", chat.Recipient()).
		Add("user_id", user.Recipient()).
		Build()

	r := NewApiRequester[map[string]any, telegram.ChatMember](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getChatMember", p)
	if err != nil {
		return nil, fmt.Errorf("failed to get chat member: %w", err)
	}
	return result, nil
}
