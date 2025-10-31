package params

import (
	"github.com/alagunto/tb/send"
)

// SendOptions has most complete control over in what way the message
// must be sent, providing an API-complete set of custom properties
// and options.
//
// Deprecated: Use send.Options instead. This type is maintained as an alias
// for backward compatibility. New code should use the send package which provides
// a more ergonomic chainable API:
//
//	send.WithParseMode(telegram.ParseModeHTML).WithoutNotification()
//
// Despite its power, SendOptions is rather inconvenient to use all
// the way through bot logic, so you might want to consider storing
// and re-using it somewhere or be using Option flags instead.
type SendOptions = send.Options

// Merge combines multiple SendOptions, with later options taking precedence.
// Deprecated: Use send.MergeOptions instead.
func Merge(opts ...SendOptions) SendOptions {
	result := NewSendOptions()
	for _, opt := range opts {
		result = result.Merge(opt)
	}
	return result
}
