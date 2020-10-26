package sdl2

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/adapters/audio"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
	. "github.com/GomeBox/sdl2/audio"
	. "github.com/GomeBox/sdl2/graphics"
	. "github.com/GomeBox/sdl2/input"
	"github.com/veandco/go-sdl2/sdl"
)

var _ adapters.System = (*System)(nil)

type System struct {
	input    InputAdapter
	graphics GraphicsAdapter
	audio    AudioAdapter
}

func NewSystem() *System {
	s := new(System)
	s.input = NewInputAdapter()
	s.graphics = NewGraphicsAdapter()
	s.audio = NewAudioAdapter()
	return s
}

func (s *System) Initialize() error {
	err := s.graphics.Init()
	if err != nil {
		return err
	}
	err = s.input.Init()
	if err != nil {
		return err
	}
	err = s.audio.Init()
	if err != nil {
		return err
	}
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "2")
	return nil
}

func (s *System) Update() error {
	//Polling all events to keep game running
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	}
	return nil
}

func (s *System) Input() input.Adapters {
	return s.input
}

func (s *System) Graphics() graphics.Adapters {
	return s.graphics
}

func (s *System) Audio() audio.Adapters {
	return s.audio
}
