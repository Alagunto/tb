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
type Group[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc] struct {
	b          *Bot[RequestType, HandlerFunc, MiddlewareFunc]
	middleware []MiddlewareFunc
}

// Use adds middleware to the chain.
func (g *Group[RequestType, HandlerFunc, MiddlewareFunc]) Use(middleware ...MiddlewareFunc) {
	g.middleware = append(g.middleware, middleware...)
}

// Handle adds endpoint handler to the bot, combining group's middleware
// with the optional given middleware.
func (g *Group[RequestType, HandlerFunc, MiddlewareFunc]) Handle(endpoint interface{}, h HandlerFunc, m ...MiddlewareFunc) {
	g.b.Handle(endpoint, h, appendMiddleware[RequestType, HandlerFunc, MiddlewareFunc](g.middleware, m)...)
}
