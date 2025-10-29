package tb

import "github.com/alagunto/tb/communications"

// SendOptions creates a new SendOptions instance for use with the builder pattern.
// Example: SendOptions().WithParseMode(ModeMarkdown).WithSilent()
func SendOptions(opts ...interface{}) *communications.SendOptions {
	return &communications.SendOptions{}
}
