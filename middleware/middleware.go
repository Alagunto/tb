package middleware

import (
	"errors"
	"log"

	"github.com/alagunto/tb"
)

// AutoRespond returns a middleware that automatically responds
// to every callback.
func AutoRespond() tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			if c.Callback() != nil {
				defer c.Respond()
			}
			return next(c)
		}
	}
}

// IgnoreVia returns a middleware that ignores all the
// "sent via" messages.
func IgnoreVia() tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			if msg := c.Message(); msg != nil && msg.Via != nil {
				return nil
			}
			return next(c)
		}
	}
}

type RecoverFunc = func(error, tb.Context)

// Recover returns a middleware that recovers a panic happened in
// the handler.
func Recover(onError ...RecoverFunc) tb.MiddlewareFunc {
	return func(next tb.HandlerFunc) tb.HandlerFunc {
		return func(c tb.Context) error {
			var f RecoverFunc
			if len(onError) > 0 {
				f = onError[0]
			} else if b, ok := c.Bot().(*tb.Bot); ok {
				f = b.OnError
			} else {
				f = func(err error, _ tb.Context) {
					log.Println("telebot/middleware/recover:", err)
				}
			}

			defer func() {
				if r := recover(); r != nil {
					if err, ok := r.(error); ok {
						f(err, c)
					} else if s, ok := r.(string); ok {
						f(errors.New(s), c)
					}
				}
			}()

			return next(c)
		}
	}
}
