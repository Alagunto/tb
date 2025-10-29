package tb

import (
	"testing"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/request"
	"github.com/alagunto/tb/telegram"
	"github.com/stretchr/testify/assert"
)

type testPoller[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc] struct {
	updates chan telegram.Update
	done    chan struct{}
}

func newTestPoller[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc]() *testPoller[RequestType, HandlerFunc, MiddlewareFunc] {
	return &testPoller[RequestType, HandlerFunc, MiddlewareFunc]{
		updates: make(chan telegram.Update),
		done:    make(chan struct{}),
	}
}

func (p *testPoller[RequestType, HandlerFunc, MiddlewareFunc]) Poll(b bot.API, updates chan telegram.Update, stop chan struct{}) {
	for {
		select {
		case upd := <-p.updates:
			updates <- upd
		case <-stop:
			return
		}
	}
}

func TestMiddlewarePoller(t *testing.T) {
	p := &testPoller[usedRequestType, usedHandlerFunc, usedMiddlewareFunc]{updates: make(chan telegram.Update), done: make(chan struct{})}
	var ids []int

	pref := defaultSettings[usedRequestType]()
	pref.Poller = p
	pref.Offline = true

	b, err := NewBot(defaultWrapBasicContext[usedRequestType, usedHandlerFunc, usedMiddlewareFunc], pref)
	if err != nil {
		t.Fatal(err)
	}

	b.Poller = NewMiddlewarePoller(p, func(u *telegram.Update) bool {
		if u.ID > 0 {
			ids = append(ids, u.ID)
			return true
		}

		p.done <- struct{}{}
		return false
	})

	go func() {
		p.updates <- telegram.Update{ID: 1}
		p.updates <- telegram.Update{ID: 2}
		p.updates <- telegram.Update{ID: 0}
	}()

	go b.Start()
	<-p.done
	b.Stop()

	assert.Contains(t, ids, 1)
	assert.Contains(t, ids, 2)
}
