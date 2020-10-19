package mocks

import "github.com/GomeBox/gome/adapters/input"

type InputAdapter struct {
	InitCnt int
}

func (adapter *InputAdapter) Init() error {
	adapter.InitCnt++
	return nil
}

func (adapter *InputAdapter) Update() {
}

func (adapter *InputAdapter) Keyboard() input.KeyboardAdapter {
	return nil
}

func (adapter *InputAdapter) ControllerCount() int {
	return 0
}

func (adapter *InputAdapter) Controller(number int) (*input.Controller, error) {
	return nil, nil
}
