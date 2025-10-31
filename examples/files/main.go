package main

import (
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
		return c.Reply("Welcome! This bot demonstrates sending media files.\n\n" +
			"Commands:\n" +
			"/photo <url> - Send a photo from URL\n" +
			"/video <url> - Send a video from URL\n" +
			"/document <url> - Send a document from URL\n" +
			"/audio <url> - Send an audio file from URL\n" +
			"/album - Send a media album with photos\n\n" +
			"Note: File uploads from disk are not yet fully implemented in v5.\n" +
			"Use URLs or file_ids for now.")
	})

	// Send photo from URL
	bot.Handle("/photo", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a photo URL. Usage: /photo <url>\n\nExample:\n/photo https://picsum.photos/800/600")
		}
		url := args[0]

		photo := &telegram.InputMediaPhoto{
			Type:    "photo",
			Media:   url,
			Caption: "📷 Photo from URL!",
		}

		return c.Send(photo)
	})

	// Send video from URL
	bot.Handle("/video", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a video URL. Usage: /video <url>\n\nExample:\n/video https://sample-videos.com/video123/mp4/240/big_buck_bunny_240p_1mb.mp4")
		}
		url := args[0]

		video := &telegram.InputMediaVideo{
			Type:    "video",
			Media:   url,
			Caption: "🎥 Video from URL!",
		}

		return c.Send(video)
	})

	// Send document from URL
	bot.Handle("/document", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a document URL. Usage: /document <url>\n\nExample:\n/document https://www.w3.org/WAI/ER/tests/xhtml/testfiles/resources/pdf/dummy.pdf")
		}
		url := args[0]

		doc := &telegram.InputMediaDocument{
			Type:    "document",
			Media:   url,
			Caption: "📄 Document from URL!",
		}

		return c.Send(doc)
	})

	// Send audio from URL
	bot.Handle("/audio", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide an audio URL. Usage: /audio <url>\n\nExample:\n/audio https://www.soundhelix.com/examples/mp3/SoundHelix-Song-1.mp3")
		}
		url := args[0]

		audio := &telegram.InputMediaAudio{
			Type:    "audio",
			Media:   url,
			Caption: "🎵 Audio from URL!",
		}

		return c.Send(audio)
	})

	// Send media album
	bot.Handle("/album", func(c *request.Native) error {
		// Create an album with multiple photos from URLs
		album := telegram.InputAlbum{
			Media: []telegram.InputMedia{
				&telegram.InputMediaPhoto{
					Type:    "photo",
					Media:   "https://picsum.photos/800/600?random=1",
					Caption: "First photo in album 📷",
				},
				&telegram.InputMediaPhoto{
					Type:  "photo",
					Media: "https://picsum.photos/800/600?random=2",
				},
				&telegram.InputMediaPhoto{
					Type:  "photo",
					Media: "https://picsum.photos/800/600?random=3",
				},
			},
		}

		if err := c.SendAlbum(album); err != nil {
			return c.Reply("Failed to send album: " + err.Error())
		}

		return c.Reply("✅ Album sent successfully!")
	})

	// Handle text messages
	bot.Handle(tb.OnText, func(c *request.Native) error {
		return c.Reply("Send /start to see available commands!")
	})

	log.Println("Bot started! Press Ctrl+C to stop.")
	log.Println("💡 Tip: Use public URLs for media files")
	bot.Start()
}
