package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// Pin pins a message in a supergroup or a channel.
//
// It supports Silent option.
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) Pin(msg bot.Editable) error {
	msgID, chatID := msg.MessageSig()

	req := telegram.PinMessageRequest{
		ChatID:    chatID,
		MessageID: msgID,
	}

	r := NewApiRequester[telegram.PinMessageRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "pinChatMessage", req)
	return err
}

// Unpin unpins a message in a supergroup or a channel.
// It supports tb.Silent option.
func (b *Bot[RequestType]) Unpin(chat bot.Recipient, messageID ...int) error {
	req := telegram.UnpinMessageRequest{
		ChatID: chat.Recipient(),
	}

	if len(messageID) > 0 {
		req.MessageID = messageID[0]
	}

	r := NewApiRequester[telegram.UnpinMessageRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "unpinChatMessage", req)
	return err
}

// UnpinAll unpins all messages in a supergroup or a channel.
// It supports tb.Silent option.
func (b *Bot[RequestType]) UnpinAll(chat bot.Recipient) error {
	req := telegram.UnpinAllMessageRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[telegram.UnpinAllMessageRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "unpinAllChatMessages", req)
	return err
}

// Notify updates the chat action for recipient.
//
// Chat action is a status message that recipient would see where
// you typically see "Harry is typing" status message. The only
// difference is that bots' chat actions live only for 5 seconds
// and die just once the client receives a message from the bot.
//
// Currently, Telegram supports only a narrow range of possible
// actions, these are aligned as constants of this package.
func (b *Bot[RequestType]) Notify(to bot.Recipient, action telegram.ChatAction, opts ...params.SendOptions) error {
	sendOpts := params.Merge(opts...)

	req := telegram.SendChatActionRequest{
		ChatID: to.Recipient(),
		Action: string(action),
	}

	if sendOpts.ThreadID != 0 {
		req.MessageThreadID = sendOpts.ThreadID
	}
	if sendOpts.BusinessConnectionID != "" {
		req.BusinessConnectionID = sendOpts.BusinessConnectionID
	}

	r := NewApiRequester[telegram.SendChatActionRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "sendChatAction", req)
	return err
}

// Leave makes bot leave a group, supergroup or channel.
func (b *Bot[RequestType]) Leave(chat bot.Recipient) error {
	req := telegram.LeaveChatRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[telegram.LeaveChatRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "leaveChat", req)
	return err
}

// MenuButton returns the current value of the bot's menu button in a private chat,
// or the default menu button.
func (b *Bot[RequestType]) MenuButton(chat *telegram.User) (*telegram.MenuButton, error) {
	req := telegram.GetChatMenuButtonRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[telegram.GetChatMenuButtonRequest, telegram.MenuButton](b.token, b.apiURL, b.client)
	return r.Request(context.Background(), "getChatMenuButton", req)
}

// SetMenuButton changes the bot's menu button in a private chat,
// or the default menu button.
//
// It accepts two kinds of menu button arguments:
//
//   - MenuButtonType for simple menu buttons (default, commands)
//   - MenuButton complete structure for web_app menu button type
func (b *Bot[RequestType]) SetMenuButton(chat *telegram.User, mb interface{}) error {
	req := telegram.SetChatMenuButtonRequest{}
	// chat_id is optional
	if chat != nil {
		req.ChatID = chat.Recipient()
	}

	switch v := mb.(type) {
	case telegram.MenuButtonType:
		req.MenuButton = &telegram.MenuButton{Type: v}
	case *telegram.MenuButton:
		req.MenuButton = v
	}

	r := NewApiRequester[telegram.SetChatMenuButtonRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setChatMenuButton", req)
	return err
}

// Logout logs out from the cloud Bot API server before launching the bot locally.
func (b *Bot[RequestType]) Logout() (bool, error) {
	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "logOut", make(map[string]any))
	if err != nil {
		return false, err
	}
	return *result, nil
}

// Close closes the bot instance before moving it from one local server to another.
func (b *Bot[RequestType]) Close() (bool, error) {
	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "close", make(map[string]any))
	if err != nil {
		return false, err
	}
	return *result, nil
}

// SetMyName change's the bot name.
func (b *Bot[RequestType]) SetMyName(name, language string) error {
	req := telegram.SetMyNameRequest{
		Name:         name,
		LanguageCode: language,
	}

	r := NewApiRequester[telegram.SetMyNameRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setMyName", req)
	return err
}

// MyName returns the current bot name for the given user language.
func (b *Bot[RequestType]) MyName(language string) (*telegram.BotName, error) {
	req := telegram.GetMyNameRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[telegram.GetMyNameRequest, telegram.BotName](b.token, b.apiURL, b.client)
	return r.Request(context.Background(), "getMyName", req)
}

// SetMyDescription change's the bot description, which is shown in the chat
// with the bot if the chat is empty.
func (b *Bot[RequestType]) SetMyDescription(desc, language string) error {
	req := telegram.SetMyDescriptionRequest{
		Description:  desc,
		LanguageCode: language,
	}

	r := NewApiRequester[telegram.SetMyDescriptionRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setMyDescription", req)
	return err
}

// MyDescription the current bot description for the given user language.
func (b *Bot[RequestType]) MyDescription(language string) (*telegram.BotDescription, error) {
	req := telegram.GetMyDescriptionRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[telegram.GetMyDescriptionRequest, telegram.BotDescription](b.token, b.apiURL, b.client)
	return r.Request(context.Background(), "getMyDescription", req)
}

// SetMyShortDescription change's the bot short description, which is shown on
// the bot's profile page and is sent together with the link when users share the bot.
func (b *Bot[RequestType]) SetMyShortDescription(desc, language string) error {
	req := telegram.SetMyShortDescriptionRequest{
		ShortDescription: desc,
		LanguageCode:     language,
	}

	r := NewApiRequester[telegram.SetMyShortDescriptionRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setMyShortDescription", req)
	return err
}

// MyShortDescription the current bot short description for the given user language.
func (b *Bot[RequestType]) MyShortDescription(language string) (*telegram.BotShortDescription, error) {
	req := telegram.GetMyShortDescriptionRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[telegram.GetMyShortDescriptionRequest, telegram.BotShortDescription](b.token, b.apiURL, b.client)
	return r.Request(context.Background(), "getMyShortDescription", req)
}
