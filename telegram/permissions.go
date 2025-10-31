package telegram

// ChatPermissions describes actions that a non-administrator user is allowed to take in a chat.
/*
	Source: https://core.telegram.org/bots/api#chatpermissions
	Fields:
	can_send_messages	Boolean	Optional. True, if the user is allowed to send text messages, contacts, giveaways, giveaway winners, invoices, locations and venues
	can_send_audios	Boolean	Optional. True, if the user is allowed to send audios
	can_send_documents	Boolean	Optional. True, if the user is allowed to send documents
	can_send_photos	Boolean	Optional. True, if the user is allowed to send photos
	can_send_videos	Boolean	Optional. True, if the user is allowed to send videos
	can_send_video_notes	Boolean	Optional. True, if the user is allowed to send video notes
	can_send_voice_notes	Boolean	Optional. True, if the user is allowed to send voice notes
	can_send_polls	Boolean	Optional. True, if the user is allowed to send polls and checklists
	can_send_other_messages	Boolean	Optional. True, if the user is allowed to send animations, games, stickers and use inline bots
	can_add_web_page_previews	Boolean	Optional. True, if the user is allowed to add web page previews to their messages
	can_change_info	Boolean	Optional. True, if the user is allowed to change the chat title, photo and other settings. Ignored in public supergroups
	can_invite_users	Boolean	Optional. True, if the user is allowed to invite new users to the chat
	can_pin_messages	Boolean	Optional. True, if the user is allowed to pin messages. Ignored in public supergroups
	can_manage_topics	Boolean	Optional. True, if the user is allowed to create forum topics. If omitted defaults to the value of can_pin_messages
*/
type ChatPermissions struct {
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
	CanManageTopics       bool `json:"can_manage_topics"`
}

// Rights represents a combined set of chat permissions and administrator rights.
// It can be used with restrictChatMember, setMyDefaultAdministratorRights and promoteChatMember.
// https://core.telegram.org/bots/api#chatpermissions
// https://core.telegram.org/bots/api#promotechatmember
type Rights struct {
	CanSendMessages         bool `json:"can_send_messages,omitempty"`
	CanSendMedia            bool `json:"can_send_media_messages,omitempty"`
	CanSendAudios           bool `json:"can_send_audios,omitempty"`
	CanSendDocuments        bool `json:"can_send_documents,omitempty"`
	CanSendPhotos           bool `json:"can_send_photos,omitempty"`
	CanSendVideos           bool `json:"can_send_videos,omitempty"`
	CanSendVideoNotes       bool `json:"can_send_video_notes,omitempty"`
	CanSendVoiceNotes       bool `json:"can_send_voice_notes,omitempty"`
	CanSendPolls            bool `json:"can_send_polls,omitempty"`
	CanSendOther            bool `json:"can_send_other_messages,omitempty"`
	CanAddPreviews          bool `json:"can_add_web_page_previews,omitempty"`
	CanChangeInfo           bool `json:"can_change_info,omitempty"`
	CanInviteUsers          bool `json:"can_invite_users,omitempty"`
	CanPinMessages          bool `json:"can_pin_messages,omitempty"`
	CanManageTopics         bool `json:"can_manage_topics,omitempty"`
	IsAnonymous             bool `json:"is_anonymous,omitempty"`
	CanManageChat           bool `json:"can_manage_chat,omitempty"`
	CanDeleteMessages       bool `json:"can_delete_messages,omitempty"`
	CanManageVideoChats     bool `json:"can_manage_video_chats,omitempty"`
	CanRestrictMembers      bool `json:"can_restrict_members,omitempty"`
	CanPromoteMembers       bool `json:"can_promote_members,omitempty"`
	CanPostMessages         bool `json:"can_post_messages,omitempty"`
	CanEditMessages         bool `json:"can_edit_messages,omitempty"`
	CanPostStories          bool `json:"can_post_stories,omitempty"`
	CanEditStories          bool `json:"can_edit_stories,omitempty"`
	CanDeleteStories        bool `json:"can_delete_stories,omitempty"`
	CanManageDirectMessages bool `json:"can_manage_direct_messages,omitempty"`
}
