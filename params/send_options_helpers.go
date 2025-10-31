package params

import "github.com/alagunto/tb/telegram"

// WithParseMode sets the parse mode for the message.
func (o SendOptions) WithParseMode(mode telegram.ParseMode) SendOptions {
	o.ParseMode = mode
	return o
}

// WithSilent sends the message silently (no notification).
func (o SendOptions) WithSilent() SendOptions {
	o.DisableNotification = true
	return o
}

// WithProtected protects the message from forwarding and saving.
func (o SendOptions) WithProtected() SendOptions {
	o.Protected = true
	return o
}

// WithNoPreview disables link preview for the message.
func (o SendOptions) WithNoPreview() SendOptions {
	o.DisableWebPagePreview = true
	return o
}

// WithReplyTo makes
// WithReplyMarkup sets the reply markup for the message.
func (o SendOptions) WithReplyMarkup(markup *telegram.ReplyMarkup) SendOptions {
	o.ReplyMarkup = markup
	return o
}

// WithEntities sets custom entities for the message.
func (o SendOptions) WithEntities(entities telegram.Entities) SendOptions {
	o.Entities = entities
	return o
}

// WithThreadID sends the message to a specific thread.
func (o SendOptions) WithThreadID(threadID int) SendOptions {
	o.ThreadID = threadID
	return o
}

// WithBusinessConnection sends the message via a business connection.
func (o SendOptions) WithBusinessConnection(id string) SendOptions {
	o.BusinessConnectionID = id
	return o
}

// WithEffectID adds a message effect (for private chats only).
func (o SendOptions) WithEffectID(id telegram.EffectID) SendOptions {
	o.EffectID = id
	return o
}

// WithReplyParams sets reply parameters for the message.
func (o SendOptions) WithReplyParams(params *telegram.ReplyParams) SendOptions {
	o.ReplyParams = params
	return o
}

// WithAllowWithoutReply allows sending messages not as a reply if the replied-to message has been deleted.
func (o SendOptions) WithAllowWithoutReply() SendOptions {
	o.AllowWithoutReply = true
	return o
}

func NewSendOptions() SendOptions {
	return SendOptions{}
}
