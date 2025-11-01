package tb

import (
	"encoding/json"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/telegram"
)

// React changes the chosen reactions on a message. Service messages can't be
// reacted to. Automatically forwarded messages from a channel to its discussion group have
// the same available reactions as messages in the channel.
func (b *Bot[RequestType]) React(to bot.Recipient, msg bot.Editable, r telegram.Reactions) error {
	if to == nil {
		return errors.WithInvalidParam(errors.ErrBadRecipient, "recipient", "nil")
	}

	msgID, _ := msg.MessageSig()
	params := map[string]string{
		"chat_id":    to.Recipient(),
		"message_id": msgID,
	}

	data, _ := json.Marshal(r)
	params["reaction"] = string(data)

	_, err := b.Raw( "setMessageReaction", params)
	return err
}
