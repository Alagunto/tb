package telegram

// MenuButton describes the bot's menu button in a private chat.
type MenuButton struct {
	Type   MenuButtonType `json:"type"`
	Text   string         `json:"text,omitempty"`
	WebApp *WebApp        `json:"web_app,omitempty"`
}

type MenuButtonType = string

const (
	MenuButtonDefault  MenuButtonType = "default"
	MenuButtonCommands MenuButtonType = "commands"
	MenuButtonWebApp   MenuButtonType = "web_app"
)
