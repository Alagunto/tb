package tb

import (
	"context"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/params"
	"github.com/alagunto/tb/telegram"
)

// SendInvoice sends an invoice.
func (b *Bot[RequestType]) SendInvoice(to bot.Recipient, invoice *telegram.Invoice, opts ...params.SendOptions) (*telegram.Message, error) {
	sendOpts := params.Merge(opts...)
	paramsMap := sendOpts.ToMap()
	paramsMap["chat_id"] = to.Recipient()

	// Copy invoice parameters
	paramsMap["title"] = invoice.Title
	paramsMap["description"] = invoice.Description
	paramsMap["payload"] = invoice.Payload
	paramsMap["provider_token"] = invoice.ProviderToken
	paramsMap["currency"] = invoice.Currency
	paramsMap["prices"] = invoice.Prices

	if invoice.MaxTipAmount > 0 {
		paramsMap["max_tip_amount"] = invoice.MaxTipAmount
	}
	if len(invoice.SuggestedTipAmounts) > 0 {
		paramsMap["suggested_tip_amounts"] = invoice.SuggestedTipAmounts
	}
	if invoice.StartParameter != "" {
		paramsMap["start_parameter"] = invoice.StartParameter
	}
	if invoice.ProviderData != "" {
		paramsMap["provider_data"] = invoice.ProviderData
	}
	if invoice.PhotoURL != "" {
		paramsMap["photo_url"] = invoice.PhotoURL
	}
	if invoice.PhotoSize > 0 {
		paramsMap["photo_size"] = invoice.PhotoSize
	}
	if invoice.PhotoWidth > 0 {
		paramsMap["photo_width"] = invoice.PhotoWidth
	}
	if invoice.PhotoHeight > 0 {
		paramsMap["photo_height"] = invoice.PhotoHeight
	}
	paramsMap["need_name"] = invoice.NeedName
	paramsMap["need_phone_number"] = invoice.NeedPhoneNumber
	paramsMap["need_email"] = invoice.NeedEmail
	paramsMap["need_shipping_address"] = invoice.NeedShippingAddress
	paramsMap["send_phone_number_to_provider"] = invoice.SendPhoneNumberToProvider
	paramsMap["send_email_to_provider"] = invoice.SendEmailToProvider
	paramsMap["is_flexible"] = invoice.IsFlexible

	r := NewApiRequester[map[string]any, telegram.Message](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "sendInvoice", paramsMap)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// AnswerShippingQuery replies to shipping queries.
// On success, True is returned.
func (b *Bot[RequestType]) AnswerShippingQuery(query *telegram.ShippingQuery, options []telegram.ShippingOption, ok bool, errorMessage string) error {
	params := make(map[string]any)
	params["shipping_query_id"] = query.ID
	params["ok"] = ok

	if ok {
		params["shipping_options"] = options
	} else {
		params["error_message"] = errorMessage
	}

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "answerShippingQuery", params)
	return err
}

// AnswerPreCheckoutQuery responds to pre-checkout queries.
// On success, True is returned.
func (b *Bot[RequestType]) AnswerPreCheckoutQuery(query *telegram.PreCheckoutQuery, ok bool, errorMessage string) error {
	params := make(map[string]any)
	params["pre_checkout_query_id"] = query.ID
	params["ok"] = ok

	if !ok {
		params["error_message"] = errorMessage
	}

	r := NewApiRequester[map[string]any, bool](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "answerPreCheckoutQuery", params)
	return err
}
