package methods

import "github.com/alagunto/tb/telegram"

// EditMessageChecklistRequest represents the request for editMessageChecklist method
type EditMessageChecklistRequest struct {
	telegram.HasReplyMarkup

	ChatID          string           `json:"chat_id,omitempty"`
	MessageID       string           `json:"message_id,omitempty"`
	InlineMessageID string           `json:"inline_message_id,omitempty"`
	Checklist       telegram.Checklist `json:"checklist"`
}

// EditMessageChecklistResponse represents the response for editMessageChecklist method
type EditMessageChecklistResponse = telegram.Message
