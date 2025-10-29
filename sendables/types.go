package sendables

type SendMethod string

const (
	SendMethodText      SendMethod = "text"
	SendMethodMessage   SendMethod = "sendMessage"
	SendMethodPhoto     SendMethod = "sendPhoto"
	SendMethodAudio     SendMethod = "sendAudio"
	SendMethodDocument  SendMethod = "sendDocument"
	SendMethodVideo     SendMethod = "sendVideo"
	SendMethodVoice     SendMethod = "sendVoice"
	SendMethodVideoNote SendMethod = "sendVideoNote"
	SendMethodAnimation SendMethod = "sendAnimation"
)

func (s SendMethod) String() string {
	return string(s)
}
