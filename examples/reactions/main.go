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

	// Create bot settings - include reaction updates
	settings := tb.Settings[*request.Native]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "message_reaction", "message_reaction_count"},
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
		return c.Reply("❤️ Welcome to Reactions Demo!\n\n" +
			"This bot demonstrates message reactions:\n\n" +
			"Commands:\n" +
			"/react - Send a message and react to it\n" +
			"/react_multiple - React with multiple emojis\n" +
			"/react_custom - Use custom emoji reactions\n" +
			"/unreact - Remove reaction from a message\n\n" +
			"💡 Tips:\n" +
			"• Reactions work in groups, channels, and private chats\n" +
			"• Reply to any message with /react to add a reaction\n" +
			"• Some reactions are premium-only (custom emojis)")
	})

	// Send and react to a message
	bot.Handle("/react", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return c.Reply("No message context")
		}

		// Check if replying to a message
		if msg.ReplyTo != nil {
			// React to the replied message
			reactions := telegram.Reactions{
				{
					Type:  telegram.ReactionTypeEmoji,
					Emoji: "👍",
				},
			}

			if err := c.API.SetMessageReaction(msg.ReplyTo, reactions, false); err != nil {
				return c.Reply("Failed to add reaction: " + err.Error())
			}

			return c.Reply("✅ Reacted with 👍 to the message!")
		}

		// Send a new message and react to it
		sent, err := c.Send("Here's a test message. Watch it get a reaction! ⏳")
		if err != nil {
			return err
		}

		// Wait a moment for effect
		time.Sleep(1 * time.Second)

		// React to the sent message
		reactions := telegram.Reactions{
			{
				Type:  telegram.ReactionTypeEmoji,
				Emoji: "❤️",
			},
		}

		if err := c.API.SetMessageReaction(sent, reactions, false); err != nil {
			return c.Reply("Failed to add reaction: " + err.Error())
		}

		return nil
	})

	// React with multiple emojis
	bot.Handle("/react_multiple", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return c.Reply("No message context")
		}

		// Check if replying to a message
		targetMessage := msg.ReplyTo
		if targetMessage == nil {
			// Send a new message to react to
			sent, err := c.Send("This message will get multiple reactions!")
			if err != nil {
				return err
			}
			targetMessage = sent
			time.Sleep(500 * time.Millisecond)
		}

		// Add multiple reactions
		reactions := telegram.Reactions{
			{Type: telegram.ReactionTypeEmoji, Emoji: "👍"},
			{Type: telegram.ReactionTypeEmoji, Emoji: "❤️"},
			{Type: telegram.ReactionTypeEmoji, Emoji: "🔥"},
		}

		if err := c.API.SetMessageReaction(targetMessage, reactions, false); err != nil {
			return c.Reply("Failed to add reactions: " + err.Error())
		}

		if msg.ReplyTo != nil {
			return c.Reply("✅ Added multiple reactions!")
		}

		return nil
	})

	// Custom emoji reactions (premium feature)
	bot.Handle("/react_custom", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return c.Reply("No message context")
		}

		// Note: Custom emoji reactions require Telegram Premium
		// This is an example with a hypothetical custom emoji ID
		reactions := telegram.Reactions{
			{
				Type:          telegram.ReactionTypeCustomEmoji,
				CustomEmojiID: "5368324170671202286", // Example custom emoji ID
			},
		}

		targetMessage := msg.ReplyTo
		if targetMessage == nil {
			sent, err := c.Send("This will get a custom emoji reaction (if you have Premium)!")
			if err != nil {
				return err
			}
			targetMessage = sent
			time.Sleep(500 * time.Millisecond)
		}

		if err := c.API.SetMessageReaction(targetMessage, reactions, false); err != nil {
			return c.Reply("⚠️ Failed to add custom emoji reaction. " +
				"Note: Custom emoji reactions require Telegram Premium!\n\n" +
				"Error: " + err.Error())
		}

		if msg.ReplyTo != nil {
			return c.Reply("✅ Custom emoji reaction added!")
		}

		return nil
	})

	// Remove reactions
	bot.Handle("/unreact", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.ReplyTo == nil {
			return c.Reply("❌ Please reply to a message with /unreact to remove its reactions")
		}

		// Empty reactions array removes all reactions
		emptyReactions := telegram.Reactions{}

		if err := c.API.SetMessageReaction(msg.ReplyTo, emptyReactions, false); err != nil {
			return c.Reply("Failed to remove reactions: " + err.Error())
		}

		return c.Reply("✅ Removed all reactions from the message")
	})

	// Handle message reaction updates
	bot.Handle(tb.OnMessageReaction, func(c *request.Native) error {
		reactionUpdate := c.MessageReactionUpdated()
		if reactionUpdate == nil {
			return nil
		}

		// Log reaction changes
		log.Printf("Reaction update in chat %d, message %d",
			reactionUpdate.Chat.ID,
			reactionUpdate.MessageID,
		)

		oldReactions := reactionUpdate.OldReaction
		newReactions := reactionUpdate.NewReaction

		log.Printf("Old reactions: %d, New reactions: %d",
			len(oldReactions),
			len(newReactions),
		)

		// Show which reactions were added/removed
		for _, r := range newReactions {
			log.Printf("  - New reaction: %s (type: %s)", r.Emoji, r.Type)
		}

		return nil
	})

	// Handle message reaction count updates (anonymous reactions)
	bot.Handle(tb.OnMessageReactionCount, func(c *request.Native) error {
		countUpdate := c.MessageReactionCountUpdated()
		if countUpdate == nil {
			return nil
		}

		// Log reaction count changes
		log.Printf("Reaction count update in chat %d, message %d",
			countUpdate.Chat.ID,
			countUpdate.MessageID,
		)

		for _, reaction := range countUpdate.Reactions {
			log.Printf("  - %s: %d reactions", reaction.Type.Emoji, reaction.TotalCount)
		}

		return nil
	})

	// Handle text messages with examples
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" || msg.Text[0] == '/' {
			return nil
		}

		// Give helpful tips
		return c.Reply("💡 Tip: Reply to any message (including this one) with /react to add a reaction!\n\n" +
			"Available commands:\n" +
			"/react - Add reaction\n" +
			"/react_multiple - Add multiple reactions\n" +
			"/unreact - Remove reactions")
	})

	log.Println("❤️ Reactions bot started! Press Ctrl+C to stop.")
	log.Println("💡 Tip: Some reactions require Telegram Premium")
	log.Println("📱 Try the bot in groups to see anonymous reaction counts!")
	bot.Start()
}

