package communications

import "github.com/alagunto/tb/telegram"

// SendOptions has most complete control over in what way the message
// must be sent, providing an API-complete set of custom properties
// and options.
//
// Despite its power, SendOptions is rather inconvenient to use all
// the way through bot logic, so you might want to consider storing
// and re-using it somewhere or be using Option flags instead.
type SendOptions struct {
	// If the message is a reply, original message.
	ReplyTo *telegram.Message

	// See ReplyMarkup struct definition.
	ReplyMarkup *telegram.ReplyMarkup

	// For text messages, disables previews for links in this message.
	DisableWebPagePreview bool

	// Sends the message silently. iOS users will not receive a notification, Android users will receive a notification with no sound.
	DisableNotification bool

	// ParseMode controls how client apps render your message.
	ParseMode telegram.ParseMode

	// Entities is a list of special entities that appear in message text, which can be specified instead of parse_mode.
	Entities telegram.Entities

	// AllowWithoutReply allows sending messages not a as reply if the replied-to message has already been deleted.
	AllowWithoutReply bool

	// Protected protects the contents of sent message from forwarding and saving.
	Protected bool

	// ThreadID supports sending messages to a thread.
	ThreadID int

	// HasSpoiler marks the message as containing a spoiler.
	HasSpoiler bool

	// ReplyParams Describes the message to reply to
	ReplyParams *telegram.ReplyParams

	// Unique identifier of the business connection
	BusinessConnectionID string

	// Unique identifier of the message effect to be added to the message; for private chats only
	EffectID telegram.EffectID
}

// Copy creates a deep copy of SendOptions with all nested structures properly copied.
func (og *SendOptions) Copy() *SendOptions {
	if og == nil {
		return nil
	}

	cp := *og

	// Deep copy ReplyTo message
	cp.ReplyTo = deepCopyMessage(og.ReplyTo)

	// Deep copy ReplyMarkup
	cp.ReplyMarkup = deepCopyReplyMarkup(og.ReplyMarkup)

	// Deep copy Entities slice
	if og.Entities != nil {
		cp.Entities = make(telegram.Entities, len(og.Entities))
		for i, entity := range og.Entities {
			cp.Entities[i] = deepCopyMessageEntity(entity)
		}
	}

	// Deep copy ReplyParams
	cp.ReplyParams = deepCopyReplyParams(og.ReplyParams)

	return &cp
}

// deepCopyMessage creates a deep copy of a Message.
// Note: This is a shallow copy of most fields as Message is typically read-only.
// We copy the pointer to avoid shared references but don't recursively copy all nested messages.
func deepCopyMessage(msg *telegram.Message) *telegram.Message {
	if msg == nil {
		return nil
	}
	msgCopy := *msg
	return &msgCopy
}

// deepCopyReplyMarkup creates a deep copy of ReplyMarkup with all nested slices and pointers.
func deepCopyReplyMarkup(markup *telegram.ReplyMarkup) *telegram.ReplyMarkup {
	if markup == nil {
		return nil
	}

	cp := *markup

	// Deep copy InlineKeyboard
	if markup.InlineKeyboard != nil {
		cp.InlineKeyboard = make([][]telegram.InlineButton, len(markup.InlineKeyboard))
		for i, row := range markup.InlineKeyboard {
			cp.InlineKeyboard[i] = make([]telegram.InlineButton, len(row))
			for j, btn := range row {
				cp.InlineKeyboard[i][j] = deepCopyInlineButton(btn)
			}
		}
	}

	// Deep copy ReplyKeyboard
	if markup.ReplyKeyboard != nil {
		cp.ReplyKeyboard = make([][]telegram.ReplyButton, len(markup.ReplyKeyboard))
		for i, row := range markup.ReplyKeyboard {
			cp.ReplyKeyboard[i] = make([]telegram.ReplyButton, len(row))
			for j, btn := range row {
				cp.ReplyKeyboard[i][j] = deepCopyReplyButton(btn)
			}
		}
	}

	return &cp
}

