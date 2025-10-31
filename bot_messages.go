package tb

import (
	"context"
	"fmt"
	"strconv"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/sendables"
	"github.com/alagunto/tb/telegram"
	"github.com/alagunto/tb/telegram/methods"
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
func (b *Bot[RequestType]) SendTo(to bot.Recipient, what interface{}, opts ...communications.SendOptions) (*telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}

	sendOpts := communications.MergeMultipleSendOptions(opts...)

	switch object := what.(type) {
	case string:
		return b.sendText(to, object, &sendOpts)
	case sendables.Sendable[telegram.Message]:
		prepared := object.PrepareForTelegram()
		msg, err := b.sendPrepared(to, prepared, &sendOpts)
		if err != nil {
			return nil, err
		}
		return msg, object.ProcessTelegramResponse(*msg)
	case outgoing.Content:
		return b.sendContent(to, object, &sendOpts)
	default:
		return nil, errors.WithInvalidParam(errors.ErrUnsupportedWhat, "what", fmt.Sprintf("%v", what))
	}
}

// sendContent handles sending content that implements the outgoing.Content interface.
// This is the new file upload system that makes upload timing explicit.
func (b *Bot[RequestType]) sendContent(to bot.Recipient, content outgoing.Content, opts *communications.SendOptions) (*telegram.Message, error) {
	method := content.ToTelegramSendMethod()

	// Start with base params from method.Params (already map[string]any)
	paramsMap := make(map[string]any)
	paramsMap["chat_id"] = to.Recipient()

	// Add content-specific params
	for k, v := range method.Params {
		paramsMap[k] = v
	}

	opts.InjectInto(paramsMap)

	// Send the request using the new FileSource format
	result, err := b.sendFiles(method.Name, method.Files, paramsMap)
	if err != nil {
		return nil, err
	}

	// Let content update itself with the response (if it implements ResponseHandler)
	if handler, ok := content.(outgoing.ResponseHandler); ok {
		if err := handler.UpdateFromResponse(result); err != nil {
			return nil, err
		}
	}

	return result, nil
}

// sendPrepared sends a prepared sendable to Telegram
func (b *Bot[RequestType]) sendPrepared(to bot.Recipient, prepared *sendables.Prepared, opts *communications.SendOptions) (*telegram.Message, error) {
	// Build params
	paramsMap := make(map[string]any)
	paramsMap["chat_id"] = to.Recipient()

	// Add prepared params (convert strings to any)
	for k, v := range prepared.Params {
		paramsMap[k] = v
	}

	opts.InjectInto(paramsMap)

	// Use sendFiles which handles ApiRequester internally
	return b.sendFiles(prepared.SendMethod.String(), prepared.Files, paramsMap)
}

// SendAlbumTo sends multiple instances of media as a single message.
// To include the caption, make sure the first Inputtable of an album has it.
// From all existing options, it only supports tele.Silent.
func (b *Bot[RequestType]) SendAlbumTo(to bot.Recipient, a telegram.InputAlbum, opts ...communications.SendOptions) ([]telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}

	sendOpts := communications.MergeMultipleSendOptions(opts...)

	p := params.New().
		Add("chat_id", to.Recipient()).
		Build()

	paramsMap := sendOpts.Inject(p)

	filesToSend := make(map[string]files.FileSource)
	for _, x := range a.Media {
		for name, source := range x.ToTelegramSendMethod().Files {
			filesToSend[name] = source
		}
	}

	// Use ApiRequester for sendMediaGroup which returns []Message

	r := NewApiRequester[map[string]any, []telegram.Message](b.token, b.apiURL, b.client)
	for _, x := range a.Media {
		for name, source := range x.ToTelegramSendMethod().Files {
			r = r.WithFileToUpload(name, source)
		}
	}
	for name, source := range filesToSend {
		r = r.WithFileToUpload(name, source)
	}

	result, err := r.Request(context.Background(), "sendMediaGroup", paramsMap)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return *result, nil
}

// ReplyTo behaves just like Send() with an exception of "reply-to" indicator.
// This function will panic upon nil Message.
func (b *Bot[RequestType]) ReplyTo(to *telegram.Message, what interface{}, opts ...communications.SendOptions) (*telegram.Message, error) {
	sendOpts := communications.MergeMultipleSendOptions(opts...).WithReplyParams(&telegram.ReplyParams{MessageID: to.ID})
	return b.SendTo(to.Chat, what, sendOpts)
}

