package tb

import (
	"errors"
	"fmt"
	"runtime/debug"
	"strings"
	"sync"
	"time"

	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/telegram"
)

// ContextInterface is the context interface - defined locally to avoid import cycles
type ContextInterface interface {
	Bot() API
	Update() Update
	Message() *telegram.Message
	Callback() *telegram.CallbackQuery
	Query() *telegram.InlineQuery
	InlineResult() *telegram.InlineResult
	ShippingQuery() *telegram.ShippingQuery
	PreCheckoutQuery() *telegram.PreCheckoutQuery
	Payment() *telegram.Payment
	Poll() *telegram.Poll
	PollAnswer() *telegram.PollAnswer
	ChatMember() *telegram.ChatMemberUpdate
	ChatJoinRequest() *telegram.ChatJoinRequest
	Migration() (int64, int64)
	Topic() *telegram.Topic
	Boost() *telegram.BoostUpdated
	BoostRemoved() *telegram.BoostRemoved
	Sender() *telegram.User
	Chat() *telegram.Chat
	Recipient() Recipient
	Text() string
	ThreadID() int
	Entities() Entities
	Data() string
	Args() []string
	Send(what interface{}, opts ...interface{}) error
	SendAlbum(a Album, opts ...interface{}) error
	Reply(what interface{}, opts ...interface{}) error
	Forward(msg Editable, opts ...interface{}) error
	ForwardTo(to Recipient, opts ...interface{}) error
	Edit(what interface{}, opts ...interface{}) error
	EditCaption(caption string, opts ...interface{}) error
	EditOrSend(what interface{}, opts ...interface{}) error
	EditOrReply(what interface{}, opts ...interface{}) error
	Delete() error
	DeleteAfter(d time.Duration) *time.Timer
	Notify(action telegram.ChatAction) error
	Ship(what ...interface{}) error
	Accept(errorMessage ...string) error
	Answer(resp *telegram.QueryResponse) error
	Respond(resp ...*telegram.CallbackResponse) error
	RespondText(text string) error
	RespondAlert(text string) error
	Get(key string) interface{}
	Set(key string, val interface{})
}

// NewContext returns a new native context object,
// field by the passed update.
func NewContext(b API, u Update) ContextInterface {
	return &nativeContext{
		b: b,
		u: u,
	}
}

// nativeContext is a native implementation of the Context interface.
// "context" is taken by context package, maybe there is a better name.
type nativeContext struct {
	b     API
	u     Update
	lock  sync.RWMutex
	store map[string]interface{}
}

func (c *nativeContext) Bot() API {
	return c.b
}

func (c *nativeContext) Update() Update {
	return c.u
}

func (c *nativeContext) Message() *telegram.Message {
	switch {
	case c.u.Message != nil:
		return c.u.Message
	case c.u.CallbackQuery != nil:
		return c.u.CallbackQuery.Message
	case c.u.EditedMessage != nil:
		return c.u.EditedMessage
	case c.u.ChannelPost != nil:
		if c.u.ChannelPost.PinnedMessage != nil {
			return c.u.ChannelPost.PinnedMessage
		}
		return c.u.ChannelPost
	case c.u.EditedChannelPost != nil:
		return c.u.EditedChannelPost
	default:
		return nil
	}
}

func (c *nativeContext) Callback() *telegram.CallbackQuery {
	return c.u.CallbackQuery
}

func (c *nativeContext) Query() *telegram.InlineQuery {
	return c.u.InlineQuery
}

func (c *nativeContext) InlineResult() *telegram.InlineResult {
	return c.u.ChosenInlineResult
}

func (c *nativeContext) ShippingQuery() *telegram.ShippingQuery {
	return c.u.ShippingQuery
}

func (c *nativeContext) PreCheckoutQuery() *telegram.PreCheckoutQuery {
	return c.u.PreCheckoutQuery
}

// Assert that this context contains a pre checkout query.
func (c *nativeContext) AssertPreCheckoutQuery() error {
	if c.PreCheckoutQuery() == nil {
		return ErrWithCurrentStack(ErrWithInvalidParam(errors.New("telebot: context pre checkout query is nil"), "pre_checkout_query", "nil"))
	}

	return nil
}

