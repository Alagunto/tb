package telegram

import "fmt"

// ParseMode defines how parse mode should be applied.
type ParseMode string

const (
	ParseModeDefault    ParseMode = ""
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
)

// IsValid checks if the parse mode is one of the supported values.
func (pm ParseMode) IsValid() bool {
	switch pm {
	case ParseModeDefault, ParseModeMarkdown, ParseModeMarkdownV2, ParseModeHTML:
		return true
	default:
		return false
	}
}

// Validate returns an error if the parse mode is not valid.
func (pm ParseMode) Validate() error {
	if !pm.IsValid() {
		return fmt.Errorf("invalid parse mode: %q (must be one of: Markdown, MarkdownV2, HTML, or empty)", pm)
	}
	return nil
}
