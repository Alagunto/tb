package tb

type PolybotSidecar[T Context] struct {
	*Bot
	contextWrapper func(Context) T
}

func NewPolyContextBot[T Context](polySettings PolySettings[T]) (*PolybotSidecar[T], error) {
	b, err := NewBot(Settings{
		URL:         polySettings.URL,
		Token:       polySettings.Token,
		Updates:     polySettings.Updates,
		Poller:      polySettings.Poller,
		OnError:     polySettings.OnError,
		ParseMode:   polySettings.ParseMode,
		Synchronous: polySettings.Synchronous,
		Verbose:     polySettings.Verbose,
		Client:      polySettings.Client,
		Offline:     polySettings.Offline,
	})
	if err != nil {
		return nil, err
	}
	pb := &PolybotSidecar[T]{
		Bot:            b,
		contextWrapper: polySettings.ContextWrapper,
	}

	return pb, nil
}
