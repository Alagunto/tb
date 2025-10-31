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

	// Create bot settings - include poll and poll_answer updates
	settings := tb.Settings[*request.Native]{
		Token: token,
		Poller: &tb.LongPoller{
			Timeout:        10 * time.Second,
			AllowedUpdates: []string{"message", "poll", "poll_answer"},
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
		return c.Reply("📊 Welcome to Polls Demo!\n\n" +
			"Create and manage Telegram polls:\n\n" +
			"Basic Polls:\n" +
			"/regular_poll - Standard poll with multiple choices\n" +
			"/quiz_poll - Quiz with correct answer\n" +
			"/anonymous_poll - Anonymous voting\n\n" +
			"Advanced Features:\n" +
			"/multi_choice - Allow multiple answers\n" +
			"/open_period - Poll with time limit\n" +
			"/close_poll - Close the last poll (reply to poll message)\n\n" +
			"💡 Tip: Reply to a poll message with /close_poll to close it")
	})

	// Regular poll - standard multiple choice
	bot.Handle("/regular_poll", func(c *request.Native) error {
		poll := &telegram.SendPollParams{
			Question: "What's your favorite programming language?",
			Options: []telegram.PollOption{
				{Text: "Go"},
				{Text: "Python"},
				{Text: "JavaScript"},
				{Text: "Rust"},
				{Text: "Other"},
			},
			IsAnonymous: false, // Show who voted
		}

		if err := c.SendPoll(poll); err != nil {
			return c.Reply("Failed to create poll: " + err.Error())
		}

		return nil
	})

	// Quiz poll - has a correct answer
	bot.Handle("/quiz_poll", func(c *request.Native) error {
		correctAnswerIndex := 2 // Zero-based index (Paris)

		poll := &telegram.SendPollParams{
			Question: "What is the capital of France?",
			Options: []telegram.PollOption{
				{Text: "London"},
				{Text: "Berlin"},
				{Text: "Paris"}, // Correct answer (index 2)
				{Text: "Madrid"},
			},
			Type:                telegram.PollTypeQuiz,
			CorrectOptionID:     &correctAnswerIndex,
			Explanation:         "Paris is the capital and largest city of France.",
			ExplanationParseMode: telegram.ParseModeHTML,
			IsAnonymous:         true,
		}

		if err := c.SendPoll(poll); err != nil {
			return c.Reply("Failed to create quiz: " + err.Error())
		}

		return c.Reply("🎓 Quiz created! Answer and see if you're correct.")
	})

	// Anonymous poll
	bot.Handle("/anonymous_poll", func(c *request.Native) error {
		poll := &telegram.SendPollParams{
			Question: "How would you rate this bot?",
			Options: []telegram.PollOption{
				{Text: "⭐️ Excellent"},
				{Text: "⭐️ Good"},
				{Text: "⭐️ Average"},
				{Text: "⭐️ Poor"},
			},
			IsAnonymous: true, // Default, but explicit for clarity
		}

		if err := c.SendPoll(poll); err != nil {
			return c.Reply("Failed to create poll: " + err.Error())
		}

		return c.Reply("🔒 Anonymous poll created. Your vote is private!")
	})

	// Multi-choice poll - allows multiple answers
	bot.Handle("/multi_choice", func(c *request.Native) error {
		poll := &telegram.SendPollParams{
			Question: "Which programming languages do you know? (Select all that apply)",
			Options: []telegram.PollOption{
				{Text: "Go"},
				{Text: "Python"},
				{Text: "JavaScript"},
				{Text: "Java"},
				{Text: "C++"},
				{Text: "Rust"},
			},
			AllowsMultipleAnswers: true,
			IsAnonymous:           false,
		}

		if err := c.SendPoll(poll); err != nil {
			return c.Reply("Failed to create poll: " + err.Error())
		}

		return c.Reply("✅ Multi-choice poll created! Select all that apply.")
	})

	// Poll with open period (time limit)
	bot.Handle("/open_period", func(c *request.Native) error {
		openPeriod := 60 // 60 seconds

		poll := &telegram.SendPollParams{
			Question: "Quick poll! What's your favorite emoji?",
			Options: []telegram.PollOption{
				{Text: "😀 Happy"},
				{Text: "🚀 Rocket"},
				{Text: "❤️ Heart"},
				{Text: "🔥 Fire"},
			},
			OpenPeriod:  &openPeriod,
			IsAnonymous: true,
		}

		if err := c.SendPoll(poll); err != nil {
			return c.Reply("Failed to create poll: " + err.Error())
		}

		return c.Reply("⏱ Poll created with 60 second time limit!")
	})

	// Close poll - must reply to a poll message
	bot.Handle("/close_poll", func(c *request.Native) error {
		msg := c.Message()
		if msg == nil {
			return c.Reply("Please use this as a command")
		}

		// Check if replying to a poll
		if msg.ReplyTo == nil || msg.ReplyTo.Poll == nil {
			return c.Reply("❌ Please reply to a poll message with /close_poll to close it")
		}

		// Stop the poll
		stoppedPoll, err := c.API.StopPoll(msg.ReplyTo)
		if err != nil {
			return c.Reply("Failed to close poll: " + err.Error())
		}

		// Show results
		result := fmt.Sprintf("📊 Poll closed!\n\nQuestion: %s\n\nResults:\n", stoppedPoll.Question)
		for i, option := range stoppedPoll.Options {
			result += fmt.Sprintf("%d. %s - %d votes\n", i+1, option.Text, option.VoterCount)
		}
		result += fmt.Sprintf("\nTotal votes: %d", stoppedPoll.TotalVoterCount)

		return c.Reply(result)
	})

	// Handle poll updates (when poll is closed)
	bot.Handle(tb.OnPoll, func(c *request.Native) error {
		poll := c.Poll()
		if poll == nil {
			return nil
		}

		// Log poll updates
		log.Printf("Poll updated: %s (Total votes: %d, Closed: %v)",
			poll.Question,
			poll.TotalVoterCount,
			poll.IsClosed,
		)

		return nil
	})

	// Handle poll answers (non-anonymous polls only)
	bot.Handle(tb.OnPollAnswer, func(c *request.Native) error {
		answer := c.PollAnswer()
		if answer == nil {
			return nil
		}

		// Log who voted
		user := answer.User
		optionIDs := answer.OptionIDs

		log.Printf("User %s (%d) voted on poll %s with options: %v",
			user.FirstName,
			user.ID,
			answer.PollID,
			optionIDs,
		)

		// Note: We can't send messages from poll_answer updates
		// as there's no chat context. This is just for logging.

		return nil
	})

	// Handle text messages
	bot.Handle(tb.OnText, func(c *request.Native) error {
		msg := c.Message()
		if msg == nil || msg.Text == "" || msg.Text[0] == '/' {
			return nil
		}

		return c.Reply("Send /start to see poll creation commands!")
	})

	log.Println("📊 Polls bot started! Press Ctrl+C to stop.")
	log.Println("💡 Tip: Use /quiz_poll for educational quizzes with correct answers")
	bot.Start()
}

