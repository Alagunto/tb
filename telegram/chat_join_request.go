package telegram

// ChatJoinRequest represents a join request sent to a chat.
type ChatJoinRequest struct {
	// Chat to which the request was sent.
	Chat *Chat `json:"chat"`

	// Sender is the user that sent the join request.
	Sender *User `json:"from"`

	// UserChatID is an ID of a private chat with the user
	// who sent the join request. The bot can use this ID
	// for 5 minutes to send messages until the join request
	// is processed, assuming no other administrator contacted the user.
	UserChatID int64 `json:"user_chat_id"`

	// Unixtime, use ChatJoinRequest.Time() to get time.Time.
	Unixtime int64 `json:"date"`

	// Bio of the user, optional.
	Bio string `json:"bio"`

	// InviteLink is the chat invite link that was used by
	//the user to send the join request, optional.
	InviteLink *ChatInviteLink `json:"invite_link"`
}
