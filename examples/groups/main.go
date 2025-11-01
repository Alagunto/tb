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

// Admin IDs - replace with your admin user IDs
var adminIDs = map[int64]bool{
	123456789: true, // Replace with actual admin ID
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
			AllowedUpdates: []string{"message", "chat_member", "my_chat_member", "chat_join_request"},
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

	// Helper: Check if user is admin
	isAdmin := func(userID int64) bool {
		return adminIDs[userID]
	}

	// Start command
	bot.Handle("/start", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil {
			return nil
		}

		msg := c.Message()
		if msg == nil {
			return nil
		}

		// Different message for groups vs private
		if msg.Chat.Type == telegram.ChatPrivate {
			return c.Reply("👥 Welcome to Group Management Demo!\n\n" +
				"Add this bot to a group and make it an administrator.\n\n" +
				"Available commands:\n" +
				"/groupinfo - Show group information\n" +
				"/rules - Display group rules\n" +
				"/setrules <text> - Set group rules (admin)\n" +
				"/ban @username - Ban a user (admin)\n" +
				"/unban @username - Unban a user (admin)\n" +
				"/mute @username - Mute a user (admin)\n" +
				"/unmute @username - Unmute a user (admin)\n" +
				"/kick @username - Kick a user (admin)\n" +
				"/promote @username - Promote to admin (admin)\n" +
				"/demote @username - Remove admin rights (admin)\n" +
				"/members - Get member count\n" +
				"/admins - List group admins")
		}

		return c.Reply("👋 Bot is active in this group! Use /groupinfo to see group details.")
	})

	// Group information
	bot.Handle("/groupinfo", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		chat := msg.Chat

		info := fmt.Sprintf("📊 Group Information\n\n"+
			"Name: %s\n"+
			"Type: %s\n"+
			"ID: %d\n",
			chat.Title,
			chat.Type,
			chat.ID,
		)

		if chat.Username != "" {
			info += fmt.Sprintf("Username: @%s\n", chat.Username)
		}

		if chat.Description != "" {
			info += fmt.Sprintf("Description: %s\n", chat.Description)
		}

		// Get member count
		count, err := c.API.GetChatMembersCount(chat)
		if err == nil {
			info += fmt.Sprintf("Members: %d\n", count)
		}

		return c.Reply(info)
	})

	// Get member count
	bot.Handle("/members", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		count, err := c.API.GetChatMembersCount(msg.Chat)
		if err != nil {
			return c.Reply("Failed to get member count: " + err.Error())
		}

		return c.Reply(fmt.Sprintf("👥 This group has %d members", count))
	})

	// List admins
	bot.Handle("/admins", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		admins, err := c.API.GetChatAdministrators(msg.Chat)
		if err != nil {
			return c.Reply("Failed to get administrators: " + err.Error())
		}

		result := "👮 Group Administrators:\n\n"
		for i, admin := range admins {
			user := admin.User
			status := admin.Status
			
			name := user.FirstName
			if user.LastName != "" {
				name += " " + user.LastName
			}
			if user.Username != "" {
				name += " (@" + user.Username + ")"
			}

			result += fmt.Sprintf("%d. %s - %s\n", i+1, name, status)
		}

		return c.Reply(result)
	})

	// Ban user (admin only)
	bot.Handle("/ban", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ This command is only available to admins")
		}

		msg := c.Message()
		if msg == nil {
			return nil
		}

		// Get user from reply or mention
		var userToBan *telegram.User
		if msg.ReplyTo != nil && msg.ReplyTo.Sender != nil {
			userToBan = msg.ReplyTo.Sender
		} else {
			return c.Reply("Please reply to a message from the user you want to ban")
		}

		// Ban the user
		member := &telegram.ChatMember{
			User: userToBan,
		}

		if err := c.API.Ban(msg.Chat, member); err != nil {
			return c.Reply("Failed to ban user: " + err.Error())
		}

		return c.Reply(fmt.Sprintf("🚫 Banned %s from the group", userToBan.FirstName))
	})

	// Unban user (admin only)
	bot.Handle("/unban", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ This command is only available to admins")
		}

		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil || msg.ReplyTo.Sender == nil {
			return c.Reply("Please reply to a message from the user you want to unban")
		}

		userToUnban := msg.ReplyTo.Sender
		member := &telegram.ChatMember{
			User: userToUnban,
		}

		if err := c.API.Unban(msg.Chat, member); err != nil {
			return c.Reply("Failed to unban user: " + err.Error())
		}

		return c.Reply(fmt.Sprintf("✅ Unbanned %s", userToUnban.FirstName))
	})

	// Restrict user (mute) - admin only
	bot.Handle("/mute", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ This command is only available to admins")
		}

		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil || msg.ReplyTo.Sender == nil {
			return c.Reply("Please reply to a message from the user you want to mute")
		}

		userToMute := msg.ReplyTo.Sender

		// Restrict permissions (cannot send messages)
		permissions := telegram.Rights{
			CanSendMessages: false,
			CanSendMedia:    false,
			CanSendPolls:    false,
			CanSendOther:    false,
			CanAddPreviews:  false,
		}

		if err := c.API.Restrict(msg.Chat, userToMute, permissions); err != nil {
			return c.Reply("Failed to mute user: " + err.Error())
		}

		return c.Reply(fmt.Sprintf("🔇 Muted %s", userToMute.FirstName))
	})

	// Unmute user (admin only)
	bot.Handle("/unmute", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ This command is only available to admins")
		}

		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil || msg.ReplyTo.Sender == nil {
			return c.Reply("Please reply to a message from the user you want to unmute")
		}

		userToUnmute := msg.ReplyTo.Sender

		// Restore permissions
		permissions := telegram.Rights{
			CanSendMessages: true,
			CanSendMedia:    true,
			CanSendPolls:    true,
			CanSendOther:    true,
			CanAddPreviews:  true,
		}

		if err := c.API.Restrict(msg.Chat, userToUnmute, permissions); err != nil {
			return c.Reply("Failed to unmute user: " + err.Error())
		}

		return c.Reply(fmt.Sprintf("🔊 Unmuted %s", userToUnmute.FirstName))
	})

	// Handle new members joining
	bot.Handle(tb.OnUserJoined, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.UserJoined == nil {
			return nil
		}

		newUser := msg.UserJoined
		welcome := fmt.Sprintf("👋 Welcome to the group, %s!\n\n"+
			"Please read /rules and enjoy your stay!",
			newUser.FirstName,
		)

		return c.Reply(welcome)
	})

	// Handle members leaving
	bot.Handle(tb.OnUserLeft, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.UserLeft == nil {
			return nil
		}

		leftUser := msg.UserLeft
		log.Printf("User left: %s (%d)", leftUser.FirstName, leftUser.ID)

		return nil // Optionally send goodbye message
	})

	// Handle chat title changes
	bot.Handle(tb.OnNewGroupTitle, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		return c.Reply(fmt.Sprintf("📝 Group title changed to: %s", msg.NewGroupTitle))
	})

	// Handle chat photo changes
	bot.Handle(tb.OnNewGroupPhoto, func(c *request.Native) error {
		return c.Reply("📷 Group photo was updated!")
	})

	// Handle pinned messages
	bot.Handle(tb.OnPinned, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.PinnedMessage == nil {
			return nil
		}

		return c.Reply("📌 A message was pinned by admin")
	})

	// Simple rules system
	var groupRules = "1. Be respectful\n2. No spam\n3. Stay on topic"

	bot.Handle("/rules", func(c *request.Native) error {
		return c.Reply("📜 Group Rules:\n\n" + groupRules)
	})

	bot.Handle("/setrules", func(c *request.Native) error {
		sender := c.Sender()
		if sender == nil || !isAdmin(sender.ID) {
			return c.Reply("⛔️ Only admins can set rules")
		}

		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /setrules <rules text>")
		}

		groupRules = joinArgs(args)
		return c.Reply("✅ Group rules updated!")
	})

	// Handle regular messages
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" || msg.Text[0] == '/' {
			return nil
		}

		// Log messages for moderation
		log.Printf("[%s] %s: %s",
			msg.Chat.Title,
			msg.Sender.FirstName,
			msg.Text,
		)

		return nil
	})

	log.Println("👥 Group management bot started! Press Ctrl+C to stop.")
	log.Println("⚠️  Add bot to a group and make it an administrator")
	log.Println("📝 Don't forget to update adminIDs in the code with your user ID!")
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}

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

