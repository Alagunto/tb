package params

import (
	"encoding/json"
	"fmt"
)

// Builder helps construct parameter maps for Telegram API calls,
// skipping empty values to keep requests clean.
type Builder struct {
	params map[string]any
}

// New creates a new parameter Builder.
func New() *Builder {
	return &Builder{
		params: make(map[string]any),
	}
}

// Add adds a string parameter, skipping if the value is empty.
func (b *Builder) Add(key string, value any) *Builder {
	if value != "" {
		b.params[key] = value
	}
	return b
}

// AddInt adds an integer parameter, skipping if the value is zero.
func (b *Builder) AddInt(key string, value int) *Builder {
	if value != 0 {
		b.params[key] = value
	}
	return b
}

// AddInt64 adds an int64 parameter, skipping if the value is zero.
func (b *Builder) AddInt64(key string, value int64) *Builder {
	if value != 0 {
		b.params[key] = value
	}
	return b
}

// AddBool adds a boolean parameter.
// Note: For Telegram API, we typically only add true values.
func (b *Builder) AddBool(key string, value bool) *Builder {
	if value {
		b.params[key] = true
	}
	return b
}

// AddFloat adds a float32 parameter, skipping if the value is zero.
func (b *Builder) AddFloat(key string, value float32) *Builder {
	if value != 0 {
		b.params[key] = value
	}
	return b
}

// AddJSON adds a parameter as JSON-encoded data.
func (b *Builder) AddJSON(key string, value interface{}) error {
	if value == nil {
		return nil
	}

	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("failed to marshal %s: %w", key, err)
	}

	b.params[key] = string(data)
	return nil
}

// Build returns the constructed parameter map.
func (b *Builder) Build() map[string]any {
	return b.params
}

// MediaParams contains common parameters for media messages.
type MediaParams struct {
	Caption      string
	ParseMode    string
	Entities     interface{} // json.Marshaler or similar
	HasSpoiler   bool
	CaptionAbove bool
}

// AddParseMode adds parse mode parameter.
func (b *Builder) AddParseMode(mode string) *Builder {
	if mode != "" {
		b.params["parse_mode"] = mode
	}
	return b
}

// AddEntities adds entities parameter.
func (b *Builder) AddEntities(entities interface{}) *Builder {
	if entities != nil {
		if data, err := json.Marshal(entities); err == nil && string(data) != "null" {
			b.params["caption_entities"] = string(data)
		}
	}
	return b
}

// Apply adds media parameters to the builder.
func (mp *MediaParams) Apply(b *Builder) error {
	b.Add("caption", mp.Caption)
	b.AddParseMode(mp.ParseMode)
	b.AddEntities(mp.Entities)
	b.AddBool("has_spoiler", mp.HasSpoiler)
	b.AddBool("show_caption_above_media", mp.CaptionAbove)
	return nil
}

// DimensionParams contains dimension parameters for media (width, height, duration).
type DimensionParams struct {
	Width    int
	Height   int
	Duration int
}

// Apply adds dimension parameters to the builder.
func (dp *DimensionParams) Apply(b *Builder) {
	b.AddInt("width", dp.Width)
	b.AddInt("height", dp.Height)
	b.AddInt("duration", dp.Duration)
}
