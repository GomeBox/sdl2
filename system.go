package sdl2

import (
	"github.com/GomeBox/gome/adapters"
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/adapters/input"
	"github.com/veandco/go-sdl2/sdl"
)

var _ adapters.System = (*System)(nil)

type System struct {
	input    *InputAdapter
	graphics *GraphicsAdapter
}

func NewSystem() *System {
	s := new(System)
	s.input = new(InputAdapter)
	s.graphics = new(GraphicsAdapter)

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
	return nil
}

func (s *System) Update() error {
	//Polling all events to keep game running
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
	}
	return nil
}

func (s *System) Input() input.Port {
	return s.input
}

func (s *System) Graphics() graphics.Port {
	return s.graphics
}
