package telegram

import (
	"strconv"
	"time"
	"unicode/utf16"

	"github.com/alagunto/tb/telegram/media"
)

// Message object represents a message.
type Message struct {
	ID int `json:"message_id"`

	// (Optional) Unique identifier of a message thread to which the message belongs; for supergroups only
	ThreadID int `json:"message_thread_id"`

	// For message sent to channels, Sender will be nil
	Sender *User `json:"from"`

	// Unixtime, use Message.Time() to get time.Time
	Unixtime int64 `json:"date"`

	// Conversation the message belongs to.
	Chat *Chat `json:"chat"`

	// Sender of the message, sent on behalf of a chat.
	SenderChat *Chat `json:"sender_chat"`

	// For forwarded messages, sender of the original message.
	OriginalSender *User `json:"forward_from"`

	// For forwarded messages, chat of the original message when
	// forwarded from a channel.
	OriginalChat *Chat `json:"forward_from_chat"`

	// For forwarded messages, identifier of the original message
	// when forwarded from a channel.
	OriginalMessageID int `json:"forward_from_message_id"`

	// For forwarded messages, signature of the post author.
	OriginalSignature string `json:"forward_signature"`

	// For forwarded messages, sender's name from users who
	// disallow adding a link to their account.
	OriginalSenderName string `json:"forward_sender_name"`

	// For forwarded messages, unixtime of the original message.
	OriginalUnixtime int `json:"forward_date"`

	// For information about the original message for forwarded messages.
	Origin *MessageOrigin `json:"forward_origin"`

	// Message is a channel post that was automatically forwarded to the connected discussion group.
	AutomaticForward bool `json:"is_automatic_forward"`

	// For replies, ReplyTo represents the original message.
	//
	// Note that the Message object in this field will not
	// contain further ReplyTo fields even if it
	// itself is a reply.
	ReplyTo *Message `json:"reply_to_message"`

	// (Optional) For replies to a story, the original story
	Story *Story `json:"story"`

	// (Optional) Information about the message that is being replied to,
	// which may come from another chat or forum topic.
	ExternalReply *ExternalReply `json:"external_reply"`

	// (Optional) For replies that quote part of the original message,
	// the quoted part of the message.
	Quote *TextQuote `json:"quote"`

	// Shows through which bot the message was sent.
	Via *User `json:"via_bot"`

	// For replies to a story, the original story.
	ReplyToStory *Story `json:"reply_to_story"`

	// (Optional) Time of last edit in Unix.
	LastEdit int64 `json:"edit_date"`

	// (Optional) True, if the message is sent to a forum topic.
	TopicMessage bool `json:"is_topic_message"`

	// (Optional) Message can't be forwarded.
	Protected bool `json:"has_protected_content,omitempty"`

	// (Optional) True, if the message was sent by an implicit action,
	// for example, as an away or a greeting business message, or as a scheduled message
	FromOffline bool `json:"is_from_offline,omitempty"`

	// AlbumID is the unique identifier of a media message group
	// this message belongs to.
	AlbumID string `json:"media_group_id"`

	// Author signature (in channels).
	Signature string `json:"author_signature"`

	// For a text message, the actual UTF-8 text of the message.
	Text string `json:"text"`

	// For registered commands, will contain the string payload:
	//
	// Ex: `/command <payload>` or `/command@botname <payload>`
	Payload string `json:"-"`

	// For a text message, special entities like usernames, URLs,
	// bot commands, etc. that appear in the text.
	Entities []MessageEntity `json:"entities,omitempty"`

	// See ParseMode and consts.
	ParseMode ParseMode `json:"parse_mode,omitempty"`

	// Media (photo, video, music, etc.) - see
	// https://core.telegram.org/bots/api#message
	Photo           *media.PhotoSize `json:"photo,omitempty"`
	Caption         string           `json:"caption,omitempty"`
	CaptionEntities []MessageEntity  `json:"caption_entities,omitempty"`

	// Entities in caption
	HasMediaSpoiler bool `json:"has_media_spoiler,omitempty"`

	// Contact is a shared contact
	Contact *Contact `json:"contact,omitempty"`

	// Location is a shared location
	Location *Location `json:"location,omitempty"`

	// Venue is a shared venue
	Venue *Venue `json:"venue,omitempty"`

	// Poll is a native poll
	Poll *Poll `json:"poll,omitempty"`

	// Dice is a dice with animated status
	Dice *Dice `json:"dice,omitempty"`

	// Game is a game message
	Game *Game `json:"game,omitempty"`

	// For a service message, represents a user that just got added to chat.
	// Might be the Bot itself.
	UserJoined *User `json:"new_chat_member,omitempty"`

	// For a service message, represents a user that just left chat.
	// Might be the Bot itself.
	UserLeft *User `json:"left_chat_member,omitempty"`

	// For a service message, new members that were added to the group or supergroup.
	UsersJoined []User `json:"new_chat_members,omitempty"`

	// NewChatTitle is new chat title
	NewChatTitle string `json:"new_chat_title,omitempty"`

	// For a service message, represents all available thumbnails of the new chat photo.
	NewChatPhoto []media.PhotoSize `json:"new_chat_photo,omitempty"`

	// For a service message, true if chat photo just got removed.
	GroupPhotoDeleted bool `json:"delete_chat_photo,omitempty"`

	// For a service message, true if group has been created.
	GroupCreated bool `json:"group_chat_created,omitempty"`

	// For a service message, true if supergroup has been created.
	SuperGroupCreated bool `json:"supergroup_chat_created,omitempty"`

	// For a service message, true if channel has been created.
	ChannelCreated bool `json:"channel_chat_created,omitempty"`

	// For a service message, the destination (supergroup) you migrated to.
	MigrateTo int64 `json:"migrate_to_chat_id,omitempty"`

	// For a service message, the Origin (normal group) you migrated from.
	MigrateFrom int64 `json:"migrate_from_chat_id,omitempty"`

	// PinnedMessage is a pinned message
	PinnedMessage *Message `json:"pinned_message,omitempty"`

	// Invoice is an invoice for a payment
	Invoice *Invoice `json:"invoice,omitempty"`

	// Payment is a service message about a successful payment
	Payment *Payment `json:"successful_payment,omitempty"`

	// RefundedPayment is a service message about a refunded payment
	RefundedPayment *RefundedPayment `json:"refunded_payment,omitempty"`

	// UserShared is a user shared
	UserShared *RecipientShared `json:"users_shared,omitempty"`

	// ChatShared is a chat shared
	ChatShared *RecipientShared `json:"chat_shared,omitempty"`

	// ConnectedWebsite is a connected website
	ConnectedWebsite string `json:"connected_website,omitempty"`

	// WriteAccessAllowed is write access allowed
	WriteAccessAllowed *WriteAccessAllowed `json:"write_access_allowed,omitempty"`

	// PassportData is passport data
	PassportData *PassportData `json:"passport_data,omitempty"`

	// ProximityAlert is proximity alert triggered
	ProximityAlert *ProximityAlert `json:"proximity_alert_triggered,omitempty"`

	// AutoDeleteTimer represents about a change in auto-delete timer settings
	AutoDeleteTimer *AutoDeleteTimer `json:"message_auto_delete_timer_changed,omitempty"`

	// ThreadCreated is forum topic created
	ThreadCreated *Thread `json:"forum_topic_created,omitempty"`

	// ThreadEdited is forum topic edited
	ThreadEdited *Thread `json:"forum_topic_edited,omitempty"`

	// ThreadClosed is forum topic closed
	ThreadClosed *struct{} `json:"forum_topic_closed,omitempty"`

	// ThreadReopened is forum topic reopened
	ThreadReopened *Thread `json:"forum_topic_reopened,omitempty"`

	// GeneralThreadHidden is general forum topic hidden
	GeneralThreadHidden *struct{} `json:"general_topic_hidden,omitempty"`

	// GeneralThreadUnhidden is general forum topic unhidden
	GeneralThreadUnhidden *struct{} `json:"general_topic_unhidden,omitempty"`

	// VideoChatScheduled is video chat scheduled
	VideoChatScheduled *VideoChatScheduled `json:"video_chat_scheduled,omitempty"`

	// VideoChatStarted is video chat started
	VideoChatStarted *VideoChatStarted `json:"video_chat_started,omitempty"`

	// VideoChatEnded is video chat ended
	VideoChatEnded *VideoChatEnded `json:"video_chat_ended,omitempty"`

	// VideoChatParticipants is video chat participants invited
	VideoChatParticipants *VideoChatParticipants `json:"video_chat_participants_invited,omitempty"`

	// WebAppData is web app data
	WebAppData *WebAppData `json:"web_app_data,omitempty"`

	// VideoNote is a video note
	VideoNote *media.VideoNote `json:"video_note,omitempty"`

	// Voice is a voice message
	Voice *media.Voice `json:"voice,omitempty"`

	// Audio is an audio file
	Audio *media.Audio `json:"audio,omitempty"`

	// Document is a general file
	Document *media.Document `json:"document,omitempty"`

	// Animation is an animation file
	Animation *media.Animation `json:"animation,omitempty"`

	// Sticker is a sticker
	Sticker *media.Sticker `json:"sticker,omitempty"`

	// Video is a video file
	Video *media.Video `json:"video,omitempty"`

	// BoostAdded is user boosted the chat
	BoostAdded *BoostAdded `json:"boost_added,omitempty"`

	// ChatBackground is chat background set
	ChatBackground ChatBackground `json:"chat_background_set,omitempty"`

	// SenderBoosts is the number of boosts added by the user
	SenderBoosts int `json:"sender_boost_count,omitempty"`

	// MaybeUnavailable is true if the message is unavailable
	MaybeUnavailable bool `json:"maybe_unavailable,omitempty"`
}

