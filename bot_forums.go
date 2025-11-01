package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// CreateForumTopic creates a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_manage_topics administrator rights.
func (b *Bot[RequestType]) CreateForumTopic(ctx context.Context, chat bot.Recipient, name string, iconColor string, iconCustomEmojiID string) (*telegram.ForumTopic, error) {
	if chat == nil {
		return nil, errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

	p := params.New().
		Add("chat_id", chat.Recipient()).
		Add("name", name).
		Add("icon_color", iconColor).
		Add("icon_custom_emoji_id", iconCustomEmojiID).
		Build()

	r := NewApiRequester[map[string]any, telegram.ForumTopic](b.token, b.apiURL, b.client)
	result, err := r.Request(ctx, "createForumTopic", p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// CreateForumTopicBackground creates a topic in a forum supergroup chat using context.Background().
func (b *Bot[RequestType]) CreateForumTopicBackground(chat bot.Recipient, name string, iconColor string, iconCustomEmojiID string) (*telegram.ForumTopic, error) {
	return b.CreateForumTopic(context.Background(), chat, name, iconColor, iconCustomEmojiID)
}

// EditForumTopic edits name and icon of a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_manage_topics administrator rights, unless it is the creator of the topic.
func (b *Bot[RequestType]) EditForumTopic(ctx context.Context, chat bot.Recipient, threadID int, name string, iconCustomEmojiID string) error {
	if chat == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

	p := params.New().
		Add("chat_id", chat.Recipient()).
		AddInt("message_thread_id", threadID).
		Add("name", name).
		Add("icon_custom_emoji_id", iconCustomEmojiID).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(ctx, "editForumTopic", p)
	return err
}

// EditForumTopicBackground edits name and icon of a topic in a forum supergroup chat using context.Background().
func (b *Bot[RequestType]) EditForumTopicBackground(chat bot.Recipient, threadID int, name string, iconCustomEmojiID string) error {
	return b.EditForumTopic(context.Background(), chat, threadID, name, iconCustomEmojiID)
}

// CloseForumTopic closes an open topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_manage_topics administrator rights, unless it is the creator of the topic.
func (b *Bot[RequestType]) CloseForumTopic(ctx context.Context, chat bot.Recipient, threadID int) error {
	if chat == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

	p := params.New().
		Add("chat_id", chat.Recipient()).
		AddInt("message_thread_id", threadID).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(ctx, "closeForumTopic", p)
	return err
}

// CloseForumTopicBackground closes an open topic in a forum supergroup chat using context.Background().
func (b *Bot[RequestType]) CloseForumTopicBackground(chat bot.Recipient, threadID int) error {
	return b.CloseForumTopic(context.Background(), chat, threadID)
}

// ReopenForumTopic reopens a closed topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_manage_topics administrator rights, unless it is the creator of the topic.
func (b *Bot[RequestType]) ReopenForumTopic(ctx context.Context, chat bot.Recipient, threadID int) error {
	if chat == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

	p := params.New().
		Add("chat_id", chat.Recipient()).
		AddInt("message_thread_id", threadID).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(ctx, "reopenForumTopic", p)
	return err
}

// ReopenForumTopicBackground reopens a closed topic in a forum supergroup chat using context.Background().
func (b *Bot[RequestType]) ReopenForumTopicBackground(chat bot.Recipient, threadID int) error {
	return b.ReopenForumTopic(context.Background(), chat, threadID)
}

// DeleteForumTopic deletes a forum topic along with all its messages in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_delete_messages administrator rights.
func (b *Bot[RequestType]) DeleteForumTopic(ctx context.Context, chat bot.Recipient, threadID int) error {
	if chat == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

	p := params.New().
		Add("chat_id", chat.Recipient()).
		AddInt("message_thread_id", threadID).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(ctx, "deleteForumTopic", p)
	return err
}

// DeleteForumTopicBackground deletes a forum topic along with all its messages in a forum supergroup chat using context.Background().
func (b *Bot[RequestType]) DeleteForumTopicBackground(chat bot.Recipient, threadID int) error {
	return b.DeleteForumTopic(context.Background(), chat, threadID)
}

// UnpinAllForumTopicMessages clears the list of pinned messages in a forum topic.
// The bot must be an administrator in the chat for this to work and must have the
// can_pin_messages administrator right in the supergroup.
func (b *Bot[RequestType]) UnpinAllForumTopicMessages(ctx context.Context, chat bot.Recipient, threadID int) error {
	if chat == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

	p := params.New().
		Add("chat_id", chat.Recipient()).
		AddInt("message_thread_id", threadID).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(ctx, "unpinAllForumTopicMessages", p)
	return err
}

// UnpinAllForumTopicMessagesBackground clears the list of pinned messages in a forum topic using context.Background().
func (b *Bot[RequestType]) UnpinAllForumTopicMessagesBackground(chat bot.Recipient, threadID int) error {
	return b.UnpinAllForumTopicMessages(context.Background(), chat, threadID)
}
