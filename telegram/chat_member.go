package telegram

type ChatMemberStatus string

const (
	ChatMemberStatusCreator       ChatMemberStatus = "creator"
	ChatMemberStatusAdministrator ChatMemberStatus = "administrator"
	ChatMemberStatusMember        ChatMemberStatus = "member"
	ChatMemberStatusRestricted    ChatMemberStatus = "restricted"
	ChatMemberStatusLeft          ChatMemberStatus = "left"
	ChatMemberStatusKicked        ChatMemberStatus = "kicked"
)

// ChatMember object contains information about one member of a chat.
/*
Source: https://core.telegram.org/bots/api#chatmember

Currently, the following 6 types of chat members are supported:
- ChatMemberOwner
- ChatMemberAdministrator
- ChatMemberMember
- ChatMemberRestricted
- ChatMemberLeft
- ChatMemberBanned
*/
type ChatMember struct {
	Status ChatMemberStatus `json:"status"`

	User *User `json:"user"`

	// ChatMemberOwner, ChatMemberAdministrator only:
	IsAnonymous bool   `json:"is_anonymous,omitempty"` // is_anonymous	Boolean	True, if the user's presence in the chat is hidden
	CustomTitle string `json:"custom_title,omitempty"` // custom_title	String	Optional. Custom title for this user

	// ChatMemberBanned, ChatMemberMember, ChatMemberRestricted only:
	// If ChatMemberRestricted: date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever
	// If ChatMemberMember: Date when the user's subscription will expire; Unix time; Optional
	// If ChatMemberBanned:	Date when restrictions will be lifted for this user; Unix time. If 0, then the user is banned forever
	UntilDate int64 `json:"until_date"` // until_date	Integer

	// ChatMemberRestricted only:
	IsMember bool `json:"is_member"` // is_member	Boolean	True, if the user is a member of the chat at the moment of the request

	// ChatMemberRestricted only, restrictable rights:
	CanSendMessages       bool `json:"can_send_messages"`
	CanSendAudios         bool `json:"can_send_audios"`
	CanSendDocuments      bool `json:"can_send_documents"`
	CanSendPhotos         bool `json:"can_send_photos"`
	CanSendVideos         bool `json:"can_send_videos"`
	CanSendVideoNotes     bool `json:"can_send_video_notes"`
	CanSendVoiceNotes     bool `json:"can_send_voice_notes"`
	CanSendPolls          bool `json:"can_send_polls"`
	CanSendOtherMessages  bool `json:"can_send_other_messages"`
	CanAddWebPagePreviews bool `json:"can_add_web_page_previews"`
	CanChangeInfo         bool `json:"can_change_info"`
	CanInviteUsers        bool `json:"can_invite_users"`
	CanPinMessages        bool `json:"can_pin_messages"`

	// ChatMemberRestricted and ChatMemberAdministrator only:
	/*
		For administrators:
		can_manage_topics	Boolean	True, if the user is allowed to create, rename, close, and reopen forum topics; for supergroups only
		For restricted users:
		can_manage_topics	Boolean	True, if the user is allowed to create forum topics
	*/
	CanManageTopics bool `json:"can_manage_topics"`

	// ChatMemberAdministrator only:
	CanBeEdited             bool `json:"can_be_edited"`              // can_be_edited	Boolean	True, if the bot is allowed to edit administrator privileges of this user
	CanPostMessages         bool `json:"can_post_messages"`          // can_post_messages	Boolean	True, if the administrator can post messages in the channel; channels only
	CanEditMessages         bool `json:"can_edit_messages"`          // can_edit_messages	Boolean	True, if the administrator can edit messages of other users and can pin messages; channels only
	CanDeleteMessages       bool `json:"can_delete_messages"`        // can_delete_messages	Boolean	True, if the administrator can delete messages of other users
	CanRestrictMembers      bool `json:"can_restrict_members"`       // can_restrict_members	Boolean	True, if the administrator can restrict, ban or unban chat members
	CanPromoteMembers       bool `json:"can_promote_members"`        // can_promote_members	Boolean	True, if the administrator can add new administrators with a subset of their own privileges or demote administrators that they have promoted, directly or indirectly (promoted by administrators that were appointed by the user)
	CanManageVideoChats     bool `json:"can_manage_video_chats"`     // can_manage_video_chats	Boolean	True, if the administrator can manage video chats
	CanManageChat           bool `json:"can_manage_chat"`            // can_manage_chat	Boolean	True, if the administrator can manage the chat; channels only
	CanPostStories          bool `json:"can_post_stories"`           // can_post_stories	Boolean	True, if the administrator can post stories to the chat; channels only
	CanEditStories          bool `json:"can_edit_stories"`           // can_edit_stories	Boolean	True, if the administrator can edit stories posted by other users; channels only
	CanDeleteStories        bool `json:"can_delete_stories"`         // can_delete_stories	Boolean	True, if the administrator can delete stories posted by other users; channels only
	CanManageDirectMessages bool `json:"can_manage_direct_messages"` // can_manage_direct_messages	Boolean	True, if the administrator can manage direct messages of the channel and decline suggested posts; for channels only
}

