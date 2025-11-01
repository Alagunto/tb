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

// Note: Web Apps require you to host a web page (HTML/JS)
// For testing, you can use these example URLs or host your own
const (
	exampleWebAppURL = "https://webappcontent.telegram.org/demo" // Official Telegram demo
	customWebAppURL  = "https://your-domain.com/webapp"          // Replace with your URL
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

	// Create bot settings - include web_app_data updates
	settings := tb.Settings[*request.Native]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "callback_query", "web_app_data"},
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
		return c.Reply("🌐 Welcome to Web Apps Demo!\n\n" +
			"Web Apps are mini applications that run inside Telegram.\n\n" +
			"Features:\n" +
			"• Full HTML/CSS/JavaScript support\n" +
			"• Access to Telegram user data\n" +
			"• Send data back to bot\n" +
			"• Responsive design\n" +
			"• Payment integration\n\n" +
			"Commands:\n" +
			"/webapp_button - Open web app via inline button\n" +
			"/webapp_keyboard - Open web app via reply keyboard\n" +
			"/webapp_menu - Set web app in menu button\n" +
			"/simple_form - Simple form web app\n" +
			"/game_demo - Web app game demo\n\n" +
			"💡 Web apps work on mobile and desktop!")
	})

	// Web app via inline button
	bot.Handle("/webapp_button", func(c *request.Native) error {
		keyboard := &telegram.InlineKeyboardMarkup{
			InlineKeyboard: [][]telegram.InlineButton{
				{
					{
						Text: "🚀 Launch Web App",
						WebApp: &telegram.WebAppInfo{
							URL: exampleWebAppURL,
						},
					},
				},
				{
					{
						Text: "📱 Another Web App",
						WebApp: &telegram.WebAppInfo{
							URL: "https://webappcontent.telegram.org/cafe",
						},
					},
				},
			},
		}

		text := "🌐 <b>Web App via Inline Button</b>\n\n" +
			"Click the button below to open a web app.\n\n" +
			"The web app opens in a full-screen overlay and can:\n" +
			"• Display rich content\n" +
			"• Use Telegram theme colors\n" +
			"• Send data back to the bot\n" +
			"• Request user information"

		return c.Reply(text, 
			tb.Send().
				WithParseMode(telegram.ParseModeHTML).
				WithReplyMarkup(keyboard))
	})

	// Web app via reply keyboard
	bot.Handle("/webapp_keyboard", func(c *request.Native) error {
		keyboard := &telegram.ReplyKeyboardMarkup{
			Keyboard: [][]telegram.ReplyButton{
				{
					{
						Text: "🎮 Open Game",
						WebApp: &telegram.WebAppInfo{
							URL: exampleWebAppURL,
						},
					},
				},
				{
					{
						Text: "📝 Fill Form",
						WebApp: &telegram.WebAppInfo{
							URL: customWebAppURL,
						},
					},
				},
				{
					{Text: "❌ Close Keyboard"},
				},
			},
			ResizeKeyboard: true,
		}

		text := "🎹 <b>Web App via Reply Keyboard</b>\n\n" +
			"Tap the buttons below to launch web apps.\n\n" +
			"Reply keyboards are persistent and great for:\n" +
			"• Quick access to web apps\n" +
			"• Frequently used features\n" +
			"• Custom interfaces"

		return c.Reply(text,
			tb.Send().
				WithParseMode(telegram.ParseModeHTML).
				WithReplyMarkup(keyboard))
	})

	// Set web app in menu button
	bot.Handle("/webapp_menu", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		// Set menu button to open web app
		menuButton := &telegram.MenuButton{
			Type: telegram.MenuButtonTypeWebApp,
			Text: "🎯 Open App",
			WebApp: &telegram.WebAppInfo{
				URL: exampleWebAppURL,
			},
		}

		if err := c.API.SetChatMenuButton(msg.Chat, menuButton); err != nil {
			return c.Reply("Failed to set menu button: " + err.Error())
		}

		return c.Reply("✅ <b>Menu button set!</b>\n\n" +
			"Look at the bottom-left corner of the chat.\n" +
			"The ☰ menu button now opens a web app!\n\n" +
			"To restore: /restore_menu",
			tb.Send().WithParseMode(telegram.ParseModeHTML))
	})

	// Restore default menu button
	bot.Handle("/restore_menu", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return nil
		}

		// Reset to default commands menu
		menuButton := &telegram.MenuButton{
			Type: telegram.MenuButtonTypeCommands,
		}

		if err := c.API.SetChatMenuButton(msg.Chat, menuButton); err != nil {
			return c.Reply("Failed to restore menu: " + err.Error())
		}

		return c.Reply("✅ Menu button restored to default (commands list)")
	})

	// Simple form web app
	bot.Handle("/simple_form", func(c *request.Native) error {
		// Note: This would require hosting your own web app
		// that uses Telegram.WebApp.sendData() to send form data
		
		keyboard := &telegram.InlineKeyboardMarkup{
			InlineKeyboard: [][]telegram.InlineButton{
				{
					{
						Text: "📝 Fill Form",
						WebApp: &telegram.WebAppInfo{
							URL: exampleWebAppURL,
						},
					},
				},
			},
		}

		text := "📝 <b>Form Web App Example</b>\n\n" +
			"A web app can collect user input and send it back.\n\n" +
			"Example form fields:\n" +
			"• Text inputs\n" +
			"• Dropdowns\n" +
			"• Date pickers\n" +
			"• File uploads\n\n" +
			"Data is sent via Telegram.WebApp.sendData()"

		return c.Reply(text,
			tb.Send().
				WithParseMode(telegram.ParseModeHTML).
				WithReplyMarkup(keyboard))
	})

	// Game demo web app
	bot.Handle("/game_demo", func(c *request.Native) error {
		keyboard := &telegram.InlineKeyboardMarkup{
			InlineKeyboard: [][]telegram.InlineButton{
				{
					{
						Text: "🎮 Play Game",
						WebApp: &telegram.WebAppInfo{
							URL: "https://webappcontent.telegram.org/demo",
						},
					},
				},
			},
		}

		text := "🎮 <b>Web App Game Demo</b>\n\n" +
			"Web apps can be used for games!\n\n" +
			"Features:\n" +
			"• HTML5 Canvas/WebGL\n" +
			"• Touch controls\n" +
			"• Leaderboards\n" +
			"• In-app purchases (via bot payments)\n" +
			"• Social sharing\n\n" +
			"Perfect for casual games, puzzles, and interactive experiences."

		return c.Reply(text,
			tb.Send().
				WithParseMode(telegram.ParseModeHTML).
				WithReplyMarkup(keyboard))
	})

	// Handle web app data (sent from web app via Telegram.WebApp.sendData())
	bot.Handle(tb.OnWebAppData, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.WebAppData == nil {
			return nil
		}

		webAppData := msg.WebAppData
		log.Printf("Received web app data: %s", webAppData.Data)

		// Process the data from web app
		response := fmt.Sprintf("✅ <b>Data received from Web App!</b>\n\n"+
			"Button text: %s\n"+
			"Data: <code>%s</code>\n\n"+
			"In a real app, you would:\n"+
			"• Parse the data (usually JSON)\n"+
			"• Validate it\n"+
			"• Process user input\n"+
			"• Save to database\n"+
			"• Send confirmation",
			webAppData.ButtonText,
			webAppData.Data,
		)

		return c.Reply(response, tb.Send().WithParseMode(telegram.ParseModeHTML))
	})

	// Handle keyboard buttons
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" {
			return nil
		}

		// Handle close keyboard button
		if msg.Text == "❌ Close Keyboard" {
			remove := &telegram.ReplyKeyboardRemove{RemoveKeyboard: true}
			return c.Reply("Keyboard removed", tb.Send().WithReplyMarkup(remove))
		}

		// Ignore commands
		if msg.Text[0] == '/' {
			return nil
		}

		return c.Reply("Send /start to see web app examples!")
	})

	log.Println("🌐 Web Apps bot started! Press Ctrl+C to stop.")
	log.Println("💡 Tips:")
	log.Println("   • Web apps require HTTPS URLs")
	log.Println("   • Test with Telegram's demo apps first")
	log.Println("   • Use Telegram.WebApp.sendData() to send data back")
	log.Println("   • Web apps work on iOS, Android, and Desktop")
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}

