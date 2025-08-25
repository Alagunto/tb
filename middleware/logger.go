package middleware

import (
	"encoding/json"
	"log"

	"github.com/alagunto/tb"
)

// Logger returns a middleware that logs incoming updates.
// If no custom logger provided, log.Default() will be used.
func Logger[Ctx tb.ContextInterface, HandlerFunc func(Ctx) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](logger ...*log.Logger) MiddlewareFunc {
	var l *log.Logger
	if len(logger) > 0 {
		l = logger[0]
	} else {
		l = log.Default()
	}

	return func(next HandlerFunc) HandlerFunc {
		return func(c Ctx) error {
			data, _ := json.MarshalIndent(c.(), "", "  ")
			l.Println(string(data))
			return next(c)
		}
	}
}
