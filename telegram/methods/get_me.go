package methods

import "github.com/alagunto/tb/telegram"

// GetMeRequest represents the request for getMe method.
type GetMeRequest struct{}

// GetMeResponse represents the response for getMe method.
type GetMeResponse = telegram.User
