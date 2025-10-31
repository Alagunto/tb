package params

import "encoding/json"

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
