package sdl2

import (
	"errors"
	"fmt"
	"github.com/GomeBox/gome/adapters/input"
)

var _ input.Port = (*InputAdapter)(nil)

type InputAdapter struct {
	keyboard *keyboard
}

func (adapter *InputAdapter) Init() error {
	keyboard, err := newKeyboard()
	if err != nil {
		return err
	}
	adapter.keyboard = keyboard
	return nil
}

func (adapter *InputAdapter) Update() {
	//Not necessary as keystate is updated by polling event in System.Update
}

func (adapter *InputAdapter) Keyboard() input.Keyboard {
	return adapter.keyboard
}

func (adapter *InputAdapter) ControllerCount() int {
	return 0
}

func (adapter *InputAdapter) Controller(number int) (*input.Controller, error) {
	return nil, errors.New("Controller number " + fmt.Sprint(number) + " not found")
}
