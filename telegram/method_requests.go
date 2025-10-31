package telegram

// HasReplyMarkup is a shared struct that can be embedded into method request structs
// that support reply markup functionality.
type HasReplyMarkup struct {
	ReplyMarkup *ReplyMarkup `json:"reply_markup,omitempty"`
}

// SetReplyMarkup sets the reply markup for the request.
func (h *HasReplyMarkup) SetReplyMarkup(markup *ReplyMarkup) {
	h.ReplyMarkup = markup
}

// SetsReplyMarkup is an interface for method request structs that support reply markup.
type SetsReplyMarkup interface {
	SetReplyMarkup(markup *ReplyMarkup)
}

// HasBusinessConnection is a shared struct that can be embedded into method request structs
// that support business connection functionality.
type HasBusinessConnection struct {
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
}

// SetBusinessConnectionID sets the business connection ID for the request.
func (h *HasBusinessConnection) SetBusinessConnectionID(id string) {
	h.BusinessConnectionID = id
}

// SetsBusinessConnection is an interface for method request structs that support business connection.
type SetsBusinessConnection interface {
	SetBusinessConnectionID(id string)
}

// HasMessageEffect is a shared struct that can be embedded into method request structs
// that support message effect functionality.
type HasMessageEffect struct {
	MessageEffectID string `json:"message_effect_id,omitempty"`
}

// SetMessageEffectID sets the message effect ID for the request.
func (h *HasMessageEffect) SetMessageEffectID(id string) {
	h.MessageEffectID = id
}

// SetsMessageEffect is an interface for method request structs that support message effects.
type SetsMessageEffect interface {
	SetMessageEffectID(id string)
}

// SetsParseMode is an interface for method request structs that support parse mode.
type SetsParseMode interface {
	SetParseMode(mode ParseMode)
}

// SetsEntities is an interface for method request structs that support entities.
type SetsEntities interface {
	SetEntities(entities Entities)
}

// ============================================================================
// Message Methods
// ============================================================================

// SendMessageRequest represents a request to send a text message.
type SendMessageRequest struct {
	HasReplyMarkup
	HasMessageEffect
	ChatID    string   `json:"chat_id"`
	Text      string   `json:"text"`
	ParseMode string   `json:"parse_mode,omitempty"`
	Entities  Entities `json:"entities,omitempty"`
}

// ForwardMessageRequest represents a request to forward a message.
type ForwardMessageRequest struct {
	ChatID              string `json:"chat_id"`
	FromChatID          int64  `json:"from_chat_id"`
	MessageID           string `json:"message_id"`
	MessageThreadID     int    `json:"message_thread_id,omitempty"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
	ProtectContent      bool   `json:"protect_content,omitempty"`
}

// ForwardMessagesRequest represents a request to forward multiple messages.
type ForwardMessagesRequest struct {
	ChatID              string   `json:"chat_id"`
	FromChatID          int64    `json:"from_chat_id"`
	MessageIDs          []string `json:"message_ids"`
	MessageThreadID     int      `json:"message_thread_id,omitempty"`
	DisableNotification bool     `json:"disable_notification,omitempty"`
	ProtectContent      bool     `json:"protect_content,omitempty"`
}

// CopyMessageRequest represents a request to copy a message.
type CopyMessageRequest struct {
	ChatID              string            `json:"chat_id"`
	FromChatID          int64             `json:"from_chat_id"`
	MessageID           string            `json:"message_id"`
	MessageThreadID     int               `json:"message_thread_id,omitempty"`
	DisableNotification bool              `json:"disable_notification,omitempty"`
	ProtectContent      bool              `json:"protect_content,omitempty"`
	ReplyParameters     *ReplyParameters  `json:"reply_parameters,omitempty"`
	ReplyMarkup         *ReplyMarkup      `json:"reply_markup,omitempty"`
}

// CopyMessagesRequest represents a request to copy multiple messages.
type CopyMessagesRequest struct {
	ChatID              string   `json:"chat_id"`
	FromChatID          int64    `json:"from_chat_id"`
	MessageIDs          []string `json:"message_ids"`
	MessageThreadID     int      `json:"message_thread_id,omitempty"`
	DisableNotification bool     `json:"disable_notification,omitempty"`
	ProtectContent      bool     `json:"protect_content,omitempty"`
}

// EditMessageTextRequest represents a request to edit message text.
type EditMessageTextRequest struct {
	ChatID          string `json:"chat_id,omitempty"`
	MessageID       string `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	Text            string `json:"text"`
}

