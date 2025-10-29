package telegram

import (
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Contact represents a phone contact.
type Contact struct {
	PhoneNumber string `json:"phone_number"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name,omitempty"`
	UserID      int64  `json:"user_id,omitempty"`
	VCard       string `json:"vcard,omitempty"`
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (c *Contact) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.Add("phone_number", c.PhoneNumber)
	b.Add("first_name", c.FirstName)
	b.Add("last_name", c.LastName)
	b.AddInt64("user_id", c.UserID)
	b.Add("vcard", c.VCard)

	return &outgoing.Method{
		Name:   "sendContact",
		Params: b.Build(),
		Files:  nil,
	}
}
