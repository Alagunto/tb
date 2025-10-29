package tb

import (
	"fmt"
	"net/http"
	"regexp"
	"sync"

	"github.com/alagunto/tb/censorship"
	"github.com/alagunto/tb/request"
	"github.com/alagunto/tb/telegram"
)

// NewBot does try to build a Bot with token `token`, which
// is a secret API key assigned to particular bot.
func NewBot[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](
	requestBuilder func(request.Interface) (RequestType, error),
	settings Settings[RequestType, HandlerFunc, MiddlewareFunc],
) (*Bot[RequestType, HandlerFunc, MiddlewareFunc], error) {
	settings.DefaultsForEmptyValues()

	bot := &Bot[RequestType, HandlerFunc, MiddlewareFunc]{
		settings: settings,

		requestBuilder: requestBuilder,
		token:          settings.Token,
		apiURL:         settings.URL,
		poller:         settings.Poller,
		onError:        settings.OnError,

		updates:          make(chan telegram.Update, settings.Updates),
		handlers:         make(map[string]HandlerFunc),
		originalHandlers: make(map[string]HandlerFunc),
		stop:             make(chan chan struct{}),

		client:   settings.Client,
		censorer: settings.Censorer,
	}

	// Instantly fetch the bot's information if not offline
	if settings.Offline {
		bot.me = &telegram.User{}
	} else {
		user, err := bot.getMe()
		if err != nil {
			return nil, err
		}
		bot.me = user
	}

	// Create the root group of handlers
	bot.group = bot.Group()

	return bot, nil
}

// Bot represents a separate Telegram bot instance.
type Bot[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc] struct {
	// Token is the bot's token, used to authenticate with the Telegram API
	token string

	// apiURL is the base URL of the Telegram API
	apiURL string

	// updates is the channel for incoming updates
	updates chan telegram.Update

	// poller is the poller for incoming updates
	poller Poller

	// onError is the callback function for errors
	onError func(error, RequestType, DebugInfo[RequestType, HandlerFunc, MiddlewareFunc])

	settings Settings[RequestType, HandlerFunc, MiddlewareFunc]

	// requestBuilder is the function to build request contexts that are passed to handlers
	requestBuilder func(request.Interface) (RequestType, error)

	// group is the root group of handlers tree
	group *Group[RequestType, HandlerFunc, MiddlewareFunc]
	// ???
	handlers         map[string]HandlerFunc
	originalHandlers map[string]HandlerFunc

	handlersWg sync.WaitGroup

	censorer censorship.Censorer
	client   *http.Client

	stopMu     sync.RWMutex
	stop       chan chan struct{}
	stopClient chan struct{}

	// me is the bot's user information, cached for performance
	me *telegram.User
}

// Group returns a new group.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Group() *Group[RequestType, HandlerFunc, MiddlewareFunc] {
	return &Group[RequestType, HandlerFunc, MiddlewareFunc]{b: b}
}

// Use adds middleware to the global bot chain.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Use(middleware ...MiddlewareFunc) {
	b.group.Use(middleware...)
}

var (
	cmdRx   = regexp.MustCompile(`^(/\w+)(@(\w+))?(\s|$)(.+)?`)
	cbackRx = regexp.MustCompile(`^\f([-\w]+)(\|(.+))?$`)
)

// Handle lets you set the handler for some command name or
// one of the supported endpoints. It also applies middleware
// if such passed to the function.
//
// Example:
//
//	b.Handle("/start", func (c tele.Context) error {
//		return c.Reply("Hello!")
//	})
//
//	b.Handle(&inlineButton, func (c tele.Context) error {
//		return c.Respond(&tele.CallbackResponse{Text: "Hello!"})
//	})
//
// Middleware usage:
//
//	b.Handle("/ban", onBan, middleware.Whitelist(ids...))
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Handle(endpoint interface{}, h HandlerFunc, m ...MiddlewareFunc) {
	end := extractEndpoint[RequestType](endpoint)
	if end == "" {
		panic("telebot: unsupported endpoint")
	}

	if len(b.group.middleware) > 0 {
		m = appendMiddleware[RequestType, HandlerFunc](b.group.middleware, m)
	}

	if _, ok := b.handlers[end]; ok {
		panic("telebot: handler is already registered for endpoint " + end + ", overriding the existing handler is almost always a bug")
	}
	handler := func(c RequestType) error {
		return applyMiddleware(h, m...)(c)
	}
	b.handlers[end] = handler
	b.originalHandlers[end] = h
}

// Trigger executes the registered handler by the endpoint.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Trigger(endpoint interface{}, c RequestType) error {
	end := extractEndpoint[RequestType](endpoint)
	if end == "" {
		return fmt.Errorf("telebot: unsupported endpoint: %v", endpoint)
	}

	handler, ok := b.handlers[end]
	if !ok {
		return fmt.Errorf("telebot: no handler found for given endpoint: %v", endpoint)
	}

	return handler(c)
}

// Start brings bot into motion by consuming incoming
// updates (see Bot.updates channel).
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Start() {
	if b.poller == nil {
		panic("telebot: can't start without a poller")
	}

	// do nothing if called twice
	b.stopMu.Lock()
	if b.stopClient != nil {
		b.stopMu.Unlock()
		return
	}

	b.stopClient = make(chan struct{})
	b.stopMu.Unlock()

	stop := make(chan struct{})
	stopConfirm := make(chan struct{})

	go func() {
		b.poller.Poll(b, b.updates, stop)
		close(stopConfirm)
	}()

	for {
		select {
		// handle incoming updates
		case upd := <-b.updates:
			ctx, err := b.NewContext(upd)
			if err != nil {
				b.onError(err, ctx, DebugInfo[RequestType, HandlerFunc, MiddlewareFunc]{})
				continue
			}
			b.ProcessUpdate(ctx, upd)
			// call to stop polling
		case confirm := <-b.stop:
			close(stop)
			<-stopConfirm
			close(confirm)
			return
		}
	}
}

// Stop gracefully shuts the poller down and waits for all handlers to complete.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Stop() {
	b.stopMu.Lock()
	ch := b.stopClient
	b.stopClient = nil
	b.stopMu.Unlock()

	if ch != nil {
		close(ch)
	}

	confirm := make(chan struct{})
	b.stop <- confirm
	<-confirm

	// Wait for all handlers to complete
	b.handlersWg.Wait()
}

// NewContext returns a new native context object,
// field by the passed update.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) NewContext(u telegram.Update) (RequestType, error) {
	return b.requestBuilder(request.NewNativeContext(b, u))
}

// CallbackEndpoint is an interface for callback buttons that have a unique identifier
type CallbackEndpoint interface {
	CallbackUnique() string
}

func extractEndpoint[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](endpoint interface{}) string {
	switch end := endpoint.(type) {
	case string:
		return end
	case CallbackEndpoint:
		return end.CallbackUnique()
	}
	return ""
}

// BotInfo represents a single object of BotName, BotDescription, BotShortDescription instances.
type BotInfo struct {
	Name             string `json:"name,omitempty"`
	Description      string `json:"description,omitempty"`
	ShortDescription string `json:"short_description,omitempty"`
}
