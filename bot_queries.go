package tb

import (
	"context"

	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/telegram"
)


// RespondToCallback sends a response for a given callback query. A callback can
// only be responded to once, subsequent attempts to respond to the same callback
// will result in an error.
//
// Example:
//
//	b.RespondToCallback(c)
//	b.RespondToCallback(c, response)
func (b *Bot[RequestType]) RespondToCallback(c *telegram.Callback, resp ...*telegram.CallbackResponse) error {
	req := telegram.AnswerCallbackQueryRequest{
		CallbackQueryID: c.ID,
	}

	if len(resp) > 0 && resp[0] != nil {
		req.Text = resp[0].Text
		req.ShowAlert = resp[0].ShowAlert
		req.URL = resp[0].URL
	}

	requester := NewApiRequester[telegram.AnswerCallbackQueryRequest, bool](b.token, b.apiURL, b.client)
	_, err := requester.Request(context.Background(), "answerCallbackQuery", req)
	return err
}

// RespondToCallbackBackground sends a response for a given callback query using background context.
func (b *Bot[RequestType]) RespondToCallbackBackground(c *telegram.Callback, resp ...*telegram.CallbackResponse) error {
	return b.RespondToCallback(c, resp...)
}

// Ship replies to the shipping query, if you sent an invoice
// requesting an address and the parameter is_flexible was specified.
//
// Example:
//
//	b.Ship(query)          // OK
//	b.Ship(query, opts...) // OK with options
//	b.Ship(query, "Oops!") // Error message
func (b *Bot[RequestType]) Ship(query *telegram.ShippingQuery, what ...interface{}) error {
	req := telegram.AnswerShippingQueryRequest{
		ShippingQueryID: query.ID,
	}

	if len(what) == 0 {
		req.Ok = true
	} else if s, ok := what[0].(string); ok {
		req.Ok = false
		req.ErrorMessage = s
	} else {
		var opts []telegram.ShippingOption
		for _, v := range what {
			opt, ok := v.(telegram.ShippingOption)
			if !ok {
				return errors.WithInvalidParam(errors.ErrUnsupportedWhat, "what", v)
			}
			opts = append(opts, opt)
		}

		req.Ok = true
		req.ShippingOptions = opts
	}

	requester := NewApiRequester[telegram.AnswerShippingQueryRequest, bool](b.token, b.apiURL, b.client)
	_, err := requester.Request(context.Background(), "answerShippingQuery", req)
	return err
}

// ShipBackground replies to the shipping query using background context.
func (b *Bot[RequestType]) ShipBackground(query *telegram.ShippingQuery, what ...interface{}) error {
	return b.Ship(query, what...)
}

// Accept finalizes the deal.
func (b *Bot[RequestType]) Accept(query *telegram.PreCheckoutQuery, errorMessage ...string) error {
	req := telegram.AnswerPreCheckoutQueryRequest{
		PreCheckoutQueryID: query.ID,
		Ok:                 len(errorMessage) == 0,
	}

	if len(errorMessage) > 0 {
		req.ErrorMessage = errorMessage[0]
	}

	requester := NewApiRequester[telegram.AnswerPreCheckoutQueryRequest, bool](b.token, b.apiURL, b.client)
	_, err := requester.Request(context.Background(), "answerPreCheckoutQuery", req)
	return err
}

// AcceptBackground finalizes the deal using background context.
func (b *Bot[RequestType]) AcceptBackground(query *telegram.PreCheckoutQuery, errorMessage ...string) error {
	return b.Accept(query, errorMessage...)
}
