package layout

import (
	"embed"
	"os"
	"testing"
	"time"

	"github.com/alagunto/tb"
	"github.com/stretchr/testify/assert"
)

//go:embed *
var fsys embed.FS

func TestLayout(t *testing.T) {
	os.Setenv("TOKEN", "TEST")

	lt, err := New("example.yml")
	if err != nil {
		t.Fatal(err)
	}

	ltfs, err := NewFromFS(fsys, "example.yml")
	if err != nil {
		t.Fatal(err)
	}

	pref := lt.Settings()
	assert.Equal(t, "TEST", pref.Token)
	assert.Equal(t, tb.ModeHTML, pref.ParseMode)
	assert.Equal(t, &tb.LongPoller{}, pref.Poller)
	assert.Equal(t, pref, ltfs.Settings())

	assert.ElementsMatch(t, []tb.Command{{
		Text:        "start",
		Description: "Start the bot",
	}, {
		Text:        "help",
		Description: "How to use the bot",
	}}, lt.Commands())

	assert.Equal(t, "string", lt.String("str"))
	assert.Equal(t, 123, lt.Int("num"))
	assert.Equal(t, int64(123), lt.Int64("num"))
	assert.Equal(t, float64(123), lt.Float("num"))
	assert.Equal(t, tb.ChatID(123), lt.ChatID("num"))

	assert.Equal(t, []string{"abc", "def"}, lt.Strings("strs"))
	assert.Equal(t, []int{123, 456}, lt.Ints("nums"))
	assert.Equal(t, []int64{123, 456}, lt.Int64s("nums"))
	assert.Equal(t, []float64{123, 456}, lt.Floats("nums"))

	obj := lt.Get("obj")
	assert.NotNil(t, obj)

	const dur = 10 * time.Minute
	assert.Equal(t, dur, obj.Duration("dur"))
	assert.True(t, lt.Duration("obj.dur") == obj.Duration("dur"))

	arr := lt.Slice("arr")
	assert.Len(t, arr, 2)

	for _, v := range arr {
		assert.Equal(t, dur, v.Duration("dur"))
	}

	assert.Equal(t, &tb.Btn{
		Unique: "pay",
		Text:   "Pay",
		Data:   "1|100.00|USD",
	}, lt.ButtonLocale("en", "pay", struct {
		UserID   int
		Amount   string
		Currency string
	}{
		UserID:   1,
		Amount:   "100.00",
		Currency: "USD",
	}))

	assert.Equal(t, &tb.ReplyMarkup{
		ReplyKeyboard: [][]tb.ReplyButton{
			{{Text: "Help"}},
			{{Text: "Settings"}},
		},
		ResizeKeyboard: true,
	}, lt.MarkupLocale("en", "reply_shortened"))

	assert.Equal(t, &tb.ReplyMarkup{
		ReplyKeyboard:   [][]tb.ReplyButton{{{Text: "Send a contact", Contact: true}}},
		ResizeKeyboard:  true,
		OneTimeKeyboard: true,
	}, lt.MarkupLocale("en", "reply_extended"))

	assert.Equal(t, &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{{
			{
				Unique: "stop",
				Text:   "Stop",
				Data:   "1",
			},
		}},
	}, lt.MarkupLocale("en", "inline", 1))

	assert.Equal(t, &tb.ReplyMarkup{
		InlineKeyboard: [][]tb.InlineButton{{
			{
				Text:   "This is a web app",
				WebApp: &tb.WebApp{URL: "https://google.com"},
			},
		}},
	}, lt.MarkupLocale("en", "web_app"))

	assert.Equal(t, &tb.ArticleResult{
		ResultBase: tb.ResultBase{
			ID:   "1853",
			Type: "article",
		},
		Title:       "Some title",
		Description: "Some description",
		ThumbURL:    "https://preview.picture",
		Text:        "This is an article.",
	}, lt.ResultLocale("en", "article", struct {
		ID          int
		Title       string
		Description string
		PreviewURL  string
	}{
		ID:          1853,
		Title:       "Some title",
		Description: "Some description",
		PreviewURL:  "https://preview.picture",
	}))

	assert.Equal(t,
		"This is an article.",
		lt.TextLocale("en", "article_message"),
	)
	assert.Equal(t,
		lt.TextLocale("en", "nested.example", "an example"),
		"This is an example.",
	)
	assert.Equal(t,
		lt.TextLocale("en", "nested.another.example", "another example"),
		"This is another example.",
	)
}
