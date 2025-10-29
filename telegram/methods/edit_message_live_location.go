package methods

import "github.com/alagunto/tb/telegram"

// EditMessageLiveLocationRequest represents the request for editMessageLiveLocation method
type EditMessageLiveLocationRequest struct {
	telegram.HasReplyMarkup

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

// EditMessageLiveLocationResponse represents the response for editMessageLiveLocation method
type EditMessageLiveLocationResponse = telegram.Message
