package middleware

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/alagunto/tb"
)

var b, _ = tb.NewBot(tb.Settings{Offline: true})

func TestRecover(t *testing.T) {
	onError := func(err error, c tb.Context) {
		require.Error(t, err, "recover test")
	}

	h := func(c tb.Context) error {
		panic("recover test")
	}

	assert.Panics(t, func() {
		h(nil)
	})

	assert.NotPanics(t, func() {
		Recover(onError)(h)(nil)
	})
}