// GetPermissionsMap returns all permission fields as a map[string]bool.
// This includes both restrictable rights (for restricted members) and administrator rights.
func (cm *ChatMember) GetPermissionsMap() map[string]bool {
	return map[string]bool{
		// Restrictable rights (ChatMemberRestricted)
		"can_send_messages":         cm.CanSendMessages,
		"can_send_audios":           cm.CanSendAudios,
		"can_send_documents":        cm.CanSendDocuments,
		"can_send_photos":           cm.CanSendPhotos,
		"can_send_videos":           cm.CanSendVideos,
		"can_send_video_notes":      cm.CanSendVideoNotes,
		"can_send_voice_notes":      cm.CanSendVoiceNotes,
		"can_send_polls":            cm.CanSendPolls,
		"can_send_other_messages":   cm.CanSendOtherMessages,
		"can_add_web_page_previews": cm.CanAddWebPagePreviews,
		"can_change_info":           cm.CanChangeInfo,
		"can_invite_users":          cm.CanInviteUsers,
		"can_pin_messages":          cm.CanPinMessages,
		"can_manage_topics":         cm.CanManageTopics,

		// Administrator rights (ChatMemberAdministrator)
		"can_be_edited":              cm.CanBeEdited,
		"can_post_messages":          cm.CanPostMessages,
		"can_edit_messages":          cm.CanEditMessages,
		"can_delete_messages":        cm.CanDeleteMessages,
		"can_restrict_members":       cm.CanRestrictMembers,
		"can_promote_members":        cm.CanPromoteMembers,
		"can_manage_video_chats":     cm.CanManageVideoChats,
		"can_manage_chat":            cm.CanManageChat,
		"can_post_stories":           cm.CanPostStories,
		"can_edit_stories":           cm.CanEditStories,
		"can_delete_stories":         cm.CanDeleteStories,
		"can_manage_direct_messages": cm.CanManageDirectMessages,
	}
}

// GetAdminRightsMap returns administrator rights as a map[string]bool.
// This includes only administrator-specific permissions.
func (cm *ChatMember) GetAdminRightsMap() map[string]bool {
	return map[string]bool{
		"is_anonymous":               cm.IsAnonymous,
		"can_manage_chat":            cm.CanManageChat,
		"can_delete_messages":        cm.CanDeleteMessages,
		"can_manage_video_chats":     cm.CanManageVideoChats,
		"can_restrict_members":       cm.CanRestrictMembers,
		"can_promote_members":        cm.CanPromoteMembers,
		"can_change_info":            cm.CanChangeInfo,
		"can_invite_users":           cm.CanInviteUsers,
		"can_post_messages":          cm.CanPostMessages,
		"can_edit_messages":          cm.CanEditMessages,
		"can_pin_messages":           cm.CanPinMessages,
		"can_post_stories":           cm.CanPostStories,
		"can_edit_stories":           cm.CanEditStories,
		"can_delete_stories":         cm.CanDeleteStories,
		"can_manage_topics":          cm.CanManageTopics,
		"can_manage_direct_messages": cm.CanManageDirectMessages,
	}
}
