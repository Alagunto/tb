package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/request"
	"github.com/alagunto/tb/telegram"
)

// Admin IDs - replace with actual admin user IDs
var adminIDs = map[int64]bool{
	975559469: true, // Replace with actual admin ID
}

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

	// Helper function to check if user is admin
	isAdmin := func(userID int64) bool {
		return adminIDs[userID]
	}

	// Start command
	bot.Handle("/start", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil {
			return c.Reply("Could not identify sender")
		}

		greeting := fmt.Sprintf("👋 Hello, %s!\n\n", sender.FirstName)

		if isAdmin(sender.ID) {
			greeting += "You are an admin. Available commands:\n" +
				"/info - Show chat information\n" +
				"/stats - Show bot statistics\n" +
				"/ban - Ban a user (reply to their message)\n" +
				"/unban - Unban a user (reply to their message)\n" +
				"/pin - Pin a message (reply to it)\n" +
				"/unpin - Unpin chat messages"
		} else {
			greeting += "Available commands:\n" +
				"/help - Show this help message\n" +
				"/echo <text> - Echo your message"
		}

		return c.Reply(greeting)
	})

	// Info command - shows chat information
	bot.Handle("/info", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return c.Reply("No message context available")
		}

		chat := msg.Chat
		sender := msg.Sender

		info := fmt.Sprintf("📊 Chat Information\n\n"+
			"Chat ID: %d\n"+
			"Chat Type: %s\n"+
			"Chat Title: %s\n\n"+
			"Your ID: %d\n"+
			"Your Username: @%s\n"+
			"Your Name: %s %s",
			chat.ID,
			chat.Type,
			chat.Title,
			sender.ID,
			sender.Username,
			sender.FirstName,
			sender.LastName,
		)

		return c.Reply(info)
	})

	// Stats command - admin only
	bot.Handle("/stats", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ This command is only available to admins")
		}

		msg := c.Message()
		if msg == nil {
			return c.Reply("No message context")
		}

		stats := fmt.Sprintf("📈 Bot Statistics\n\n"+
			"Message ID: %d\n"+
			"Chat ID: %d\n"+
			"Sender ID: %d",
			msg.ID,
			msg.Chat.ID,
			sender.ID,
		)

		return c.Reply(stats)
	})

	// Ban command - admin only
	bot.Handle("/ban", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ This command is only available to admins")
		}

		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil || msg.ReplyTo.Sender == nil {
			return c.Reply("Please reply to a message from the user you want to ban")
		}

		userToBan := msg.ReplyTo.Sender
		if isAdmin(userToBan.ID) {
			return c.Reply("❌ Cannot ban an admin")
		}

		// Note: In v5, ban/kick methods would need to be implemented in bot.go
		// This is a placeholder showing the intended API usage
		return c.Reply(fmt.Sprintf("⚠️ Ban command received for user @%s (ID: %d)\n\n"+
			"Note: Ban functionality needs to be implemented in bot.go",
			userToBan.Username, userToBan.ID))
	})

	// Pin command - admin only
	bot.Handle("/pin", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ This command is only available to admins")
		}

		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil {
			return c.Reply("Please reply to a message you want to pin")
		}

		// Pin the replied-to message
		if err := c.API.Pin(msg.ReplyTo); err != nil {
			return c.Reply("Failed to pin message: " + err.Error())
		}

		return c.Reply("✅ Message pinned!")
	})

	// Unpin command - admin only
	bot.Handle("/unpin", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ This command is only available to admins")
		}

		msg := c.Message()
		if msg == nil {
			return c.Reply("No message context")
		}

		if err := c.API.Unpin(msg.Chat); err != nil {
			return c.Reply("Failed to unpin: " + err.Error())
		}

		return c.Reply("✅ Messages unpinned!")
	})

	// Echo command - available to everyone
	bot.Handle("/echo", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /echo <text>")
		}

		text := strings.Join(args, " ")
		opts := tb.SendOptions().WithParseMode(telegram.ParseModeHTML)

		response := fmt.Sprintf("🔊 <b>Echo:</b>\n%s", text)
		return c.Reply(response, opts)
	})

	// Handle all text messages
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		// Don't respond to commands
		if strings.HasPrefix(msg.Text, "/") {
			return nil
		}

		sender := c.Sender()
		if sender == nil {
			return nil
		}

		response := fmt.Sprintf("💬 %s said: %s", sender.FirstName, msg.Text)
		return c.Reply(response)
	})

	// Handle callback queries (for inline buttons)
	bot.Handle(tb.OnCallback, func(c *request.Native) error {
		callback := c.CallbackQuery()
		if callback == nil {
			return nil
		}

		// Respond to the callback
		response := &telegram.CallbackResponse{
			Text:      fmt.Sprintf("You clicked: %s", callback.Data),
			ShowAlert: false,
		}

		return c.API.RespondToCallback(callback, response)
	})

	log.Println("Bot started! Press Ctrl+C to stop.")
	log.Printf("Admins: %v", getAdminIDs())
	bot.Start()
}

func getAdminIDs() []int64 {
	ids := make([]int64, 0, len(adminIDs))
	for id := range adminIDs {
		ids = append(ids, id)
	}
	return ids
}
