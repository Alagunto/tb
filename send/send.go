// Package send provides a fluent, chainable API for configuring message send options.
//
// The send package offers two styles of building send options:
//
// # Standalone Facade Functions
//
// Use standalone functions as entry points for building option chains:
//
//	send.WithParseMode(telegram.ParseModeHTML).WithoutNotification().WithEffect(effectID)
//
// # Method Chaining
//
// Start with New() and chain methods:
//
//	opts := send.New().
//		WithParseMode(telegram.ParseModeHTML).
//		WithSilent().
//		WithProtected()
//
// # Usage Examples
//
// Send a message with HTML formatting and no notification:
//
//	bot.Send("Hello!", send.WithParseMode(telegram.ParseModeHTML).WithSilent())
//
// Send a message with an effect and custom markup:
//
//	c.Reply(response, send.WithEffect(effectID).WithReplyMarkup(markup))
//
// Combine multiple options:
//
//	opts := send.WithParseMode(telegram.ParseModeMarkdown).
//		WithProtected().
//		WithThreadID(threadID).
//		WithNoPreview()
//	bot.Send("Protected message", opts)
//
// # Migration from params Package
//
// If you're migrating from the params package:
//
//	// Old way
//	opts := params.SendOptions{}.WithParseMode(mode).WithSilent()
//
//	// New way (preferred)
//	opts := send.WithParseMode(mode).WithSilent()
//
// The old params.SendOptions is still available as a type alias for backward compatibility.
package send

