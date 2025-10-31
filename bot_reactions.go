package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/telegram"
)

// SetMessageReaction changes reactions on a message. Returns true on success.
func (b *Bot[RequestType]) SetMessageReaction(msg bot.Editable, reactions telegram.Reactions, isBig bool) error {
	msgID, chatID := msg.MessageSig()

	params := make(map[string]any)
	params["chat_id"] = chatID
	params["message_id"] = msgID
	params["reaction"] = reactions
	params["is_big"] = isBig

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setMessageReaction", params)
	return err
}
