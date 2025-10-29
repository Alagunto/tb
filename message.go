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
	ModeDefault    = telegram.ModeDefault
	ModeMarkdown   = telegram.ModeMarkdown
	ModeMarkdownV2 = telegram.ModeMarkdownV2
	ModeHTML       = telegram.ModeHTML
)

// Type aliases for reactions
type (
	MessageReaction      = telegram.MessageReaction
	MessageReactionCount = telegram.MessageReactionCount
)
