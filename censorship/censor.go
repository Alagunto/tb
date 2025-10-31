package censorship

import (
	ahocorasick "github.com/pgavlin/aho-corasick"
)

type Censorer interface {
	CensorText(text string) string
}

type SpecificSubstringsCensorer struct {
	ac              ahocorasick.AhoCorasick
	censoredWords   []string
	caseInsensitive bool
}

func NewSpecificSubstringsCensorer(censoredWords []string) Censorer {
	return NewSpecificSubstringsCensorerWithOptions(censoredWords, true)
}

func NewSpecificSubstringsCensorerWithOptions(censoredWords []string, caseInsensitive bool) Censorer {
	if len(censoredWords) == 0 {
		return &SpecificSubstringsCensorer{
			censoredWords:   censoredWords,
			caseInsensitive: caseInsensitive,
		}
	}

	// Build the Aho-Corasick automaton
	builder := ahocorasick.NewAhoCorasickBuilder(ahocorasick.Opts{
		AsciiCaseInsensitive: caseInsensitive,
		MatchOnlyWholeWords:  true,
		MatchKind:            ahocorasick.LeftMostLongestMatch,
		DFA:                  true, // Use DFA for O(N) performance
	})

	ac := builder.Build(censoredWords)

	return &SpecificSubstringsCensorer{
		ac:              ac,
		censoredWords:   censoredWords,
		caseInsensitive: caseInsensitive,
	}
}

// CensorText filters out censored words from the given text.
// It replaces censored words with asterisks of the same length.
func (c *SpecificSubstringsCensorer) CensorText(text string) string {
	if len(c.censoredWords) == 0 {
		return text
	}

	// Find all matches using Aho-Corasick
	matches := c.ac.FindAll(text)

	if len(matches) == 0 {
		return text
	}

	// Build result by replacing matches with asterisks
	// Work with bytes since Aho-Corasick returns byte indices
	result := []byte(text)
	for _, match := range matches {
		for i := match.Start(); i < match.End(); i++ {
			result[i] = '*'
		}
	}

	return string(result)
}
