package tb

import (
	"strings"

	"github.com/alagunto/tb/telegram"
)

// ProcessUpdate processes a single incoming update.
// A started bot calls this function automatically.
func (b *Bot[RequestType]) ProcessUpdate(c RequestType, u telegram.Update) {
	if u.Message != nil {
		m := u.Message

		if m.PinnedMessage != nil {
			b.runHandler(c, OnPinned)
			return
		}

		// Commands
		if m.Text != "" {
			// Filtering malicious messages
			if m.Text[0] == '\a' {
				return
			}

			match := cmdRx.FindAllStringSubmatch(m.Text, -1)
			if match != nil {
				// Syntax: "</command>@<bot> <payload>"
				command, botName := match[0][1], match[0][3]

				if botName != "" && !strings.EqualFold(b.me.Username, botName) {
					return
				}

				m.Payload = match[0][5]
				if b.runHandler(c, command) {
					return
				}
			}

			// 1:1 satisfaction
			if b.runHandler(c, m.Text) {
				return
			}

			b.runHandler(c, OnText)
			return
		}

		// Edits
		if edited := u.EditedMessage; edited != nil {
			b.runHandler(c, OnEdited)
			return
		}

		// Channel posts
		if channelPost := u.ChannelPost; channelPost != nil {
			m := channelPost

			if m.PinnedMessage != nil {
				b.runHandler(c, OnPinned)
				return
			}

			b.runHandler(c, OnChannelPost)
			return
		}

		// Edited channel posts
		if channelPost := u.EditedChannelPost; channelPost != nil {
			b.runHandler(c, OnEditedChannelPost)
			return
		}

		// Inline buttons
		if query := u.CallbackQuery; query != nil {
			if data := query.Data; data != "" && data[0] == '\f' {
				match := cbackRx.FindAllStringSubmatch(data, -1)
				if match != nil {
					unique, payload := match[0][1], match[0][3]
					if _, ok := b.handlers["\f"+unique]; ok {
						query.Data = payload
						b.runHandler(c, "\f"+unique)
						return
					}
				}
			}

			b.runHandler(c, OnCallback)
			return
		}

		// Inline queries
		if query := u.InlineQuery; query != nil {
			b.runHandler(c, OnQuery)
			return
		}

		// Chosen inline results
		if result := u.ChosenInlineResult; result != nil {
			b.runHandler(c, OnInlineResult)
			return
		}

		// Shipping queries
		if query := u.ShippingQuery; query != nil {
			b.runHandler(c, OnShipping)
			return
		}

		// Pre checkout queries
		if query := u.PreCheckoutQuery; query != nil {
			b.runHandler(c, OnCheckout)
			return
		}

		// Polls
		if poll := u.Poll; poll != nil {
			b.runHandler(c, OnPoll)
			return
		}

		// Poll answers
		if answer := u.PollAnswer; answer != nil {
			b.runHandler(c, OnPollAnswer)
			return
		}

		// My chat member updated
		if upd := u.MyChatMember; upd != nil {
			b.runHandler(c, OnMyChatMember)
			return
		}

		// Chat member updated
		if upd := u.ChatMember; upd != nil {
			b.runHandler(c, OnChatMember)
			return
		}

		// Chat join request
		if upd := u.ChatJoinRequest; upd != nil {
			b.runHandler(c, OnChatJoinRequest)
			return
		}

		// Message reactions
		if react := u.MessageReaction; react != nil {
			b.runHandler(c, OnReaction)
			return
		}

		// Message reaction count
		if react := u.MessageReactionCount; react != nil {
			b.runHandler(c, OnReactionCount)
			return
		}

		// Boost updated
		if boost := u.ChatBoost; boost != nil {
			b.runHandler(c, OnBoostUpdated)
			return
		}

		// Boost removed
		if boost := u.RemovedChatBoost; boost != nil {
			b.runHandler(c, OnBoostRemoved)
			return
		}

	// Media
	if len(m.Photo) > 0 {
		b.runHandler(c, OnPhoto)
		return
	}
		if m.Voice != nil {
			b.runHandler(c, OnVoice)
			return
		}
		if m.Audio != nil {
			b.runHandler(c, OnAudio)
			return
		}
		if m.Animation != nil {
			b.runHandler(c, OnAnimation)
			return
		}
		if m.Document != nil {
			b.runHandler(c, OnDocument)
			return
		}
		if m.Sticker != nil {
			b.runHandler(c, OnSticker)
			return
		}
		if m.Video != nil {
			b.runHandler(c, OnVideo)
			return
		}
		if m.VideoNote != nil {
			b.runHandler(c, OnVideoNote)
			return
		}
		if m.Contact != nil {
			b.runHandler(c, OnContact)
			return
		}
		if m.Location != nil {
			b.runHandler(c, OnLocation)
			return
		}
		if m.Venue != nil {
			b.runHandler(c, OnVenue)
			return
		}
		if m.Dice != nil {
			b.runHandler(c, OnDice)
			return
		}
		if m.Invoice != nil {
			b.runHandler(c, OnInvoice)
			return
		}
		if m.Payment != nil {
			b.runHandler(c, OnPayment)
			return
		}
		if m.Game != nil {
			b.runHandler(c, OnGame)
			return
		}
		if m.Poll != nil {
			b.runHandler(c, OnPoll)
			return
		}

		// Topics
		if m.ThreadCreated != nil {
			b.runHandler(c, OnTopicCreated)
			return
		}
		if m.ThreadReopened != nil {
			b.runHandler(c, OnTopicReopened)
			return
		}
		if m.ThreadClosed != nil {
			b.runHandler(c, OnTopicClosed)
			return
		}
		if m.ThreadEdited != nil {
			b.runHandler(c, OnTopicEdited)
			return
		}
		if m.GeneralThreadHidden != nil {
			b.runHandler(c, OnGeneralTopicHidden)
			return
		}
		if m.GeneralThreadUnhidden != nil {
			b.runHandler(c, OnGeneralTopicUnhidden)
			return
		}

		// Service messages
		if len(m.UsersJoined) > 0 {
			for _, user := range m.UsersJoined {
				m.UserJoined = &user
				b.runHandler(c, OnUserJoined)
			}
			return
		}
		if m.UserLeft != nil {
			b.runHandler(c, OnUserLeft)
			return
		}
		if m.NewChatTitle != "" {
			b.runHandler(c, OnNewGroupTitle)
			return
		}
		if m.NewChatPhoto != nil {
			b.runHandler(c, OnNewGroupPhoto)
			return
		}
		if m.GroupPhotoDeleted {
			b.runHandler(c, OnGroupPhotoDeleted)
			return
		}
		if m.GroupCreated {
			b.runHandler(c, OnGroupCreated)
			return
		}
		if m.SuperGroupCreated {
			b.runHandler(c, OnSuperGroupCreated)
			return
		}
		if m.ChannelCreated {
			b.runHandler(c, OnChannelCreated)
			return
		}

		// Migration
		if m.MigrateTo != 0 {
			b.runHandler(c, OnMigration)
			return
		}

		if m.VideoChatStarted != nil {
			b.runHandler(c, OnVideoChatStarted)
			return
		}
		if m.VideoChatEnded != nil {
			b.runHandler(c, OnVideoChatEnded)
			return
		}
		if m.VideoChatScheduled != nil {
			b.runHandler(c, OnVideoChatScheduled)
			return
		}
		if m.VideoChatParticipants != nil {
			b.runHandler(c, OnVideoChatParticipantsInvited)
			return
		}

		if m.WebAppData != nil {
			b.runHandler(c, OnWebApp)
			return
		}

		if m.WriteAccessAllowed != nil {
			b.runHandler(c, OnWriteAccessAllowed)
			return
		}

		if m.ProximityAlert != nil {
			b.runHandler(c, OnProximityAlert)
			return
		}

		if m.AutoDeleteTimer != nil {
			b.runHandler(c, OnAutoDeleteTimer)
			return
		}
	}
}

func (b *Bot[RequestType]) runHandler(c RequestType, endpoint string) bool {
	handler, ok := b.handlers[endpoint]
	if !ok {
		return false
	}

	// Execute handler directly (middleware is handled by Group)
	if err := handler(c); err != nil && b.onError != nil {
		b.onError(err, c)
	}

	return true
}
