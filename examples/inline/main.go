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

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		log.Fatal("BOT_TOKEN environment variable is required")
	}

	// Create a request builder function
	requestBuilder := func(req request.Interface) (*request.Native, error) {
		return request.NewNativeFromRequest(req), nil
	}

	// Create bot settings - enable inline_query updates
	settings := tb.Settings[*request.Native]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "inline_query", "chosen_inline_result"},
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
		botInfo := bot.Me()
		
		return c.Reply(fmt.Sprintf("🔍 Welcome to Inline Mode Demo!\n\n"+
			"Inline mode lets you use the bot in ANY chat by typing:\n\n"+
			"@%s <query>\n\n"+
			"Try these in any chat:\n"+
			"• @%s hello - Get greeting results\n"+
			"• @%s help - Get help articles\n"+
			"• @%s gif cats - Search for cat gifs\n"+
			"• @%s article - Get article results\n\n"+
			"Features demonstrated:\n"+
			"✅ Article results (with HTML content)\n"+
			"✅ Photo results (from URLs)\n"+
			"✅ GIF results\n"+
			"✅ Result descriptions and thumbnails\n"+
			"✅ Inline keyboards on results\n\n"+
			"⚠️ Note: Inline mode must be enabled in @BotFather:\n"+
			"1. Send /setinline to @BotFather\n"+
			"2. Select your bot\n"+
			"3. Set a placeholder text",
			botInfo.Username, botInfo.Username, botInfo.Username, botInfo.Username, botInfo.Username))
	})

	// Handle inline queries
	bot.Handle(tb.OnQuery, func(c *request.Native) error {
		query := c.InlineQuery()
		if query == nil {
			return nil
		}

		log.Printf("Inline query from %s: %s", query.Sender.FirstName, query.Query)

		// Prepare results based on query
		results := make(telegram.InlineResults, 0)

		// Convert query to lowercase for matching
		q := strings.ToLower(query.Query)

		// If query is empty, show default results
		if q == "" || strings.Contains(q, "hello") || strings.Contains(q, "hi") {
			// Article result with HTML content
			results = append(results, &telegram.InlineResultArticle{
				Type:  "article",
				ID:    "hello_1",
				Title: "👋 Say Hello",
				Description: "Send a friendly greeting",
				InputMessageContent: &telegram.InputTextMessageContent{
					MessageText: "👋 <b>Hello!</b>\n\nGreetings from the inline bot!",
					ParseMode: telegram.ParseModeHTML,
				},
				ThumbURL: "https://via.placeholder.com/100/FF6B6B/ffffff?text=Hi",
			})

			results = append(results, &telegram.InlineResultArticle{
				Type:  "article",
				ID:    "hello_2",
				Title: "🎉 Enthusiastic Hello",
				Description: "Send an exciting greeting",
				InputMessageContent: &telegram.InputTextMessageContent{
					MessageText: "🎉 <b>HELLO THERE!</b>\n\n<i>So excited to meet you!</i>",
					ParseMode: telegram.ParseModeHTML,
				},
			})
		}

		// Help articles
		if strings.Contains(q, "help") {
			results = append(results, &telegram.InlineResultArticle{
				Type:  "article",
				ID:    "help_1",
				Title: "📚 How to use inline mode",
				Description: "Learn about inline bots",
				InputMessageContent: &telegram.InputTextMessageContent{
					MessageText: "📚 <b>Using Inline Mode</b>\n\n" +
						"Type <code>@botname query</code> in any chat to use inline mode!\n\n" +
						"The bot will show results you can share instantly.",
					ParseMode: telegram.ParseModeHTML,
				},
				// Add inline keyboard to the result
				ReplyMarkup: &telegram.InlineKeyboardMarkup{
					InlineKeyboard: [][]telegram.InlineButton{
						{
							{Text: "📖 Telegram Docs", URL: "https://core.telegram.org/bots/inline"},
						},
					},
				},
			})

			results = append(results, &telegram.InlineResultArticle{
				Type:  "article",
				ID:    "help_2",
				Title: "🤖 About this bot",
				Description: "Information about the inline bot",
				InputMessageContent: &telegram.InputTextMessageContent{
					MessageText: "🤖 <b>Inline Bot Demo</b>\n\n" +
						"This bot demonstrates various inline query result types:\n" +
						"• Articles with formatted text\n" +
						"• Photos and GIFs\n" +
						"• Custom thumbnails\n" +
						"• Inline keyboards on results",
					ParseMode: telegram.ParseModeHTML,
				},
			})
		}

		// Photo results
		if strings.Contains(q, "photo") || strings.Contains(q, "pic") {
			results = append(results, &telegram.InlineResultPhoto{
				Type:     "photo",
				ID:       "photo_1",
				PhotoURL: "https://picsum.photos/800/600?random=1",
				ThumbURL: "https://picsum.photos/100/100?random=1",
				Title: "🖼 Random Photo 1",
				Description: "A beautiful random photo",
				Caption: "📷 Random photo from Lorem Picsum",
			})

			results = append(results, &telegram.InlineResultPhoto{
				Type:     "photo",
				ID:       "photo_2",
				PhotoURL: "https://picsum.photos/800/600?random=2",
				ThumbURL: "https://picsum.photos/100/100?random=2",
				Title: "🖼 Random Photo 2",
				Description: "Another beautiful random photo",
			})
		}

		// GIF results
		if strings.Contains(q, "gif") || strings.Contains(q, "anim") {
			results = append(results, &telegram.InlineResultGIF{
				Type:   "gif",
				ID:     "gif_1",
				GIFURL: "https://media.giphy.com/media/JIX9t2j0ZTN9S/giphy.gif",
				ThumbURL: "https://media.giphy.com/media/JIX9t2j0ZTN9S/200.gif",
				Title: "😺 Cute Cat GIF",
			})

			results = append(results, &telegram.InlineResultGIF{
				Type:   "gif",
				ID:     "gif_2",
				GIFURL: "https://media.giphy.com/media/mlvseq9yvZhba/giphy.gif",
				ThumbURL: "https://media.giphy.com/media/mlvseq9yvZhba/200.gif",
				Title: "🐕 Dog GIF",
			})
		}

		// Article with buttons
		if strings.Contains(q, "article") || strings.Contains(q, "button") {
			keyboard := &telegram.InlineKeyboardMarkup{
				InlineKeyboard: [][]telegram.InlineButton{
					{
						{Text: "👍 Like", CallbackData: "like"},
						{Text: "👎 Dislike", CallbackData: "dislike"},
					},
					{
						{Text: "🔗 Visit Website", URL: "https://telegram.org"},
					},
				},
			}

			results = append(results, &telegram.InlineResultArticle{
				Type:  "article",
				ID:    "article_button",
				Title: "📄 Article with Buttons",
				Description: "Send an article with inline keyboard",
				InputMessageContent: &telegram.InputTextMessageContent{
					MessageText: "📄 <b>Interactive Article</b>\n\n" +
						"This message includes inline buttons for interaction!",
					ParseMode: telegram.ParseModeHTML,
				},
				ReplyMarkup: keyboard,
			})
		}

		// If no results matched, provide a default
		if len(results) == 0 {
			results = append(results, &telegram.InlineResultArticle{
				Type:  "article",
				ID:    "default",
				Title: "🔍 No specific results",
				Description: fmt.Sprintf("Share your query: %s", query.Query),
				InputMessageContent: &telegram.InputTextMessageContent{
					MessageText: fmt.Sprintf("🔍 You searched for: <b>%s</b>\n\n" +
						"Try queries like: hello, help, photo, gif, article", query.Query),
					ParseMode: telegram.ParseModeHTML,
				},
			})
		}

		// Answer the inline query
		// Cache results for 5 minutes (300 seconds)
		response := &telegram.InlineQueryResponse{
			Results:   results,
			CacheTime: 300,
			IsPersonal: false, // Set true to give personalized results
		}

		return c.API.AnswerInlineQuery(query, response)
	})

	// Handle chosen inline results (when user sends a result)
	bot.Handle(tb.OnChosenInlineResult, func(c *request.Native) error {
		chosen := c.ChosenInlineResult()
		if chosen == nil {
			return nil
		}

		log.Printf("User %s chose inline result: %s (query: %s)",
			chosen.Sender.FirstName,
			chosen.ResultID,
			chosen.Query,
		)

		// You can track which results are most popular
		// Note: We can't send messages here as there's no chat context

		return nil
	})

	// Handle regular messages
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" || msg.Text[0] == '/' {
			return nil
		}

		botInfo := bot.Me()
		return c.Reply(fmt.Sprintf("Try using inline mode!\n\n"+
			"In any chat, type: @%s <query>", botInfo.Username))
	})

	botInfo := bot.Me()
	log.Println("🔍 Inline mode bot started! Press Ctrl+C to stop.")
	log.Printf("Bot username: @%s", botInfo.Username)
	log.Println("💡 Try it: Type @" + botInfo.Username + " hello in any chat!")
	log.Println("⚠️  Make sure inline mode is enabled in @BotFather (use /setinline)")
	if err := bot.Start(); err != nil {
		log.Fatal(err)
	}
}

