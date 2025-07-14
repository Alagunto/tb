package handlers

import "github.com/alagunto/tb"

func NewReplyWithHTMLHandler(html string) tb.HandlerFunc {
	return func(c tb.Context) error {
		return c.Reply(html, &tb.SendOptions{ParseMode: tb.ModeHTML})
	}
}

func NewReplyWithMarkdownHandler(markdown string) tb.HandlerFunc {
	return func(c tb.Context) error {
		return c.Reply(markdown, &tb.SendOptions{ParseMode: tb.ModeMarkdown})
	}
}

func NewTextReplyHandler(text string, opts ...interface{}) tb.HandlerFunc {
	return func(c tb.Context) error {
		return c.Reply(text, opts...)
	}
}
