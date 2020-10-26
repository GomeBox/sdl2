package audio

import (
	"github.com/veandco/go-sdl2/mix"
)

type sound struct {
	chunk *mix.Chunk
}

func newSound(chunk *mix.Chunk) *sound {
	return &sound{chunk: chunk}
}

func (sound *sound) Play() error {
	_, err := sound.chunk.Play(-1, 0)
	return err
}
