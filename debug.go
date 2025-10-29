package tb

import "github.com/alagunto/tb/request"

type DebugInfo[RequestType request.Interface, HandlerFunc func(RequestType) error, MiddlewareFunc func(HandlerFunc) HandlerFunc] struct {
	Handler  HandlerFunc
	Stack    string
	Endpoint string
}
