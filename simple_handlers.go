package tb

func NewTextReplyAsHTMLHandler(html string, opts ...interface{}) HandlerFunc {
	opts = append(opts, &SendOptions{ParseMode: ModeHTML})
	return func(c Context) error {
		return c.Reply(html, opts...)
	}
}

func NewTextReplyAsMarkdownHandler(markdown string, opts ...interface{}) HandlerFunc {
	opts = append(opts, &SendOptions{ParseMode: ModeMarkdown})
	return func(c Context) error {
		return c.Reply(markdown, opts...)
	}
}

func NewTextReplyHandler(text string, opts ...interface{}) HandlerFunc {
	return func(c Context) error {
		return c.Reply(text, opts...)
	}
}

func NewReplyHandler(what interface{}, opts ...interface{}) HandlerFunc {
	return func(c Context) error {
		return c.Reply(what, opts...)
	}
}

func NewEditHandler(what interface{}, opts ...interface{}) HandlerFunc {
	return func(c Context) error {
		return c.Edit(what, opts...)
	}
}
