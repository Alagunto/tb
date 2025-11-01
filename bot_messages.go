package tb

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// SendTo accepts 2+ arguments, starting with destination chat, followed by
// some Sendable (or string!) and optional send options.
//
// NOTE:
//
//	Since most arguments are of type interface{}, but have pointer
//	method receivers, make sure to pass them by-pointer, NOT by-value.
//
// What is a send option exactly? It can be one of the following types:
//
//   - *SendOptions (the actual object accepted by Telegram API)
//   - *ReplyMarkup (a component of SendOptions)
//   - Option (a shortcut flag for popular options)
//   - ParseMode (HTML, Markdown, etc)
func (b *Bot[RequestType]) SendTo(to bot.Recipient, what interface{}, opts ...params.SendOptions) (*telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}

	sendOpts := params.Merge(opts...)

	switch object := what.(type) {
	case string:
		return b.sendText(context.Background(), to, object, &sendOpts)
	case telegram.InputMedia:
		return b.sendMedia(context.Background(), to, object, &sendOpts)
	default:
		return nil, errors.WithInvalidParam(errors.ErrUnsupportedWhat, "what", fmt.Sprintf("%v", what))
	}
}

// SendToBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) SendToBackground(to bot.Recipient, what interface{}, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.SendTo(to, what, opts...)
}

func (b *Bot[RequestType]) sendText(ctx context.Context, to bot.Recipient, text string, opts *params.SendOptions) (*telegram.Message, error) {
	req := telegram.SendMessageRequest{
		ChatID: to.Recipient(),
		Text:   b.CensorText(text),
	}

	if opts != nil {
		req.ParseMode = string(opts.ParseMode)
		req.Entities = opts.Entities
		opts.InjectIntoMethodRequest(&req)
	}

	r := NewApiRequester[telegram.SendMessageRequest, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(ctx, "sendMessage", req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// sendMedia handles sending content that implements the telegram.InputMedia interface.
func (b *Bot[RequestType]) sendMedia(ctx context.Context, to bot.Recipient, media telegram.InputMedia, opts *params.SendOptions) (*telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}

	mediaParams, err := media.Field()
	if err != nil {
		return nil, err
	}

	// Get and validate media type for the API method
	mediaType, ok := mediaParams["type"].(string)
	if !ok || mediaType == "" {
		return nil, errors.WithInvalidParam(errors.ErrUnsupportedWhat, "type", "media type is missing or invalid")
	}

	// Prepare send options
	var sendOpts params.SendOptions
	if opts != nil {
		sendOpts = params.Merge(*opts)
	} else {
		sendOpts = params.NewSendOptions()
	}
	paramsMap := sendOpts.ToMap()

	// Merge media parameters into the request
	for k, v := range mediaParams {
		// For single media sends, rename "media" to match the type (photo, video, etc.)
		// because sendPhoto expects "photo" parameter, sendVideo expects "video", etc.
		if k == "media" {
			paramsMap[mediaType] = v
		} else if k != "type" {
			// Skip "type" field as it's not part of the send API parameters
			paramsMap[k] = v
		}
	}

	r := NewApiRequester[map[string]any, telegram.Message](b.token, b.apiURL, b.client)

	// TODO: Handle file uploads - need to implement proper file upload detection
	// The media.Field() method should already set media parameter correctly
	// (file_id for existing files, URL for remote files, or attach:// for uploads)

	paramsMap["chat_id"] = to.Recipient()
	if opts != nil {
		opts.InjectIntoMap(paramsMap)
	}

	// Capitalize first letter for proper Telegram API method name (sendPhoto, sendVideo, etc.)
	methodName := "send" + strings.ToUpper(string(mediaType[0])) + mediaType[1:]
	result, err := r.Request(ctx, methodName, paramsMap)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SendAlbumTo sends multiple instances of media as a single message.
// To include the caption, make sure the first Inputtable of an album has it.
// From all existing options, it only supports tele.Silent.
func (b *Bot[RequestType]) SendAlbumTo(to bot.Recipient, a telegram.InputAlbum, opts ...params.SendOptions) ([]telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}

	sendOpts := params.Merge(opts...)
	paramsMap := sendOpts.ToMap()
	paramsMap["chat_id"] = to.Recipient()

	// TODO: Handle file uploads - need to implement proper file upload detection
	mediaArray := make([]map[string]any, len(a.Media))
	for i, x := range a.Media {
		mediaParams, err := x.Field()
		if err != nil {
			return nil, err
		}
		mediaArray[i] = mediaParams
	}
	paramsMap["media"] = mediaArray

	r := NewApiRequester[map[string]any, []telegram.Message](b.token, b.apiURL, b.client)

	result, err := r.Request(context.Background(), "sendMediaGroup", paramsMap)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return *result, nil
}

// SendAlbumToBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) SendAlbumToBackground(to bot.Recipient, a telegram.InputAlbum, opts ...params.SendOptions) ([]telegram.Message, error) {
	return b.SendAlbumTo(to, a, opts...)
}

// ReplyTo behaves just like Send() with an exception of "reply-to" indicator.
// This function will panic upon nil Message.
func (b *Bot[RequestType]) ReplyTo(to *telegram.Message, what interface{}, opts ...params.SendOptions) (*telegram.Message, error) {
	sendOpts := params.Merge(opts...).WithReplyParams(&telegram.ReplyParameters{MessageID: to.ID})

	var recipient bot.Recipient
	if to.Chat.Type == telegram.ChatPrivate {
		recipient = to.Sender
	} else {
		recipient = to.Chat
		if to.ThreadID != 0 {
			sendOpts = sendOpts.WithThreadID(to.ThreadID)
		}
	}

	return b.SendTo(recipient, what, sendOpts)
}

// ReplyToBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) ReplyToBackground(to *telegram.Message, what interface{}, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.ReplyTo(to, what, opts...)
}

