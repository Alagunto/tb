package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
	"github.com/alagunto/tb/telegram/methods"
)

// Pin pins a message in a supergroup or a channel.
//
// It supports Silent option.
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) Pin(msg bot.Editable) error {
	msgID, chatID := msg.MessageSig()

	req := methods.PinChatMessageRequest{
		ChatID:    chatID,
		MessageID: msgID,
	}

	r := NewApiRequester[methods.PinChatMessageRequest, methods.PinChatMessageResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "pinChatMessage", req)
	return err
}

// Unpin unpins a message in a supergroup or a channel.
// It supports tb.Silent option.
func (b *Bot[RequestType]) Unpin(chat bot.Recipient, messageID ...int) error {
	req := methods.UnpinChatMessageRequest{
		ChatID: chat.Recipient(),
	}

	if len(messageID) > 0 {
		req.MessageID = messageID[0]
	}

	r := NewApiRequester[methods.UnpinChatMessageRequest, methods.UnpinChatMessageResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "unpinChatMessage", req)
	return err
}

// UnpinAll unpins all messages in a supergroup or a channel.
// It supports tb.Silent option.
func (b *Bot[RequestType]) UnpinAll(chat bot.Recipient) error {
	req := methods.UnpinAllChatMessagesRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[methods.UnpinAllChatMessagesRequest, methods.UnpinAllChatMessagesResponse](b.token, b.apiURL, b.client)
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
func (b *Bot[RequestType]) Notify(to bot.Recipient, action telegram.ChatAction, opts ...communications.SendOptions) error {
	sendOpts := params.Merge(opts...)

	req := methods.SendChatActionRequest{
		ChatID: to.Recipient(),
		Action: string(action),
	}

	if sendOpts.ThreadID != 0 {
		req.MessageThreadID = sendOpts.ThreadID
	}
	if sendOpts.BusinessConnectionID != "" {
		req.BusinessConnectionID = sendOpts.BusinessConnectionID
	}

	r := NewApiRequester[methods.SendChatActionRequest, methods.SendChatActionResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "sendChatAction", req)
	return err
}

// Leave makes bot leave a group, supergroup or channel.
func (b *Bot[RequestType]) Leave(chat bot.Recipient) error {
	req := methods.LeaveChatRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[methods.LeaveChatRequest, methods.LeaveChatResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "leaveChat", req)
	return err
}

// MenuButton returns the current value of the bot's menu button in a private chat,
// or the default menu button.
func (b *Bot[RequestType]) MenuButton(chat *telegram.User) (*MenuButton, error) {
	req := methods.GetChatMenuButtonRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[methods.GetChatMenuButtonRequest, methods.GetChatMenuButtonResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getChatMenuButton", req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SetMenuButton changes the bot's menu button in a private chat,
// or the default menu button.
//
// It accepts two kinds of menu button arguments:
//
//   - MenuButtonType for simple menu buttons (default, commands)
//   - MenuButton complete structure for web_app menu button type
func (b *Bot[RequestType]) SetMenuButton(chat *telegram.User, mb interface{}) error {
	req := methods.SetChatMenuButtonRequest{}

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

	r := NewApiRequester[methods.SetChatMenuButtonRequest, methods.SetChatMenuButtonResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setChatMenuButton", req)
	return err
}

// Logout logs out from the cloud Bot API server before launching the bot locally.
func (b *Bot[RequestType]) Logout() (bool, error) {
	r := NewApiRequester[methods.LogOutRequest, methods.LogOutResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "logOut", methods.LogOutRequest{})
	if err != nil {
		return false, err
	}
	return bool(*result), nil
}

// Close closes the bot instance before moving it from one local server to another.
func (b *Bot[RequestType]) Close() (bool, error) {
	r := NewApiRequester[methods.CloseRequest, methods.CloseResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "close", methods.CloseRequest{})
	if err != nil {
		return false, err
	}
	return bool(*result), nil
}

// SetMyName change's the bot name.
func (b *Bot[RequestType]) SetMyName(name, language string) error {
	req := methods.SetMyNameRequest{
		Name:         name,
		LanguageCode: language,
	}

	r := NewApiRequester[methods.SetMyNameRequest, methods.SetMyNameResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setMyName", req)
	return err
}

// MyName returns the current bot name for the given user language.
func (b *Bot[RequestType]) MyName(language string) (*methods.GetMyNameResponse, error) {
	req := methods.GetMyNameRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[methods.GetMyNameRequest, methods.GetMyNameResponse](b.token, b.apiURL, b.client)
	return r.Request(context.Background(), "getMyName", req)
}

// SetMyDescription change's the bot description, which is shown in the chat
// with the bot if the chat is empty.
func (b *Bot[RequestType]) SetMyDescription(desc, language string) error {
	req := methods.SetMyDescriptionRequest{
		Description:  desc,
		LanguageCode: language,
	}

	r := NewApiRequester[methods.SetMyDescriptionRequest, methods.SetMyDescriptionResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setMyDescription", req)
	return err
}

// MyDescription the current bot description for the given user language.
func (b *Bot[RequestType]) MyDescription(language string) (*methods.GetMyDescriptionResponse, error) {
	req := methods.GetMyDescriptionRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[methods.GetMyDescriptionRequest, methods.GetMyDescriptionResponse](b.token, b.apiURL, b.client)
	return r.Request(context.Background(), "getMyDescription", req)
}

// SetMyShortDescription change's the bot short description, which is shown on
// the bot's profile page and is sent together with the link when users share the bot.
func (b *Bot[RequestType]) SetMyShortDescription(desc, language string) error {
	req := methods.SetMyShortDescriptionRequest{
		ShortDescription: desc,
		LanguageCode:     language,
	}

	r := NewApiRequester[methods.SetMyShortDescriptionRequest, methods.SetMyShortDescriptionResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setMyShortDescription", req)
	return err
}

// MyShortDescription the current bot short description for the given user language.
func (b *Bot[RequestType]) MyShortDescription(language string) (*methods.GetMyShortDescriptionResponse, error) {
	req := methods.GetMyShortDescriptionRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[methods.GetMyShortDescriptionRequest, methods.GetMyShortDescriptionResponse](b.token, b.apiURL, b.client)
	return r.Request(context.Background(), "getMyShortDescription", req)
}
