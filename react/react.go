package react

import (
	"github.com/alagunto/tb"
)

type Reaction = tb.Reaction

func React(r ...Reaction) tb.Reactions {
	return tb.Reactions{Reactions: r}
}

// Currently available emojis.
var (
	ThumbUp                   = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "👍"}
	ThumbDown                 = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "👎"}
	Heart                     = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "❤"}
	Fire                      = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🔥"}
	HeartEyes                 = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😍"}
	ClappingHands             = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "👏"}
	GrinningFace              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😁"}
	ThinkingFace              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤔"}
	ExplodingHead             = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤯"}
	ScreamingFace             = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😱"}
	SwearingFace              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤬"}
	CryingFace                = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😢"}
	PartyPopper               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🎉"}
	StarStruck                = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤩"}
	VomitingFace              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤮"}
	PileOfPoo                 = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "💩"}
	PrayingHands              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🙏"}
	OkHand                    = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "👌"}
	DoveOfPeace               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🕊"}
	ClownFace                 = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤡"}
	YawningFace               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🥱"}
	WoozyFace                 = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🥴"}
	Whale                     = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🐳"}
	HeartOnFire               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "❤‍🔥"}
	MoonFace                  = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🌚"}
	HotDog                    = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🌭"}
	HundredPoints             = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "💯"}
	RollingOnTheFloorLaughing = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤣"}
	Lightning                 = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "⚡"}
	Banana                    = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🍌"}
	Trophy                    = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🏆"}
	BrokenHeart               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "💔"}
	FaceWithRaisedEyebrow     = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤨"}
	NeutralFace               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😐"}
	Strawberry                = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🍓"}
	Champagne                 = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🍾"}
	KissMark                  = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "💋"}
	MiddleFinger              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🖕"}
	EvilFace                  = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😈"}
	SleepingFace              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😴"}
	LoudlyCryingFace          = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😭"}
	NerdFace                  = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤓"}
	Ghost                     = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "👻"}
	Engineer                  = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "👨‍💻"}
	Eyes                      = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "👀"}
	JackOLantern              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🎃"}
	NoMonkey                  = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🙈"}
	SmilingFaceWithHalo       = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😇"}
	FearfulFace               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😨"}
	Handshake                 = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤝"}
	WritingHand               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "✍"}
	HuggingFace               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤗"}
	Brain                     = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🫡"}
	SantaClaus                = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🎅"}
	ChristmasTree             = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🎄"}
	Snowman                   = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "☃"}
	NailPolish                = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "💅"}
	ZanyFace                  = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤪"}
	Moai                      = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🗿"}
	Cool                      = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🆒"}
	HeartWithArrow            = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "💘"}
	HearMonkey                = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🙉"}
	Unicorn                   = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🦄"}
	FaceBlowingKiss           = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😘"}
	Pill                      = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "💊"}
	SpeaklessMonkey           = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🙊"}
	Sunglasses                = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😎"}
	AlienMonster              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "👾"}
	ManShrugging              = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤷‍♂️"}
	PersonShrugging           = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤷"}
	WomanShrugging            = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "🤷‍♀️"}
	PoutingFace               = Reaction{Type: tb.ReactionTypeEmoji, Emoji: "😡"}
)
