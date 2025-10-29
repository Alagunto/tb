package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/files"
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

// InputPhoto wraps telegram.Photo to implement Inputtable interface
type InputPhoto struct {
	*telegram.Photo
}

// InputMedia implements Inputtable interface
func (p *InputPhoto) InputMedia() *telegram.InputMedia {
	im := &telegram.InputMedia{
		Type: "photo",
		// Media will be set by SendAlbumTo to "attach://fileN"
	}
	if p.Photo != nil {
		im.Caption = p.Caption
		im.CaptionAbove = p.CaptionAbove
		im.HasSpoiler = p.HasSpoiler
	}
	return im
}

// MediaFile implements Inputtable interface
func (p *InputPhoto) MediaFile() interface{} {
	if p.Photo != nil {
		return p.Source
	}
	return nil
}

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Create a request builder function
	requestBuilder := func(req request.Interface) (*Context, error) {
		return &Context{Interface: req}, nil
	}

	// Create bot settings
	settings := tb.Settings[*Context, func(*Context) error, func(func(*Context) error) func(*Context) error]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "callback_query", "inline_query"},
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

	// Start command with inline keyboard
	bot.Handle("/start", func(c *Context) error {
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
	bot.Handle("/keyboard", func(c *Context) error {
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
	bot.Handle("\fbtn1", func(c *Context) error {
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

	bot.Handle("\fbtn2", func(c *Context) error {
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

	bot.Handle("\foption_a", func(c *Context) error {
		callback := c.CallbackQuery()
		if callback != nil {
			resp := &telegram.CallbackResponse{
				Text: "Selected Option A",
			}
			c.RespondToCallback(callback, resp)
		}
		return c.EditLast("You selected Option A!")
	})

	bot.Handle("\foption_b", func(c *Context) error {
		callback := c.CallbackQuery()
		if callback != nil {
			resp := &telegram.CallbackResponse{
				Text: "Selected Option B",
			}
			c.RespondToCallback(callback, resp)
		}
		return c.EditLast("You selected Option B!")
	})

	bot.Handle("\foption_c", func(c *Context) error {
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
	bot.Handle("/reply_keyboard", func(c *Context) error {
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
	bot.Handle("/album", func(c *Context) error {
		album := telegram.Album{
			&InputPhoto{Photo: &telegram.Photo{
				Source:  files.UseURL("https://via.placeholder.com/300x200?text=Photo+1"),
				Caption: "First photo in the album",
			}},
			&InputPhoto{Photo: &telegram.Photo{
				Source: files.UseURL("https://via.placeholder.com/300x200?text=Photo+2"),
			}},
			&InputPhoto{Photo: &telegram.Photo{
				Source: files.UseURL("https://via.placeholder.com/300x200?text=Photo+3"),
			}},
		}
		return c.SendAlbum(album)
	})

	// Message editing example
	bot.Handle("/edit", func(c *Context) error {
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
	bot.Handle("/delete", func(c *Context) error {
		msg, err := c.SendTo(c.Recipient(), "This message will be deleted in 3 seconds...")
		if err != nil {
			return err
		}

		time.Sleep(3 * time.Second)
		return c.Delete(msg)
	})

	// Forward message example
	bot.Handle("/forward", func(c *Context) error {
		// Forward the user's message back to them
		msg := c.Message()
		if msg == nil {
			return c.Reply("Please reply to a message to forward it.")
		}
		_, err := c.ForwardTo(c.Recipient(), msg)
		return err
	})

	// Copy message example
	bot.Handle("/copy", func(c *Context) error {
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
	bot.Handle("/location", func(c *Context) error {
		location := &telegram.Location{
			Lat: 37.7749,
			Lng: -122.4194,
		}
		return c.Send(location)
	})

	// Poll example
	bot.Handle("/poll", func(c *Context) error {
		// Note: Poll creation requires specific Telegram API methods
		// This is a placeholder showing the concept
		return c.Reply("Poll creation would go here. Check Telegram API documentation for poll creation.")
	})

	// Reaction example
	bot.Handle("/reaction", func(c *Context) error {
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
	bot.Handle(tb.OnQuery, func(c *Context) error {
		query := c.InlineQuery()
		if query == nil {
			return nil
		}

		results := []telegram.InlineQueryResult{
			{
				Type:  "article",
				ID:    "1",
				Title: "Result 1",
				InputMessageContent: &telegram.InputText{
					Text:      "You selected Result 1",
					ParseMode: telegram.ParseModeHTML,
				},
			},
			{
				Type:  "article",
				ID:    "2",
				Title: "Result 2",
				InputMessageContent: &telegram.InputText{
					Text:      "You selected Result 2",
					ParseMode: telegram.ParseModeHTML,
				},
			},
		}

		response := &telegram.QueryResponse{
			Results: results,
		}

		return c.AnswerInlineQuery(query, response)
	})

	// Handle callback queries (generic handler)
	bot.Handle(tb.OnCallback, func(c *Context) error {
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
