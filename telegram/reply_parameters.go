package telegram

// ReplyParameters describes reply parameters for the message that is being sent.
// https://core.telegram.org/bots/api#replyparameters
type ReplyParameters struct {
	MessageID                int             `json:"message_id"`
	ChatID                   interface{}     `json:"chat_id,omitempty"`
	AllowSendingWithoutReply bool            `json:"allow_sending_without_reply,omitempty"`
	Quote                    string          `json:"quote,omitempty"`
	QuoteParseMode           ParseMode       `json:"quote_parse_mode,omitempty"`
	QuoteEntities            []MessageEntity `json:"quote_entities,omitempty"`
	QuotePosition            int             `json:"quote_position,omitempty"`
}
