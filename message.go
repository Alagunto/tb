package tb

import "github.com/alagunto/tb/telegram"

// Message is an alias for telegram.Message
type Message = telegram.Message

// Type aliases for reactions
type (
	MessageReaction      = telegram.MessageReaction
	MessageReactionCount = telegram.MessageReactionCount
)

// Type aliases for menu and callback
type (
	MenuButton       = telegram.MenuButton
	CallbackResponse = telegram.CallbackResponse
	Location         = telegram.Location
	ShippingOption   = telegram.ShippingOption
)