// EditMessageChecklistRequest represents a request to edit message checklist.
type EditMessageChecklistRequest struct {
	ChatID          string    `json:"chat_id,omitempty"`
	MessageID       string    `json:"message_id,omitempty"`
	InlineMessageID string    `json:"inline_message_id,omitempty"`
	Checklist       Checklist `json:"checklist"`
}

// EditMessageLiveLocationRequest represents a request to edit live location.
type EditMessageLiveLocationRequest struct {
	ChatID               string   `json:"chat_id,omitempty"`
	MessageID            string   `json:"message_id,omitempty"`
	InlineMessageID      string   `json:"inline_message_id,omitempty"`
	Latitude             float64  `json:"latitude"`
	Longitude            float64  `json:"longitude"`
	HorizontalAccuracy   *float64 `json:"horizontal_accuracy,omitempty"`
	Heading              int      `json:"heading,omitempty"`
	ProximityAlertRadius int      `json:"proximity_alert_radius,omitempty"`
	LivePeriod           int      `json:"live_period,omitempty"`
}

// EditMessageReplyMarkupRequest represents a request to edit message reply markup.
type EditMessageReplyMarkupRequest struct {
	ChatID          string       `json:"chat_id,omitempty"`
	MessageID       string       `json:"message_id,omitempty"`
	InlineMessageID string       `json:"inline_message_id,omitempty"`
	ReplyMarkup     *ReplyMarkup `json:"reply_markup,omitempty"`
}

// EditMessageCaptionRequest represents a request to edit message caption.
type EditMessageCaptionRequest struct {
	ChatID          string          `json:"chat_id,omitempty"`
	MessageID       string          `json:"message_id,omitempty"`
	InlineMessageID string          `json:"inline_message_id,omitempty"`
	Caption         string          `json:"caption"`
	ParseMode       string          `json:"parse_mode,omitempty"`
	CaptionEntities []MessageEntity `json:"caption_entities,omitempty"`
}

