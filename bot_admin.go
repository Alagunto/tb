package tb

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// Pin pins a message in a supergroup or a channel.
//
// It supports Silent option.
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) Pin(msg bot.Editable) error {
	if msg == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "message", nil)
	}

	msgID, chatID := msg.MessageSig()

	req := telegram.PinMessageRequest{
		ChatID:    chatID,
		MessageID: msgID,
	}

	r := NewApiRequester[telegram.PinMessageRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "pinChatMessage", req)
	return err
}

// PinBackground pins a message in a supergroup or a channel using background context.
//
// It supports Silent option.
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) PinBackground(msg bot.Editable) error {
	return b.Pin(msg)
}

// Unpin unpins a message in a supergroup or a channel.
// It supports tb.Silent option.
func (b *Bot[RequestType]) Unpin(chat bot.Recipient, messageID ...int) error {
	if chat == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

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

// UnpinBackground unpins a message in a supergroup or a channel using background context.
// It supports tb.Silent option.
func (b *Bot[RequestType]) UnpinBackground(chat bot.Recipient, messageID ...int) error {
	return b.Unpin(chat, messageID...)
}

// UnpinAll unpins all messages in a supergroup or a channel.
// It supports tb.Silent option.
func (b *Bot[RequestType]) UnpinAll(chat bot.Recipient) error {
	if chat == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

	req := telegram.UnpinAllMessageRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[telegram.UnpinAllMessageRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "unpinAllChatMessages", req)
	return err
}

// UnpinAllBackground unpins all messages in a supergroup or a channel using background context.
// It supports tb.Silent option.
func (b *Bot[RequestType]) UnpinAllBackground(chat bot.Recipient) error {
	return b.UnpinAll(chat)
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
	if to == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "recipient", nil)
	}

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

// NotifyBackground updates the chat action for recipient using background context.
//
// Chat action is a status message that recipient would see where
// you typically see "Harry is typing" status message. The only
// difference is that bots' chat actions live only for 5 seconds
// and die just once the client receives a message from the bot.
//
// Currently, Telegram supports only a narrow range of possible
// actions, these are aligned as constants of this package.
func (b *Bot[RequestType]) NotifyBackground(to bot.Recipient, action telegram.ChatAction, opts ...params.SendOptions) error {
	return b.Notify(to, action, opts...)
}

// Leave makes bot leave a group, supergroup or channel.
func (b *Bot[RequestType]) Leave(chat bot.Recipient) error {
	if chat == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "chat", nil)
	}

	req := telegram.LeaveChatRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[telegram.LeaveChatRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "leaveChat", req)
	return err
}

// LeaveBackground makes bot leave a group, supergroup or channel using background context.
func (b *Bot[RequestType]) LeaveBackground(chat bot.Recipient) error {
	return b.Leave(chat)
}

// MenuButton returns the current value of the bot's menu button in a private chat,
// or the default menu button.
func (b *Bot[RequestType]) MenuButton(ctx context.Context, chat *telegram.User) (*telegram.MenuButton, error) {
	req := telegram.GetChatMenuButtonRequest{
		ChatID: chat.Recipient(),
	}

	r := NewApiRequester[telegram.GetChatMenuButtonRequest, telegram.MenuButton](b.token, b.apiURL, b.client)
	return r.Request(ctx, "getChatMenuButton", req)
}

