package tb

import (
	"encoding/json"

	"github.com/alagunto/tb/communications"
	"github.com/alagunto/tb/telegram"
)

// React changes the chosen reactions on a message. Service messages can't be
// reacted to. Automatically forwarded messages from a channel to its discussion group have
// the same available reactions as messages in the channel.
func (b *Bot[Ctx, HandlerFunc, MiddlewareFunc]) React(to communications.Recipient, msg communications.Editable, r telegram.Reactions) error {
	if to == nil {
		return ErrWithCurrentStack(ErrWithInvalidParam(ErrBadRecipient, "recipient", "nil"))
	}

	msgID, _ := msg.MessageSig()
	params := map[string]string{
		"chat_id":    to.Recipient(),
		"message_id": msgID,
	}

	data, _ := json.Marshal(r.Reactions)
	params["reaction"] = string(data)

	if r.Big {
		params["is_big"] = "true"
	}

	_, err := b.Raw("setMessageReaction", params)
	return err
}
