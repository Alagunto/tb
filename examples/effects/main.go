package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/communications"
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
			"/effect_rainbow - Send message with rainbow effect\n" +
			"/effect_snow - Send message with snow effect\n" +
			"/effect_hearts - Send message with hearts effect\n" +
			"/effect_celebrate - Send message with celebration effect\n" +
			"/effect_fireworks - Send message with fireworks effect\n" +
			"/effect_shake - Send message with shake effect\n" +
			"/effect_explosion - Send message with explosion effect\n" +
			"/effect_boom - Send message with boom effect\n" +
			"/effect_all - Send messages with all available effects")
	})

	// Helper function to send message with effect
	sendWithEffect := func(c *request.Native, effectID telegram.EffectID, description string) error {
		text := fmt.Sprintf("Message with effect: %s", description)
		opts := communications.NewSendOptions().WithEffectID(effectID)
		return c.Reply(text, opts)
	}

	// Rainbow effect
	bot.Handle("/effect_rainbow", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectID("rainbow"), "Rainbow")
	})

	// Snow effect
	bot.Handle("/effect_snow", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectID("snow"), "Snow")
	})

	// Hearts effect
	bot.Handle("/effect_hearts", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectID("hearts"), "Hearts")
	})

	// Celebration effect
	bot.Handle("/effect_celebrate", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectID("celebrate"), "Celebration")
	})

	// Fireworks effect
	bot.Handle("/effect_fireworks", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectID("fireworks"), "Fireworks")
	})

	// Shake effect
	bot.Handle("/effect_shake", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectID("shake"), "Shake")
	})

	// Explosion effect
	bot.Handle("/effect_explosion", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectID("explosion"), "Explosion")
	})

	// Boom effect
	bot.Handle("/effect_boom", func(c *request.Native) error {
		return sendWithEffect(c, telegram.EffectID("boom"), "Boom")
	})

	// Send all effects
	bot.Handle("/effect_all", func(c *request.Native) error {
		effects := []struct {
			id          telegram.EffectID
			description string
		}{
			{telegram.EffectID("rainbow"), "Rainbow"},
			{telegram.EffectID("snow"), "Snow"},
			{telegram.EffectID("hearts"), "Hearts"},
			{telegram.EffectID("celebrate"), "Celebration"},
			{telegram.EffectID("fireworks"), "Fireworks"},
			{telegram.EffectID("shake"), "Shake"},
			{telegram.EffectID("explosion"), "Explosion"},
			{telegram.EffectID("boom"), "Boom"},
		}

		for _, effect := range effects {
			text := fmt.Sprintf("Effect: %s", effect.description)
			opts := communications.NewSendOptions().WithEffectID(effect.id)
			if err := c.Reply(text, opts); err != nil {
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
	bot.Start()
}