// MenuButtonBackground returns the current value of the bot's menu button in a private chat,
// or the default menu button using background context.
func (b *Bot[RequestType]) MenuButtonBackground(chat *telegram.User) (*telegram.MenuButton, error) {
	return b.MenuButton(context.Background(), chat)
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

// SetMenuButtonBackground changes the bot's menu button in a private chat,
// or the default menu button using background context.
//
// It accepts two kinds of menu button arguments:
//
//   - MenuButtonType for simple menu buttons (default, commands)
//   - MenuButton complete structure for web_app menu button type
func (b *Bot[RequestType]) SetMenuButtonBackground(chat *telegram.User, mb interface{}) error {
	return b.SetMenuButton(chat, mb)
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

// LogoutBackground logs out from the cloud Bot API server before launching the bot locally using background context.
func (b *Bot[RequestType]) LogoutBackground() (bool, error) {
	return b.Logout()
}

// Close closes the bot instance before moving it from one local server to another.
func (b *Bot[RequestType]) Close(ctx context.Context) (bool, error) {
	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	result, err := r.Request(ctx, "close", make(map[string]any))
	if err != nil {
		return false, err
	}
	return *result, nil
}

// CloseBackground closes the bot instance before moving it from one local server to another using background context.
func (b *Bot[RequestType]) CloseBackground() (bool, error) {
	return b.Close(context.Background())
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

// SetMyNameBackground change's the bot name using background context.
func (b *Bot[RequestType]) SetMyNameBackground(name, language string) error {
	return b.SetMyName(name, language)
}

// MyName returns the current bot name for the given user language.
func (b *Bot[RequestType]) MyName(ctx context.Context, language string) (*telegram.BotName, error) {
	req := telegram.GetMyNameRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[telegram.GetMyNameRequest, telegram.BotName](b.token, b.apiURL, b.client)
	return r.Request(ctx, "getMyName", req)
}

// MyNameBackground returns the current bot name for the given user language using background context.
func (b *Bot[RequestType]) MyNameBackground(language string) (*telegram.BotName, error) {
	return b.MyName(context.Background(), language)
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

// SetMyDescriptionBackground change's the bot description, which is shown in the chat
// with the bot if the chat is empty using background context.
func (b *Bot[RequestType]) SetMyDescriptionBackground(desc, language string) error {
	return b.SetMyDescription(desc, language)
}

// MyDescription the current bot description for the given user language.
func (b *Bot[RequestType]) MyDescription(ctx context.Context, language string) (*telegram.BotDescription, error) {
	req := telegram.GetMyDescriptionRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[telegram.GetMyDescriptionRequest, telegram.BotDescription](b.token, b.apiURL, b.client)
	return r.Request(ctx, "getMyDescription", req)
}

// MyDescriptionBackground the current bot description for the given user language using background context.
func (b *Bot[RequestType]) MyDescriptionBackground(language string) (*telegram.BotDescription, error) {
	return b.MyDescription(context.Background(), language)
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

// SetMyShortDescriptionBackground change's the bot short description, which is shown on
// the bot's profile page and is sent together with the link when users share the bot using background context.
func (b *Bot[RequestType]) SetMyShortDescriptionBackground(desc, language string) error {
	return b.SetMyShortDescription(desc, language)
}

// MyShortDescription the current bot short description for the given user language.
func (b *Bot[RequestType]) MyShortDescription(ctx context.Context, language string) (*telegram.BotShortDescription, error) {
	req := telegram.GetMyShortDescriptionRequest{
		LanguageCode: language,
	}

	r := NewApiRequester[telegram.GetMyShortDescriptionRequest, telegram.BotShortDescription](b.token, b.apiURL, b.client)
	return r.Request(ctx, "getMyShortDescription", req)
}

// MyShortDescriptionBackground the current bot short description for the given user language using background context.
func (b *Bot[RequestType]) MyShortDescriptionBackground(language string) (*telegram.BotShortDescription, error) {
	return b.MyShortDescription(context.Background(), language)
}

// ProfilePhotosOf returns the profile photos of the given user.
func (b *Bot[RequestType]) ProfilePhotosOf(user *telegram.User) ([]telegram.PhotoSize, error) {
	if user == nil {
		return nil, errors.WithInvalidParam(errors.ErrTelebot, "user", nil)
	}

	params := map[string]string{
		"user_id": strconv.FormatInt(user.ID, 10),
	}

	data, err := b.Raw("getUserProfilePhotos", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result [][]telegram.PhotoSize `json:"photos"`
	}

	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrap(err)
	}

	// Flatten the photos array or return the first page
	// The API returns photos in different sizes, grouped by photo
	if len(resp.Result) == 0 {
		return []telegram.PhotoSize{}, nil
	}

	// Return all photos from all sizes, flattened
	var allPhotos []telegram.PhotoSize
	for _, photoGroup := range resp.Result {
		allPhotos = append(allPhotos, photoGroup...)
	}

	return allPhotos, nil
}
