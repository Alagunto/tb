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

func (b *Builder) With(opts ...SendOptions) *Builder {
	Merge(opts...).InjectIntoMap(b.params)
	return b
}

// Build returns the constructed parameter map.
func (b *Builder) Build() map[string]any {
	return b.params
}
