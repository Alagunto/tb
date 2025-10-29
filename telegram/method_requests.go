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

// SetsParseMode is an interface for method request structs that support parse mode.
type SetsParseMode interface {
	SetParseMode(mode ParseMode)
}

// SetsEntities is an interface for method request structs that support entities.
type SetsEntities interface {
	SetEntities(entities Entities)
}
