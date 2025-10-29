package tb

import (
	"context"

	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/telegram"
	"github.com/alagunto/tb/telegram/methods"
)

// AnswerInlineQuery sends a response to an inline query.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) AnswerInlineQuery(query *telegram.InlineQuery, resp *telegram.QueryResponse) error {
	if query == nil {
		return errors.WithInvalidParam(errors.ErrTelebot, "query", nil)
	}

	req := methods.AnswerInlineQueryRequest{
		InlineQueryID: query.ID,
	}

	if resp != nil {
		req.Results = resp.Results
		req.CacheTime = resp.CacheTime
		req.IsPersonal = resp.IsPersonal
		req.NextOffset = resp.NextOffset
		req.Button = resp.Button
	}

	r := NewApiRequester[methods.AnswerInlineQueryRequest, methods.AnswerInlineQueryResponse](b.token, b.apiURL, b.client)
	_, err := r.Request(context.Background(), "answerInlineQuery", req)
	return err
}

// RespondToCallback sends a response for a given callback query. A callback can
// only be responded to once, subsequent attempts to respond to the same callback
// will result in an error.
//
// Example:
//
//	b.RespondToCallback(c)
//	b.RespondToCallback(c, response)
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) RespondToCallback(c *telegram.CallbackQuery, resp ...*telegram.CallbackResponse) error {
	req := methods.AnswerCallbackQueryRequest{
		CallbackQueryID: c.ID,
	}

	if len(resp) > 0 && resp[0] != nil {
		req.Text = resp[0].Text
		req.ShowAlert = resp[0].ShowAlert
		req.URL = resp[0].URL
	}

	requester := NewApiRequester[methods.AnswerCallbackQueryRequest, methods.AnswerCallbackQueryResponse](b.token, b.apiURL, b.client)
	_, err := requester.Request(context.Background(), "answerCallbackQuery", req)
	return err
}

// Answer sends a response for a given inline query. A query can only
// be responded to once, subsequent attempts to respond to the same query
// will result in an error.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Answer(query *telegram.InlineQuery, resp *telegram.QueryResponse) error {
	return b.AnswerInlineQuery(query, resp)
}

// AnswerWebApp sends a response for a query from Web App and returns
// information about an inline message sent by a Web App on behalf of a user
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) AnswerWebApp(query *telegram.InlineQuery, r telegram.Result) (*telegram.WebAppMessage, error) {
	req := methods.AnswerWebAppQueryRequest{
		WebAppQueryID: query.ID,
		Result:        r,
	}

	requester := NewApiRequester[methods.AnswerWebAppQueryRequest, methods.AnswerWebAppQueryResponse](b.token, b.apiURL, b.client)
	result, err := requester.Request(context.Background(), "answerWebAppQuery", req)
	if err != nil {
		return nil, err
	}

	return result, nil
}

// Ship replies to the shipping query, if you sent an invoice
// requesting an address and the parameter is_flexible was specified.
//
// Example:
//
//	b.Ship(query)          // OK
//	b.Ship(query, opts...) // OK with options
//	b.Ship(query, "Oops!") // Error message
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Ship(query *telegram.ShippingQuery, what ...interface{}) error {
	req := methods.AnswerShippingQueryRequest{
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

	requester := NewApiRequester[methods.AnswerShippingQueryRequest, methods.AnswerShippingQueryResponse](b.token, b.apiURL, b.client)
	_, err := requester.Request(context.Background(), "answerShippingQuery", req)
	return err
}

// Accept finalizes the deal.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) Accept(query *telegram.PreCheckoutQuery, errorMessage ...string) error {
	req := methods.AnswerPreCheckoutQueryRequest{
		PreCheckoutQueryID: query.ID,
		Ok:                 len(errorMessage) == 0,
	}

	if len(errorMessage) > 0 {
		req.ErrorMessage = errorMessage[0]
	}

	requester := NewApiRequester[methods.AnswerPreCheckoutQueryRequest, methods.AnswerPreCheckoutQueryResponse](b.token, b.apiURL, b.client)
	_, err := requester.Request(context.Background(), "answerPreCheckoutQuery", req)
	return err
}
