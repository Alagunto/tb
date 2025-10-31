package tb

import "github.com/alagunto/tb/params"

// SendOptions creates a new SendOptions instance for use with the builder pattern.
// Example: SendOptions().WithParseMode(ModeMarkdown).WithSilent()
func SendOptions(opts ...interface{}) *params.SendOptions {
	return &params.SendOptions{}
}
