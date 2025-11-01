package telegram

// Update object represents an incoming update.
// At most one of the optional parameters can be present in any given update.
//
// Source: https://core.telegram.org/bots/api#update
type Update struct {
	ID int `json:"update_id"` // update_id	Integer	The update's unique identifier. Update identifiers start from a certain positive number and increase sequentially.

	Message           *Message `json:"message,omitempty"`             // message	Message	Optional. New incoming message of any kind - text, photo, sticker, etc.
	EditedMessage     *Message `json:"edited_message,omitempty"`      // edited_message	Message	Optional. New version of a message that is known to the bot and was edited
	ChannelPost       *Message `json:"channel_post,omitempty"`        // channel_post	Message	Optional. New incoming channel post of any kind - text, photo, sticker, etc.
	EditedChannelPost *Message `json:"edited_channel_post,omitempty"` // edited_channel_post	Message	Optional. New version of a channel post that is known to the bot and was edited

	// The new status of a reaction to a message.
	MessageReaction *MessageReactionUpdated `json:"message_reaction,omitempty"`
	// The new count of reactions to a message.
	MessageReactionCount *MessageReactionCountUpdated `json:"message_reaction_count,omitempty"`

	CallbackQuery      *Callback      `json:"callback_query,omitempty"`       // callback_query	CallbackQuery	Optional. New incoming callback query
	InlineQuery        *InlineQuery        `json:"inline_query,omitempty"`         // inline_query	InlineQuery	Optional. New incoming inline query
	ChosenInlineResult *ChosenInlineResult `json:"chosen_inline_result,omitempty"` // chosen_inline_result	ChosenInlineResult	Optional. The result of an inline query that was chosen by a user and sent to their chat partner

	ShippingQuery    *ShippingQuery    `json:"shipping_query,omitempty"`     // shipping_query	ShippingQuery	Optional. New incoming shipping query. Only for invoices with flexible price
	PreCheckoutQuery *PreCheckoutQuery `json:"pre_checkout_query,omitempty"` // pre_checkout_query	PreCheckoutQuery	Optional. New incoming pre-checkout query. Contains full information about checkout

	Poll       *Poll       `json:"poll,omitempty"`        // poll	Poll	Optional. New poll state. Bots receive only updates about manually stopped polls and polls, which are sent by the bot itself.
	PollAnswer *PollAnswer `json:"poll_answer,omitempty"` // poll_answer	PollAnswer	Optional. A user changed their answer in a non-anonymous poll. Bots receive new votes only in polls that were sent by the bot itself.

	MyChatMember    *ChatMember      `json:"my_chat_member,omitempty"`    // my_chat_member	ChatMemberUpdated	Optional. The bot's chat member status was updated in a chat. For private chats, this update is received only when the bot is blocked or unblocked by the user.
	ChatMember      *ChatMember      `json:"chat_member,omitempty"`       // chat_member	ChatMemberUpdated	Optional. A chat member's status was updated in a chat. The bot must be an administrator in the chat and must explicitly specify "chat_member" in the list of allowed_updates to receive these updates.
	ChatJoinRequest *ChatJoinRequest `json:"chat_join_request,omitempty"` // chat_join_request	ChatJoinRequest	Optional. A request to join the chat has been sent. The bot must have the can_invite_users administrator right in the chat to receive these updates.

	ChatBoost        *BoostUpdated `json:"chat_boost,omitempty"`         // chat_boost	ChatBoostUpdated	Optional. A chat boost was added or changed. The bot must be an administrator in the chat to receive these updates.
	RemovedChatBoost *BoostRemoved `json:"removed_chat_boost,omitempty"` // removed_chat_boost	ChatBoostRemoved	Optional. A boost was removed from a chat. The bot must be an administrator in the chat to receive these updates.

	BusinessConnection      *BusinessConnection      `json:"business_connection,omitempty"`       // business_connection	BusinessConnection	Optional. The bot was connected to or disconnected from a business account, or a user edited an existing connection with the bot
	BusinessMessage         *Message                 `json:"business_message,omitempty"`          // business_message	Message	Optional. New message from a connected business account
	EditedBusinessMessage   *Message                 `json:"edited_business_message,omitempty"`   // edited_business_message	Message	Optional. New version of a message from a connected business account
	DeletedBusinessMessages *BusinessMessagesDeleted `json:"deleted_business_messages,omitempty"` // deleted_business_messages	BusinessMessagesDeleted	Optional. Messages were deleted from a connected business account
}
