package telegram

import "encoding/json"

// PollType defines poll types.
type PollType string

const (
	// NOTE:
	// Despite "any" type isn't described in documentation,
	// it needed for proper KeyboardButtonPollType marshaling.
	PollAny PollType = "any"

	PollQuiz    PollType = "quiz"
	PollRegular PollType = "regular"
)

// MarshalJSON implements json.Marshaler. It allows passing PollType as a
// keyboard's poll type instead of KeyboardButtonPollType object.
func (pt PollType) MarshalJSON() ([]byte, error) {
	return json.Marshal(&struct {
		Type string `json:"type"`
	}{
		Type: string(pt),
	})
}
