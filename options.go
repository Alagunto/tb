package tb

import (
	"encoding/json"
	"strconv"

	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/telegram"
)

// Option is a shortcut flag type for certain message features
// (so-called options). It means that instead of passing
// fully-fledged SendOptions* to Send(), you can use these
// flags instead.
//
// Supported options are defined as iota-constants.
type Option int

const (
	// NoPreview = SendOptions.DisableWebPagePreview
	NoPreview Option = iota

	// Silent = SendOptions.DisableNotification
	Silent

	// AllowWithoutReply = SendOptions.AllowWithoutReply
	AllowWithoutReply

	// Protected = SendOptions.Protected
	Protected

	// ForceReply = ReplyMarkup.ForceReply
	ForceReply

	// OneTimeKeyboard = ReplyMarkup.OneTimeKeyboard
	OneTimeKeyboard

	// RemoveKeyboard = ReplyMarkup.RemoveKeyboard
	RemoveKeyboard

	// IgnoreThread is used to ignore the thread when responding to a message via context.
	IgnoreThread
)

// Placeholder is used to set input field placeholder as a send option.
func Placeholder(text string) *communications.SendOptions {
	return &communications.SendOptions{
		ReplyMarkup: &telegram.ReplyMarkup{
			ForceReply:  true,
			Placeholder: text,
		},
	}
}

func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) extractOptions(how []interface{}) *communications.SendOptions {
	opts := &communications.SendOptions{
		ParseMode: b.parseMode,
	}

	for _, prop := range how {
		switch opt := prop.(type) {
		case *communications.SendOptions:
			opts = opt.Copy()
		case *telegram.ReplyMarkup:
			if opt != nil {
				// Create a copy of the ReplyMarkup
				markupCopy := *opt
				opts.ReplyMarkup = &markupCopy
			}
		case *telegram.ReplyParams:
			opts.ReplyParams = opt
		case *telegram.Topic:
			opts.ThreadID = opt.ThreadID
		case Option:
			switch opt {
			case NoPreview:
				opts.DisableWebPagePreview = true
			case Silent:
				opts.DisableNotification = true
			case AllowWithoutReply:
				opts.AllowWithoutReply = true
			case ForceReply:
				if opts.ReplyMarkup == nil {
					opts.ReplyMarkup = &telegram.ReplyMarkup{}
				}
				opts.ReplyMarkup.ForceReply = true
			case OneTimeKeyboard:
				if opts.ReplyMarkup == nil {
					opts.ReplyMarkup = &telegram.ReplyMarkup{}
				}
				opts.ReplyMarkup.OneTimeKeyboard = true
			case RemoveKeyboard:
				if opts.ReplyMarkup == nil {
					opts.ReplyMarkup = &telegram.ReplyMarkup{}
				}
				opts.ReplyMarkup.RemoveKeyboard = true
			case Protected:
				opts.Protected = true
			default:
				panic("telebot: unsupported flag-option")
			}
		case ParseMode:
			opts.ParseMode = opt
		case Entities:
			opts.Entities = opt
		default:
			panic("telebot: unsupported send-option")
		}
	}

	return opts
}

func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) RawEmbedSendOptions(params map[string]string, opt *communications.SendOptions) {
	if opt == nil {
		return
	}

	// Apply censoring to text content in params
	textFields := []string{
		"text", "caption", "question", "explanation", "title", "description",
		"performer", "file_name", "address", // Audio artist, filenames, venue addresses
	}
	for _, field := range textFields {
		if text, exists := params[field]; exists && text != "" {
			params[field] = b.CensorText(text)
		}
	}

	// Handle ReplyParams (takes precedence over ReplyTo)
	if opt.ReplyParams != nil {
		replyParams, _ := json.Marshal(opt.ReplyParams)
		params["reply_parameters"] = string(replyParams)
	} else if opt.ReplyTo != nil && opt.ReplyTo.ID != 0 {
		// Fallback to old reply_to_message_id for backward compatibility
		params["reply_to_message_id"] = strconv.Itoa(opt.ReplyTo.ID)
	}

	if opt.DisableWebPagePreview {
		params["disable_web_page_preview"] = "true"
	}

	if opt.DisableNotification {
		params["disable_notification"] = "true"
	}

	if opt.ParseMode != ModeDefault {
		params["parse_mode"] = string(opt.ParseMode)
	}

	if len(opt.Entities) > 0 {
		delete(params, "parse_mode")
		entities, _ := json.Marshal(opt.Entities)

		if params["caption"] != "" {
			params["caption_entities"] = string(entities)
		} else {
			params["entities"] = string(entities)
		}
	}

	if opt.AllowWithoutReply {
		params["allow_sending_without_reply"] = "true"
	}

	if opt.ReplyMarkup != nil {
		processButtons(opt.ReplyMarkup.InlineKeyboard)
		replyMarkup, _ := json.Marshal(opt.ReplyMarkup)
		params["reply_markup"] = string(replyMarkup)
	}

	if opt.Protected {
		params["protect_content"] = "true"
	}

	if opt.ThreadID != 0 {
		params["message_thread_id"] = strconv.Itoa(opt.ThreadID)
	}

	if opt.HasSpoiler {
		params["has_spoiler"] = "true"
	}

	if opt.BusinessConnectionID != "" {
		params["business_connection_id"] = opt.BusinessConnectionID
	}

	if opt.EffectID != "" {
		params["message_effect_id"] = string(opt.EffectID)
	}
}

func processButtons(keys [][]telegram.InlineButton) {
	if len(keys) < 1 || len(keys[0]) < 1 {
		return
	}

	for i := range keys {
		for j := range keys[i] {
			key := &keys[i][j]
			if key.Unique != "" {
				// Format: "\f<callback_name>|<data>"
				data := key.Data
				if data == "" {
					key.Data = "\f" + key.Unique
				} else {
					key.Data = "\f" + key.Unique + "|" + data
				}
			}
		}
	}
}

// PreviewOptions describes the options used for link preview generation.
type PreviewOptions struct {
	// (Optional) True, if the link preview is disabled.
	Disabled bool `json:"is_disabled"`

	// (Optional) URL to use for the link preview. If empty, then the first URL
	// found in the message text will be used.
	URL string `json:"url"`

	// (Optional) True, if the media in the link preview is supposed to be shrunk;
	// ignored if the URL isn't explicitly specified or media size change.
	// isn't supported for the preview.
	SmallMedia bool `json:"prefer_small_media"`

	// (Optional) True, if the media in the link preview is supposed to be enlarged;
	// ignored if the URL isn't explicitly specified or media size change.
	// isn't supported for the preview.
	LargeMedia bool `json:"prefer_large_media"`

	// (Optional) True, if the link preview must be shown above the message text;
	// otherwise, the link preview will be shown below the message text.
	AboveText bool `json:"show_above_text"`
}

func embedMessages(params map[string]string, msgs []communications.Editable) {
	ids := make([]string, 0, len(msgs))

	_, chatID := msgs[0].MessageSig()
	for _, msg := range msgs {
		msgID, _ := msg.MessageSig()
		ids = append(ids, msgID)
	}

	data, err := json.Marshal(ids)
	if err != nil {
		return
	}

	params["message_ids"] = string(data)
	params["chat_id"] = strconv.FormatInt(chatID, 10)
}
