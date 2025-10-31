package tb

import (
	"context"
	"strconv"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/telegram"
	"github.com/alagunto/tb/telegram/methods"
)

// ChatByID fetches chat info of its ID.
//
// Including current name of the user for one-on-one conversations,
// current username of a user, group or channel, etc.
func (b *Bot[RequestType]) ChatByID(id int64) (*telegram.Chat, error) {
	return b.ChatByUsername(strconv.FormatInt(id, 10))
}

// ChatByUsername fetches chat info by its username.
func (b *Bot[RequestType]) ChatByUsername(name string) (*telegram.Chat, error) {
	req := methods.GetChatRequest{
		ChatID: name,
	}

	r := NewApiRequester[methods.GetChatRequest, methods.GetChatResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getChat", req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// ProfilePhotosOf returns list of profile pictures for a user.
func (b *Bot[RequestType]) ProfilePhotosOf(user *telegram.User) ([]telegram.Photo, error) {
	req := methods.GetUserProfilePhotosRequest{
		UserID: user.Recipient(),
	}

	r := NewApiRequester[methods.GetUserProfilePhotosRequest, methods.GetUserProfilePhotosResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getUserProfilePhotos", req)
	if err != nil {
		return nil, err
	}
	return result.Photos, nil
}

// ChatMemberOf returns information about a member of a chat.
func (b *Bot[RequestType]) ChatMemberOf(chat, user bot.Recipient) (*telegram.ChatMember, error) {
	req := methods.GetChatMemberRequest{
		ChatID: chat.Recipient(),
		UserID: user.Recipient(),
	}

	r := NewApiRequester[methods.GetChatMemberRequest, methods.GetChatMemberResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getChatMember", req)
	if err != nil {
		return nil, err
	}
	return result, nil
}

// GetMe returns the bot's information.
func (b *Bot[RequestType]) GetMe() (*telegram.User, error) {
	return b.getMe()
}

// StarTransactions returns the bot's star transactions.
func (b *Bot[RequestType]) StarTransactions(offset, limit int) ([]telegram.StarTransaction, error) {
	req := methods.GetStarTransactionsRequest{
		Offset: offset,
		Limit:  limit,
	}

	type starTransactionsResponse struct {
		Transactions []telegram.StarTransaction `json:"transactions"`
	}

	r := NewApiRequester[methods.GetStarTransactionsRequest, starTransactionsResponse](b.token, b.apiURL, b.client)
	result, err := r.Request(context.Background(), "getStarTransactions", req)
	if err != nil {
		return nil, err
	}
	return result.Transactions, nil
}
