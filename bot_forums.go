package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/telegram"
)

// CreateForumTopic creates a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_manage_topics administrator rights.
func (b *Bot[RequestType]) CreateForumTopic(chat bot.Recipient, name string, iconColor string, iconCustomEmojiID string) (*telegram.ForumTopic, error) {
	params := make(map[string]any)
	params["chat_id"] = chat.Recipient()
	params["name"] = name

	if iconColor != "" {
		params["icon_color"] = iconColor
	}
	if iconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = iconCustomEmojiID
	}

	r := NewApiRequester[map[string]any, telegram.ForumTopic](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "createForumTopic", params)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// EditForumTopic edits name and icon of a topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_manage_topics administrator rights, unless it is the creator of the topic.
func (b *Bot[RequestType]) EditForumTopic(chat bot.Recipient, threadID int, name string, iconCustomEmojiID string) error {
	params := make(map[string]any)
	params["chat_id"] = chat.Recipient()
	params["message_thread_id"] = threadID

	if name != "" {
		params["name"] = name
	}
	if iconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = iconCustomEmojiID
	}

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "editForumTopic", params)
	return err
}

// CloseForumTopic closes an open topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_manage_topics administrator rights, unless it is the creator of the topic.
func (b *Bot[RequestType]) CloseForumTopic(chat bot.Recipient, threadID int) error {
	params := make(map[string]any)
	params["chat_id"] = chat.Recipient()
	params["message_thread_id"] = threadID

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "closeForumTopic", params)
	return err
}

// ReopenForumTopic reopens a closed topic in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_manage_topics administrator rights, unless it is the creator of the topic.
func (b *Bot[RequestType]) ReopenForumTopic(chat bot.Recipient, threadID int) error {
	params := make(map[string]any)
	params["chat_id"] = chat.Recipient()
	params["message_thread_id"] = threadID

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "reopenForumTopic", params)
	return err
}

// DeleteForumTopic deletes a forum topic along with all its messages in a forum supergroup chat.
// The bot must be an administrator in the chat for this to work and must have the
// can_delete_messages administrator rights.
func (b *Bot[RequestType]) DeleteForumTopic(chat bot.Recipient, threadID int) error {
	params := make(map[string]any)
	params["chat_id"] = chat.Recipient()
	params["message_thread_id"] = threadID

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "deleteForumTopic", params)
	return err
}

// UnpinAllForumTopicMessages clears the list of pinned messages in a forum topic.
// The bot must be an administrator in the chat for this to work and must have the
// can_pin_messages administrator right in the supergroup.
func (b *Bot[RequestType]) UnpinAllForumTopicMessages(chat bot.Recipient, threadID int) error {
	params := make(map[string]any)
	params["chat_id"] = chat.Recipient()
	params["message_thread_id"] = threadID

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "unpinAllForumTopicMessages", params)
	return err
}