// ForwardTo behaves just like SendTo() but of all options it only supports Silent (see Bots API).
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) ForwardTo(to bot.Recipient, msg bot.Editable, opts ...params.SendOptions) (*telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}
	msgID, chatID := msg.MessageSig()

	sendOpts := params.Merge(opts...)

	req := telegram.ForwardMessageRequest{
		ChatID:     to.Recipient(),
		FromChatID: chatID,
		MessageID:  msgID,
	}

	if sendOpts.ThreadID != 0 {
		req.MessageThreadID = sendOpts.ThreadID
	}
	if sendOpts.DisableNotification {
		req.DisableNotification = true
	}
	if sendOpts.Protected {
		req.ProtectContent = true
	}

	r := NewApiRequester[telegram.ForwardMessageRequest, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "forwardMessage", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// ForwardToBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) ForwardToBackground(to bot.Recipient, msg bot.Editable, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.ForwardTo(to, msg, opts...)
}

// ForwardManyTo method forwards multiple messages of any kind.
// If some of the specified messages can't be found or forwarded, they are skipped.
// Service messages and messages with protected content can't be forwarded.
// Album grouping is kept for forwarded messages.
func (b *Bot[RequestType]) ForwardManyTo(to bot.Recipient, msgs []bot.Editable, opts ...params.SendOptions) ([]telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}
	if len(msgs) == 0 {
		return nil, nil
	}

	sendOpts := params.Merge(opts...)

	// Extract message IDs and from_chat_id
	messageIDs := make([]string, len(msgs))
	var fromChatID int64

	for i, msg := range msgs {
		msgID, chatID := msg.MessageSig()
		messageIDs[i] = msgID

		if i == 0 {
			fromChatID = chatID
		} else if fromChatID != chatID {
			// All messages must be from the same chat
			return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "messages", "all messages must be from the same chat")
		}
	}

	req := telegram.ForwardMessagesRequest{
		ChatID:     to.Recipient(),
		FromChatID: fromChatID,
		MessageIDs: messageIDs,
	}

	if sendOpts.ThreadID != 0 {
		req.MessageThreadID = sendOpts.ThreadID
	}
	if sendOpts.DisableNotification {
		req.DisableNotification = true
	}
	if sendOpts.Protected {
		req.ProtectContent = true
	}

	r := NewApiRequester[telegram.ForwardMessagesRequest, []telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "forwardMessages", req)
	if err != nil {
		return nil, err
	}

	return *result, nil
}

// ForwardManyToBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) ForwardManyToBackground(to bot.Recipient, msgs []bot.Editable, opts ...params.SendOptions) ([]telegram.Message, error) {
	return b.ForwardManyTo(to, msgs, opts...)
}

// CopyTo behaves just like ForwardTo() but the copied message doesn't have a link to the original message (see Bots API).
//
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) CopyTo(to bot.Recipient, msg bot.Editable, opts ...params.SendOptions) (*telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}
	msgID, chatID := msg.MessageSig()

	sendOpts := params.Merge(opts...)

	req := telegram.CopyMessageRequest{
		ChatID:     to.Recipient(),
		FromChatID: chatID,
		MessageID:  msgID,
	}

	if sendOpts.ThreadID != 0 {
		req.MessageThreadID = sendOpts.ThreadID
	}
	if sendOpts.DisableNotification {
		req.DisableNotification = true
	}
	if sendOpts.Protected {
		req.ProtectContent = true
	}
	if sendOpts.ReplyParams != nil {
		req.ReplyParameters = sendOpts.ReplyParams
	}
	if sendOpts.ReplyMarkup != nil {
		req.ReplyMarkup = sendOpts.ReplyMarkup
	}

	r := NewApiRequester[telegram.CopyMessageRequest, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "copyMessage", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// CopyToBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) CopyToBackground(to bot.Recipient, msg bot.Editable, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.CopyTo(to, msg, opts...)
}

// CopyManyTo this method makes a copy of messages of any kind.
// If some of the specified messages can't be found or copied, they are skipped.
// Service messages, giveaway messages, giveaway winners messages, and
// invoice messages can't be copied. A quiz poll can be copied only if the value of the field
// correct_option_id is known to the bot. The method is analogous
// to the method forwardMessages, but the copied messages don't have a link to the original message.
// Album grouping is kept for copied messages.
func (b *Bot[RequestType]) CopyManyTo(to bot.Recipient, msgs []bot.Editable, opts ...params.SendOptions) ([]telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}
	if len(msgs) == 0 {
		return nil, nil
	}

	sendOpts := params.Merge(opts...)

	// Extract message IDs and from_chat_id
	messageIDs := make([]string, len(msgs))
	var fromChatID int64

	for i, msg := range msgs {
		msgID, chatID := msg.MessageSig()
		messageIDs[i] = msgID

		if i == 0 {
			fromChatID = chatID
		} else if fromChatID != chatID {
			// All messages must be from the same chat
			return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "messages", "all messages must be from the same chat")
		}
	}

	req := telegram.CopyMessagesRequest{
		ChatID:     to.Recipient(),
		FromChatID: fromChatID,
		MessageIDs: messageIDs,
	}

	if sendOpts.ThreadID != 0 {
		req.MessageThreadID = sendOpts.ThreadID
	}
	if sendOpts.DisableNotification {
		req.DisableNotification = true
	}
	if sendOpts.Protected {
		req.ProtectContent = true
	}

	r := NewApiRequester[telegram.CopyMessagesRequest, []telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "copyMessages", req)
	if err != nil {
		return nil, err
	}

	return *result, nil
}

// CopyManyToBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) CopyManyToBackground(to bot.Recipient, msgs []bot.Editable, opts ...params.SendOptions) ([]telegram.Message, error) {
	return b.CopyManyTo(to, msgs, opts...)
}

