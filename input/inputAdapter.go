package input

import (
	"errors"
	"fmt"
	"github.com/GomeBox/gome/adapters/input"
)

var _ input.Adapters = (*inputAdapter)(nil)

type InputAdapter interface {
	input.Adapters
	Init() error
}

func NewInputAdapter() InputAdapter {
	return new(inputAdapter)
}

type inputAdapter struct {
	keyboard *keyboard
}

func (adapter *inputAdapter) Init() error {
	keyboard, err := newKeyboard()
	if err != nil {
		return err
	}
	adapter.keyboard = keyboard
	return nil
}

func (adapter *inputAdapter) Update() {
	//Not necessary as keystate is updated by polling event in System.Update
}

func (adapter *inputAdapter) Keyboard() input.KeyboardAdapter {
	return adapter.keyboard
}

func (adapter *inputAdapter) ControllerCount() int {
	return 0
}

func (adapter *inputAdapter) Controller(number int) (*input.Controller, error) {
	return nil, errors.New("Controller number " + fmt.Sprint(number) + " not found")
}
