package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/request"
	"github.com/alagunto/tb/telegram"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Create a request builder function
	requestBuilder := func(req request.Interface) (*request.Native, error) {
		return request.NewNativeFromRequest(req), nil
	}

	// Create bot settings
	settings := tb.Settings[*request.Native]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "callback_query", "inline_query"},
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

	// Start command with inline keyboard
	bot.Handle("/start", func(c *request.Native) error {
		keyboard := &telegram.ReplyMarkup{
			InlineKeyboard: [][]telegram.InlineButton{
				{
					{Text: "Button 1", Unique: "btn1"},
					{Text: "Button 2", Unique: "btn2"},
				},
				{
					{Text: "URL Button", URL: "https://telegram.org"},
					{Text: "Inline Query", InlineQuery: "test query"},
				},
			},
		}
		opts := communications.NewSendOptions().WithReplyMarkup(keyboard)
		return c.Reply("Welcome! This bot demonstrates advanced features.\n\n"+
			"Commands:\n"+
			"/keyboard - Show inline keyboard\n"+
			"/reply_keyboard - Show reply keyboard\n"+
			"/album - Send media album\n"+
			"/edit - Test message editing\n"+
			"/delete - Test message deletion\n"+
			"/forward - Forward a message\n"+
			"/copy - Copy a message\n"+
			"/location - Send location\n"+
			"/poll - Create a poll\n"+
			"/reaction - Add reaction to message", opts)
	})

	// Inline keyboard example
	bot.Handle("/keyboard", func(c *request.Native) error {
		keyboard := &telegram.ReplyMarkup{
			InlineKeyboard: [][]telegram.InlineButton{
				{
					{Text: "Option A", Unique: "option_a"},
					{Text: "Option B", Unique: "option_b"},
				},
				{
					{Text: "Option C", Unique: "option_c"},
				},
			},
		}
		opts := communications.NewSendOptions().WithReplyMarkup(keyboard)
		return c.Reply("Choose an option:", opts)
	})

	// Handle callback queries from inline buttons
	bot.Handle("\fbtn1", func(c *request.Native) error {
		// Answer the callback query
		callback := c.CallbackQuery()
		if callback != nil {
			resp := &telegram.CallbackResponse{
				Text:      "You clicked Button 1!",
				ShowAlert: false,
			}
			c.RespondToCallback(callback, resp)
		}
		return c.EditLastCaption("You selected Button 1!")
	})

	bot.Handle("\fbtn2", func(c *request.Native) error {
		callback := c.CallbackQuery()
		if callback != nil {
			resp := &telegram.CallbackResponse{
				Text:      "You clicked Button 2!",
				ShowAlert: false,
			}
			c.RespondToCallback(callback, resp)
		}
		return c.EditLastCaption("You selected Button 2!")
	})

	bot.Handle("\foption_a", func(c *request.Native) error {
		callback := c.CallbackQuery()
		if callback != nil {
			resp := &telegram.CallbackResponse{
				Text: "Selected Option A",
			}
			c.RespondToCallback(callback, resp)
		}
		return c.EditLast("You selected Option A!")
	})

	bot.Handle("\foption_b", func(c *request.Native) error {
		callback := c.CallbackQuery()
		if callback != nil {
			resp := &telegram.CallbackResponse{
				Text: "Selected Option B",
			}
			c.RespondToCallback(callback, resp)
		}
		return c.EditLast("You selected Option B!")
	})

	bot.Handle("\foption_c", func(c *request.Native) error {
		callback := c.CallbackQuery()
		if callback != nil {
			resp := &telegram.CallbackResponse{
				Text:      "Selected Option C",
				ShowAlert: true,
			}
			c.RespondToCallback(callback, resp)
		}
		return c.EditLast("You selected Option C!")
	})

	// Reply keyboard example
	bot.Handle("/reply_keyboard", func(c *request.Native) error {
		keyboard := &telegram.ReplyMarkup{
			ReplyKeyboard: [][]telegram.ReplyButton{
				{
					{Text: "Button 1"},
					{Text: "Button 2"},
				},
				{
					{Text: "Button 3"},
					{Text: "Request Location", Location: true},
				},
			},
			ResizeKeyboard:  true,
			OneTimeKeyboard: true,
		}
		opts := communications.NewSendOptions().WithReplyMarkup(keyboard)
		return c.Reply("Reply keyboard activated! Press any button.", opts)
	})

	// Media album example
	bot.Handle("/album", func(c *request.Native) error {
		album := telegram.InputAlbum{
			Media: []outgoing.Content{
				&telegram.Photo{
					Source:  files.UseURL("https://via.placeholder.com/300x200?text=Photo+1"),
					Caption: "First photo in the album",
				},
				&telegram.Photo{
					Source: files.UseURL("https://via.placeholder.com/300x200?text=Photo+2"),
				},
				&telegram.Photo{
					Source: files.UseURL("https://via.placeholder.com/300x200?text=Photo+3"),
				},
			},
		}
		return c.SendAlbum(album)
	})

	// Message editing example
	bot.Handle("/edit", func(c *request.Native) error {
		// Send initial message
		msg, err := c.SendTo(c.Recipient(), "This message will be edited in 2 seconds...")
		if err != nil {
			return err
		}

		// Wait and edit
		time.Sleep(2 * time.Second)
		opts := communications.NewSendOptions().WithParseMode(telegram.ParseModeHTML)
		_, err = c.Edit(msg, "Message edited! <b>Success!</b>", opts)
		return err
	})

	// Message deletion example
	bot.Handle("/delete", func(c *request.Native) error {
		msg, err := c.SendTo(c.Recipient(), "This message will be deleted in 3 seconds...")
		if err != nil {
			return err
		}

		time.Sleep(3 * time.Second)
		return c.Delete(msg)
	})

	// Forward message example
	bot.Handle("/forward", func(c *request.Native) error {
		// Forward the user's message back to them
		msg := c.Message()
		if msg == nil {
			return c.Reply("Please reply to a message to forward it.")
		}
		_, err := c.ForwardTo(c.Recipient(), msg)
		return err
	})

	// Copy message example
	bot.Handle("/copy", func(c *request.Native) error {
		// Copy the user's message
		msg := c.Message()
		if msg == nil {
			return c.Reply("Please reply to a message to copy it.")
		}
		opts := communications.NewSendOptions().WithParseMode(telegram.ParseModeHTML)
		_, err := c.CopyTo(c.Recipient(), msg, opts)
		return err
	})

	// Location example
	bot.Handle("/location", func(c *request.Native) error {
		location := &telegram.Location{
			Lat: 37.7749,
			Lng: -122.4194,
		}
		return c.Send(location)
	})

	// Poll example
	bot.Handle("/poll", func(c *request.Native) error {
		// Note: Poll creation requires specific Telegram API methods
		// This is a placeholder showing the concept
		return c.Reply("Poll creation would go here. Check Telegram API documentation for poll creation.")
	})

	// Reaction example
	bot.Handle("/reaction", func(c *request.Native) error {
		// Send a message first
		msg, err := c.SendTo(c.Recipient(), "React to this message!")
		if err != nil {
			return err
		}

		// Note: Reactions are typically added by users, but you can react programmatically
		// using the Telegram API. This is a placeholder.
		return c.Reply(fmt.Sprintf("Message sent with ID: %d. You can react to it!", msg.ID))
	})

	// Handle inline queries
	bot.Handle(tb.OnQuery, func(c *request.Native) error {
		query := c.InlineQuery()
		if query == nil {
			return nil
		}

		article1 := &telegram.InlineQueryArticleResult{
			InlineQueryResultBase: telegram.InlineQueryResultBase{
				ID: "1",
			},
			Title:       "Result 1",
			Text:        "You selected Result 1",
			Description: "This is the first result",
		}
		article2 := &telegram.InlineQueryArticleResult{
			InlineQueryResultBase: telegram.InlineQueryResultBase{
				ID: "2",
			},
			Title:       "Result 2",
			Text:        "You selected Result 2",
			Description: "This is the second result",
		}
		results := []telegram.InlineQueryResult{
			article1,
			article2,
		}

		response := &telegram.InlineQueryResponse{
			Results: results,
		}

		return c.AnswerInlineQuery(query, response)
	})

	// Handle callback queries (generic handler)
	bot.Handle(tb.OnCallback, func(c *request.Native) error {
		callback := c.CallbackQuery()
		if callback != nil {
			resp := &telegram.CallbackResponse{
				Text: "Callback received!",
			}
			return c.RespondToCallback(callback, resp)
		}
		return nil
	})

	log.Println("Bot started! Press Ctrl+C to stop.")
	bot.Start()
}
