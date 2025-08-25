package tb

func RespondWithHTML[Ctx ContextInterface, HandlerFunc func(c Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](html string, opts ...interface{}) HandlerFunc {
	opts = append(opts, ParseMode(ModeHTML))
	return func(c Ctx) error {
		return c.Reply(html, opts...)
	}
}

func EditWithHTML[Ctx ContextInterface, HandlerFunc func(c Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](html string, opts ...interface{}) HandlerFunc {
	opts = append(opts, ParseMode(ModeHTML))
	return func(c Ctx) error {
		return c.Edit(html, opts...)
	}
}

func ReplyWithMarkdown[Ctx ContextInterface, HandlerFunc func(c Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](markdown string, opts ...interface{}) HandlerFunc {
	opts = append(opts, ParseMode(ModeMarkdown))
	return func(c Ctx) error {
		return c.Reply(markdown, opts...)
	}
}

func EditWithMarkdown[Ctx ContextInterface, HandlerFunc func(c Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](markdown string, opts ...interface{}) HandlerFunc {
	opts = append(opts, ParseMode(ModeMarkdown))
	return func(c Ctx) error {
		return c.Edit(markdown, opts...)
	}
}

func ReplyWithText[Ctx ContextInterface, HandlerFunc func(c Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](text string, opts ...interface{}) HandlerFunc {
	return func(c Ctx) error {
		return c.Reply(text, opts...)
	}
}

func EditWithText[Ctx ContextInterface, HandlerFunc func(c Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](text string, opts ...interface{}) HandlerFunc {
	return func(c Ctx) error {
		return c.Edit(text, opts...)
	}
}

func ReplyWith[Ctx ContextInterface, HandlerFunc func(c Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](what interface{}, opts ...interface{}) HandlerFunc {
	return func(c Ctx) error {
		return c.Reply(what, opts...)
	}
}

func EditWith[Ctx ContextInterface, HandlerFunc func(c Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](what interface{}, opts ...interface{}) HandlerFunc {
	return func(c Ctx) error {
		return c.Edit(what, opts...)
	}
}
