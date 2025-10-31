package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/request"
	"github.com/alagunto/tb/telegram"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Create a request builder function that wraps native context
	requestBuilder := func(req request.Interface) (*request.Native, error) {
		return request.NewNativeFromRequest(req), nil
	}

	// Create bot settings
	settings := tb.Settings[*request.Native]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "edited_message"},
		},
		OnError: func(err error, ctx *request.Native) {
			log.Printf("Error: %v", err)
		},
	}

	// Create the bot
	bot, err := tb.NewBot(requestBuilder, settings)
	if err != nil {
		log.Fatalf("Failed to create bot: %v", err)
	}

	// Register handlers for different parse modes
	bot.Handle("/start", func(c *request.Native) error {
		return c.Reply("Welcome! This bot demonstrates different parse modes.\n\n" +
			"Commands:\n" +
			"/markdown - Test Markdown parse mode\n" +
			"/markdownv2 - Test MarkdownV2 parse mode\n" +
			"/html - Test HTML parse mode\n" +
			"/entities - Test custom entities\n" +
			"/default - Test default (no parse mode)")
	})

	bot.Handle("/markdown", func(c *request.Native) error {
		text := "*Bold text*\n_Italic text_\n`Code inline`\n[Link](https://telegram.org)"
		return c.Reply(text, tb.SendOptions().WithParseMode(telegram.ParseModeMarkdown))
	})

	bot.Handle("/markdownv2", func(c *request.Native) error {
		text := "*Bold\\* text*\n_Italic\\_ text_\n`Code\\` inline`\n[Link](https://telegram\\.org)"
		return c.Reply(text, tb.SendOptions().WithParseMode(telegram.ParseModeMarkdownV2))
	})

	bot.Handle("/html", func(c *request.Native) error {
		text := "<b>Bold text</b>\n<i>Italic text</i>\n<code>Code inline</code>\n<a href=\"https://telegram.org\">Link</a>"
		opts := tb.SendOptions().WithParseMode(telegram.ParseModeHTML)
		return c.Reply(text, opts)
	})

	bot.Handle("/entities", func(c *request.Native) error {
		// Send a message with custom entities (more control than parse modes)
		text := "Bold Italic Code Link"
		entities := telegram.Entities{
			{Type: telegram.EntityBold, Offset: 0, Length: 4},
			{Type: telegram.EntityItalic, Offset: 5, Length: 6},
			{Type: telegram.EntityCode, Offset: 12, Length: 4},
			{Type: telegram.EntityTextLink, Offset: 17, Length: 4, URL: "https://telegram.org"},
		}
		opts := tb.SendOptions().WithEntities(entities)
		return c.Reply(text, opts)
	})

	bot.Handle("/default", func(c *request.Native) error {
		// No parse mode - plain text
		return c.Reply("This is plain text with *no* formatting applied.")
	})

	// Handle all text messages
	bot.Handle(tb.OnText, func(c *request.Native) error {
		// Echo back with some formatting
		text := fmt.Sprintf("You said: %s", c.Text())
		opts := tb.SendOptions().WithParseMode(telegram.ParseModeHTML)
		return c.Reply(text, opts)
	})

	log.Println("Bot started! Press Ctrl+C to stop.")
	bot.Start()
}
