package telegram

import (
	"github.com/alagunto/tb/outgoing"
	"github.com/alagunto/tb/params"
)

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

// ToTelegramSendMethod implements the outgoing.Content interface.
func (l *Location) ToTelegramSendMethod() *outgoing.Method {
	b := params.New()
	b.AddFloat("latitude", l.Lat)
	b.AddFloat("longitude", l.Lng)

	if l.HorizontalAccuracy != nil {
		b.AddFloat("horizontal_accuracy", *l.HorizontalAccuracy)
	}

	b.AddInt("heading", l.Heading)
	b.AddInt("proximity_alert_radius", l.AlertRadius)
	b.AddInt("live_period", l.LivePeriod)
	b.Add("business_connection_id", l.BusinessConnectionID)

	return &outgoing.Method{
		Name:   "sendLocation",
		Params: b.Build(),
		Files:  nil,
	}
}
