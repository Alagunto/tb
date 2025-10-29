package censorship

import (
	"regexp"
	"strings"
)

type Censorer interface {
	CensorText(text string) string
}

type SpecificSubstringsCensorer struct {
	censoredWords []string
}

func NewSpecificSubstringsCensorer(censoredWords []string) Censorer {
	return &SpecificSubstringsCensorer{censoredWords: censoredWords}
}

// censorText filters out censored words from the given text (case-insensitive).
// It replaces censored words with asterisks of the same length.
func (c *SpecificSubstringsCensorer) censorText(text string) string {
	if len(c.censoredWords) == 0 {
		return text
	}

	result := text
	for _, word := range c.censoredWords {
		if word == "" {
			continue
		}

		// Create a case-insensitive regex to find all variations of the word
		// Using \b for word boundaries to match complete words only
		pattern := `(?i)\b` + regexp.QuoteMeta(word) + `\b`
		re := regexp.MustCompile(pattern)

		// Replace with asterisks of the same length
		result = re.ReplaceAllStringFunc(result, func(match string) string {
			return strings.Repeat("*", len(match))
		})
	}

	return result
}

func (c *SpecificSubstringsCensorer) CensorText(text string) string {
	return c.censorText(text)
}
