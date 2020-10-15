package input

import (
	"errors"
	"fmt"
	"github.com/GomeBox/gome/adapters/input"
	"github.com/veandco/go-sdl2/sdl"
)

type keyboard struct {
	keystate []uint8
	keymap   map[input.KeyType]uint8
}

func newKeyboard() (*keyboard, error) {
	keyboard := new(keyboard)
	keyboard.keystate = sdl.GetKeyboardState()
	keymap, err := createKeymap()
	if err != nil {
		return nil, err
	}
	keyboard.keymap = keymap
	return keyboard, nil
}

func (keyboard keyboard) KeyPressed(key input.KeyType) (bool, error) {
	sdlKey, ok := keyboard.keymap[key]
	if !ok {
		return false, errors.New("Key " + fmt.Sprint(key) + " was not found")
	}
	return keyboard.keystate[sdlKey] != 0, nil
}

func createKeymap() (map[input.KeyType]uint8, error) {
	keymap := make(map[input.KeyType]uint8)
	if err := addKeymapKey(input.KeyEsc, sdl.SCANCODE_ESCAPE, keymap); err != nil {
		return nil, err
	}
	//keymap[input.KeyEsc,sdl.SCANCODE_ESCAPE
	if err := addKeymapKey(input.KeyA, sdl.SCANCODE_A, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyB, sdl.SCANCODE_B, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyC, sdl.SCANCODE_C, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyD, sdl.SCANCODE_D, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyE, sdl.SCANCODE_E, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF, sdl.SCANCODE_F, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyG, sdl.SCANCODE_G, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyH, sdl.SCANCODE_H, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyI, sdl.SCANCODE_I, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyJ, sdl.SCANCODE_J, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyK, sdl.SCANCODE_K, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyL, sdl.SCANCODE_L, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyM, sdl.SCANCODE_M, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyN, sdl.SCANCODE_N, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyO, sdl.SCANCODE_O, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyP, sdl.SCANCODE_P, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyQ, sdl.SCANCODE_Q, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyR, sdl.SCANCODE_R, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyS, sdl.SCANCODE_S, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyT, sdl.SCANCODE_T, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyU, sdl.SCANCODE_U, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyV, sdl.SCANCODE_V, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyW, sdl.SCANCODE_W, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyX, sdl.SCANCODE_X, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyY, sdl.SCANCODE_Y, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyZ, sdl.SCANCODE_Z, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyUp, sdl.SCANCODE_UP, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyDown, sdl.SCANCODE_DOWN, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyLeft, sdl.SCANCODE_LEFT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyRight, sdl.SCANCODE_RIGHT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key0, sdl.SCANCODE_0, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key1, sdl.SCANCODE_1, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key2, sdl.SCANCODE_2, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key3, sdl.SCANCODE_3, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key4, sdl.SCANCODE_4, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key5, sdl.SCANCODE_5, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key6, sdl.SCANCODE_6, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key7, sdl.SCANCODE_7, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key8, sdl.SCANCODE_8, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.Key9, sdl.SCANCODE_9, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyRightCtrl, sdl.SCANCODE_RCTRL, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyLeftCtrl, sdl.SCANCODE_LCTRL, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyRightAlt, sdl.SCANCODE_RALT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyLeftAlt, sdl.SCANCODE_LALT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeySpace, sdl.SCANCODE_SPACE, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyReturn, sdl.SCANCODE_RETURN, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyBackslash, sdl.SCANCODE_BACKSLASH, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyBackspace, sdl.SCANCODE_BACKSPACE, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyCapslock, sdl.SCANCODE_CAPSLOCK, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyComma, sdl.SCANCODE_COMMA, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyDelete, sdl.SCANCODE_DELETE, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyEnd, sdl.SCANCODE_END, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyEquals, sdl.SCANCODE_EQUALS, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF1, sdl.SCANCODE_F1, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF2, sdl.SCANCODE_F2, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF3, sdl.SCANCODE_F3, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF4, sdl.SCANCODE_F4, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF5, sdl.SCANCODE_F5, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF6, sdl.SCANCODE_F6, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF7, sdl.SCANCODE_F7, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF8, sdl.SCANCODE_F8, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF9, sdl.SCANCODE_F9, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF10, sdl.SCANCODE_F10, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF11, sdl.SCANCODE_F11, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF12, sdl.SCANCODE_F12, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyF13, sdl.SCANCODE_F13, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyHome, sdl.SCANCODE_HOME, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyInsert, sdl.SCANCODE_INSERT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum0, sdl.SCANCODE_KP_0, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum1, sdl.SCANCODE_KP_1, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum2, sdl.SCANCODE_KP_2, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum3, sdl.SCANCODE_KP_3, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum4, sdl.SCANCODE_KP_4, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum5, sdl.SCANCODE_KP_5, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum6, sdl.SCANCODE_KP_6, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum7, sdl.SCANCODE_KP_7, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum8, sdl.SCANCODE_KP_8, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyNum9, sdl.SCANCODE_KP_9, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyPageDown, sdl.SCANCODE_PAGEDOWN, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyPageUp, sdl.SCANCODE_PAGEUP, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyPause, sdl.SCANCODE_PAUSE, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyPeriod, sdl.SCANCODE_PERIOD, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyRightBracket, sdl.SCANCODE_RIGHTBRACKET, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyLeftBracket, sdl.SCANCODE_LEFTBRACKET, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeySemicolon, sdl.SCANCODE_SEMICOLON, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeySlash, sdl.SCANCODE_SLASH, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(input.KeyTab, sdl.SCANCODE_TAB, keymap); err != nil {
		return nil, err
	}
	return keymap, nil
}

func addKeymapKey(key input.KeyType, sdlScanCode uint8, keymap map[input.KeyType]uint8) error {
	_, keyFound := keymap[key]
	if keyFound {
		return errors.New("Key " + fmt.Sprint(key) + " already in keymap")
	}
	keymap[key] = sdlScanCode
	return nil
}