// Callback returns callback query if it exists.
//
// Assert that this context contains a callback query.
func (c *nativeContext) Payment() *telegram.Payment {
	if c.u.Message == nil {
		return nil
	}
	return c.u.Message.Payment
}

func (c *nativeContext) ChatMember() *telegram.ChatMemberUpdate {
	switch {
	case c.u.ChatMember != nil:
		return c.u.ChatMember
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember
	default:
		return nil
	}
}

func (c *nativeContext) ChatJoinRequest() *telegram.ChatJoinRequest {
	return c.u.ChatJoinRequest
}

func (c *nativeContext) Poll() *telegram.Poll {
	return c.u.Poll
}

func (c *nativeContext) PollAnswer() *telegram.PollAnswer {
	return c.u.PollAnswer
}

func (c *nativeContext) Migration() (int64, int64) {
	m := c.u.Message
	if m == nil {
		return 0, 0
	}
	return m.MigrateFrom, m.MigrateTo
}

func (c *nativeContext) Topic() *Topic {
	m := c.u.Message
	if m == nil {
		return nil
	}
	switch {
	case m.TopicCreated != nil:
		return m.TopicCreated
	case m.TopicReopened != nil:
		return m.TopicReopened
	case m.TopicEdited != nil:
		return m.TopicEdited
	}
	return nil
}

func (c *nativeContext) Boost() *telegram.BoostUpdated {
	return c.u.ChatBoost
}

func (c *nativeContext) BoostRemoved() *telegram.BoostRemoved {
	return c.u.RemovedChatBoost
}

func (c *nativeContext) Sender() *telegram.User {
	switch {
	case c.u.CallbackQuery != nil:
		return c.u.CallbackQuery.Sender
	case c.Message() != nil:
		return c.Message().Sender
	case c.u.InlineQuery != nil:
		return c.u.InlineQuery.Sender
	case c.u.ChosenInlineResult != nil:
		return c.u.ChosenInlineResult.Sender
	case c.u.ShippingQuery != nil:
		return c.u.ShippingQuery.Sender
	case c.u.PreCheckoutQuery != nil:
		return c.u.PreCheckoutQuery.Sender
	case c.u.PollAnswer != nil:
		return c.u.PollAnswer.Sender
	case c.u.MyChatMember != nil:
		return c.u.MyChatMember.User
	case c.u.ChatMember != nil:
		return c.u.ChatMember.User
	case c.u.ChatJoinRequest != nil:
		return c.u.ChatJoinRequest.Sender
	case c.u.ChatBoost != nil:
		if b := c.u.ChatBoost.Boost; b != nil && b.Source != nil {
			return b.Source.Booster
		}
	case c.u.RemovedChatBoost != nil:
		if b := c.u.RemovedChatBoost; b.Source != nil {
			return b.Source.Booster
		}
	}
	return nil
}

func (c *nativeContext) Chat() *telegram.Chat {
	switch {
	case c.Message() != nil:
		return c.Message().Chat
	case c.u.ChatJoinRequest != nil:
		return c.u.ChatJoinRequest.Chat
	default:
		return nil
	}
}

func (c *nativeContext) Recipient() communications.Recipient {
	chat := c.Chat()
	if chat != nil {
		return chat
	}
	return c.Sender()
}

func (c *nativeContext) Text() string {
	m := c.Message()
	if m == nil {
		return ""
	}
	if m.Caption != "" {
		return m.Caption
	}
	return m.Text
}

func (c *nativeContext) Entities() Entities {
	m := c.Message()
	if m == nil {
		return nil
	}
	if len(m.CaptionEntities) > 0 {
		return m.CaptionEntities
	}
	return m.Entities
}

func (c *nativeContext) Data() string {
	switch {
	case c.u.Message != nil:
		m := c.u.Message
		if m.Payment != nil {
			return m.Payment.Payload
		}
		return m.Payload
	case c.u.CallbackQuery != nil:
		return c.u.CallbackQuery.Data
	case c.u.InlineQuery != nil:
		return c.u.InlineQuery.Text
	case c.u.ChosenInlineResult != nil:
		return c.u.ChosenInlineResult.Query
	case c.u.ShippingQuery != nil:
		return c.u.ShippingQuery.Payload
	case c.u.PreCheckoutQuery != nil:
		return c.u.PreCheckoutQuery.Payload
	default:
		return ""
	}
}

