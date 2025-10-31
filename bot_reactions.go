package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// SetMessageReaction changes reactions on a message. Returns true on success.
func (b *Bot[RequestType]) SetMessageReaction(msg bot.Editable, reactions telegram.Reactions, isBig bool) error {
	msgID, chatID := msg.MessageSig()

	p := params.New().
		Add("chat_id", chatID).
		Add("message_id", msgID).
		Add("reaction", reactions).
		AddBool("is_big", isBig).
		Build()

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "setMessageReaction", p)
	return err
}
