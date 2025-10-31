package send

import "github.com/alagunto/tb/telegram"

// Chainable facade functions that provide a fluent API for building send options.
// These functions can be used as entry points for building Options chains:
//
//	send.WithParseMode(telegram.ParseModeHTML).WithSilent().WithEffect(effectID)

// WithParseMode creates Options with the specified parse mode.
func WithParseMode(mode telegram.ParseMode) Options {
	return New().WithParseMode(mode)
}

// WithSilent creates Options with silent mode enabled (no notification).
func WithSilent() Options {
	return New().WithSilent()
}

// WithoutNotification creates Options with notifications disabled.
// This is an alias for WithSilent() for better readability.
func WithoutNotification() Options {
	return New().WithSilent()
}

// WithProtected creates Options with content protection enabled.
func WithProtected() Options {
	return New().WithProtected()
}

// WithNoPreview creates Options with link preview disabled.
func WithNoPreview() Options {
	return New().WithNoPreview()
}

// WithReplyMarkup creates Options with the specified reply markup.
func WithReplyMarkup(markup *telegram.ReplyMarkup) Options {
	return New().WithReplyMarkup(markup)
}

// WithEntities creates Options with custom entities.
func WithEntities(entities telegram.Entities) Options {
	return New().WithEntities(entities)
}

// WithThreadID creates Options for sending to a specific thread.
func WithThreadID(threadID int) Options {
	return New().WithThreadID(threadID)
}

// WithBusinessConnection creates Options for sending via a business connection.
func WithBusinessConnection(id string) Options {
	return New().WithBusinessConnection(id)
}

// WithEffect creates Options with the specified message effect.
// This is an alias for WithEffectID() for better readability.
func WithEffect(id telegram.EffectID) Options {
	return New().WithEffectID(id)
}

// WithEffectID creates Options with the specified message effect ID.
func WithEffectID(id telegram.EffectID) Options {
	return New().WithEffectID(id)
}

// WithReplyParams creates Options with reply parameters.
func WithReplyParams(params *telegram.ReplyParameters) Options {
	return New().WithReplyParams(params)
}

// WithAllowWithoutReply creates Options allowing messages to be sent even if
// the replied-to message was deleted.
func WithAllowWithoutReply() Options {
	return New().WithAllowWithoutReply()
}

// MergeOptions combines multiple Options instances, with later options taking precedence.
// This is a convenience function for merging options from different sources.
func MergeOptions(opts ...Options) Options {
	result := New()
	for _, opt := range opts {
		result = result.Merge(opt)
	}
	return result
}

