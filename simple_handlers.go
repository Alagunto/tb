package tb

func HTMLReply(html string, opts ...interface{}) HandlerFunc {
	opts = append(opts, ParseMode(ModeHTML))
	return func(c Context) error {
		return c.Reply(html, opts...)
	}
}

func MarkdownReply(markdown string, opts ...interface{}) HandlerFunc {
	opts = append(opts, ParseMode(ModeMarkdown))
	return func(c Context) error {
		return c.Reply(markdown, opts...)
	}
}

func TextReply(text string, opts ...interface{}) HandlerFunc {
	return func(c Context) error {
		return c.Reply(text, opts...)
	}
}

func ReplyHandler(what interface{}, opts ...interface{}) HandlerFunc {
	return func(c Context) error {
		return c.Reply(what, opts...)
	}
}

func EditHandler(what interface{}, opts ...interface{}) HandlerFunc {
	return func(c Context) error {
		return c.Edit(what, opts...)
	}
}
