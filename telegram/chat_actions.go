package telegram

// ChatAction is a client-side status indicating bot activity.
type ChatAction string

const (
	Typing             ChatAction = "typing"
	UploadingPhoto     ChatAction = "upload_photo"
	UploadingVideo     ChatAction = "upload_video"
	UploadingAudio     ChatAction = "upload_audio"
	UploadingDocument  ChatAction = "upload_document"
	UploadingVoice     ChatAction = "upload_voice"
	UploadingVideoNote ChatAction = "upload_video_note"
	FindingLocation    ChatAction = "find_location"
	RecordingVideo     ChatAction = "record_video"
	RecordingVoice     ChatAction = "record_voice"
	RecordingVideoNote ChatAction = "record_video_note"
	ChoosingSticker    ChatAction = "choose_sticker"
)
