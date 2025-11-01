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

	// Create a request builder function
	requestBuilder := func(req request.Interface) (*request.Native, error) {
		return request.NewNativeFromRequest(req), nil
	}

	// Create bot settings
	settings := tb.Settings[*request.Native]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "callback_query"},
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

	// Start command
	bot.Handle("/start", func(c *request.Native) error {
		return c.Reply("🎹 Welcome to Keyboards Demo!\n\n" +
			"This bot demonstrates different keyboard types:\n\n" +
			"Reply Keyboards (fixed buttons):\n" +
			"/reply_keyboard - Show reply keyboard\n" +
			"/contact_location - Request contact/location\n" +
			"/remove_keyboard - Remove reply keyboard\n\n" +
			"Inline Keyboards (attached to messages):\n" +
			"/inline_keyboard - Show inline keyboard\n" +
			"/inline_url - Keyboard with URL buttons\n" +
			"/inline_grid - Grid layout keyboard")
	})

	// Reply keyboard - fixed buttons that appear below chat input
	bot.Handle("/reply_keyboard", func(c *request.Native) error {
		// Create reply keyboard with multiple rows
		keyboard := &telegram.ReplyMarkup{
			ReplyKeyboard: [][]telegram.ReplyButton{
				{
					{Text: "🔴 Red"},
					{Text: "🟢 Green"},
					{Text: "🔵 Blue"},
				},
				{
					{Text: "📊 Option A"},
					{Text: "📈 Option B"},
				},
				{
					{Text: "❌ Cancel"},
				},
			},
			ResizeKeyboard:  true,  // Make keyboard smaller
			OneTimeKeyboard: false, // Keep keyboard visible after use
		}

		text := "Choose an option from the keyboard below:\n\n" +
			"Reply keyboards appear at the bottom of the chat. " +
			"When you tap a button, it sends that text as a message."

		return c.Reply(text, tb.Send().WithReplyMarkup(keyboard))
	})

	// Contact and location request buttons
	bot.Handle("/contact_location", func(c *request.Native) error {
		keyboard := &telegram.ReplyMarkup{
			ReplyKeyboard: [][]telegram.ReplyButton{
				{
					{Text: "📱 Share Contact", Contact: true},
				},
				{
					{Text: "📍 Share Location", Location: true},
				},
				{
					{Text: "❌ Cancel"},
				},
			},
			ResizeKeyboard:  true,
			OneTimeKeyboard: true, // Hide keyboard after one use
		}

		text := "Special buttons can request permissions:\n\n" +
			"📱 Contact - Shares your phone number\n" +
			"📍 Location - Shares your current location\n\n" +
			"These work only in private chats for privacy reasons."

		return c.Reply(text, tb.Send().WithReplyMarkup(keyboard))
	})

	// Remove reply keyboard
	bot.Handle("/remove_keyboard", func(c *request.Native) error {
		// RemoveKeyboard field hides the keyboard
		remove := &telegram.ReplyMarkup{
			RemoveKeyboard: true,
		}

		return c.Reply("✅ Reply keyboard removed!", tb.Send().WithReplyMarkup(remove))
	})

	// Inline keyboard - buttons attached to specific messages
	bot.Handle("/inline_keyboard", func(c *request.Native) error {
		// Create inline keyboard with callback data
		keyboard := &telegram.ReplyMarkup{
			InlineKeyboard: [][]telegram.InlineButton{
				{
					{Text: "👍 Like", CallbackData: "like"},
					{Text: "👎 Dislike", CallbackData: "dislike"},
				},
				{
					{Text: "⭐️ Rate 1", CallbackData: "rate_1"},
					{Text: "⭐️ Rate 2", CallbackData: "rate_2"},
					{Text: "⭐️ Rate 3", CallbackData: "rate_3"},
				},
				{
					{Text: "ℹ️ Info", CallbackData: "info"},
				},
			},
		}

		text := "Inline keyboards are attached to messages!\n\n" +
			"Features:\n" +
			"• Don't clutter the chat\n" +
			"• Can have callback actions\n" +
			"• Can update the message\n" +
			"• Work in groups and channels\n\n" +
			"Try clicking the buttons below:"

		return c.Reply(text, tb.Send().WithReplyMarkup(keyboard))
	})

	// Inline keyboard with URL buttons
	bot.Handle("/inline_url", func(c *request.Native) error {
		keyboard := &telegram.ReplyMarkup{
			InlineKeyboard: [][]telegram.InlineButton{
				{
					{Text: "📚 Telegram Bot API", URL: "https://core.telegram.org/bots/api"},
				},
				{
					{Text: "💬 Telegram", URL: "https://telegram.org"},
					{Text: "🐙 GitHub", URL: "https://github.com"},
				},
				{
					// SwitchInlineQuery allows users to share bot in other chats
					{Text: "🔗 Share Bot", SwitchInlineQuery: "Check out this bot!"},
				},
			},
		}

		text := "Inline keyboards can open URLs!\n\n" +
			"Types of inline buttons:\n" +
			"• URL - Opens a web page\n" +
			"• CallbackData - Triggers callback query\n" +
			"• SwitchInlineQuery - Shares bot with query\n" +
			"• WebApp - Opens mini app"

		return c.Reply(text, tb.Send().WithReplyMarkup(keyboard))
	})

	// Grid layout keyboard
	bot.Handle("/inline_grid", func(c *request.Native) error {
		// Create a grid of number buttons
		var rows [][]telegram.InlineButton
		for i := 1; i <= 9; i += 3 {
			row := []telegram.InlineButton{
				{Text: fmt.Sprintf("%d", i), CallbackData: fmt.Sprintf("num_%d", i)},
				{Text: fmt.Sprintf("%d", i+1), CallbackData: fmt.Sprintf("num_%d", i+1)},
				{Text: fmt.Sprintf("%d", i+2), CallbackData: fmt.Sprintf("num_%d", i+2)},
			}
			rows = append(rows, row)
		}
		// Add zero and clear buttons
		rows = append(rows, []telegram.InlineButton{
			{Text: "0", CallbackData: "num_0"},
			{Text: "Clear", CallbackData: "clear"},
		})

		keyboard := &telegram.ReplyMarkup{
			InlineKeyboard: rows,
		}

		return c.Reply("Calculator-style grid keyboard:", tb.Send().WithReplyMarkup(keyboard))
	})

	// Handle callback queries from inline keyboards
	bot.Handle(tb.OnCallback, func(c *request.Native) error {
		callback := c.CallbackQuery()
		if callback == nil {
			return nil
		}

		// Handle different callback data
		var responseText string
		showAlert := false

		switch callback.Data {
		case "like":
			responseText = "👍 You liked this!"
		case "dislike":
			responseText = "👎 You disliked this!"
		case "info":
			responseText = "ℹ️ This is callback data: " + callback.Data
			showAlert = true // Show as popup instead of notification
		default:
			if len(callback.Data) > 4 && callback.Data[:4] == "num_" {
				responseText = "You pressed: " + callback.Data[4:]
			} else if callback.Data[:5] == "rate_" {
				responseText = "⭐️ You rated: " + callback.Data[5:] + " stars"
			} else if callback.Data == "clear" {
				responseText = "🗑️ Cleared!"
			} else {
				responseText = "Button clicked: " + callback.Data
			}
		}

		// Respond to callback (required to remove loading state)
		response := &telegram.CallbackResponse{
			Text:      responseText,
			ShowAlert: showAlert,
		}

		return c.API.RespondToCallback(callback, response)
	})

	// Handle reply keyboard button presses (they arrive as regular messages)
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		// Handle reply keyboard buttons
		switch msg.Text {
		case "🔴 Red", "🟢 Green", "🔵 Blue":
			return c.Reply(fmt.Sprintf("You selected: %s", msg.Text))
		case "📊 Option A":
			return c.Reply("✅ You chose Option A")
		case "📈 Option B":
			return c.Reply("✅ You chose Option B")
		case "❌ Cancel":
			remove := &telegram.ReplyMarkup{RemoveKeyboard: true}
			return c.Reply("❌ Cancelled", tb.Send().WithReplyMarkup(remove))
		default:
			// Not a keyboard button, show help
			if msg.Text != "" && msg.Text[0] != '/' {
				return c.Reply("Send /start to see keyboard examples")
			}
		}

		return nil
	})

	// Handle contact sharing
	bot.Handle(tb.OnContact, func(c *request.Native) error {
		contact := c.Message().Contact
		if contact == nil {
			return nil
		}

		response := fmt.Sprintf("📱 Thank you for sharing your contact!\n\n"+
			"Name: %s %s\n"+
			"Phone: %s",
			contact.FirstName,
			contact.LastName,
			contact.PhoneNumber,
		)

		return c.Reply(response)
	})

	// Handle location sharing
	bot.Handle(tb.OnLocation, func(c *request.Native) error {
		location := c.Message().Location
		if location == nil {
			return nil
		}

		response := fmt.Sprintf("📍 Thank you for sharing your location!\n\n"+
			"Latitude: %.6f\n"+
			"Longitude: %.6f\n\n"+
			"Google Maps: https://www.google.com/maps?q=%.6f,%.6f",
			location.Latitude,
			location.Longitude,
			location.Latitude,
			location.Longitude,
		)

		return c.Reply(response)
	})

	log.Println("🎹 Keyboards bot started! Press Ctrl+C to stop.")
	log.Println("💡 Tip: Reply keyboards work in all chat types, inline keyboards work everywhere")
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
