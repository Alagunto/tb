package telegram

// InputMedia is an interface for all input media types.
type InputMedia interface {
	Field() (map[string]any, error)
}