// EditMessageMediaRequest represents a request to edit message media.
type EditMessageMediaRequest struct {
	ChatID          string `json:"chat_id,omitempty"`
	MessageID       string `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
	Media           string `json:"media"`
}

// DeleteMessageRequest represents a request to delete a message.
type DeleteMessageRequest struct {
	ChatID    int64  `json:"chat_id"`
	MessageID string `json:"message_id"`
}

// DeleteMessagesRequest represents a request to delete multiple messages.
type DeleteMessagesRequest struct {
	ChatID     int64    `json:"chat_id"`
	MessageIDs []string `json:"message_ids"`
}

// StopMessageLiveLocationRequest represents a request to stop live location updates.
type StopMessageLiveLocationRequest struct {
	ChatID          string `json:"chat_id,omitempty"`
	MessageID       string `json:"message_id,omitempty"`
	InlineMessageID string `json:"inline_message_id,omitempty"`
}

// StopPollRequest represents a request to stop a poll.
type StopPollRequest struct {
	ChatID    int64  `json:"chat_id"`
	MessageID string `json:"message_id"`
}

// ============================================================================
// Admin Methods
// ============================================================================

// PinMessageRequest represents a request to pin a message.
type PinMessageRequest struct {
	ChatID              int64  `json:"chat_id"`
	MessageID           string `json:"message_id"`
	DisableNotification bool   `json:"disable_notification,omitempty"`
}

// UnpinMessageRequest represents a request to unpin a message.
type UnpinMessageRequest struct {
	ChatID    string `json:"chat_id"`
	MessageID int    `json:"message_id,omitempty"`
}

// UnpinAllMessageRequest represents a request to unpin all messages.
type UnpinAllMessageRequest struct {
	ChatID string `json:"chat_id"`
}

// SendChatActionRequest represents a request to send a chat action.
type SendChatActionRequest struct {
	ChatID               string `json:"chat_id"`
	Action               string `json:"action"`
	MessageThreadID      int    `json:"message_thread_id,omitempty"`
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
}

// LeaveChatRequest represents a request to leave a chat.
type LeaveChatRequest struct {
	ChatID string `json:"chat_id"`
}

// GetChatMenuButtonRequest represents a request to get chat menu button.
type GetChatMenuButtonRequest struct {
	ChatID string `json:"chat_id"`
}

// SetChatMenuButtonRequest represents a request to set chat menu button.
type SetChatMenuButtonRequest struct {
	ChatID     string      `json:"chat_id,omitempty"`
	MenuButton *MenuButton `json:"menu_button"`
}

// SetMyNameRequest represents a request to set bot name.
type SetMyNameRequest struct {
	Name         string `json:"name"`
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyNameRequest represents a request to get bot name.
type GetMyNameRequest struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyDescriptionRequest represents a request to set bot description.
type SetMyDescriptionRequest struct {
	Description  string `json:"description"`
	LanguageCode string `json:"language_code,omitempty"`
}

// GetMyDescriptionRequest represents a request to get bot description.
type GetMyDescriptionRequest struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// SetMyShortDescriptionRequest represents a request to set bot short description.
type SetMyShortDescriptionRequest struct {
	ShortDescription string `json:"short_description"`
	LanguageCode     string `json:"language_code,omitempty"`
}

// GetMyShortDescriptionRequest represents a request to get bot short description.
type GetMyShortDescriptionRequest struct {
	LanguageCode string `json:"language_code,omitempty"`
}

// ============================================================================
// Chat Methods
// ============================================================================

// GetChatRequest represents a request to get chat information.
type GetChatRequest struct {
	ChatID string `json:"chat_id"`
}

// GetUserProfilePhotosRequest represents a request to get user profile photos.
type GetUserProfilePhotosRequest struct {
	UserID string `json:"user_id"`
}

// GetChatMemberRequest represents a request to get chat member information.
type GetChatMemberRequest struct {
	ChatID string `json:"chat_id"`
	UserID string `json:"user_id"`
}

// GetStarTransactionsRequest represents a request to get star transactions.
type GetStarTransactionsRequest struct {
	Offset int `json:"offset,omitempty"`
	Limit  int `json:"limit,omitempty"`
}

// ============================================================================
// Query Methods
// ============================================================================

// AnswerInlineQueryRequest represents a request to answer an inline query.
type AnswerInlineQueryRequest struct {
	InlineQueryID string                   `json:"inline_query_id"`
	Results       []InlineQueryResult      `json:"results,omitempty"`
	CacheTime     int                      `json:"cache_time,omitempty"`
	IsPersonal    bool                     `json:"is_personal,omitempty"`
	NextOffset    string                   `json:"next_offset,omitempty"`
	Button        *InlineQueryResultsButton `json:"button,omitempty"`
	SwitchPMText  string                   `json:"switch_pm_text,omitempty"`
	SwitchPMParameter string               `json:"switch_pm_parameter,omitempty"`
}

// AnswerCallbackQueryRequest represents a request to answer a callback query.
type AnswerCallbackQueryRequest struct {
	CallbackQueryID string `json:"callback_query_id"`
	Text            string `json:"text,omitempty"`
	ShowAlert       bool   `json:"show_alert,omitempty"`
	URL             string `json:"url,omitempty"`
}

// AnswerShippingQueryRequest represents a request to answer a shipping query.
type AnswerShippingQueryRequest struct {
	ShippingQueryID string           `json:"shipping_query_id"`
	Ok              bool             `json:"ok"`
	ShippingOptions []ShippingOption `json:"shipping_options,omitempty"`
	ErrorMessage    string           `json:"error_message,omitempty"`
}

// AnswerPreCheckoutQueryRequest represents a request to answer a pre-checkout query.
type AnswerPreCheckoutQueryRequest struct {
	PreCheckoutQueryID string `json:"pre_checkout_query_id"`
	Ok                 bool   `json:"ok"`
	ErrorMessage       string `json:"error_message,omitempty"`
}

// ============================================================================
// File Methods
// ============================================================================

// GetFileRequest represents a request to get file information.
type GetFileRequest struct {
	FileID string `json:"file_id"`
}
