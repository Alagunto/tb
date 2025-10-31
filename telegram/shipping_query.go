package telegram

// ShippingQuery contains information about an incoming shipping query.
type ShippingQuery struct {
	ID              string          `json:"id"`
	From            *User           `json:"from"`
	InvoicePayload  string          `json:"invoice_payload"`
	ShippingAddress ShippingAddress `json:"shipping_address"`
}
