package tb

// CensorText is the public interface method that implements RawBotInterface.
func (b *Bot[RequestType, HandlerFunc, MiddlewareFunc]) CensorText(text string) string {
	if b.censorer == nil {
		return text
	}
	return b.censorer.CensorText(text)
}
