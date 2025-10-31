package tb

import (
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/send"
)

// SendOptions creates a new SendOptions instance for use with the builder pattern.
// Deprecated: Use Send() for the new chainable API:
//
//	send.WithParseMode(telegram.ParseModeHTML).WithSilent()
//
// Or use the Send() facade:
//
//	tb.Send().WithParseMode(mode).WithSilent()
func SendOptions(opts ...interface{}) *params.SendOptions {
	opt := params.SendOptions{}
	return &opt
}

// Send creates a new send.Options instance for use with the modern chainable API.
// Example:
//
//	tb.Send().WithParseMode(telegram.ParseModeHTML).WithSilent().WithEffect(effectID)
func Send() send.Options {
	return send.New()
}
