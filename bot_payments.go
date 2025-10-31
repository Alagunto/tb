package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// SendInvoice sends an invoice.
func (b *Bot[RequestType]) SendInvoice(to bot.Recipient, invoice *telegram.SendInvoiceParams, opts ...params.SendOptions) (*telegram.Message, error) {
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
	result, err := r.Request(context.Background(), "sendInvoice", p)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AnswerShippingQuery replies to shipping queries.
// On success, True is returned.
func (b *Bot[RequestType]) AnswerShippingQuery(query *telegram.ShippingQuery, options []telegram.ShippingOption, ok bool, errorMessage string) error {
	p := params.New().
		Add("shipping_query_id", query.ID).
		AddBool("ok", ok)

	if ok {
		p.Add("shipping_options", options)
	} else {
		p.Add("error_message", errorMessage)
	}

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "answerShippingQuery", p.Build())
	return err
}

// AnswerPreCheckoutQuery responds to pre-checkout queries.
// On success, True is returned.
func (b *Bot[RequestType]) AnswerPreCheckoutQuery(query *telegram.PreCheckoutQuery, ok bool, errorMessage string) error {
	p := params.New().
		Add("pre_checkout_query_id", query.ID).
		AddBool("ok", ok)

	if !ok {
		p.Add("error_message", errorMessage)
	}

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "answerPreCheckoutQuery", p.Build())
	return err
}
