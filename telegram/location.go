package telegram

// Location object represents geographic position.
type Location struct {
	Lat                float32  `json:"latitude"`
	Lng                float32  `json:"longitude"`
	HorizontalAccuracy *float32 `json:"horizontal_accuracy,omitempty"`
	Heading            int      `json:"heading,omitempty"`
	AlertRadius        int      `json:"proximity_alert_radius,omitempty"`

	// Period in seconds for which the location will be updated
	// (see Live Locations, should be between 60 and 86400.)
	LivePeriod int `json:"live_period,omitempty"`

	// (Optional) Unique identifier of the business connection
	// on behalf of which the message to be edited was sent
	BusinessConnectionID string `json:"business_connection_id,omitempty"`
}
