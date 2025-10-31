package params

// NewSendOptions creates a new SendOptions instance.
// Deprecated: Use send.New() instead for the new chainable API.
func NewSendOptions() SendOptions {
	return SendOptions{}
}
