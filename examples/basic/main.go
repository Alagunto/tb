package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/request"
	"github.com/alagunto/tb/telegram"
)

// Context wraps request.Interface to provide convenient access to native context methods
type Context struct {
	request.Interface
}

// Send is a convenience method that uses SendTo internally
func (c *Context) Send(what interface{}, opts ...communications.SendOptions) error {
	opt := communications.MergeMultipleSendOptions(opts...)
	_, err := c.SendTo(c.Recipient(), what, opt)
	return err
}

// Reply is a convenience method that uses ReplyTo internally
func (c *Context) Reply(what interface{}, opts ...communications.SendOptions) error {
	msg := c.Message()
	if msg == nil {
		return fmt.Errorf("no message to reply to")
	}
	opt := communications.MergeMultipleSendOptions(opts...)
	_, err := c.ReplyTo(msg, what, opt)
	return err
}

// SendAlbum is a convenience method
func (c *Context) SendAlbum(a telegram.Album, opts ...communications.SendOptions) error {
	_, err := c.SendAlbumTo(c.Recipient(), a, opts...)
	return err
}

// EditLast edits the last message from callback or inline query
func (c *Context) EditLast(what interface{}, opts ...communications.SendOptions) error {
	update := c.Update()
	if update.ChosenInlineResult != nil {
		_, err := c.Edit(update.ChosenInlineResult, what, opts...)
		return err
	}
	if update.CallbackQuery != nil {
		_, err := c.Edit(update.CallbackQuery, what, opts...)
		return err
	}
	return fmt.Errorf("nothing to edit")
}

// EditLastCaption edits the caption of the last message
func (c *Context) EditLastCaption(caption string, opts ...communications.SendOptions) error {
	update := c.Update()
	if update.ChosenInlineResult != nil {
		_, err := c.EditCaption(update.ChosenInlineResult, caption, opts...)
		return err
	}
	if update.CallbackQuery != nil {
		_, err := c.EditCaption(update.CallbackQuery, caption, opts...)
		return err
	}
	return fmt.Errorf("nothing to edit")
}

// DeleteLast deletes the current message
func (c *Context) DeleteLast() error {
	msg := c.Message()
	if msg == nil {
		return fmt.Errorf("no message to delete")
	}
	return c.Delete(msg)
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Create a request builder function that wraps native context
	requestBuilder := func(req request.Interface) (*Context, error) {
		return &Context{Interface: req}, nil
	}

	// Create bot settings
	settings := tb.Settings[*Context, func(*Context) error, func(func(*Context) error) func(*Context) error]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:      10 * time.Second,
			AllowedUpdates: []string{"message", "edited_message"},
		},
		OnError: func(err error, ctx *Context, info tb.DebugInfo[*Context, func(*Context) error, func(func(*Context) error) func(*Context) error]) {
			log.Printf("Error: %v", err)
		},
	}

	// Create the bot
	bot, err := tb.NewBot(requestBuilder, settings)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Register handlers for different parse modes
	bot.Handle("/start", func(c *Context) error {
		return c.Reply("Welcome! This bot demonstrates different parse modes.\n\n" +
			"Commands:\n" +
			"/markdown - Test Markdown parse mode\n" +
			"/markdownv2 - Test MarkdownV2 parse mode\n" +
			"/html - Test HTML parse mode\n" +
			"/entities - Test custom entities\n" +
			"/default - Test default (no parse mode)")
	})

	bot.Handle("/markdown", func(c *Context) error {
		text := "*Bold text*\n_Italic text_\n`Code inline`\n[Link](https://telegram.org)"
		opts := communications.NewSendOptions().WithParseMode(telegram.ParseModeMarkdown)
		return c.Reply(text, opts)
	})

	bot.Handle("/markdownv2", func(c *Context) error {
		text := "*Bold\\* text*\n_Italic\\_ text_\n`Code\\` inline`\n[Link](https://telegram\\.org)"
		opts := communications.NewSendOptions().WithParseMode(telegram.ParseModeMarkdownV2)
		return c.Reply(text, opts)
	})

	bot.Handle("/html", func(c *Context) error {
		text := "<b>Bold text</b>\n<i>Italic text</i>\n<code>Code inline</code>\n<a href=\"https://telegram.org\">Link</a>"
		opts := communications.NewSendOptions().WithParseMode(telegram.ParseModeHTML)
		return c.Reply(text, opts)
	})

	bot.Handle("/entities", func(c *Context) error {
		// Send a message with custom entities (more control than parse modes)
		text := "Bold Italic Code Link"
		entities := telegram.Entities{
			{Type: telegram.EntityBold, Offset: 0, Length: 4},
			{Type: telegram.EntityItalic, Offset: 5, Length: 6},
			{Type: telegram.EntityCode, Offset: 12, Length: 4},
			{Type: telegram.EntityTextLink, Offset: 17, Length: 4, URL: "https://telegram.org"},
		}
		opts := communications.NewSendOptions().WithEntities(entities)
		return c.Reply(text, opts)
	})

	bot.Handle("/default", func(c *Context) error {
		// No parse mode - plain text
		return c.Reply("This is plain text with *no* formatting applied.")
	})

	// Handle all text messages
	bot.Handle(tb.OnText, func(c *Context) error {
		// Echo back with some formatting
		text := fmt.Sprintf("You said: %s", c.Text())
		opts := communications.NewSendOptions().WithParseMode(telegram.ParseModeHTML)
		return c.Reply(text, opts)
	})

	log.Println("Bot started! Press Ctrl+C to stop.")
	bot.Start()
}