func (c *nativeContext) Args() []string {
	m := c.u.Message
	switch {
	case m != nil && m.Payment != nil:
		return strings.Split(m.Payment.Payload, "|")
	case m != nil:
		payload := strings.Trim(m.Payload, " ")
		if payload != "" {
			return strings.Fields(payload)
		}
	case c.u.CallbackQuery != nil:
		return strings.Split(c.u.CallbackQuery.Data, "|")
	case c.u.InlineQuery != nil:
		return strings.Split(c.u.InlineQuery.Text, " ")
	case c.u.ChosenInlineResult != nil:
		return strings.Split(c.u.ChosenInlineResult.Query, " ")
	}
	return nil
}

func (c *nativeContext) ThreadID() int {
	switch {
	case c.Message() != nil:
		return c.Message().ThreadID
	default:
		return 0
	}
}

func (c *nativeContext) Send(what interface{}, opts ...interface{}) error {
	opts = c.inheritOpts(opts...)
	_, err := c.b.Send(c.Recipient(), what, opts...)
	return err
}

func (c *nativeContext) inheritOpts(opts ...interface{}) []interface{} {
	var (
		ignoreThread bool
	)

	if opts == nil {
		opts = make([]interface{}, 0)
	}

	for _, opt := range opts {
		switch opt.(type) {
		case Option:
			switch opt {
			case IgnoreThread:
				ignoreThread = true
			default:
			}
		}
	}

	switch {
	case !ignoreThread && c.ThreadID() != 0:
		opts = append(opts, &Topic{ThreadID: c.ThreadID()})
	}

	return opts
}

func (c *nativeContext) SendAlbum(a Album, opts ...interface{}) error {
	opts = c.inheritOpts(opts...)

	_, err := c.b.SendAlbum(c.Recipient(), a, opts...)
	return err
}

func (c *nativeContext) Reply(what interface{}, opts ...interface{}) error {
	msg := c.Message()
	if msg == nil {
		return ErrWithCurrentStack(ErrWithInvalidParam(ErrBadContext, "message", "nil"))
	}
	opts = c.inheritOpts(opts...)
	_, err := c.b.Reply(msg, what, opts...)
	return err
}

func (c *nativeContext) Forward(msg Editable, opts ...interface{}) error {
	_, err := c.b.Forward(c.Recipient(), msg, opts...)
	return err
}

func (c *nativeContext) ForwardTo(to Recipient, opts ...interface{}) error {
	msg := c.Message()
	if msg == nil {
		return ErrBadContext
	}
	_, err := c.b.Forward(to, msg, opts...)
	return err
}

func (c *nativeContext) Edit(what interface{}, opts ...interface{}) error {
	opts = c.inheritOpts(opts...)

	if c.u.ChosenInlineResult != nil {
		_, err := c.b.Edit(c.u.ChosenInlineResult, what, opts...)
		return err
	}
	if c.u.CallbackQuery != nil {
		_, err := c.b.Edit(c.u.CallbackQuery, what, opts...)
		return err
	}
	return ErrBadContext
}

func (c *nativeContext) EditCaption(caption string, opts ...interface{}) error {
	opts = c.inheritOpts(opts...)

	if c.u.ChosenInlineResult != nil {
		_, err := c.b.EditCaption(c.u.ChosenInlineResult, caption, opts...)
		return err
	}
	if c.u.CallbackQuery != nil {
		_, err := c.b.EditCaption(c.u.CallbackQuery, caption, opts...)
		return err
	}
	return ErrBadContext
}

func (c *nativeContext) EditOrSend(what interface{}, opts ...interface{}) error {
	err := c.Edit(what, opts...)
	if err == ErrBadContext {
		return c.Send(what, opts...)
	}
	return err
}

