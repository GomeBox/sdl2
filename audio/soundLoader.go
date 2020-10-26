package audio

import (
	"github.com/GomeBox/gome/adapters/audio"
	"github.com/veandco/go-sdl2/mix"
)

type soundLoader struct {
}

func (soundLoader *soundLoader) Load(fileName string) (audio.Sound, error) {
	chunk, err := mix.LoadWAV(fileName)
	if err != nil {
		return nil, err
	}
	return newSound(chunk), nil
}
