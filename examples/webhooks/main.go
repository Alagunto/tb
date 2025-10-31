package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/alagunto/tb"
	"github.com/alagunto/tb/request"
	"github.com/alagunto/tb/telegram"
)

// Webhook configuration
var (
	webhookURL  = os.Getenv("WEBHOOK_URL")       // e.g., https://your-domain.com:8443
	webhookPort = os.Getenv("WEBHOOK_PORT")      // e.g., 8443
	certFile    = os.Getenv("WEBHOOK_CERT_FILE") // Optional: path to cert.pem
	keyFile     = os.Getenv("WEBHOOK_KEY_FILE")  // Optional: path to key.pem
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Validate webhook configuration
	if webhookURL == "" {
		log.Fatal("WEBHOOK_URL environment variable is required (e.g., https://your-domain.com:8443)")
	}

	if webhookPort == "" {
		webhookPort = "8443" // Default webhook port
		log.Printf("Using default webhook port: %s", webhookPort)
	}

	// Create a request builder function
	requestBuilder := func(req request.Interface) (*request.Native, error) {
		return request.NewNativeFromRequest(req), nil
	}

	// Create webhook poller instead of long poller
	webhook := &tb.Webhook{
		Listen:   ":" + webhookPort,
		Endpoint: &tb.WebhookEndpoint{PublicURL: webhookURL},
		// Optional: If you have SSL certificate files
		// TLS: &tb.WebhookTLS{
		// 	Key:  keyFile,
		// 	Cert: certFile,
		// },
	}

	// Create bot settings with webhook
	settings := tb.Settings[*request.Native]{
		Token:  token,
		Poller: webhook,
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
		return c.Reply("🔗 Welcome to Webhooks Demo!\n\n" +
			"This bot uses webhooks instead of long polling.\n\n" +
			"<b>Webhooks vs Long Polling:</b>\n\n" +
			"<b>Long Polling (default):</b>\n" +
			"• Bot asks Telegram for updates every few seconds\n" +
			"• Easy to set up (no server needed)\n" +
			"• Works behind firewalls/NAT\n" +
			"• Higher latency\n\n" +
			"<b>Webhooks (this bot):</b>\n" +
			"• Telegram sends updates to your server instantly\n" +
			"• Requires public HTTPS URL\n" +
			"• Lower latency\n" +
			"• Better for high-traffic bots\n" +
			"• Scales better\n\n" +
			"Commands:\n" +
			"/webhook_info - Show webhook information\n" +
			"/test_webhook - Test webhook connection\n" +
			"/switch_to_polling - Switch to long polling\n\n" +
			"<b>Requirements:</b>\n" +
			"✅ Public domain with HTTPS\n" +
			"✅ Valid SSL certificate\n" +
			"✅ Open port (8443, 443, 80, or 88)",
			tb.Send().WithParseMode(telegram.ParseModeHTML))
	})

	// Show webhook information
	bot.Handle("/webhook_info", func(c *request.Native) error {
		// Get webhook info from Telegram
		info, err := c.API.GetWebhookInfo()
		if err != nil {
			return c.Reply("Failed to get webhook info: " + err.Error())
		}

		response := fmt.Sprintf("🔗 <b>Webhook Information</b>\n\n"+
			"<b>URL:</b> <code>%s</code>\n"+
			"<b>Has Custom Certificate:</b> %v\n"+
			"<b>Pending Updates:</b> %d\n"+
			"<b>Max Connections:</b> %d\n",
			info.URL,
			info.HasCustomCertificate,
			info.PendingUpdateCount,
			info.MaxConnections,
		)

		if info.LastErrorDate != 0 {
			errorTime := time.Unix(int64(info.LastErrorDate), 0)
			response += fmt.Sprintf("\n<b>⚠️ Last Error:</b>\n"+
				"Date: %s\n"+
				"Message: %s\n",
				errorTime.Format("2006-01-02 15:04:05"),
				info.LastErrorMessage,
			)
		}

		if len(info.AllowedUpdates) > 0 {
			response += fmt.Sprintf("\n<b>Allowed Updates:</b> %v\n", info.AllowedUpdates)
		}

		return c.Reply(response, tb.Send().WithParseMode(telegram.ParseModeHTML))
	})

	// Test webhook
	bot.Handle("/test_webhook", func(c *request.Native) error {
		return c.Reply("✅ Webhook is working!\n\n" +
			"If you see this message, your webhook is properly configured.\n\n" +
			"Server: " + webhookURL + "\n" +
			"Port: " + webhookPort)
	})

	// Switch to long polling (removes webhook)
	bot.Handle("/switch_to_polling", func(c *request.Native) error {
		// To switch to polling, you need to delete the webhook
		if err := c.API.DeleteWebhook(true); err != nil {
			return c.Reply("Failed to delete webhook: " + err.Error())
		}

		return c.Reply("✅ Webhook deleted!\n\n" +
			"To use long polling:\n" +
			"1. Stop this bot\n" +
			"2. Use LongPoller instead of Webhook in settings\n" +
			"3. Restart the bot\n\n" +
			"<b>Note:</b> You'll need to restart the bot for changes to take effect.",
			tb.Send().WithParseMode(telegram.ParseModeHTML))
	})

	// Echo command
	bot.Handle("/echo", func(c *request.Native) error {
		args := c.Args()
		if len(args) == 0 {
			return c.Reply("Usage: /echo <text>")
		}

		text := ""
		for i, arg := range args {
			if i > 0 {
				text += " "
			}
			text += arg
		}

		return c.Reply("📢 " + text)
	})

	// Handle all text messages
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" || msg.Text[0] == '/' {
			return nil
		}

		// Log the message
		log.Printf("Received via webhook: %s", msg.Text)

		return c.Reply("Received via webhook: " + msg.Text + "\n\nSend /start for help")
	})

	// Additional HTTP routes (optional)
	// You can add custom routes to your webhook server
	http.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == "/" {
			w.Write([]byte("Telegram Bot Webhook Server is running!"))
		} else {
			http.NotFound(w, r)
		}
	})

	// Start the webhook server
	log.Println("🔗 Webhook bot starting...")
	log.Printf("Webhook URL: %s", webhookURL)
	log.Printf("Listening on port: %s", webhookPort)
	log.Println("")
	log.Println("💡 Tips:")
	log.Println("   • Make sure your domain points to this server")
	log.Println("   • Ensure port " + webhookPort + " is open in firewall")
	log.Println("   • Use a valid SSL certificate (Let's Encrypt recommended)")
	log.Println("   • Test webhook with /webhook_info command")
	log.Println("")

	if certFile != "" && keyFile != "" {
		log.Println("🔒 Using SSL certificate files")
		log.Printf("   Cert: %s", certFile)
		log.Printf("   Key: %s", keyFile)
	} else {
		log.Println("⚠️  No SSL certificate files provided")
		log.Println("   Make sure you're behind a reverse proxy (nginx/caddy) with SSL")
		log.Println("   Or provide WEBHOOK_CERT_FILE and WEBHOOK_KEY_FILE environment variables")
	}

	log.Println("")
	log.Println("Starting bot... Press Ctrl+C to stop.")

	// For self-signed certificates or testing, you might need to skip verification
	// WARNING: Only use this for testing, not in production!
	if os.Getenv("WEBHOOK_INSECURE_SKIP_VERIFY") == "true" {
		log.Println("⚠️  WARNING: SSL verification disabled (insecure)")
		http.DefaultTransport.(*http.Transport).TLSClientConfig = &tls.Config{
			InsecureSkipVerify: true,
		}
	}

	// Start the bot (this will start the webhook HTTP server)
	bot.Start()
}

