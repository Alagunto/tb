package telegram

import (
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

// Dice represents an animated emoji that displays a random value.
type Dice struct {
	Emoji string `json:"emoji"`
	Value int    `json:"value"`
}

// ToTelegramSendMethod implements the outgoing.Content interface.
func (d *Dice) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.Add("emoji", d.Emoji)

	return &outgoing.Method{
		Name:   "sendDice",
		Params: b.Build(),
		Files:  nil,
	}
}