// ForwardTo behaves just like SendTo() but of all options it only supports Silent (see Bots API).
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) ForwardTo(to bot.Recipient, msg bot.Editable, opts ...communications.SendOptions) (*telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}
	msgID, chatID := msg.MessageSig()

	sendOpts := communications.MergeMultipleSendOptions(opts...)

	req := methods.ForwardMessageRequest{
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

	r := NewApiRequester[methods.ForwardMessageRequest, methods.ForwardMessageResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "forwardMessage", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// ForwardManyTo method forwards multiple messages of any kind.
// If some of the specified messages can't be found or forwarded, they are skipped.
// Service messages and messages with protected content can't be forwarded.
// Album grouping is kept for forwarded messages.
func (b *Bot[RequestType]) ForwardManyTo(to bot.Recipient, msgs []bot.Editable, opts ...communications.SendOptions) ([]telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}
	if len(msgs) == 0 {
		return nil, nil
	}

	sendOpts := communications.MergeMultipleSendOptions(opts...)

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

	req := methods.ForwardMessagesRequest{
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

	r := NewApiRequester[methods.ForwardMessagesRequest, methods.ForwardMessagesResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "forwardMessages", req)
	if err != nil {
		return nil, err
	}

	return *result, nil
}

// CopyTo behaves just like ForwardTo() but the copied message doesn't have a link to the original message (see Bots API).
//
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) CopyTo(to bot.Recipient, msg bot.Editable, opts ...communications.SendOptions) (*telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}
	msgID, chatID := msg.MessageSig()

	sendOpts := communications.MergeMultipleSendOptions(opts...)

	req := methods.CopyMessageRequest{
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

	r := NewApiRequester[methods.CopyMessageRequest, methods.CopyMessageResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "copyMessage", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// CopyManyTo this method makes a copy of messages of any kind.
// If some of the specified messages can't be found or copied, they are skipped.
// Service messages, giveaway messages, giveaway winners messages, and
// invoice messages can't be copied. A quiz poll can be copied only if the value of the field
// correct_option_id is known to the bot. The method is analogous
// to the method forwardMessages, but the copied messages don't have a link to the original message.
// Album grouping is kept for copied messages.
func (b *Bot[RequestType]) CopyManyTo(to bot.Recipient, msgs []bot.Editable, opts ...communications.SendOptions) ([]telegram.Message, error) {
	if to == nil {
		return nil, errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", nil)
	}
	if len(msgs) == 0 {
		return nil, nil
	}

	sendOpts := communications.MergeMultipleSendOptions(opts...)

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

	req := methods.CopyMessagesRequest{
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

	r := NewApiRequester[methods.CopyMessagesRequest, methods.CopyMessagesResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "copyMessages", req)
	if err != nil {
		return nil, err
	}

	return *result, nil
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
func (b *Bot[RequestType]) Edit(msg bot.Editable, what interface{}, opts ...communications.SendOptions) (*telegram.Message, error) {
	sendOpts := communications.MergeMultipleSendOptions(opts...)
	msgID, chatID := msg.MessageSig()

	switch v := what.(type) {
	case *telegram.ReplyMarkup:
		return b.EditReplyMarkup(msg, v)
	case outgoing.Content:
		return b.EditMedia(msg, v, opts...)
	case string:
		req := methods.EditMessageTextRequest{
			Text: b.CensorText(v),
		}

		if chatID == 0 { // if inline message
			req.InlineMessageID = msgID
		} else {
			req.ChatID = strconv.FormatInt(chatID, 10)
			req.MessageID = msgID
		}

		sendOpts.InjectIntoMethodRequest(&req)

		r := NewApiRequester[methods.EditMessageTextRequest, methods.EditMessageTextResponse](b.token, b.apiURL, b.client)
		result, err := r.Request(context.Background(), "editMessageText", req)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		return result, nil

	case Checklist:
		req := methods.EditMessageChecklistRequest{
			Checklist: v,
		}

		if chatID == 0 { // if inline message
			req.InlineMessageID = msgID
		} else {
			req.ChatID = strconv.FormatInt(chatID, 10)
			req.MessageID = msgID
		}

		sendOpts.InjectIntoMethodRequest(&req)

		r := NewApiRequester[methods.EditMessageChecklistRequest, methods.EditMessageChecklistResponse](b.token, b.apiURL, b.client)
		result, err := r.Request(context.Background(), "editMessageChecklist", req)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		return result, nil

	case Location:
		req := methods.EditMessageLiveLocationRequest{
			Latitude:  float64(v.Lat),
			Longitude: float64(v.Lng),
		}

		if chatID == 0 { // if inline message
			req.InlineMessageID = msgID
		} else {
			req.ChatID = strconv.FormatInt(chatID, 10)
			req.MessageID = msgID
		}

		if v.HorizontalAccuracy != nil {
			accuracy := float64(*v.HorizontalAccuracy)
			req.HorizontalAccuracy = &accuracy
		}
		if v.Heading != 0 {
			req.Heading = v.Heading
		}
		if v.AlertRadius != 0 {
			req.ProximityAlertRadius = v.AlertRadius
		}
		if v.LivePeriod != 0 {
			req.LivePeriod = v.LivePeriod
		}
		sendOpts.InjectIntoMethodRequest(&req)

		r := NewApiRequester[methods.EditMessageLiveLocationRequest, methods.EditMessageLiveLocationResponse](b.token, b.apiURL, b.client)
		result, err := r.Request(context.Background(), "editMessageLiveLocation", req)
		if err != nil {
			return nil, errors.Wrap(err)
		}
		return result, nil

	default:
		return nil, errors.WithInvalidParam(errors.ErrUnsupportedWhat, "what", fmt.Sprintf("%v", what))
	}
}

// EditReplyMarkup edits reply markup of already sent message.
// This function will panic upon nil Editable.
// Pass nil or empty ReplyMarkup to delete it from the message.
//
// If edited message is sent by the bot, returns it,
// otherwise returns nil and ErrTrueResult.
func (b *Bot[RequestType]) EditReplyMarkup(msg bot.Editable, markup *telegram.ReplyMarkup) (*telegram.Message, error) {
	msgID, chatID := msg.MessageSig()

	req := methods.EditMessageReplyMarkupRequest{}

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

	// Use SendOptions to prepare buttons (handles Unique field processing)
	opt := communications.SendOptions{ReplyMarkup: markup}
	if markup != nil {
		markup.InlineKeyboard = opt.PrepareButtons(markup.InlineKeyboard)
	}

	req.ReplyMarkup = markup

	r := NewApiRequester[methods.EditMessageReplyMarkupRequest, methods.EditMessageReplyMarkupResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "editMessageReplyMarkup", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// EditCaption edits already sent photo caption with known recipient and message id.
// This function will panic upon nil Editable.
//
// If edited message is sent by the bot, returns it,
// otherwise returns nil and ErrTrueResult.
func (b *Bot[RequestType]) EditCaption(msg bot.Editable, caption string, opts ...communications.SendOptions) (*telegram.Message, error) {
	msgID, chatID := msg.MessageSig()

	sendOpts := communications.MergeMultipleSendOptions(opts...)

	req := methods.EditMessageCaptionRequest{
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

	sendOpts.InjectIntoMethodRequest(&req)

	r := NewApiRequester[methods.EditMessageCaptionRequest, methods.EditMessageCaptionResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "editMessageCaption", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
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
func (b *Bot[RequestType]) EditMedia(msg bot.Editable, media outgoing.Content, opts ...communications.SendOptions) (*telegram.Message, error) {
	// Handle media file
	mediaFile := media.ToTelegramSendMethod().Files["media"]

	repr := "attach://media"

	msgID, chatID := msg.MessageSig()

	sendOpts := communications.MergeMultipleSendOptions(opts...)

	req := methods.EditMessageMediaRequest{
		Media: repr,
	}

	if chatID == 0 { // if inline message
		req.InlineMessageID = msgID
	} else {
		req.ChatID = strconv.FormatInt(chatID, 10)
		req.MessageID = msgID
	}

	sendOpts.InjectIntoMethodRequest(&req)

	// Create the requester
	r := NewApiRequester[methods.EditMessageMediaRequest, methods.EditMessageMediaResponse](b.token, b.apiURL, b.client)
	r = r.WithFileToUpload("media", mediaFile)

	result, err := r.Request(context.Background(), "editMessageMedia", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
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

	req := methods.DeleteMessageRequest{
		ChatID:    chatID,
		MessageID: msgID,
	}

	r := NewApiRequester[methods.DeleteMessageRequest, methods.DeleteMessageResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "deleteMessage", req)
	return err
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

	req := methods.DeleteMessagesRequest{
		ChatID:     chatID,
		MessageIDs: messageIDs,
	}

	r := NewApiRequester[methods.DeleteMessagesRequest, methods.DeleteMessagesResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "deleteMessages", req)
	return err
}

// StopLiveLocation stops broadcasting live message location
// before Location.LivePeriod expires.
//
// It supports ReplyMarkup.
// This function will panic upon nil Editable.
//
// If the message is sent by the bot, returns it,
// otherwise returns nil and ErrTrueResult.
func (b *Bot[RequestType]) StopLiveLocation(msg bot.Editable, opts ...communications.SendOptions) (*Message, error) {
	msgID, chatID := msg.MessageSig()

	sendOpts := communications.MergeMultipleSendOptions(opts...)

	req := methods.StopMessageLiveLocationRequest{}

	if chatID == 0 { // if inline message
		req.InlineMessageID = msgID
	} else {
		req.ChatID = strconv.FormatInt(chatID, 10)
		req.MessageID = msgID
	}

	sendOpts.InjectIntoMethodRequest(&req)

	r := NewApiRequester[methods.StopMessageLiveLocationRequest, methods.StopMessageLiveLocationResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "stopMessageLiveLocation", req)
	if err != nil {
		return nil, errors.Wrap(err)
	}

	return result, nil
}

// StopPoll stops a poll which was sent by the bot and returns
// the stopped Poll object with the final results.
//
// It supports ReplyMarkup.
// This function will panic upon nil Editable.
func (b *Bot[RequestType]) StopPoll(msg bot.Editable, opts ...communications.SendOptions) (*Poll, error) {
	msgID, chatID := msg.MessageSig()

	sendOpts := communications.MergeMultipleSendOptions(opts...)

	req := methods.StopPollRequest{
		ChatID:    chatID,
		MessageID: msgID,
	}

	sendOpts.InjectIntoMethodRequest(&req)

	r := NewApiRequester[methods.StopPollRequest, methods.StopPollResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "stopPoll", req)
	if err != nil {
		return nil, err
	}
	return result, nil
}
