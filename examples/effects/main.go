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
			AllowedUpdates: []string{"message"},
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
		return c.Reply("Welcome! This bot demonstrates message effects.\n\n" +
			"⚠️ Note: Message effects only work in private chats!\n\n" +
			"Commands:\n" +
			"/effect_fire - Send message with fire effect 🔥\n" +
			"/effect_like - Send message with like effect 👍\n" +
			"/effect_dislike - Send message with dislike effect 👎\n" +
			"/effect_heart - Send message with heart effect ❤️\n" +
			"/effect_celebrate - Send message with celebration effect 🎉\n" +
			"/effect_poop - Send message with poop effect 💩\n" +
			"/effect_all - Send messages with all available effects")
	})

	// Helper function to send message with effect
	sendWithEffect := func(c *request.Native, effectID telegram.EffectID, description string) error {
		msg := c.Message()
		if msg == nil {
			return c.Reply("No message context available")
		}

		// Check if this is a private chat
		if msg.Chat.Type != telegram.ChatPrivate {
			return c.Reply("⚠️ Message effects only work in private chats! Please use this command in a direct message with the bot.")
		}

		text := fmt.Sprintf("Message with effect: %s", description)
		// New chainable API - WithEffect is an alias for WithEffectID
		return c.Reply(text, tb.Send().WithEffect(effectID))
	}

	// Fire effect
	bot.Handle("/effect_fire", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectFire, "Fire 🔥")
	})

	// Like effect
	bot.Handle("/effect_like", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectLike, "Like 👍")
	})

	// Dislike effect
	bot.Handle("/effect_dislike", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectDislike, "Dislike 👎")
	})

	// Heart effect
	bot.Handle("/effect_heart", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectHeart, "Heart ❤️")
	})

	// Celebration effect
	bot.Handle("/effect_celebrate", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectCelebration, "Celebration 🎉")
	})

	// Poop effect
	bot.Handle("/effect_poop", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectPoop, "Poop 💩")
	})

	// Send all effects
	bot.Handle("/effect_all", func(c *request.Native) error {
		// These are the six core effects (numeric IDs 0-5)
		// Use getAvailableMessageEffects API for dynamic discovery of all effects
		effects := []struct {
			id          telegram.EffectID
			description string
		}{
			{telegram.EffectFire, "Fire 🔥"},
			{telegram.EffectLike, "Like 👍"},
			{telegram.EffectDislike, "Dislike 👎"},
			{telegram.EffectHeart, "Heart ❤️"},
			{telegram.EffectCelebration, "Celebration 🎉"},
			{telegram.EffectPoop, "Poop 💩"},
		}

		for _, effect := range effects {
			text := fmt.Sprintf("Effect: %s", effect.description)
			// New chainable API
			if err := c.Reply(text, tb.Send().WithEffect(effect.id)); err != nil {
				log.Printf("Error sending effect %s: %v", effect.description, err)
			}
			// Small delay between messages
			time.Sleep(500 * time.Millisecond)
		}

		return c.Reply("All effects sent! Check your private chat.")
	})

	// Handle messages sent with effects
	bot.Handle(tb.OnText, func(c *request.Native) error {
		return c.Reply("Send /start to see available effects!")
	})

	log.Println("Bot started! Press Ctrl+C to stop.")
	log.Println("⚠️  Remember: Message effects only work in private chats!")
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}