// MessageSig satisfies Editable interface (see Editable.)
func (m *Message) MessageSig() (string, int64) {
	return strconv.Itoa(m.ID), m.Chat.ID
}

// Inaccessible shows whether the message is InaccessibleMessage object.
func (m *Message) Inaccessible() bool {
	return m.Sender == nil
}

// Time returns the moment of message creation in local time.
func (m *Message) Time() time.Time {
	return time.Unix(m.Unixtime, 0)
}

// LastEdited returns time.Time of last edit.
func (m *Message) LastEdited() time.Time {
	return time.Unix(m.LastEdit, 0)
}

// IsForwarded says whether message is forwarded copy of another message or not.
func (m *Message) IsForwarded() bool {
	return m.OriginalSender != nil || m.OriginalChat != nil
}

// IsReply says whether message is a reply to another message.
func (m *Message) IsReply() bool {
	return m.ReplyTo != nil
}

// IsService returns true, if message is a service message, returns false otherwise.
//
// Service messages are automatically sent messages, which typically occur on some global action.
// For instance, when anyone leaves the chat or chat title changes.
func (m *Message) IsService() bool {
	return m.UserJoined != nil ||
		len(m.UsersJoined) > 0 ||
		m.UserLeft != nil ||
		m.NewChatTitle != "" ||
		len(m.NewChatPhoto) > 0 ||
		m.GroupPhotoDeleted ||
		m.GroupCreated ||
		m.SuperGroupCreated ||
		m.MigrateTo != m.MigrateFrom
}

// EntityText returns the substring of the message identified by the given MessageEntity.
//
// It's safer than manually slicing Text because Telegram uses UTF-16 indices whereas Go string are []byte.
func (m *Message) EntityText(e MessageEntity) string {
	text := m.Text
	if text == "" {
		text = m.Caption
	}

	a := utf16.Encode([]rune(text))
	off, end := e.Offset, e.Offset+e.Length

	if off < 0 || end > len(a) {
		return ""
	}

	return string(utf16.Decode(a[off:end]))
}

// Placeholder types that will be defined later
type PassportData struct{}
type ExternalReply struct{}
type TextQuote struct{}
type MessageOrigin struct{}
