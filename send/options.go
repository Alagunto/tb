package send

import (
	"github.com/alagunto/tb/telegram"
)

// Options provides complete control over how messages are sent,
// offering an API-complete set of custom properties and options.
//
// Options supports a chainable builder pattern for ease of use:
//
//	send.WithParseMode(telegram.ParseModeHTML).WithoutNotification().WithEffect(effectID)
//
// Or use the method chaining approach:
//
//	opts := send.New().
//		WithParseMode(telegram.ParseModeHTML).
//		WithSilent().
//		WithProtected()
type Options struct {
	// ReplyMarkup describes the reply markup for the message
	ReplyMarkup *telegram.ReplyMarkup

	// DisableWebPagePreview disables previews for links in text messages
	DisableWebPagePreview bool

	// DisableNotification sends the message silently (no notification)
	DisableNotification bool

	// ParseMode controls how client apps render the message
	ParseMode telegram.ParseMode

	// Entities is a list of special entities that appear in message text
	Entities telegram.Entities

	// AllowWithoutReply allows sending messages not as a reply if the replied-to message was deleted
	AllowWithoutReply bool

	// Protected protects the contents of sent message from forwarding and saving
	Protected bool

	// ThreadID supports sending messages to a thread
	ThreadID int

	// ReplyParams describes the message to reply to
	ReplyParams *telegram.ReplyParameters

	// BusinessConnectionID is the unique identifier of the business connection
	BusinessConnectionID string

	// EffectID is the unique identifier of the message effect (private chats only)
	EffectID telegram.EffectID
}

// New creates a new Options instance with default values.
func New() Options {
	return Options{}
}

// Merge combines multiple Options instances, with later options taking precedence.
func (o Options) Merge(others ...Options) Options {
	result := o
	for _, other := range others {
		result = result.merge(other)
	}
	return result
}

func (o Options) merge(other Options) Options {
	// Selectively merge non-zero values from other into o
	// This ensures we don't override existing values with zero values
	
	if other.ReplyMarkup != nil {
		o.ReplyMarkup = other.ReplyMarkup
	}
	
	if other.DisableWebPagePreview {
		o.DisableWebPagePreview = true
	}
	
	if other.DisableNotification {
		o.DisableNotification = true
	}
	
	if other.ParseMode != telegram.ParseModeDefault {
		o.ParseMode = other.ParseMode
	}
	
	if len(other.Entities) > 0 {
		o.Entities = other.Entities
	}
	
	if other.AllowWithoutReply {
		o.AllowWithoutReply = true
	}
	
	if other.Protected {
		o.Protected = true
	}
	
	if other.ThreadID != 0 {
		o.ThreadID = other.ThreadID
	}
	
	if other.ReplyParams != nil {
		o.ReplyParams = other.ReplyParams
	}
	
	if other.BusinessConnectionID != "" {
		o.BusinessConnectionID = other.BusinessConnectionID
	}
	
	if other.EffectID != "" {
		o.EffectID = other.EffectID
	}
	
	return o
}

// InjectIntoMap adds Options parameters directly into the provided params map.
func (o Options) InjectIntoMap(params map[string]any) error {
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
		delete(params, "parse_mode")
		params["entities"] = o.Entities
	}

	if o.AllowWithoutReply {
		params["allow_sending_without_reply"] = true
	}

	if o.ReplyMarkup != nil {
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

// ToMap converts Options to a map representation.
func (o Options) ToMap() map[string]any {
	params := make(map[string]any)
	o.InjectIntoMap(params)
	return params
}

// InjectIntoMethodRequest injects Options fields into a method request
// using type assertions and setter interfaces.
func (o Options) InjectIntoMethodRequest(request interface{}) {
	if request == nil {
		return
	}

	if o.ReplyMarkup != nil {
		if setter, ok := request.(telegram.SetsReplyMarkup); ok {
			setter.SetReplyMarkup(o.ReplyMarkup)
		}
	}

	if o.ParseMode != telegram.ParseModeDefault {
		if setter, ok := request.(telegram.SetsParseMode); ok {
			setter.SetParseMode(o.ParseMode)
		}
	}

	if len(o.Entities) > 0 {
		if setter, ok := request.(telegram.SetsEntities); ok {
			setter.SetEntities(o.Entities)
		}
	}

	if o.BusinessConnectionID != "" {
		if setter, ok := request.(telegram.SetsBusinessConnection); ok {
			setter.SetBusinessConnectionID(o.BusinessConnectionID)
		}
	}

	if o.EffectID != "" {
		if setter, ok := request.(telegram.SetsMessageEffect); ok {
			setter.SetMessageEffectID(string(o.EffectID))
		}
	}
}

// WithParseMode sets the parse mode for the message.
func (o Options) WithParseMode(mode telegram.ParseMode) Options {
	o.ParseMode = mode
	return o
}

// WithSilent sends the message silently (no notification).
func (o Options) WithSilent() Options {
	o.DisableNotification = true
	return o
}

// WithoutNotification sends the message silently (no notification).
// This is an alias for WithSilent() for better readability in chains.
func (o Options) WithoutNotification() Options {
	o.DisableNotification = true
	return o
}

// WithProtected protects the message from forwarding and saving.
func (o Options) WithProtected() Options {
	o.Protected = true
	return o
}

// WithNoPreview disables link preview for the message.
func (o Options) WithNoPreview() Options {
	o.DisableWebPagePreview = true
	return o
}

// WithReplyMarkup sets the reply markup for the message.
func (o Options) WithReplyMarkup(markup *telegram.ReplyMarkup) Options {
	o.ReplyMarkup = markup
	return o
}

// WithEntities sets custom entities for the message.
func (o Options) WithEntities(entities telegram.Entities) Options {
	o.Entities = entities
	return o
}

// WithThreadID sends the message to a specific thread.
func (o Options) WithThreadID(threadID int) Options {
	o.ThreadID = threadID
	return o
}

// WithBusinessConnection sends the message via a business connection.
func (o Options) WithBusinessConnection(id string) Options {
	o.BusinessConnectionID = id
	return o
}

// WithEffectID adds a message effect (for private chats only).
func (o Options) WithEffectID(id telegram.EffectID) Options {
	o.EffectID = id
	return o
}

// WithEffect adds a message effect (for private chats only).
// This is an alias for WithEffectID for better readability.
func (o Options) WithEffect(id telegram.EffectID) Options {
	o.EffectID = id
	return o
}

// WithReplyParams sets reply parameters for the message.
func (o Options) WithReplyParams(params *telegram.ReplyParameters) Options {
	o.ReplyParams = params
	return o
}

// WithAllowWithoutReply allows sending messages not as a reply if the replied-to message was deleted.
func (o Options) WithAllowWithoutReply() Options {
	o.AllowWithoutReply = true
	return o
}

