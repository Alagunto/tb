# Rules of project

- All telegram models must be in ./telegram. That includes any model that is declared by the telegram api.
- All the bot.go methods that interact with telegram api must use ./api_requester.go
- bot.go methods must not declare their own structs. instead use ./telegram/ package structs as type params for generic api_requester.go
- Clean architecture. I really mean it.
- No repeating yourself. Create abstractions when required.
- For sending/editing messages methods in Bot, we must support ./params/ as sending options to allow interfaces like bot.Send(chat, photo, tb.SendOptions().WithoutNotification().WithEffectID(effectID))
