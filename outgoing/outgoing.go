package outgoing

import "github.com/alagunto/tb/files"

// Content is an interface for all objects that can be sent via Telegram.
// It describes how the object should be sent (method name, parameters, files).
//
// This is the core abstraction that makes file uploads explicit:
// - Content describes what to send and how
// - The bot's Send() method handles the actual upload/sending
// - ResponseHandler (if implemented) updates the object with Telegram's response
//
// Example:
//
//	photo := &Photo{
//	    Source:  files.UseLocalFile("image.jpg"),
//	    Caption: "Hello!",
//	}
//	msg, err := bot.Send(chat, photo, nil)
type Content interface {
	// ToTelegramSendMethod describes how to send this content to Telegram.
	// It returns the method name, parameters, and file sources.
	ToTelegramSendMethod() *Method
}

// Method describes a Telegram API method call with its parameters and files.
type Method struct {
	// Name is the Telegram API method name (e.g., "sendPhoto", "sendMessage")
	Name string

	// Params are the text parameters to send with the request
	Params map[string]any

	// Files are the file sources to upload (if any)
	Files map[string]files.FileSource
}

// ResponseHandler is an optional interface that Content types can implement
// to update themselves with data from Telegram's response.
//
// This is typically used to populate FileRef fields after upload.
// The interface uses interface{} to avoid import cycles - the actual type
// passed will be *telegram.Message.
type ResponseHandler interface {
	// UpdateFromResponse updates the object with data from Telegram's response.
	// The msg parameter will be of type *telegram.Message.
	UpdateFromResponse(msg interface{}) error
}
