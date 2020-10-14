package sdl2

import (
	"github.com/GomeBox/sdl2/mocks"
	"testing"
)

func TestNewSystem(t *testing.T) {
	system := NewSystem()
	if system == nil {
		t.Errorf("NewSystem returned nil")
		return
	}
	if system.input == nil {
		t.Errorf("system.input is nil")
		return
	}
	if system.graphics == nil {
		t.Errorf("system.graphics is nil")
	}
}

func TestInitialize(t *testing.T) {
	system := NewSystem()
	inputMock := new(mocks.InputAdapter)
	graphicsMock := new(mocks.GraphicsAdapter)
	system.input = inputMock
	system.graphics = graphicsMock
	err := system.Initialize()
	if err != nil {
		t.Errorf("system.Initialize() returned unexpected error '%s'", err.Error())
		return
	}
	want := 1
	got := inputMock.InitCnt
	if got != want {
		t.Errorf("input.Init() was not called correct number of times. Got: %d,  want: %d", got, want)
		return
	}
	got = graphicsMock.InitCnt
	if got != want {
		t.Errorf("graphics.Init() was not called correct number of times. Got: %d,  want: %d", got, want)
		return
	}
}
