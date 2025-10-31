package telegram

// ParseMode defines how parse mode should be applied.
type ParseMode string

const (
	ParseModeDefault    ParseMode = ""
	ParseModeMarkdown   ParseMode = "Markdown"
	ParseModeMarkdownV2 ParseMode = "MarkdownV2"
	ParseModeHTML       ParseMode = "HTML"
)
