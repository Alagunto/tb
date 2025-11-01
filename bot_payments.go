package tb

import (
	"context"
	"fmt"
	"regexp"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// validatePaymentAmount ensures amounts are positive and within reasonable limits
func validatePaymentAmount(amount int) error {
	if amount <= 0 {
		return fmt.Errorf("payment amount must be positive, got %d", amount)
	}
	// Telegram uses smallest currency unit (e.g., cents)
	// Max reasonable amount: $1,000,000 = 100,000,000 cents
	if amount > 100000000 {
		return fmt.Errorf("payment amount too large: %d", amount)
	}
	return nil
}

// validateInvoicePayload checks invoice payload format
func validateInvoicePayload(payload string) error {
	if payload == "" {
		return fmt.Errorf("invoice payload cannot be empty")
	}
	if len(payload) > 128 {
		return fmt.Errorf("invoice payload too long: %d bytes (max 128)", len(payload))
	}
	return nil
}

// validateCurrencyCode ensures currency code follows ISO 4217 format
func validateCurrencyCode(currency string) error {
	if currency == "" {
		return fmt.Errorf("currency code cannot be empty")
	}
	// ISO 4217 currency codes are exactly 3 uppercase letters
	matched, err := regexp.MatchString(`^[A-Z]{3}$`, currency)
	if err != nil {
		return fmt.Errorf("currency validation error: %v", err)
	}
	if !matched {
		return fmt.Errorf("invalid currency code format: %s (must be 3 uppercase letters)", currency)
	}
	return nil
}

// validateInvoiceDescription checks description length and content
func validateInvoiceDescription(description string) error {
	if description == "" {
		return fmt.Errorf("invoice description cannot be empty")
	}
	if len(description) > 255 {
		return fmt.Errorf("invoice description too long: %d characters (max 255)", len(description))
	}
	return nil
}

// validatePrices ensures prices array is valid and amounts are reasonable
func validatePrices(prices []telegram.LabeledPrice) error {
	if len(prices) == 0 {
		return fmt.Errorf("prices array cannot be empty")
	}

	totalAmount := 0
	for _, price := range prices {
		if price.Label == "" {
			return fmt.Errorf("price label cannot be empty")
		}
		if len(price.Label) > 32 {
			return fmt.Errorf("price label too long: %d characters (max 32)", len(price.Label))
		}
		if err := validatePaymentAmount(price.Amount); err != nil {
			return fmt.Errorf("invalid amount for price '%s': %v", price.Label, err)
		}
		totalAmount += price.Amount
	}

	// Also validate total amount is reasonable
	if totalAmount > 100000000 {
		return fmt.Errorf("total invoice amount too large: %d (max 100,000,000)", totalAmount)
	}

	return nil
}

// validatePreCheckoutQuery ensures pre-checkout query data is valid
func validatePreCheckoutQuery(query *telegram.PreCheckoutQuery) error {
	if query == nil {
		return fmt.Errorf("pre-checkout query cannot be nil")
	}
	if query.ID == "" {
		return fmt.Errorf("pre-checkout query ID cannot be empty")
	}
	if err := validateCurrencyCode(query.Currency); err != nil {
		return fmt.Errorf("invalid currency in pre-checkout query: %v", err)
	}
	if err := validatePaymentAmount(query.TotalAmount); err != nil {
		return fmt.Errorf("invalid total amount in pre-checkout query: %v", err)
	}
	if err := validateInvoicePayload(query.InvoicePayload); err != nil {
		return fmt.Errorf("invalid invoice payload in pre-checkout query: %v", err)
	}
	return nil
}

// SendInvoice sends an invoice.
func (b *Bot[RequestType]) SendInvoice(ctx context.Context, to bot.Recipient, invoice *telegram.SendInvoiceParams, opts ...params.SendOptions) (*telegram.Message, error) {
	if invoice == nil {
		return nil, fmt.Errorf("invoice parameters cannot be nil")
	}

	// Validate invoice data
	if invoice.Title == "" {
		return nil, fmt.Errorf("invoice title cannot be empty")
	}
	if len(invoice.Title) > 32 {
		return nil, fmt.Errorf("invoice title too long: %d characters (max 32)", len(invoice.Title))
	}

	if err := validateInvoiceDescription(invoice.Description); err != nil {
		return nil, fmt.Errorf("invalid invoice description: %v", err)
	}

	if err := validateInvoicePayload(invoice.Payload); err != nil {
		return nil, fmt.Errorf("invalid invoice payload: %v", err)
	}

	if invoice.ProviderToken == "" {
		return nil, fmt.Errorf("provider token cannot be empty")
	}

	if err := validateCurrencyCode(invoice.Currency); err != nil {
		return nil, fmt.Errorf("invalid currency: %v", err)
	}

	if err := validatePrices(invoice.Prices); err != nil {
		return nil, fmt.Errorf("invalid prices: %v", err)
	}

	// Validate optional fields
	if invoice.MaxTipAmount < 0 {
		return nil, fmt.Errorf("max tip amount cannot be negative: %d", invoice.MaxTipAmount)
	}

	for i, tipAmount := range invoice.SuggestedTipAmounts {
		if tipAmount <= 0 {
			return nil, fmt.Errorf("suggested tip amount at index %d must be positive: %d", i, tipAmount)
		}
		if tipAmount > invoice.MaxTipAmount {
			return nil, fmt.Errorf("suggested tip amount at index %d (%d) exceeds max tip amount (%d)", i, tipAmount, invoice.MaxTipAmount)
		}
	}

	if len(invoice.StartParameter) > 64 {
		return nil, fmt.Errorf("start parameter too long: %d characters (max 64)", len(invoice.StartParameter))
	}

	sendOpts := params.Merge(opts...)

	p := params.New().
		Add("chat_id", to.Recipient()).
		Add("title", invoice.Title).
		Add("description", invoice.Description).
		Add("payload", invoice.Payload).
		Add("provider_token", invoice.ProviderToken).
		Add("currency", invoice.Currency).
		Add("prices", invoice.Prices).
		AddInt("max_tip_amount", invoice.MaxTipAmount).
		Add("suggested_tip_amounts", invoice.SuggestedTipAmounts).
		Add("start_parameter", invoice.StartParameter).
		Add("provider_data", invoice.ProviderData).
		Add("photo_url", invoice.PhotoURL).
		AddInt("photo_size", invoice.PhotoSize).
		AddInt("photo_width", invoice.PhotoWidth).
		AddInt("photo_height", invoice.PhotoHeight).
		AddBool("need_name", invoice.NeedName).
		AddBool("need_phone_number", invoice.NeedPhoneNumber).
		AddBool("need_email", invoice.NeedEmail).
		AddBool("need_shipping_address", invoice.NeedShippingAddress).
		AddBool("send_phone_number_to_provider", invoice.SendPhoneNumberToProvider).
		AddBool("send_email_to_provider", invoice.SendEmailToProvider).
		AddBool("is_flexible", invoice.IsFlexible).
		With(sendOpts).
		Build()

	r := NewApiRequester[map[string]any, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(ctx, "sendInvoice", p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// SendInvoiceBackground sends an invoice using context.Background().
func (b *Bot[RequestType]) SendInvoiceBackground(to bot.Recipient, invoice *telegram.SendInvoiceParams, opts ...params.SendOptions) (*telegram.Message, error) {
	return b.SendInvoice(context.Background(), to, invoice, opts...)
}

// AnswerShippingQuery replies to shipping queries.
// On success, True is returned.
func (b *Bot[RequestType]) AnswerShippingQuery(ctx context.Context, query *telegram.ShippingQuery, options []telegram.ShippingOption, ok bool, errorMessage string) error {
	if query == nil {
		return fmt.Errorf("shipping query cannot be nil")
	}

	p := params.New().
		Add("shipping_query_id", query.ID).
		AddBool("ok", ok)

	if ok {
		p.Add("shipping_options", options)
	} else {
		p.Add("error_message", errorMessage)
	}

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(ctx, "answerShippingQuery", p.Build())
	return err
}

// AnswerShippingQueryBackground replies to shipping queries using context.Background().
// On success, True is returned.
func (b *Bot[RequestType]) AnswerShippingQueryBackground(query *telegram.ShippingQuery, options []telegram.ShippingOption, ok bool, errorMessage string) error {
	return b.AnswerShippingQuery(context.Background(), query, options, ok, errorMessage)
}

// AnswerPreCheckoutQuery responds to pre-checkout queries.
// On success, True is returned.
func (b *Bot[RequestType]) AnswerPreCheckoutQuery(ctx context.Context, query *telegram.PreCheckoutQuery, ok bool, errorMessage string) error {
	if query == nil {
		return fmt.Errorf("pre-checkout query cannot be nil")
	}

	p := params.New().
		Add("pre_checkout_query_id", query.ID).
		AddBool("ok", ok)

	if !ok {
		p.Add("error_message", errorMessage)
	}

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(ctx, "answerPreCheckoutQuery", p.Build())
	return err
}

// AnswerPreCheckoutQueryBackground responds to pre-checkout queries using context.Background().
// On success, True is returned.
func (b *Bot[RequestType]) AnswerPreCheckoutQueryBackground(query *telegram.PreCheckoutQuery, ok bool, errorMessage string) error {
	return b.AnswerPreCheckoutQuery(context.Background(), query, ok, errorMessage)
}
