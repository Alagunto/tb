package tb

import "github.com/alagunto/tb/telegram"

// Type aliases for Poll types
type (
	Poll       = telegram.Poll
	PollType   = telegram.PollType
	PollOption = telegram.PollOption
	PollAnswer = telegram.PollAnswer
)

// Re-export constants
const (
	PollAny     = telegram.PollAny
	PollQuiz    = telegram.PollQuiz
	PollRegular = telegram.PollRegular
)
