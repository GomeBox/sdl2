package audio

import (
	"github.com/GomeBox/gome/adapters/audio"
	"github.com/veandco/go-sdl2/mix"
)

type songLoader struct {
}

func (songLoader *songLoader) Load(fileName string) (audio.Song, error) {
	mus, err := mix.LoadMUS(fileName)
	if err != nil {
		return nil, err
	}
	return newSong(mus), nil
}
