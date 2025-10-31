package main

import (
	"log"
	"os"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/files"
	"github.com/alagunto/tb/outgoing"
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
		return c.Reply("Welcome! This bot demonstrates file sending.\n\n" +
			"Commands:\n" +
			"/photo <url> - Send a photo from URL\n" +
			"/photo_local <path> - Send a local photo file\n" +
			"/document <url> - Send a document from URL\n" +
			"/document_local <path> - Send a local document\n" +
			"/video <url> - Send a video from URL\n" +
			"/video_local <path> - Send a local video\n" +
			"/audio <url> - Send an audio file\n" +
			"/voice <url> - Send a voice message\n" +
			"/sticker <file_id> - Send a sticker by file_id\n" +
			"/album - Send a media album")
	})

	// Send photo from URL
	bot.Handle("/photo", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a photo URL. Usage: /photo <url>")
		}
		url := args[0]
		photo := &telegram.Photo{
			Source:  files.UseURL(url),
			Caption: "This is a photo sent from URL!",
		}
		return c.Send(photo)
	})

	// Send local photo
	bot.Handle("/photo_local", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a file path. Usage: /photo_local <path>")
		}
		path := args[0]
		photo := &telegram.Photo{
			Source:  files.UseLocalFile(path),
			Caption: "This is a local photo file!",
		}
		opts := communications.NewSendOptions().WithParseMode(telegram.ParseModeHTML)
		return c.Send(photo, opts)
	})

	// Send document from URL
	bot.Handle("/document", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a document URL. Usage: /document <url>")
		}
		url := args[0]
		doc := &telegram.Document{
			Source:  files.UseURL(url),
			Caption: "This is a document sent from URL!",
		}
		return c.Send(doc)
	})

	// Send local document
	bot.Handle("/document_local", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a file path. Usage: /document_local <path>")
		}
		path := args[0]
		doc := &telegram.Document{
			Source:  files.UseLocalFile(path),
			Caption: "This is a local document file!",
		}
		return c.Send(doc)
	})

	// Send video from URL
	bot.Handle("/video", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a video URL. Usage: /video <url>")
		}
		url := args[0]
		video := &telegram.Video{
			Source:  files.UseURL(url),
			Caption: "This is a video sent from URL!",
		}
		return c.Send(video)
	})

	// Send local video
	bot.Handle("/video_local", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a file path. Usage: /video_local <path>")
		}
		path := args[0]
		video := &telegram.Video{
			Source:  files.UseLocalFile(path),
			Caption: "This is a local video file!",
		}
		return c.Send(video)
	})

	// Send audio from URL
	bot.Handle("/audio", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide an audio URL. Usage: /audio <url>")
		}
		url := args[0]
		audio := &telegram.Audio{
			Source:  files.UseURL(url),
			Caption: "This is an audio file!",
		}
		return c.Send(audio)
	})

	// Send voice message from URL
	bot.Handle("/voice", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a voice URL. Usage: /voice <url>")
		}
		url := args[0]
		voice := &telegram.Voice{
			Source: files.UseURL(url),
		}
		return c.Send(voice)
	})

	// Send sticker by file_id
	bot.Handle("/sticker", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Please provide a sticker file_id. Usage: /sticker <file_id>")
		}
		fileID := args[0]
		sticker := &telegram.Sticker{
			Source: files.UseTelegramFile(fileID),
		}
		return c.Send(sticker)
	})

	// Send media album
	bot.Handle("/album", func(c *request.Native) error {
		// Create an album with multiple photos
		// Note: You need to provide actual URLs or file paths for this to work
		album := telegram.InputAlbum{
			Media: []outgoing.Content{
				&telegram.Photo{
					Source:  files.UseURL("https://via.placeholder.com/150?text=Photo1"),
					Caption: "First photo in album",
				},
				&telegram.Photo{
					Source: files.UseURL("https://via.placeholder.com/150?text=Photo2"),
				},
				&telegram.Photo{
					Source: files.UseURL("https://via.placeholder.com/150?text=Photo3"),
				},
			},
		}
		return c.SendAlbum(album)
	})

	log.Println("Bot started! Press Ctrl+C to stop.")
	bot.Start()
}
