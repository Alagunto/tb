package telegram

// ReplyButton represents a button displayed in reply-keyboard.
//
// Set either Contact or Location to true in order to request
// sensitive info, such as user's phone number or current location.
type ReplyButton struct {
	Text string `json:"text"`

	Contact      bool                      `json:"request_contact,omitempty"`
	Location     bool                      `json:"request_location,omitempty"`
	Poll         *KeyboardButtonPollType   `json:"request_poll,omitempty"`
	User         *KeyboardButtonRequestUser `json:"request_users,omitempty"`
	Chat         *KeyboardButtonRequestChat `json:"request_chat,omitempty"`
	WebApp       *WebAppInfo               `json:"web_app,omitempty"`
}