func (c *nativeContext) EditOrReply(what interface{}, opts ...interface{}) error {
	err := c.Edit(what, opts...)
	if err == ErrBadContext {
		return c.Reply(what, opts...)
	}
	return err
}

func (c *nativeContext) Delete() error {
	msg := c.Message()
	if msg == nil {
		return ErrWithCurrentStack(ErrWithInvalidParam(ErrBadContext, "message", "nil"))
	}
	return c.b.Delete(msg)
}

func (c *nativeContext) DeleteAfter(d time.Duration) *time.Timer {
	return time.AfterFunc(d, func() {
		if err := c.Delete(); err != nil {
			if b, ok := c.b.(*Bot[*nativeContext, func(*nativeContext) error, func(func(*nativeContext) error) func(*nativeContext) error]); ok {
				debugInfo := DebugInfo[*nativeContext, func(*nativeContext) error, func(func(*nativeContext) error) func(*nativeContext) error]{
					Handler:  nil,
					Endpoint: "",
					Stack:    string(debug.Stack()),
				}
				b.OnError(err, c, debugInfo)
			}
		}
	})
}

func (c *nativeContext) Notify(action ChatAction) error {
	return c.b.Notify(c.Recipient(), action, c.ThreadID())
}

func (c *nativeContext) Ship(what ...interface{}) error {
	if c.u.ShippingQuery == nil {
		return ErrWithCurrentStack(ErrWithInvalidParam(errors.New("telebot: context shipping query is nil"), "shipping_query", "nil"))
	}
	return c.b.Ship(c.u.ShippingQuery, what...)
}

func (c *nativeContext) Accept(errorMessage ...string) error {
	if c.u.PreCheckoutQuery == nil {
		return ErrWithCurrentStack(ErrWithInvalidParam(errors.New("telebot: context pre checkout query is nil"), "pre_checkout_query", "nil"))
	}
	return c.b.Accept(c.u.PreCheckoutQuery, errorMessage...)
}

func (c *nativeContext) Respond(resp ...*CallbackResponse) error {
	if c.u.CallbackQuery == nil {
		return ErrWithCurrentStack(ErrWithInvalidParam(errors.New("telebot: context callback is nil"), "callback", "nil"))
	}
	return c.b.Respond(c.u.CallbackQuery, resp...)
}

func (c *nativeContext) RespondText(text string) error {
	return c.Respond(&CallbackResponse{Text: text})
}

func (c *nativeContext) RespondAlert(text string) error {
	return c.Respond(&CallbackResponse{Text: text, ShowAlert: true})
}

func (c *nativeContext) Answer(resp *telegram.QueryResponse) error {
	if c.u.InlineQuery == nil {
		return ErrWithCurrentStack(ErrWithInvalidParam(errors.New("telebot: context inline query is nil"), "inline_query", "nil"))
	}
	return c.b.Answer(c.u.InlineQuery, resp)
}

func (c *nativeContext) Set(key string, value interface{}) {
	c.lock.Lock()
	defer c.lock.Unlock()

	if c.store == nil {
		c.store = make(map[string]interface{})
	}

	c.store[key] = value
}

func (c *nativeContext) Get(key string) interface{} {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.store[key]
}

// GetTyped retrieves a value from the context with type safety.
// Returns the value and a boolean indicating whether the key exists and type matches.
func GetTyped[T any](c ContextInterface, key string) (T, bool) {
	val := c.Get(key)
	if val == nil {
		var zero T
		return zero, false
	}
	typed, ok := val.(T)
	return typed, ok
}

// MustGetTyped retrieves a value from the context and panics if not found or wrong type.
// Use this when you're certain the value exists and has the correct type.
func MustGetTyped[T any](c ContextInterface, key string) T {
	val, ok := GetTyped[T](c, key)
	if !ok {
		panic(fmt.Sprintf("telebot: key %q not found or has wrong type", key))
	}
	return val
}

// GetTypedOr retrieves a value from the context or returns a default value.
func GetTypedOr[T any](c ContextInterface, key string, defaultVal T) T {
	val, ok := GetTyped[T](c, key)
	if !ok {
		return defaultVal
	}
	return val
}
