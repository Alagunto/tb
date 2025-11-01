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

// IMPORTANT: To use this example, your bot must be an administrator in a channel
// Add your channel ID here (e.g., @yourchannel or -1001234567890)
var targetChannelID = os.Getenv("CHANNEL_ID")

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	if targetChannelID == "" {
		log.Fatal("CHANNEL_ID environment variable is required (e.g., @yourchannel or -1001234567890)")
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
			AllowedUpdates: []string{"message", "channel_post", "edited_channel_post"},
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

	// Get channel recipient
	channel, err := bot.ChatByID(targetChannelID)
	if err != nil {
		log.Fatalf("Failed to get channel (make sure bot is admin): %v", err)
	}

	log.Printf("Connected to channel: %s (ID: %s)", channel.Title, targetChannelID)

	// Start command (works in private chat with bot)
	bot.Handle("/start", func(c *request.Native) error {
		return c.Reply("📢 Welcome to Channel Management Demo!\n\n" +
			"This bot manages the channel: " + channel.Title + "\n\n" +
			"Commands:\n" +
			"/post <text> - Post a message to the channel\n" +
			"/post_html - Post with HTML formatting\n" +
			"/post_buttons - Post with inline buttons\n" +
			"/post_photo <url> - Post a photo\n" +
			"/edit - Edit last message (reply to it)\n" +
			"/delete - Delete a message (reply to it)\n" +
			"/pin - Pin a channel message (reply to it)\n" +
			"/unpin - Unpin all messages\n" +
			"/stats - Get channel statistics\n\n" +
			"⚠️ Bot must be an admin in the channel!")
	})

	// Post text message to channel
	bot.Handle("/post", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /post <your message>\n\nExample:\n/post Hello, channel members!")
		}

		text := joinArgs(args)
		
		// Send to channel
		_, err := bot.Send(channel, text)
		if err != nil {
			return c.Reply("Failed to post: " + err.Error())
		}

		return c.Reply("✅ Message posted to channel!")
	})

	// Post with HTML formatting
	bot.Handle("/post_html", func(c *request.Native) error {
		text := "<b>Important Announcement</b>\n\n" +
			"This is a <i>formatted</i> message with:\n" +
			"• <b>Bold text</b>\n" +
			"• <i>Italic text</i>\n" +
			"• <code>Code blocks</code>\n" +
			"• <a href=\"https://telegram.org\">Links</a>\n\n" +
			"<blockquote>Quoted text for emphasis</blockquote>"

		opts := tb.Send().WithParseMode(telegram.ParseModeHTML)
		_, err := bot.Send(channel, text, opts)
		if err != nil {
			return c.Reply("Failed to post: " + err.Error())
		}

		return c.Reply("✅ Formatted message posted to channel!")
	})

	// Post with inline buttons
	bot.Handle("/post_buttons", func(c *request.Native) error {
		text := "📢 Check out our resources!"

		keyboard := &telegram.InlineKeyboardMarkup{
			InlineKeyboard: [][]telegram.InlineButton{
				{
					{Text: "🌐 Website", URL: "https://telegram.org"},
					{Text: "📚 Docs", URL: "https://core.telegram.org/bots"},
				},
				{
					{Text: "💬 Community", URL: "https://t.me/telegram"},
				},
			},
		}

		opts := tb.Send().WithReplyMarkup(keyboard)
		_, err := bot.Send(channel, text, opts)
		if err != nil {
			return c.Reply("Failed to post: " + err.Error())
		}

		return c.Reply("✅ Message with buttons posted to channel!")
	})

	// Post photo to channel
	bot.Handle("/post_photo", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /post_photo <image_url>\n\nExample:\n/post_photo https://picsum.photos/800/600")
		}

		photoURL := args[0]
		photo := &telegram.InputMediaPhoto{
			Type:    "photo",
			Media:   photoURL,
			Caption: "📷 Posted by bot",
		}

		_, err := bot.Send(channel, photo)
		if err != nil {
			return c.Reply("Failed to post photo: " + err.Error())
		}

		return c.Reply("✅ Photo posted to channel!")
	})

	// Edit channel message
	bot.Handle("/edit", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil {
			return c.Reply("❌ Please reply to a channel message forward with /edit <new text>")
		}

		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: Reply to a message with /edit <new text>")
		}

		newText := joinArgs(args)
		
		// Edit the message
		// Note: msg.ReplyTo should be a forwarded channel message
		if err := c.API.Edit(msg.ReplyTo, newText); err != nil {
			return c.Reply("Failed to edit: " + err.Error() + "\n\nMake sure you reply to a forwarded channel message.")
		}

		return c.Reply("✅ Message edited in channel!")
	})

	// Delete channel message
	bot.Handle("/delete", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil {
			return c.Reply("❌ Please reply to a channel message forward with /delete")
		}

		// Delete the message
		if err := c.API.Delete(msg.ReplyTo); err != nil {
			return c.Reply("Failed to delete: " + err.Error() + "\n\nMake sure you reply to a forwarded channel message.")
		}

		return c.Reply("✅ Message deleted from channel!")
	})

	// Pin channel message
	bot.Handle("/pin", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil {
			return c.Reply("❌ Please reply to a channel message forward with /pin")
		}

		// Pin the message
		if err := c.API.Pin(msg.ReplyTo); err != nil {
			return c.Reply("Failed to pin: " + err.Error())
		}

		return c.Reply("✅ Message pinned in channel!")
	})

	// Unpin all channel messages
	bot.Handle("/unpin", func(c *request.Native) error {
		// Unpin all messages in channel
		if err := c.API.Unpin(channel); err != nil {
			return c.Reply("Failed to unpin: " + err.Error())
		}

		return c.Reply("✅ All messages unpinned in channel!")
	})

	// Get channel statistics
	bot.Handle("/stats", func(c *request.Native) error {
		// Get chat info
		chatInfo, err := bot.ChatByID(targetChannelID)
		if err != nil {
			return c.Reply("Failed to get channel info: " + err.Error())
		}

		stats := fmt.Sprintf("📊 Channel Statistics\n\n"+
			"Name: %s\n"+
			"Username: @%s\n"+
			"Type: %s\n"+
			"ID: %d\n",
			chatInfo.Title,
			chatInfo.Username,
			chatInfo.Type,
			chatInfo.ID,
		)

		if chatInfo.Description != "" {
			stats += fmt.Sprintf("Description: %s\n", chatInfo.Description)
		}

		return c.Reply(stats)
	})

	// Handle channel posts (messages posted in the channel)
	bot.Handle(tb.OnChannelPost, func(c *request.Native) error {
		channelPost := c.Message()
		if channelPost == nil {
			return nil
		}

		log.Printf("New channel post in %s: %s",
			channelPost.Chat.Title,
			channelPost.Text,
		)

		return nil
	})

	// Handle edited channel posts
	bot.Handle(tb.OnEditedChannelPost, func(c *request.Native) error {
		editedPost := c.Message()
		if editedPost == nil {
			return nil
		}

		log.Printf("Channel post edited in %s: %s",
			editedPost.Chat.Title,
			editedPost.Text,
		)

		return nil
	})

	// Handle regular text messages (in private chat with bot)
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" || msg.Text[0] == '/' {
			return nil
		}

		return c.Reply("Send /start to see channel management commands!")
	})

	log.Println("📢 Channel management bot started! Press Ctrl+C to stop.")
	log.Printf("Managing channel: %s", channel.Title)
	log.Println("⚠️  Make sure the bot is an administrator in the channel")
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}

// Helper function to join command arguments
func joinArgs(args []string) string {
	result := ""
	for i, arg := range args {
		if i > 0 {
			result += " "
		}
		result += arg
	}
	return result
}