// deepCopyInlineButton creates a deep copy of an InlineButton.
func deepCopyInlineButton(btn telegram.InlineButton) telegram.InlineButton {
	cp := btn

	if btn.InlineQueryChosenChat != nil {
		chatCopy := *btn.InlineQueryChosenChat
		cp.InlineQueryChosenChat = &chatCopy
	}

	if btn.Login != nil {
		loginCopy := *btn.Login
		cp.Login = &loginCopy
	}

	if btn.WebApp != nil {
		webAppCopy := *btn.WebApp
		cp.WebApp = &webAppCopy
	}

	if btn.CallbackGame != nil {
		gameCopy := *btn.CallbackGame
		cp.CallbackGame = &gameCopy
	}

	return cp
}

// deepCopyReplyButton creates a deep copy of a ReplyButton.
func deepCopyReplyButton(btn telegram.ReplyButton) telegram.ReplyButton {
	cp := btn

	if btn.User != nil {
		userCopy := deepCopyReplyRecipient(*btn.User)
		cp.User = &userCopy
	}

	if btn.Chat != nil {
		chatCopy := deepCopyReplyRecipient(*btn.Chat)
		cp.Chat = &chatCopy
	}

	if btn.WebApp != nil {
		webAppCopy := *btn.WebApp
		cp.WebApp = &webAppCopy
	}

	return cp
}

// deepCopyReplyRecipient creates a deep copy of ReplyRecipient with all pointer fields.
func deepCopyReplyRecipient(rr telegram.ReplyRecipient) telegram.ReplyRecipient {
	cp := rr

	// Copy *bool fields
	if rr.Bot != nil {
		val := *rr.Bot
		cp.Bot = &val
	}
	if rr.Premium != nil {
		val := *rr.Premium
		cp.Premium = &val
	}
	if rr.Forum != nil {
		val := *rr.Forum
		cp.Forum = &val
	}
	if rr.WithUsername != nil {
		val := *rr.WithUsername
		cp.WithUsername = &val
	}
	if rr.Created != nil {
		val := *rr.Created
		cp.Created = &val
	}
	if rr.BotMember != nil {
		val := *rr.BotMember
		cp.BotMember = &val
	}
	if rr.RequestTitle != nil {
		val := *rr.RequestTitle
		cp.RequestTitle = &val
	}
	if rr.RequestName != nil {
		val := *rr.RequestName
		cp.RequestName = &val
	}
	if rr.RequestUsername != nil {
		val := *rr.RequestUsername
		cp.RequestUsername = &val
	}
	if rr.RequestPhoto != nil {
		val := *rr.RequestPhoto
		cp.RequestPhoto = &val
	}

	// Copy ChatPermissions pointers (shallow copy as it's typically read-only)
	if rr.UserRights != nil {
		rightsCopy := *rr.UserRights
		cp.UserRights = &rightsCopy
	}
	if rr.BotRights != nil {
		rightsCopy := *rr.BotRights
		cp.BotRights = &rightsCopy
	}

	return cp
}

// deepCopyMessageEntity creates a deep copy of a MessageEntity.
func deepCopyMessageEntity(entity telegram.MessageEntity) telegram.MessageEntity {
	cp := entity

	// Copy User pointer if present
	if entity.User != nil {
		userCopy := *entity.User
		cp.User = &userCopy
	}

	return cp
}

// deepCopyReplyParams creates a deep copy of ReplyParams.
func deepCopyReplyParams(params *telegram.ReplyParams) *telegram.ReplyParams {
	if params == nil {
		return nil
	}

	cp := *params

	// Deep copy QuoteEntities slice
	if params.QuoteEntities != nil {
		cp.QuoteEntities = make([]telegram.MessageEntity, len(params.QuoteEntities))
		for i, entity := range params.QuoteEntities {
			cp.QuoteEntities[i] = deepCopyMessageEntity(entity)
		}
	}

	return &cp
}
