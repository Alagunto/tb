package tb

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/alagunto/tb/bot"
	"github.com/alagunto/tb/errors"
	"github.com/alagunto/tb/telegram"
)

// Forever is a ExpireUnixtime of "forever" banning.
func Forever() int64 {
	return time.Now().Add(367 * 24 * time.Hour).Unix()
}

// Ban will ban user from chat until `member.UntilDate`.
func (b *Bot[RequestType]) Ban(chat *telegram.Chat, member *telegram.ChatMember, revokeMessages ...bool) error {
	params := map[string]string{
		"chat_id":    chat.Recipient(),
		"user_id":    member.User.Recipient(),
		"until_date": strconv.FormatInt(member.UntilDate, 10),
	}
	if len(revokeMessages) > 0 {
		params["revoke_messages"] = strconv.FormatBool(revokeMessages[0])
	}

	_, err := b.Raw( "kickChatMember", params)
	return err
}

// Unban will unban user from chat, who would have thought eh?
// forBanned does nothing if the user is not banned.
func (b *Bot[RequestType]) Unban(chat *telegram.Chat, user *telegram.User, forBanned ...bool) error {
	params := map[string]string{
		"chat_id": chat.Recipient(),
		"user_id": user.Recipient(),
	}

	if len(forBanned) > 0 {
		params["only_if_banned"] = strconv.FormatBool(forBanned[0])
	}

	_, err := b.Raw( "unbanChatMember", params)
	return err
}

// Restrict lets you restrict a subset of member's rights until
// member.UntilDate, such as:
//
//   - can send messages
//   - can send media
//   - can send other
//   - can add web page previews
func (b *Bot[RequestType]) Restrict(chat *telegram.Chat, member *telegram.ChatMember) error {
	params := map[string]interface{}{
		"chat_id":     chat.Recipient(),
		"user_id":     member.User.Recipient(),
		"until_date":  strconv.FormatInt(member.UntilDate, 10),
		"permissions": member.GetPermissionsMap(),
	}

	_, err := b.Raw( "restrictChatMember", params)
	return err
}

// Promote lets you update member's admin rights, such as:
//
//   - can change info
//   - can post messages
//   - can edit messages
//   - can delete messages
//   - can invite users
//   - can restrict members
//   - can pin messages
//   - can promote members
func (b *Bot[RequestType]) Promote(chat *telegram.Chat, member *telegram.ChatMember) error {
	params := map[string]interface{}{
		"chat_id":      chat.Recipient(),
		"user_id":      member.User.Recipient(),
		"is_anonymous": member.IsAnonymous,
	}
	for key, value := range member.GetAdminRightsMap() {
		if key != "is_anonymous" { // Skip is_anonymous as it's already set
			params[key] = value
		}
	}

	_, err := b.Raw( "promoteChatMember", params)
	return err
}

// AdminsOf returns a member list of chat admins.
//
// On success, returns an Array of ChatMember objects that
// contains information about all chat administrators except other bots.
//
// If the chat is a group or a supergroup and
// no administrators were appointed, only the creator will be returned.
func (b *Bot[RequestType]) AdminsOf(chat *telegram.Chat) ([]telegram.ChatMember, error) {
	params := map[string]string{
		"chat_id": chat.Recipient(),
	}

	data, err := b.Raw( "getChatAdministrators", params)
	if err != nil {
		return nil, err
	}

	var resp struct {
		Result []telegram.ChatMember
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return nil, errors.Wrap(err)
	}
	return resp.Result, nil
}

// Len returns the number of members in a chat.
func (b *Bot[RequestType]) Len(chat *telegram.Chat) (int, error) {
	params := map[string]string{
		"chat_id": chat.Recipient(),
	}

	data, err := b.Raw( "getChatMembersCount", params)
	if err != nil {
		return 0, err
	}

	var resp struct {
		Result int
	}
	if err := json.Unmarshal(data, &resp); err != nil {
		return 0, errors.Wrap(err)
	}
	return resp.Result, nil
}

// SetAdminTitle sets a custom title for an administrator.
// A title should be 0-16 characters length, emoji are not allowed.
func (b *Bot[RequestType]) SetAdminTitle(chat *telegram.Chat, user *telegram.User, title string) error {
	params := map[string]string{
		"chat_id":      chat.Recipient(),
		"user_id":      user.Recipient(),
		"custom_title": title,
	}

	_, err := b.Raw( "setChatAdministratorCustomTitle", params)
	return err
}

// BanSenderChat will use this method to ban a channel chat in a supergroup or a channel.
// Until the chat is unbanned, the owner of the banned chat won't be able
// to send messages on behalf of any of their channels.
func (b *Bot[RequestType]) BanSenderChat(chat *telegram.Chat, sender bot.Recipient) error {
	params := map[string]string{
		"chat_id":        chat.Recipient(),
		"sender_chat_id": sender.Recipient(),
	}

	_, err := b.Raw( "banChatSenderChat", params)
	return err
}

// UnbanSenderChat will use this method to unban a previously banned channel chat in a supergroup or channel.
// The bot must be an administrator for this to work and must have the appropriate administrator rights.
func (b *Bot[RequestType]) UnbanSenderChat(chat *telegram.Chat, sender bot.Recipient) error {
	params := map[string]string{
		"chat_id":        chat.Recipient(),
		"sender_chat_id": sender.Recipient(),
	}

	_, err := b.Raw( "unbanChatSenderChat", params)
	return err
}

// // // DefaultRights returns the current default administrator rights of the bot.
// // func (b *Bot[RequestType]) DefaultRights(forChannels bool) (*telegram.Rights, error) {
// // 	params := map[string]bool{
// // 		"for_channels": forChannels,
// // 	}

// // 	data, err := b.Raw( "getMyDefaultAdministratorRights", params)
// // 	if err != nil {
// // 		return nil, err
// // 	}

// // 	var resp struct {
// // 		Result *telegram.
// // 	}
// // 	if err := json.Unmarshal(data, &resp); err != nil {
// // 		return nil, wrapError(err)
// // 	}
// // 	return resp.Result, nil
// // }

// // SetDefaultRights changes the default administrator rights requested by the bot
// // when it's added as an administrator to groups or channels.
// func (b *Bot[RequestType]) SetDefaultRights(rights telegram.Rights, forChannels bool) error {
// 	params := map[string]interface{}{
// 		"rights":       rights,
// 		"for_channels": forChannels,
// 	}

// 	_, err := b.Raw( "setMyDefaultAdministratorRights", params)
// 	return err
// }
