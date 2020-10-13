package sdl2

import (
	"errors"
	"github.com/GomeBox/gome/adapters/input"
	"github.com/veandco/go-sdl2/sdl"
)

var _ input.Port = (*InputAdapter)(nil)

type InputAdapter struct {
	keyboard keyboard
}

type keyboard struct {
	keystate []uint8
	keymap   map[input.KeyType]uint8
}

func (keyboard keyboard) KeyPressed(key input.KeyType) (bool, error) {
	sdlKey, ok := keyboard.keymap[key]
	if !ok {
		return false, errors.New("Key " + string(key) + " was not found")
	}
	return keyboard.keystate[sdlKey] != 0, nil
}

func CreateKeyMapping() map[input.KeyType]uint8 {
	keymap := make(map[input.KeyType]uint8)
	keymap[input.KeyEsc] = sdl.SCANCODE_ESCAPE
	return keymap
}

func (adapter *InputAdapter) Init() {
	adapter.keyboard.keystate = sdl.GetKeyboardState()
	adapter.keyboard.keymap = CreateKeyMapping()
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
	return nil, nil
}
