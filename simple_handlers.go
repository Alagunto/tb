package tb

func NewReplyWithHTMLHandler(html string) HandlerFunc {
	return func(c Context) error {
		return c.Reply(html, &SendOptions{ParseMode: ModeHTML})
	}
}

func NewReplyWithMarkdownHandler(markdown string) HandlerFunc {
	return func(c Context) error {
		return c.Reply(markdown, &SendOptions{ParseMode: ModeMarkdown})
	}
}

func NewTextReplyHandler(text string, opts ...interface{}) HandlerFunc {
	return func(c Context) error {
		return c.Reply(text, opts...)
	}
}
