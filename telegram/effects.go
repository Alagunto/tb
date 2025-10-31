package telegram

// EffectID represents a unique identifier for a message effect.
// Used to add visual effects to messages (for private chats only).
//
// Note: Effect IDs are obtained from Telegram's getAvailableMessageEffects method.
// The IDs below are common effects, but you should verify current IDs via the API
// as they may change or new effects may be added.
//
// Core Effects (numeric IDs 0-5):
// These are the six standard effects available to all users.
// Use getAvailableMessageEffects API method for dynamic discovery of all effects,
// including premium effects with undocumented IDs.
type EffectID string

// Common message effect IDs.
// These represent the six core effects with numeric IDs 0-5.
const (
	EffectFire        EffectID = "5104841245755180586" // 🔥 Fire effect (ID: 0)
	EffectLike        EffectID = "5107584321108051014" // 👍 Like/thumbs up effect (ID: 1)
	EffectDislike     EffectID = "5104858069142078462" // 👎 Dislike effect (ID: 2)
	EffectHeart       EffectID = "5159385139981059251" // ❤️ Heart effect (ID: 3)
	EffectCelebration EffectID = "5046509860389126442" // 🎉 Celebration effect (ID: 4)
	EffectPoop        EffectID = "5046589136895476101" // 💩 Poop effect (ID: 5)
)
