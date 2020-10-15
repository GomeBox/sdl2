package input

import (
	"fmt"
	"github.com/GomeBox/gome/adapters/input"
	"testing"
)

func TestKeyPressed(t *testing.T) {
	var keys = []input.KeyType{
		input.KeyA,
		input.KeyB,
		input.KeyC,
		input.KeyD,
		input.KeyE,
		input.KeyF,
		input.KeyG,
		input.KeyH,
		input.KeyI,
		input.KeyJ,
		input.KeyK,
		input.KeyL,
		input.KeyM,
		input.KeyN,
		input.KeyO,
		input.KeyP,
		input.KeyQ,
		input.KeyR,
		input.KeyS,
		input.KeyT,
		input.KeyU,
		input.KeyV,
		input.KeyW,
		input.KeyX,
		input.KeyY,
		input.KeyZ,
		input.KeyUp,
		input.KeyDown,
		input.KeyLeft,
		input.KeyRight,
		input.Key0,
		input.Key1,
		input.Key2,
		input.Key3,
		input.Key4,
		input.Key5,
		input.Key6,
		input.Key7,
		input.Key8,
		input.Key9,
		input.KeyLeftCtrl,
		input.KeyRightAlt,
		input.KeyRightAlt,
		input.KeyLeftAlt,
		input.KeySpace,
		input.KeyReturn,
		input.KeyEsc,
		input.KeyBackslash,
		input.KeyBackspace,
		input.KeyCapslock,
		input.KeyComma,
		input.KeyDelete,
		input.KeyEnd,
		input.KeyEquals,
		input.KeyF1,
		input.KeyF2,
		input.KeyF3,
		input.KeyF4,
		input.KeyF5,
		input.KeyF6,
		input.KeyF7,
		input.KeyF8,
		input.KeyF9,
		input.KeyF10,
		input.KeyF11,
		input.KeyF12,
		input.KeyF13,
		input.KeyHome,
		input.KeyInsert,
		input.KeyNum0,
		input.KeyNum1,
		input.KeyNum2,
		input.KeyNum3,
		input.KeyNum4,
		input.KeyNum5,
		input.KeyNum6,
		input.KeyNum7,
		input.KeyNum8,
		input.KeyNum9,
		input.KeyPageDown,
		input.KeyPageUp,
		input.KeyPause,
		input.KeyPeriod,
		input.KeyRightBracket,
		input.KeyLeftBracket,
		input.KeySemicolon,
		input.KeySlash,
		input.KeyTab,
	}
	keyboard, err := newKeyboard()
	if err != nil {
		t.Error(err)
		return
	}
	for _, key := range keys {
		_, err := keyboard.KeyPressed(key)
		if err != nil {
			t.Errorf("Key " + fmt.Sprint(key) + " not handled")
		}
	}
}
