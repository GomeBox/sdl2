package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type FontLoader interface {
	Load(fileName string, size int) (graphics.TextCreator, error)
}

type fontLoader struct {
	renderer *sdl.Renderer
}

func NewFontLoader(renderer *sdl.Renderer) FontLoader {
	fontLoader := new(fontLoader)
	fontLoader.renderer = renderer
	return fontLoader
}

func (fontLoader *fontLoader) Load(fileName string, size int) (graphics.TextCreator, error) {
	font, err := ttf.OpenFont(fileName, size)
	if err != nil {
		return nil, err
	}
	textCreator := &TextCreator{
		font:     font,
		renderer: fontLoader.renderer,
	}
	return textCreator, nil
}
