package mocks

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

type GraphicsAdapter struct {
	InitCnt int
}

func (adapter *GraphicsAdapter) Init() error {
	adapter.InitCnt++
	return nil
}

func (adapter *GraphicsAdapter) OpenWindow(windowSettings *graphics.WindowSettings) error {
	return nil
}

func (adapter *GraphicsAdapter) Size() primitives.Dimensions {
	return primitives.Dimensions{}
}

func (adapter *GraphicsAdapter) IsOpen() bool {
	return false
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

func (adapter *GraphicsAdapter) TextureCreator() graphics.TextureCreator {
	return adapter
}

func (adapter *GraphicsAdapter) Create(dimensions *primitives.Dimensions, color *primitives.Color) (graphics.Texture, error) {
	return nil, nil
}

func (adapter *GraphicsAdapter) Load(fileName string) (graphics.Texture, error) {
	return nil, nil
}

func (adapter *GraphicsAdapter) Present() error {
	return nil
}

func (adapter *GraphicsAdapter) Clear() error {
	return nil
}
