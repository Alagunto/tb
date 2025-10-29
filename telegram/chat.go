package telegram

import "strconv"

// Chat object represents a Telegram user, bot, group or a channel.
type Chat struct {
	ID int64 `json:"id"`

	// See ChatType and consts.
	Type ChatType `json:"type"`

	// Won't be there for ChatPrivate.
	Title string `json:"title"`

	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
}

// Recipient returns chat id as string for sending messages.
// If chat has a username, returns @username instead.
func (c *Chat) Recipient() string {
	if c.Username != "" {
		return "@" + c.Username
	}
	return strconv.FormatInt(c.ID, 10)
}

// ChatType represents one of the possible chat types.
type ChatType string

const (
	ChatPrivate        ChatType = "private"
	ChatGroup          ChatType = "group"
	ChatSuperGroup     ChatType = "supergroup"
	ChatChannel        ChatType = "channel"
	ChatChannelPrivate ChatType = "privatechannel"
)

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

// ChatInviteLink object represents an invite for a chat.
type ChatInviteLink struct {
	// The invite link.
	InviteLink string `json:"invite_link"`

	// Invite link name.
	Name string `json:"name"`

	// The creator of the link.
	Creator *User `json:"creator"`

	// If the link is primary.
	IsPrimary bool `json:"is_primary"`

	// If the link is revoked.
	IsRevoked bool `json:"is_revoked"`

	// (Optional) Point in time when the link will expire,
	// use ExpireDate() to get time.Time.
	ExpireUnixtime int64 `json:"expire_date,omitempty"`

	// (Optional) Maximum number of users that can be members of
	// the chat simultaneously.
	MemberLimit int `json:"member_limit,omitempty"`

	// (Optional) True, if users joining the chat via the link need to
	// be approved by chat administrators. If True, member_limit can't be specified.
	JoinRequest bool `json:"creates_join_request"`

	// (Optional) Number of pending join requests created using this link.
	PendingCount int `json:"pending_join_request_count"`
}

// ChatLocation represents a location to which a chat is connected.
type ChatLocation struct {
	Location Location `json:"location,omitempty"`
	Address  string   `json:"address,omitempty"`
}

// ChatPhoto object represents a chat photo.
type ChatPhoto struct {
	// File identifiers of small (160x160) chat photo
	SmallFileID   string `json:"small_file_id"`
	SmallUniqueID string `json:"small_file_unique_id"`

	// File identifiers of big (640x640) chat photo
	BigFileID   string `json:"big_file_id"`
	BigUniqueID string `json:"big_file_unique_id"`
}
