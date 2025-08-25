package tb

import (
	"encoding/json"
	"strconv"
)

type Topic struct {
	Name              string `json:"name"`
	IconColor         int    `json:"icon_color"`
	IconCustomEmojiID string `json:"icon_custom_emoji_id"`
	ThreadID          int    `json:"message_thread_id"`
}

// CreateTopic creates a topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) CreateTopic(chat *Chat, topic *Topic) (*Topic, error) {
	params := map[string]string{
		"chat_id": chat.Recipient(),
		"name":    topic.Name,
	}

	if topic.IconColor != 0 {
		params["icon_color"] = strconv.Itoa(topic.IconColor)
	}
	if topic.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = topic.IconCustomEmojiID
	}

	data, err := b.Raw("createForumTopic", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result *Topic
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, err
}

// EditTopic edits name and icon of a topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) EditTopic(chat *Chat, topic *Topic) error {
	params := map[string]interface{}{
		"chat_id":           chat.Recipient(),
		"message_thread_id": topic.ThreadID,
	}

	if topic.Name != "" {
		params["name"] = topic.Name
	}
	if topic.IconCustomEmojiID != "" {
		params["icon_custom_emoji_id"] = topic.IconCustomEmojiID
	}

	_, err := b.Raw("editForumTopic", params)
	return err
}

// CloseTopic closes an open topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) CloseTopic(chat *Chat, topic *Topic) error {
	params := map[string]interface{}{
		"chat_id":           chat.Recipient(),
		"message_thread_id": topic.ThreadID,
	}

	_, err := b.Raw("closeForumTopic", params)
	return err
}

// ReopenTopic reopens a closed topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) ReopenTopic(chat *Chat, topic *Topic) error {
	params := map[string]interface{}{
		"chat_id":           chat.Recipient(),
		"message_thread_id": topic.ThreadID,
	}

	_, err := b.Raw("reopenForumTopic", params)
	return err
}

// DeleteTopic deletes a forum topic along with all its messages in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) DeleteTopic(chat *Chat, topic *Topic) error {
	params := map[string]interface{}{
		"chat_id":           chat.Recipient(),
		"message_thread_id": topic.ThreadID,
	}

	_, err := b.Raw("deleteForumTopic", params)
	return err
}

// UnpinAllTopicMessages clears the list of pinned messages in a forum topic. The bot must be an administrator in the chat for this to work and must have the can_pin_messages administrator right in the supergroup.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) UnpinAllTopicMessages(chat *Chat, topic *Topic) error {
	params := map[string]interface{}{
		"chat_id":           chat.Recipient(),
		"message_thread_id": topic.ThreadID,
	}

	_, err := b.Raw("unpinAllForumTopicMessages", params)
	return err
}

// TopicIconStickers gets custom emoji stickers, which can be used as a forum topic icon by any user.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) TopicIconStickers() ([]Sticker, error) {
	params := map[string]string{}

	data, err := b.Raw("getForumTopicIconStickers", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []Sticker
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, wrapError(err)
	}
	return resp.Result, nil
}

// EditGeneralTopic edits name of the 'General' topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) EditGeneralTopic(chat *Chat, topic *Topic) error {
	params := map[string]interface{}{
		"chat_id": chat.Recipient(),
		"name":    topic.Name,
	}

	_, err := b.Raw("editGeneralForumTopic", params)
	return err
}

// CloseGeneralTopic closes an open 'General' topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) CloseGeneralTopic(chat *Chat) error {
	params := map[string]interface{}{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("closeGeneralForumTopic", params)
	return err
}

// ReopenGeneralTopic reopens a closed 'General' topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) ReopenGeneralTopic(chat *Chat) error {
	params := map[string]interface{}{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("reopenGeneralForumTopic", params)
	return err
}

// HideGeneralTopic hides the 'General' topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) HideGeneralTopic(chat *Chat) error {
	params := map[string]interface{}{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("hideGeneralForumTopic", params)
	return err
}

// UnhideGeneralTopic unhides the 'General' topic in a forum supergroup chat.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) UnhideGeneralTopic(chat *Chat) error {
	params := map[string]interface{}{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("unhideGeneralForumTopic", params)
	return err
}

// UnpinAllGeneralTopicMessages clears the list of pinned messages in a General forum topic.
// The bot must be an administrator in the chat for this to work and must have the
// can_pin_messages administrator right in the supergroup.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) UnpinAllGeneralTopicMessages(chat *Chat) error {
	params := map[string]interface{}{
		"chat_id": chat.Recipient(),
	}

	_, err := b.Raw("unpinAllGeneralForumTopicMessages", params)
	return err
}
