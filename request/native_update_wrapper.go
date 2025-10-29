package request

import (
	"strings"
	"sync"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/telegram"
)

// NewContext returns a new native context object,
// field by the passed update.
func NewNativeContext(b bot.API, u telegram.Update) *nativeContext {
	return &nativeContext{
		API: b,
		u:   u,
	}
}

// nativeContext is a native implementation of the Context interface.
// "context" is taken by context package, maybe there is a better name.
type nativeContext struct {
	bot.API
	u     telegram.Update
	lock  sync.RWMutex
	store map[string]interface{}
}

var _ Interface = (*nativeContext)(nil)

func (c *nativeContext) Update() *telegram.Update {
	return &c.u
}

func (c *nativeContext) Bot() interface{} {
	return c.API
}

func (c *nativeContext) contextualSendOptions() communications.SendOptions {
	// Add thread ID to the send options if we are responding to a thread
	return communications.NewSendOptions().
		WithThreadID(c.ThreadID())
}

func (c *nativeContext) Message() *telegram.Message {
	switch {
	case c.u.Message != nil:
		return c.u.Message
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

func (c *nativeContext) CallbackMessage() *telegram.Message {
	if c.u.CallbackQuery == nil {
		return nil
	}
	return c.u.CallbackQuery.Message
}

func (c *nativeContext) CallbackQuery() *telegram.CallbackQuery {
	return c.u.CallbackQuery
}

func (c *nativeContext) InlineQuery() *telegram.InlineQuery {
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

func (c *nativeContext) Payment() *telegram.Payment {
	if c.u.Message == nil {
		return nil
	}
	return c.u.Message.Payment
}

func (c *nativeContext) ChatMember() *telegram.ChatMember {
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

func (c *nativeContext) Thread() *telegram.Thread {
	m := c.u.Message
	if m == nil {
		return nil
	}
	switch {
	case m.ThreadCreated != nil:
		return m.ThreadCreated
	case m.ThreadReopened != nil:
		return m.ThreadReopened
	case m.ThreadEdited != nil:
		return m.ThreadEdited
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
	case c.u.ChatBoost != nil:
		return c.u.ChatBoost.Chat
	case c.u.RemovedChatBoost != nil:
		return c.u.RemovedChatBoost.Chat
	// TODO: verify if other cases are available
	default:
		return nil
	}
}

func (c *nativeContext) Recipient() bot.Recipient {
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

func (c *nativeContext) Entities() telegram.Entities {
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

func (c *nativeContext) Send(what interface{}, opts ...communications.SendOptions) error {
	opt := c.contextualSendOptions().MergeWithMany(opts...)
	_, err := c.API.SendTo(c.Recipient(), what, opt)
	return err
}

func (c *nativeContext) SendAlbum(a telegram.Album, opts ...communications.SendOptions) error {
	_, err := c.API.SendAlbumTo(c.Recipient(), a, opts...)
	return err
}

func (c *nativeContext) Reply(what interface{}, opts ...communications.SendOptions) error {
	msg := c.Message()
	if msg == nil {
		return errors.WithMissingEntity(errors.ErrContextInsufficient, errors.MissingEntityMessage)
	}
	opt := c.contextualSendOptions().MergeWithMany(opts...)
	_, err := c.API.ReplyTo(msg, what, opt)
	return err
}

func (c *nativeContext) Forward(msg bot.Editable, opts ...communications.SendOptions) error {
	opt := c.contextualSendOptions().MergeWithMany(opts...)
	_, err := c.API.ForwardTo(c.Recipient(), msg, opt)
	return err
}

func (c *nativeContext) EditLast(what interface{}, opts ...communications.SendOptions) error {
	if c.u.ChosenInlineResult != nil {
		_, err := c.Edit(c.u.ChosenInlineResult, what, opts...)
		return err
	}
	if c.u.CallbackQuery != nil {
		_, err := c.Edit(c.u.CallbackQuery, what, opts...)
		return err
	}
	return errors.WithMissingEntity(errors.ErrContextInsufficient, errors.MissingEntityMessage)
}

func (c *nativeContext) EditLastCaption(caption string, opts ...communications.SendOptions) error {
	if c.u.ChosenInlineResult != nil {
		_, err := c.API.EditCaption(c.u.ChosenInlineResult, caption, opts...)
		return err
	}
	if c.u.CallbackQuery != nil {
		_, err := c.API.EditCaption(c.u.CallbackQuery, caption, opts...)
		return err
	}
	return errors.ErrNothingToEdit
}

func (c *nativeContext) DeleteLast() error {
	msg := c.Message()
	if msg == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "message", nil)
	}
	return c.API.Delete(msg)
}

func (c *nativeContext) AnswerInlineQuery(query *telegram.InlineQuery, resp *telegram.QueryResponse) error {
	// If query is not provided, use the one from the update
	if query == nil {
		if c.u.InlineQuery == nil {
			return errors.WithInvalidParam(errors.ErrTelebot, "inline_query", nil)
		}
		query = c.u.InlineQuery
	}
	return c.AnswerInlineQuery(query, resp)
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
