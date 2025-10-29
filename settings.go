package tb

import (
	"net/http"
	"time"

	"github.com/alagunto/tb/censorship"
)

// Settings represents a utility struct for passing certain
// properties of a bot around and is required to make bots.
type Settings[Ctx ContextInterface, HandlerFunc func(Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc] struct {
	URL   string
	Token string

	// Updates channel capacity, defaulted to 100.
	Updates int

	// Poller is the provider of Updates.
	Poller Poller

	// Synchronous prevents handlers from running in parallel.
	// It makes ProcessUpdate return after the handler is finished.
	Synchronous bool

	// Verbose forces bot to log all upcoming requests.
	// Use for debugging purposes only.
	Verbose bool

	// ParseMode used to set default parse mode of all sent messages.
	// It attaches to every send, edit or whatever method. You also
	// will be able to override the default mode by passing a new one.
	ParseMode ParseMode

	// OnError is a callback function that will get called on errors
	// resulted from the handler. It is used as post-middleware function.
	// Notice that context can be nil. Receives error, context and stack trace.
	OnError func(error, Ctx, DebugInfo[Ctx, HandlerFunc, MiddlewareFunc])

	// HTTP Client used to make requests to telegram api
	Client *http.Client

	// Offline allows to create a bot without network for testing purposes.
	Offline bool

	// RetryPolicy configures automatic retry behavior for failed API calls.
	// If nil, no automatic retry is performed.
	// Use DefaultRetryPolicy() for sensible defaults.
	RetryPolicy *RetryPolicy

	// RateLimiter limits the rate of API calls to prevent hitting Telegram limits.
	// If nil, no rate limiting is applied.
	// Recommended: NewRateLimiter(30, time.Second) for 30 requests per second.
	RateLimiter *RateLimiter

	// Censorer censors text sent by the bot.
	// .CensorText(string) string receives the original text and returns the censored text.
	// Use censorship.NewSpecificSubstringsCensorer([]string{"bad word", "another bad word"}) to create a new censorer that filters out the specified phrases.
	// Can be used to filter out tokens, keys, usernames, etc.
	Censorer censorship.Censorer
}

func (s *Settings[Ctx, HandlerFunc, MiddlewareFunc]) DefaultsForEmptyValues() {
	if s.Updates == 0 {
		s.Updates = 100
	}

	if s.Client == nil {
		s.Client = &http.Client{Timeout: time.Minute}
	}

	if s.URL == "" {
		s.URL = TelegramAPIURL
	}
	if s.Poller == nil {
		s.Poller = &LongPoller{}
	}
	if s.OnError == nil {
		s.OnError = defaultOnError
	}

}
