package tb

type OutgoingMessage struct {
	Message string
	Options []interface{}
}

var _ Sendable = (*OutgoingMessage)(nil)
