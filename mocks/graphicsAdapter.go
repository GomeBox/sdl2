package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
)

type GraphicsAdapter struct {
	InitCnt int
}

func (adapter *GraphicsAdapter) Init() error {
	adapter.InitCnt++
	return nil
}

func (adapter *GraphicsAdapter) ShowWindow(windowSettings *graphics.WindowSettings) error {
	return nil
}

func (adapter *GraphicsAdapter) TextureLoader() graphics.TextureLoader {
	return adapter
}

func (adapter *GraphicsAdapter) WindowAdapter() graphics.WindowAdapter {
	return adapter
}

func (adapter *GraphicsAdapter) Load(fileName string) (graphics.TextureDrawer, error) {
	return nil, nil
}