// Edit is magic, it lets you change already sent message.
// This function will panic upon nil Editable.
//
// If edited message is sent by the bot, returns it,
// otherwise returns nil and ErrTrueResult.
//
// Use cases:
//
//	b.Edit(m, m.Text, newMarkup)
//	b.Edit(m, "new <b>text</b>", tele.ModeHTML)
//	b.Edit(m, &tele.ReplyMarkup{...})
//	b.Edit(m, &tele.Photo{File: ...})
//	b.Edit(m, tele.Location{42.1337, 69.4242})
//	b.Edit(c, "edit inline message from the callback")
//	b.Edit(r, "edit message from chosen inline result")
func (b *Bot[RequestType]) Edit(msg bot.Editable, what interface{}, opts ...params.SendOptions) (*telegram.Message, error) {
	sendOpts := params.Merge(opts...)
	msgID, chatID := msg.MessageSig()

	switch v := what.(type) {
	case *telegram.ReplyMarkup:
		return b.EditReplyMarkup(msg, v)
	case telegram.InputMedia:
		return b.EditMedia(msg, v, sendOpts)
	case string:
		req := telegram.EditMessageTextRequest{
			Text: b.CensorText(v),
		}

		if chatID == 0 { // if inline message
			req.InlineMessageID = msgID
		} else {
			req.ChatID = strconv.FormatInt(chatID, 10)
			req.MessageID = msgID
		}

		params.New().With(sendOpts).Build()

		r := NewApiRequester[telegram.EditMessageTextRequest, telegram.Message](b.token, b.apiURL, b.client)
		result, err := r.Request(context.Background(), "editMessageText", req)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		return result, nil

	case telegram.Checklist:
		req := telegram.EditMessageChecklistRequest{
			Checklist: v,
		}

		if chatID == 0 { // if inline message
			req.InlineMessageID = msgID
		} else {
			req.ChatID = strconv.FormatInt(chatID, 10)
			req.MessageID = msgID
		}

		params.New().With(sendOpts).Build()

		r := NewApiRequester[telegram.EditMessageChecklistRequest, telegram.Message](b.token, b.apiURL, b.client)
		result, err := r.Request(context.Background(), "editMessageChecklist", req)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		return result, nil

	case telegram.Location:
		req := telegram.EditMessageLiveLocationRequest{
			Latitude:  v.Latitude,
			Longitude: v.Longitude,
		}

		if chatID == 0 { // if inline message
			req.InlineMessageID = msgID
		} else {
			req.ChatID = strconv.FormatInt(chatID, 10)
			req.MessageID = msgID
		}

		if v.HorizontalAccuracy != 0 {
			req.HorizontalAccuracy = &v.HorizontalAccuracy
		}
		if v.Heading != 0 {
			req.Heading = v.Heading
		}
		if v.ProximityAlertRadius != 0 {
			req.ProximityAlertRadius = v.ProximityAlertRadius
		}
		if v.LivePeriod != 0 {
			req.LivePeriod = v.LivePeriod
		}
		params.New().With(sendOpts).Build()

		r := NewApiRequester[telegram.EditMessageLiveLocationRequest, telegram.Message](b.token, b.apiURL, b.client)
		result, err := r.Request(context.Background(), "editMessageLiveLocation", req)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		return result, nil

	default:
		return nil, errors.WithInvalidParam(errors.ErrUnsupportedWhat, "what", fmt.Sprintf("%v", what))
	}
}

// EditBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) EditBackground(msg bot.Editable, what interface{}, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.Edit(msg, what, opts...)
}

