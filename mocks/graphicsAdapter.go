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

func (adapter *GraphicsAdapter) ScreenPresenter() graphics.ScreenPresenter {
	return adapter
}

func (adapter *GraphicsAdapter) FontLoader() graphics.FontLoader {
	return new(FontLoader)
}

func (adapter *GraphicsAdapter) Load(fileName string) (graphics.TextureDrawer, error) {
	return nil, nil
}

func (adapter *GraphicsAdapter) Present() error {
	return nil
}

func (adapter *GraphicsAdapter) Clear() error {
	return nil
}
