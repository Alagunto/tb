package tb

import "github.com/alagunto/tb/telegram"

// Message is an alias for telegram.Message
type Message = telegram.Message

// Type aliases for message-related types
type (
	MessageEntity = telegram.MessageEntity
	EntityType    = telegram.EntityType
	Entities      = telegram.Entities
	ParseMode     = telegram.ParseMode
	Album         = telegram.Album
	PaidAlbum     = telegram.PaidAlbum
	Inputtable    = telegram.Inputtable
	Media         = telegram.Media
)

// Re-export entity type constants
const (
	EntityMention       = telegram.EntityMention
	EntityTMention      = telegram.EntityTMention
	EntityHashtag       = telegram.EntityHashtag
	EntityCashtag       = telegram.EntityCashtag
	EntityCommand       = telegram.EntityCommand
	EntityURL           = telegram.EntityURL
	EntityEmail         = telegram.EntityEmail
	EntityPhone         = telegram.EntityPhone
	EntityBold          = telegram.EntityBold
	EntityItalic        = telegram.EntityItalic
	EntityUnderline     = telegram.EntityUnderline
	EntityStrikethrough = telegram.EntityStrikethrough
	EntityCode          = telegram.EntityCode
	EntityCodeBlock     = telegram.EntityCodeBlock
	EntityTextLink      = telegram.EntityTextLink
	EntitySpoiler       = telegram.EntitySpoiler
	EntityCustomEmoji   = telegram.EntityCustomEmoji
	EntityBlockquote    = telegram.EntityBlockquote
	EntityEBlockquote   = telegram.EntityEBlockquote
)

// Re-export parse mode constants
const (
	ModeDefault    = telegram.ParseModeDefault
	ModeMarkdown   = telegram.ParseModeMarkdown
	ModeMarkdownV2 = telegram.ParseModeMarkdownV2
	ModeHTML       = telegram.ParseModeHTML
)

// Type aliases for reactions
type (
	MessageReaction      = telegram.MessageReaction
	MessageReactionCount = telegram.MessageReactionCount
)

// Type aliases for menu and callback
type (
	MenuButton       = telegram.MenuButton
	CallbackResponse = telegram.CallbackResponse
	Location         = telegram.Location
	ShippingOption   = telegram.ShippingOption
)
