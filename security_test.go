package tb

import (
	"testing"

	"github.com/alagunto/tb/telegram"
)

func TestValidatePaymentAmount(t *testing.T) {
	tests := []struct {
		name    string
		amount  int
		wantErr bool
	}{
		{"negative amount", -100, true},
		{"zero amount", 0, true},
		{"positive small amount", 100, false},
		{"positive large amount", 1000000, false},
		{"maximum valid amount", 100000000, false},
		{"exceeds maximum", 100000001, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePaymentAmount(tt.amount)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePaymentAmount() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateInvoicePayload(t *testing.T) {
	tests := []struct {
		name    string
		payload string
		wantErr bool
	}{
		{"empty payload", "", true},
		{"valid short payload", "order_123", false},
		{"maximum valid payload", string(make([]byte, 128)), false},
		{"too long payload", string(make([]byte, 129)), true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateInvoicePayload(tt.payload)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateInvoicePayload() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateCurrencyCode(t *testing.T) {
	tests := []struct {
		name    string
		currency string
		wantErr bool
	}{
		{"empty currency", "", true},
		{"valid USD", "USD", false},
		{"valid EUR", "EUR", false},
		{"lowercase", "usd", true},
		{"too short", "US", true},
		{"too long", "USDD", true},
		{"contains numbers", "US1", true},
		{"contains symbols", "U$D", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validateCurrencyCode(tt.currency)
			if (err != nil) != tt.wantErr {
				t.Errorf("validateCurrencyCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePrices(t *testing.T) {
	tests := []struct {
		name    string
		prices  []telegram.LabeledPrice
		wantErr bool
	}{
		{"empty prices", []telegram.LabeledPrice{}, true},
		{"nil prices", nil, true},
		{"valid single price", []telegram.LabeledPrice{{Label: "Item", Amount: 100}}, false},
		{"valid multiple prices", []telegram.LabeledPrice{
			{Label: "Item 1", Amount: 100},
			{Label: "Item 2", Amount: 200},
		}, false},
		{"empty label", []telegram.LabeledPrice{{Label: "", Amount: 100}}, true},
		{"label too long", []telegram.LabeledPrice{{Label: string(make([]byte, 33)), Amount: 100}}, true},
		{"negative amount", []telegram.LabeledPrice{{Label: "Item", Amount: -100}}, true},
		{"zero amount", []telegram.LabeledPrice{{Label: "Item", Amount: 0}}, true},
		{"amount too large", []telegram.LabeledPrice{{Label: "Item", Amount: 100000001}}, true},
		{"total amount too large", []telegram.LabeledPrice{
			{Label: "Item 1", Amount: 50000001},
			{Label: "Item 2", Amount: 50000001},
		}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePrices(tt.prices)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePrices() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidatePreCheckoutQuery(t *testing.T) {
	tests := []struct {
		name  string
		query *telegram.PreCheckoutQuery
		wantErr bool
	}{
		{"nil query", nil, true},
		{"empty ID", &telegram.PreCheckoutQuery{ID: "", Currency: "USD", TotalAmount: 100, InvoicePayload: "test"}, true},
		{"invalid currency", &telegram.PreCheckoutQuery{ID: "123", Currency: "usd", TotalAmount: 100, InvoicePayload: "test"}, true},
		{"negative amount", &telegram.PreCheckoutQuery{ID: "123", Currency: "USD", TotalAmount: -100, InvoicePayload: "test"}, true},
		{"zero amount", &telegram.PreCheckoutQuery{ID: "123", Currency: "USD", TotalAmount: 0, InvoicePayload: "test"}, true},
		{"amount too large", &telegram.PreCheckoutQuery{ID: "123", Currency: "USD", TotalAmount: 100000001, InvoicePayload: "test"}, true},
		{"empty payload", &telegram.PreCheckoutQuery{ID: "123", Currency: "USD", TotalAmount: 100, InvoicePayload: ""}, true},
		{"payload too long", &telegram.PreCheckoutQuery{ID: "123", Currency: "USD", TotalAmount: 100, InvoicePayload: string(make([]byte, 129))}, true},
		{"valid query", &telegram.PreCheckoutQuery{ID: "123", Currency: "USD", TotalAmount: 100, InvoicePayload: "test"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validatePreCheckoutQuery(tt.query)
			if (err != nil) != tt.wantErr {
				t.Errorf("validatePreCheckoutQuery() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}