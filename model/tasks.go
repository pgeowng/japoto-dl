package model

type AudioHLS interface {
	Playlist(body string) (tsaudio []File, err error)
	TSAudio(tsaudio File) (keys []File, audio []File, err error)
	Validate(file File) error
	Image(file File) error
	CheckAlreadyLoaded(filename string) bool
}

type Tasks interface {
	AudioHLS() AudioHLS
}
