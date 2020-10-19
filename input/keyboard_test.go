package input

import (
	"fmt"
	"github.com/GomeBox/gome/primitives"
	"testing"
)

func TestKeyPressed(t *testing.T) {
	var keys = []primitives.KeyType{
		primitives.KeyA,
		primitives.KeyB,
		primitives.KeyC,
		primitives.KeyD,
		primitives.KeyE,
		primitives.KeyF,
		primitives.KeyG,
		primitives.KeyH,
		primitives.KeyI,
		primitives.KeyJ,
		primitives.KeyK,
		primitives.KeyL,
		primitives.KeyM,
		primitives.KeyN,
		primitives.KeyO,
		primitives.KeyP,
		primitives.KeyQ,
		primitives.KeyR,
		primitives.KeyS,
		primitives.KeyT,
		primitives.KeyU,
		primitives.KeyV,
		primitives.KeyW,
		primitives.KeyX,
		primitives.KeyY,
		primitives.KeyZ,
		primitives.KeyUp,
		primitives.KeyDown,
		primitives.KeyLeft,
		primitives.KeyRight,
		primitives.Key0,
		primitives.Key1,
		primitives.Key2,
		primitives.Key3,
		primitives.Key4,
		primitives.Key5,
		primitives.Key6,
		primitives.Key7,
		primitives.Key8,
		primitives.Key9,
		primitives.KeyLeftCtrl,
		primitives.KeyRightAlt,
		primitives.KeyRightAlt,
		primitives.KeyLeftAlt,
		primitives.KeySpace,
		primitives.KeyReturn,
		primitives.KeyEsc,
		primitives.KeyBackslash,
		primitives.KeyBackspace,
		primitives.KeyCapslock,
		primitives.KeyComma,
		primitives.KeyDelete,
		primitives.KeyEnd,
		primitives.KeyEquals,
		primitives.KeyF1,
		primitives.KeyF2,
		primitives.KeyF3,
		primitives.KeyF4,
		primitives.KeyF5,
		primitives.KeyF6,
		primitives.KeyF7,
		primitives.KeyF8,
		primitives.KeyF9,
		primitives.KeyF10,
		primitives.KeyF11,
		primitives.KeyF12,
		primitives.KeyF13,
		primitives.KeyHome,
		primitives.KeyInsert,
		primitives.KeyNum0,
		primitives.KeyNum1,
		primitives.KeyNum2,
		primitives.KeyNum3,
		primitives.KeyNum4,
		primitives.KeyNum5,
		primitives.KeyNum6,
		primitives.KeyNum7,
		primitives.KeyNum8,
		primitives.KeyNum9,
		primitives.KeyPageDown,
		primitives.KeyPageUp,
		primitives.KeyPause,
		primitives.KeyPeriod,
		primitives.KeyRightBracket,
		primitives.KeyLeftBracket,
		primitives.KeySemicolon,
		primitives.KeySlash,
		primitives.KeyTab,
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
