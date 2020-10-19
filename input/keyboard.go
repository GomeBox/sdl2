package input

import (
	"errors"
	"fmt"
	"github.com/GomeBox/gome/primitives"
	"github.com/veandco/go-sdl2/sdl"
)

type keyboard struct {
	keystate []uint8
	keymap   map[primitives.KeyType]uint8
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

func (keyboard keyboard) KeyPressed(key primitives.KeyType) (bool, error) {
	sdlKey, ok := keyboard.keymap[key]
	if !ok {
		return false, errors.New("Key " + fmt.Sprint(key) + " was not found")
	}
	return keyboard.keystate[sdlKey] != 0, nil
}

func createKeymap() (map[primitives.KeyType]uint8, error) {
	keymap := make(map[primitives.KeyType]uint8)
	if err := addKeymapKey(primitives.KeyEsc, sdl.SCANCODE_ESCAPE, keymap); err != nil {
		return nil, err
	}
	//keymap[primitives.KeyEsc,sdl.SCANCODE_ESCAPE
	if err := addKeymapKey(primitives.KeyA, sdl.SCANCODE_A, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyB, sdl.SCANCODE_B, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyC, sdl.SCANCODE_C, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyD, sdl.SCANCODE_D, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyE, sdl.SCANCODE_E, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF, sdl.SCANCODE_F, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyG, sdl.SCANCODE_G, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyH, sdl.SCANCODE_H, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyI, sdl.SCANCODE_I, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyJ, sdl.SCANCODE_J, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyK, sdl.SCANCODE_K, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyL, sdl.SCANCODE_L, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyM, sdl.SCANCODE_M, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyN, sdl.SCANCODE_N, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyO, sdl.SCANCODE_O, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyP, sdl.SCANCODE_P, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyQ, sdl.SCANCODE_Q, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyR, sdl.SCANCODE_R, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyS, sdl.SCANCODE_S, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyT, sdl.SCANCODE_T, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyU, sdl.SCANCODE_U, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyV, sdl.SCANCODE_V, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyW, sdl.SCANCODE_W, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyX, sdl.SCANCODE_X, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyY, sdl.SCANCODE_Y, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyZ, sdl.SCANCODE_Z, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyUp, sdl.SCANCODE_UP, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyDown, sdl.SCANCODE_DOWN, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyLeft, sdl.SCANCODE_LEFT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyRight, sdl.SCANCODE_RIGHT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key0, sdl.SCANCODE_0, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key1, sdl.SCANCODE_1, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key2, sdl.SCANCODE_2, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key3, sdl.SCANCODE_3, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key4, sdl.SCANCODE_4, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key5, sdl.SCANCODE_5, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key6, sdl.SCANCODE_6, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key7, sdl.SCANCODE_7, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key8, sdl.SCANCODE_8, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.Key9, sdl.SCANCODE_9, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyRightCtrl, sdl.SCANCODE_RCTRL, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyLeftCtrl, sdl.SCANCODE_LCTRL, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyRightAlt, sdl.SCANCODE_RALT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyLeftAlt, sdl.SCANCODE_LALT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeySpace, sdl.SCANCODE_SPACE, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyReturn, sdl.SCANCODE_RETURN, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyBackslash, sdl.SCANCODE_BACKSLASH, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyBackspace, sdl.SCANCODE_BACKSPACE, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyCapslock, sdl.SCANCODE_CAPSLOCK, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyComma, sdl.SCANCODE_COMMA, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyDelete, sdl.SCANCODE_DELETE, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyEnd, sdl.SCANCODE_END, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyEquals, sdl.SCANCODE_EQUALS, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF1, sdl.SCANCODE_F1, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF2, sdl.SCANCODE_F2, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF3, sdl.SCANCODE_F3, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF4, sdl.SCANCODE_F4, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF5, sdl.SCANCODE_F5, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF6, sdl.SCANCODE_F6, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF7, sdl.SCANCODE_F7, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF8, sdl.SCANCODE_F8, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF9, sdl.SCANCODE_F9, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF10, sdl.SCANCODE_F10, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF11, sdl.SCANCODE_F11, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF12, sdl.SCANCODE_F12, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyF13, sdl.SCANCODE_F13, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyHome, sdl.SCANCODE_HOME, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyInsert, sdl.SCANCODE_INSERT, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum0, sdl.SCANCODE_KP_0, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum1, sdl.SCANCODE_KP_1, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum2, sdl.SCANCODE_KP_2, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum3, sdl.SCANCODE_KP_3, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum4, sdl.SCANCODE_KP_4, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum5, sdl.SCANCODE_KP_5, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum6, sdl.SCANCODE_KP_6, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum7, sdl.SCANCODE_KP_7, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum8, sdl.SCANCODE_KP_8, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyNum9, sdl.SCANCODE_KP_9, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyPageDown, sdl.SCANCODE_PAGEDOWN, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyPageUp, sdl.SCANCODE_PAGEUP, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyPause, sdl.SCANCODE_PAUSE, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyPeriod, sdl.SCANCODE_PERIOD, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyRightBracket, sdl.SCANCODE_RIGHTBRACKET, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyLeftBracket, sdl.SCANCODE_LEFTBRACKET, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeySemicolon, sdl.SCANCODE_SEMICOLON, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeySlash, sdl.SCANCODE_SLASH, keymap); err != nil {
		return nil, err
	}
	if err := addKeymapKey(primitives.KeyTab, sdl.SCANCODE_TAB, keymap); err != nil {
		return nil, err
	}
	return keymap, nil
}

func addKeymapKey(key primitives.KeyType, sdlScanCode uint8, keymap map[primitives.KeyType]uint8) error {
	_, keyFound := keymap[key]
	if keyFound {
		return errors.New("Key " + fmt.Sprint(key) + " already in keymap")
	}
	keymap[key] = sdlScanCode
	return nil
}
