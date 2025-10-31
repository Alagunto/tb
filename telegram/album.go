package telegram

// Album is a media group.
type Album []InputMedia

// InputAlbum is a media group to be sent.
type InputAlbum struct {
	Media []InputMedia
}
