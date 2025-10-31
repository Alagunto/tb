package params

import (
	"fmt"
	"reflect"

	"github.com/alagunto/tb/telegram"
)

// SendOptions has most complete control over in what way the message
// must be sent, providing an API-complete set of custom properties
// and options.
//
// Despite its power, SendOptions is rather inconvenient to use all
// the way through bot logic, so you might want to consider storing
// and re-using it somewhere or be using Option flags instead.
type SendOptions struct {
	// Describes the reply markup for the message
	ReplyMarkup *telegram.ReplyMarkup

	// For text messages, disables previews for links in this message.
	DisableWebPagePreview bool

	// Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
	DisableNotification bool

	// ParseMode controls how client apps render your message.
	ParseMode telegram.ParseMode

	// Entities is a list of special entities that appear in message text, which can be specified instead of parse_mode.
	Entities telegram.Entities

	// AllowWithoutReply allows sending messages not a as reply if the replied-to message has already been deleted.
	AllowWithoutReply bool

	// Protected protects the contents of sent message from forwarding and saving.
	Protected bool

	// ThreadID supports sending messages to a thread.
	ThreadID int

	// ReplyParams Describes the message to reply to
	ReplyParams *telegram.ReplyParams

	// Unique identifier of the business connection
	BusinessConnectionID string

	// Unique identifier of the message effect to be added to the message; for private chats only
	EffectID telegram.EffectID
}

func (o SendOptions) Merge(others ...SendOptions) SendOptions {
	result := o
	for _, other := range others {
		result = result.merge(other)
	}
	return result
}

func (o SendOptions) merge(other SendOptions) SendOptions {
	srcVal := reflect.ValueOf(&other)
	dstVal := reflect.ValueOf(&o)

	for i := 0; i < srcVal.Type().Elem().NumField(); i++ {
		sFld := srcVal.Elem().Field(i)
		dFld := dstVal.Elem().Field(i)

		if dFld.Kind() == reflect.Ptr {
			if dFld.IsNil() {
				// dFld is a nil pointer, we need to allocate a new value for it
				newVal := reflect.New(dFld.Type().Elem())
				dFld.Set(newVal)
			}
		}

		if !dFld.CanSet() {
			if dFld.CanAddr() {
				dFld = dFld.Addr()
			} else {
				panic(fmt.Errorf("field is not settable: %s", dFld.String()))
			}
		}

		dFld.Set(sFld)
	}

	return o
}

// InjectInto adds SendOptions parameters directly into the provided params map.
func (o SendOptions) InjectIntoMap(params map[string]any) error {
	// Handle ReplyParams (takes precedence over ReplyTo)
	if o.ReplyParams != nil {
		params["reply_parameters"] = o.ReplyParams
	}

	if o.DisableWebPagePreview {
		params["disable_web_page_preview"] = true
	}

	if o.DisableNotification {
		params["disable_notification"] = true
	}

	if o.ParseMode != telegram.ParseModeDefault {
		params["parse_mode"] = o.ParseMode
	}

	if len(o.Entities) > 0 {
		// if we have entities specified, parse_mode is not being respected by telegram
		delete(params, "parse_mode")
		params["entities"] = o.Entities
	}

	if o.AllowWithoutReply {
		// Optional. Pass True if the message should be sent even if the specified message to be replied to is not found.
		// Always False for replies in another chat or forum topic.
		// Always True for messages sent on behalf of a business account.
		params["allow_sending_without_reply"] = true
	}

	if o.ReplyMarkup != nil {
		o.ReplyMarkup.InlineKeyboard = o.PrepareButtons(o.ReplyMarkup.InlineKeyboard)
		params["reply_markup"] = o.ReplyMarkup
	}

	if o.Protected {
		params["protect_content"] = true
	}

	if o.ThreadID != 0 {
		params["message_thread_id"] = o.ThreadID
	}

	if o.BusinessConnectionID != "" {
		params["business_connection_id"] = o.BusinessConnectionID
	}

	if o.EffectID != "" {
		params["message_effect_id"] = o.EffectID
	}

	return nil
}

func (o SendOptions) ToMap() map[string]any {
	params := make(map[string]any)
	o.InjectIntoMap(params)
	return params
}

// PrepareButtons processes InlineButtons with Unique field for callback handling
func (o SendOptions) PrepareButtons(keys [][]telegram.InlineButton) [][]telegram.InlineButton {
	if len(keys) < 1 || len(keys[0]) < 1 {
		return keys
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
	return keys
}

// InjectIntoMethodRequest injects SendOptions fields into a method request
// using type assertions and setter interfaces.
func (o SendOptions) InjectIntoMethodRequest(request interface{}) {
	if request == nil {
		return
	}

	// Inject reply markup
	if o.ReplyMarkup != nil {
		if setter, ok := request.(telegram.SetsReplyMarkup); ok {
			o.ReplyMarkup.InlineKeyboard = o.PrepareButtons(o.ReplyMarkup.InlineKeyboard)
			setter.SetReplyMarkup(o.ReplyMarkup)
		}
	}

	// Inject parse mode
	if o.ParseMode != telegram.ParseModeDefault {
		if setter, ok := request.(telegram.SetsParseMode); ok {
			setter.SetParseMode(o.ParseMode)
		}
	}

	// Inject entities
	if len(o.Entities) > 0 {
		if setter, ok := request.(telegram.SetsEntities); ok {
			setter.SetEntities(o.Entities)
		}
	}

	// Inject business connection ID
	if o.BusinessConnectionID != "" {
		if setter, ok := request.(telegram.SetsBusinessConnection); ok {
			setter.SetBusinessConnectionID(o.BusinessConnectionID)
		}
	}
}

func Merge(opts ...SendOptions) SendOptions {
	result := NewSendOptions()
	for _, opt := range opts {
		result = result.Merge(opt)
	}
	return result
}
