package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

type TextCreator struct {
	font     *ttf.Font
	renderer *sdl.Renderer
}

func (creator *TextCreator) Create(text string, color primitives.Color) (graphics.TextDrawer, error) {
	sdlColor := sdl.Color{R: color.R, G: color.G, B: color.B, A: color.A}

	fontSurface, err := creator.font.RenderUTF8Blended(text, sdlColor)
	if err != nil {
		return nil, err
	}

	texture, err := creator.renderer.CreateTextureFromSurface(fontSurface)
	fontSurface.Free()
	if err != nil {
		return nil, err
	}

	_, _, w, h, err := texture.Query()
	if err != nil {
		return nil, err
	}

	textureDrawer := textureDrawer{
		sdlTexture: texture,
		renderer:   creator.renderer,
		width:      int(w),
		height:     int(h),
	}

	drawer := NewTextDrawer(textureDrawer)
	return drawer, nil
}
