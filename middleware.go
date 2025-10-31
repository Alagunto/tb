package tb

import "github.com/alagunto/tb/request"

func appendMiddleware[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](a, b []MiddlewareFunc) []MiddlewareFunc {
	if len(a) == 0 {
		return b
	}

	m := make([]MiddlewareFunc, 0, len(a)+len(b))
	return append(m, append(a, b...)...)
}

func applyMiddleware[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc](h HandlerFunc, m ...MiddlewareFunc) HandlerFunc {
	for i := len(m) - 1; i >= 0; i-- {
		h = m[i](h)
	}
	return h
}

// Group is a separated group of handlers, united by the general middleware.
type Group[RequestType request.Interface] struct {
	b          *Bot[RequestType]
	middleware []func(func(RequestType) error) func(RequestType) error
}

// Use adds middleware to the chain.
func (g *Group[RequestType]) Use(middleware ...func(func(RequestType) error) func(RequestType) error) {
	g.middleware = append(g.middleware, middleware...)
}

// Handle adds endpoint handler to the bot, combining group's middleware
// with the optional given middleware.
func (g *Group[RequestType]) Handle(endpoint interface{}, h func(RequestType) error, m ...func(func(RequestType) error) func(RequestType) error) {
	g.b.Handle(endpoint, h, appendMiddleware[RequestType](g.middleware, m)...)
}
