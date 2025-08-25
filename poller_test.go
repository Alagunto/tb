package tb

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type testPoller[Ctx ContextInterface, HandlerFunc func(Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc] struct {
	updates chan Update
	done    chan struct{}
}

func newTestPoller[Ctx ContextInterface, HandlerFunc func(Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc]() *testPoller[Ctx, HandlerFunc, MiddlewareFunc] {
	return &testPoller[Ctx, HandlerFunc, MiddlewareFunc]{
		updates: make(chan Update),
		done:    make(chan struct{}),
	}
}

func (p *testPoller[Ctx, HandlerFunc, MiddlewareFunc]) Poll(b *Bot[Ctx, HandlerFunc, MiddlewareFunc], updates chan Update, stop chan struct{}) {
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
	p := newTestPoller[usedCtx, usedHandlerFunc, usedMiddlewareFunc]()
	var ids []int

	pref := defaultSettings[usedCtx, usedHandlerFunc, usedMiddlewareFunc]()
	pref.Poller = p
	pref.Offline = true

	b, err := NewBot[usedCtx, usedHandlerFunc, usedMiddlewareFunc](pref)
	if err != nil {
		t.Fatal(err)
	}

	b.Poller = NewMiddlewarePoller[usedCtx, usedHandlerFunc, usedMiddlewareFunc](p, func(u *Update) bool {
		if u.ID > 0 {
			ids = append(ids, u.ID)
			return true
		}

		p.done <- struct{}{}
		return false
	})

	go func() {
		p.updates <- Update{ID: 1}
		p.updates <- Update{ID: 2}
		p.updates <- Update{ID: 0}
	}()

	go b.Start()
	<-p.done
	b.Stop()

	assert.Contains(t, ids, 1)
	assert.Contains(t, ids, 2)
}
