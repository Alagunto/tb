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

// IMPORTANT: This example requires a supergroup with Topics (Forum) enabled
// Set your forum chat ID here
var forumChatID = os.Getenv("FORUM_CHAT_ID")

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	if forumChatID == "" {
		log.Fatal("FORUM_CHAT_ID environment variable is required\n" +
			"Create a supergroup, enable Topics in group settings, and add the bot as admin")
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

	// Get forum chat
	forumChat, err := bot.ChatByID(forumChatID)
	if err != nil {
		log.Fatalf("Failed to get forum chat: %v", err)
	}

	log.Printf("Connected to forum: %s", forumChat.Title)

	// Start command
	bot.Handle("/start", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		// Check if in forum or private chat
		if msg.Chat.Type == telegram.ChatPrivate {
			return c.Reply("🗂 Welcome to Forum/Topics Demo!\n\n" +
				"This bot manages forum topics (threads) in supergroups.\n\n" +
				"Commands:\n" +
				"/create_topic <name> - Create a new topic\n" +
				"/edit_topic <name> - Edit topic name (reply to topic message)\n" +
				"/close_topic - Close a topic (reply to topic message)\n" +
				"/reopen_topic - Reopen a topic (reply to topic message)\n" +
				"/delete_topic - Delete a topic (reply to topic message)\n" +
				"/pin_topic - Pin a message in topic (reply to it)\n" +
				"/unpin_topic - Unpin message (reply to it)\n" +
				"/post_to_topic <topic_id> <text> - Post to specific topic\n\n" +
				"⚠️ Add bot to forum and make it admin with topic management rights!")
		}

		return c.Reply("🗂 Forum bot is active! Use commands to manage topics.\n\n" +
			"Try /create_topic <name> to create a new discussion thread!")
	})

	// Create new forum topic
	bot.Handle("/create_topic", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /create_topic <topic name>\n\nExample:\n/create_topic General Discussion")
		}

		topicName := joinArgs(args)

		// Create forum topic with default icon
		topic, err := c.API.CreateForumTopic(forumChat, topicName, "", "")
		if err != nil {
			return c.Reply("Failed to create topic: " + err.Error() + "\n\nMake sure:\n" +
				"1. This is a supergroup with Topics enabled\n" +
				"2. Bot is an administrator\n" +
				"3. Bot has 'Manage Topics' permission")
		}

		response := fmt.Sprintf("✅ Created topic: %s\n\nTopic ID: %d\n\n"+
			"You can now post messages to this topic!",
			topicName,
			topic.MessageThreadID,
		)

		return c.Reply(response)
	})

	// Edit forum topic (must reply to a message in that topic)
	bot.Handle("/edit_topic", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.MessageThreadID == 0 {
			return c.Reply("❌ Please use this command in a forum topic thread")
		}

		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /edit_topic <new name>\n\nExample:\n/edit_topic Updated Topic Name")
		}

		newName := joinArgs(args)

		// Edit forum topic name
		if err := c.API.EditForumTopic(msg.Chat, msg.MessageThreadID, newName, ""); err != nil {
			return c.Reply("Failed to edit topic: " + err.Error())
		}

		return c.Reply(fmt.Sprintf("✅ Topic renamed to: %s", newName))
	})

	// Close forum topic
	bot.Handle("/close_topic", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.MessageThreadID == 0 {
			return c.Reply("❌ Please use this command in a forum topic thread")
		}

		if err := c.API.CloseForumTopic(msg.Chat, msg.MessageThreadID); err != nil {
			return c.Reply("Failed to close topic: " + err.Error())
		}

		return c.Reply("🔒 Topic closed. Only admins can post now.")
	})

	// Reopen forum topic
	bot.Handle("/reopen_topic", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.MessageThreadID == 0 {
			return c.Reply("❌ Please use this command in a forum topic thread")
		}

		if err := c.API.ReopenForumTopic(msg.Chat, msg.MessageThreadID); err != nil {
			return c.Reply("Failed to reopen topic: " + err.Error())
		}

		return c.Reply("🔓 Topic reopened. Everyone can post again!")
	})

	// Delete forum topic
	bot.Handle("/delete_topic", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.MessageThreadID == 0 {
			return c.Reply("❌ Please use this command in a forum topic thread")
		}

		if err := c.API.DeleteForumTopic(msg.Chat, msg.MessageThreadID); err != nil {
			return c.Reply("Failed to delete topic: " + err.Error())
		}

		// Note: This message might not be visible as topic is being deleted
		return c.Reply("🗑 Topic deleted!")
	})

	// Pin message in topic
	bot.Handle("/pin_topic", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil {
			return c.Reply("❌ Please reply to a message you want to pin")
		}

		if err := c.API.Pin(msg.ReplyTo); err != nil {
			return c.Reply("Failed to pin message: " + err.Error())
		}

		return c.Reply("📌 Message pinned in topic!")
	})

	// Unpin message in topic
	bot.Handle("/unpin_topic", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil {
			return c.Reply("❌ Please reply to a pinned message to unpin it")
		}

		if err := c.API.Unpin(msg.ReplyTo); err != nil {
			return c.Reply("Failed to unpin message: " + err.Error())
		}

		return c.Reply("📌 Message unpinned from topic!")
	})

	// Post to specific topic by ID
	bot.Handle("/post_to_topic", func(c *request.Native) error {
		args := c.Args()
		if len(args) < 2 {
			return c.Reply("Usage: /post_to_topic <topic_id> <message>\n\n" +
				"Example:\n/post_to_topic 123 Hello, forum members!")
		}

		// Parse topic ID
		var topicID int
		if _, err := fmt.Sscanf(args[0], "%d", &topicID); err != nil {
			return c.Reply("❌ Invalid topic ID. Must be a number.")
		}

		// Get message text (everything after topic ID)
		messageText := ""
		for i := 1; i < len(args); i++ {
			if i > 1 {
				messageText += " "
			}
			messageText += args[i]
		}

		// Send message to specific topic
		// Note: This requires using send options with message_thread_id
		opts := tb.Send().WithMessageThreadID(topicID)
		_, err := bot.Send(forumChat, messageText, opts)
		if err != nil {
			return c.Reply("Failed to post to topic: " + err.Error())
		}

		return c.Reply(fmt.Sprintf("✅ Posted message to topic %d", topicID))
	})

	// Handle messages in forum topics
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" || msg.Text[0] == '/' {
			return nil
		}

		// Check if message is in a topic
		if msg.MessageThreadID != 0 {
			log.Printf("Message in topic %d: %s", msg.MessageThreadID, msg.Text)
		}

		return nil
	})

	log.Println("🗂 Forum topics bot started! Press Ctrl+C to stop.")
	log.Printf("Managing forum: %s", forumChat.Title)
	log.Println("⚠️  Requirements:")
	log.Println("   1. Supergroup with Topics enabled")
	log.Println("   2. Bot must be administrator")
	log.Println("   3. Bot needs 'Manage Topics' permission")
	bot.Start()
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

