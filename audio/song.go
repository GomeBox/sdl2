package audio

import "github.com/veandco/go-sdl2/mix"

type song struct {
	music *mix.Music
}

func newSong(music *mix.Music) *song {
	return &song{music: music}
}

func (song *song) Play() error {
	return song.music.Play(1)
}
