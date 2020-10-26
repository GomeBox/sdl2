package audio

import (
	"github.com/GomeBox/gome/adapters/audio"
	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

type AudioAdapter interface {
	Init() error
	SongLoader() audio.SongLoader
	SoundLoader() audio.SoundLoader
}

type audioAdapter struct {
	songLoader  *songLoader
	soundLoader *soundLoader
}

func NewAudioAdapter() AudioAdapter {
	a := new(audioAdapter)
	a.soundLoader = new(soundLoader)
	a.songLoader = new(songLoader)
	return a
}

func (adapter *audioAdapter) Init() error {
	err := sdl.Init(sdl.INIT_AUDIO)
	if err != nil {
		return err
	}
	return mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, mix.DEFAULT_CHUNKSIZE)
}

func (adapter *audioAdapter) SongLoader() audio.SongLoader {
	return adapter.songLoader
}

func (adapter *audioAdapter) SoundLoader() audio.SoundLoader {
	return adapter.soundLoader
}