// EditReplyMarkup edits reply markup of already sent message.
// This function will panic upon nil Editable.
// Pass nil or empty ReplyMarkup to delete it from the message.
//
// If edited message is sent by the bot, returns it,
// otherwise returns nil and ErrTrueResult.
func (b *Bot[RequestType]) EditReplyMarkup(msg bot.Editable, markup *telegram.ReplyMarkup) (*telegram.Message, error) {
	msgID, chatID := msg.MessageSig()

	req := telegram.EditMessageReplyMarkupRequest{}

	if chatID == 0 { // if inline message
		req.InlineMessageID = msgID
	} else {
		req.ChatID = strconv.FormatInt(chatID, 10)
		req.MessageID = msgID
	}

	if markup == nil {
		// will delete reply markup
		markup = &telegram.ReplyMarkup{}
	}

	// TODO: Implement button preparation logic
	// The PrepareButtons method doesn't exist in SendOptions
	// Need to implement logic to process button unique identifiers

	req.ReplyMarkup = markup

	r := NewApiRequester[telegram.EditMessageReplyMarkupRequest, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "editMessageReplyMarkup", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// EditReplyMarkupBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) EditReplyMarkupBackground(msg bot.Editable, markup *telegram.ReplyMarkup) (*telegram.Message, error) {
	return b.EditReplyMarkup(msg, markup)
}

// EditCaption edits already sent photo caption with known recipient and message id.
// This function will panic upon nil Editable.
//
// If edited message is sent by the bot, returns it,
// otherwise returns nil and ErrTrueResult.
func (b *Bot[RequestType]) EditCaption(msg bot.Editable, caption string, opts ...params.SendOptions) (*telegram.Message, error) {
	msgID, chatID := msg.MessageSig()

	sendOpts := params.Merge(opts...)

	req := telegram.EditMessageCaptionRequest{
		Caption:         b.CensorText(caption),
		ParseMode:       string(sendOpts.ParseMode),
		CaptionEntities: sendOpts.Entities,
	}

	if chatID == 0 { // if inline message
		req.InlineMessageID = msgID
	} else {
		req.ChatID = strconv.FormatInt(chatID, 10)
		req.MessageID = msgID
	}

	params.New().With(sendOpts).Build()

	r := NewApiRequester[telegram.EditMessageCaptionRequest, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "editMessageCaption", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// EditCaptionBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) EditCaptionBackground(msg bot.Editable, caption string, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.EditCaption(msg, caption, opts...)
}

// EditMedia edits already sent media with known recipient and message id.
// This function will panic upon nil Editable.
//
// If edited message is sent by the bot, returns it,
// otherwise returns nil and ErrTrueResult.
//
// Use cases:
//
//	b.EditMedia(m, &tele.Photo{File: tele.FromDisk("chicken.jpg")})
//	b.EditMedia(m, &tele.Video{File: tele.FromURL("http://video.mp4")})
func (b *Bot[RequestType]) EditMedia(msg bot.Editable, media telegram.InputMedia, opts ...params.SendOptions) (*telegram.Message, error) {
	msgID, chatID := msg.MessageSig()

	sendOpts := params.Merge(opts...)

	req := telegram.EditMessageMediaRequest{}

	if chatID == 0 { // if inline message
		req.InlineMessageID = msgID
	} else {
		req.ChatID = strconv.FormatInt(chatID, 10)
		req.MessageID = msgID
	}

	params.New().With(sendOpts).Build()

	// TODO: Handle file uploads - need to implement proper file upload detection
	// Get media parameters from the InputMedia
	mediaParams, err := media.Field()
	if err != nil {
		return nil, err
	}
	if mediaStr, ok := mediaParams["media"].(string); ok {
		req.Media = mediaStr
	}

	r := NewApiRequester[telegram.EditMessageMediaRequest, telegram.Message](b.token, b.apiURL, b.client)

	result, err := r.Request(context.Background(), "editMessageMedia", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// EditMediaBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) EditMediaBackground(msg bot.Editable, media telegram.InputMedia, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.EditMedia(msg, media, opts...)
}

// Delete removes the message, including service messages.
// This function will panic upon nil Editable.
//
//   - A message can only be deleted if it was sent less than 48 hours ago.
//   - A dice message in a private chat can only be deleted if it was sent more than 24 hours ago.
//   - Bots can delete outgoing messages in private chats, groups, and supergroups.
//   - Bots can delete incoming messages in private chats.
//   - Bots granted can_post_messages permissions can delete outgoing messages in channels.
//   - If the bot is an administrator of a group, it can delete any message there.
//   - If the bot has can_delete_messages permission in a supergroup or a
//     channel, it can delete any message there.
func (b *Bot[RequestType]) Delete(msg bot.Editable) error {
	msgID, chatID := msg.MessageSig()

	req := telegram.DeleteMessageRequest{
		ChatID:    chatID,
		MessageID: msgID,
	}

	r := NewApiRequester[telegram.DeleteMessageRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "deleteMessage", req)
	return err
}

// DeleteBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) DeleteBackground(msg bot.Editable) error {
	return b.Delete(msg)
}

// DeleteMany deletes multiple messages simultaneously.
// If some of the specified messages can't be found, they are skipped.
func (b *Bot[RequestType]) DeleteMany(msgs []bot.Editable) error {
	if len(msgs) == 0 {
		return nil
	}

	// Extract message IDs and chat_id
	messageIDs := make([]string, len(msgs))
	var chatID int64

	for i, msg := range msgs {
		msgID, cID := msg.MessageSig()
		messageIDs[i] = msgID

		if i == 0 {
			chatID = cID
		} else if chatID != cID {
			// All messages must be from the same chat
			return errors.WithInvalidParam(errors.ErrBadRecipient, "messages", "all messages must be from the same chat")
		}
	}

	req := telegram.DeleteMessagesRequest{
		ChatID:     chatID,
		MessageIDs: messageIDs,
	}

	r := NewApiRequester[telegram.DeleteMessagesRequest, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "deleteMessages", req)
	return err
}

// DeleteManyBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) DeleteManyBackground(msgs []bot.Editable) error {
	return b.DeleteMany(msgs)
}

// StopLiveLocation stops broadcasting live message location
// before Location.LivePeriod expires.
//
// It supports ReplyMarkup.
// This function will panic upon nil Editable.
//
// If the message is sent by the bot, returns it,
// otherwise returns nil and ErrTrueResult.
func (b *Bot[RequestType]) StopLiveLocation(msg bot.Editable, opts ...params.SendOptions) (*telegram.Message, error) {
	msgID, chatID := msg.MessageSig()

	sendOpts := params.Merge(opts...)

	req := telegram.StopMessageLiveLocationRequest{}

	if chatID == 0 { // if inline message
		req.InlineMessageID = msgID
	} else {
		req.ChatID = strconv.FormatInt(chatID, 10)
		req.MessageID = msgID
	}

	params.New().With(sendOpts).Build()

	r := NewApiRequester[telegram.StopMessageLiveLocationRequest, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "stopMessageLiveLocation", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// StopLiveLocationBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) StopLiveLocationBackground(msg bot.Editable, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.StopLiveLocation(msg, opts...)
}

// StopPoll stops a poll which was sent by the bot and returns
// the stopped Poll object with the final results.
//
// It supports ReplyMarkup.
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) StopPoll(msg bot.Editable, opts ...params.SendOptions) (*telegram.Poll, error) {
	msgID, chatID := msg.MessageSig()

	sendOpts := params.Merge(opts...)

	req := telegram.StopPollRequest{
		ChatID:    chatID,
		MessageID: msgID,
	}

	params.New().With(sendOpts).Build()

	r := NewApiRequester[telegram.StopPollRequest, telegram.Poll](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "stopPoll", req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// StopPollBackground is a convenience wrapper using context.Background()
func (b *Bot[RequestType]) StopPollBackground(msg bot.Editable, opts ...params.SendOptions) (*telegram.Poll, error) {
	return b.StopPoll(msg, opts...)
}
