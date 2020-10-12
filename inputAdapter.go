package game_sdl2

import (
	"github.com/GomeBox/gome/adapters/input"
)

var _ input.Interface = (*InputAdapter)(nil)

type InputAdapter struct {
}

func (adapter *InputAdapter) Init() {

}

func (adapter *InputAdapter) Update() {

}

func (adapter *InputAdapter) Keyboard() *input.Keyboard {
	return nil
}

func (adapter *InputAdapter) ControllerCount() int {
	return 0
}

func (adapter *InputAdapter) Controller(number int) (*input.Controller, error) {
	return nil, nil
}
