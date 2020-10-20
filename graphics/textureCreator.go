package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
	"github.com/veandco/go-sdl2/sdl"
)

var _ graphics.TextureCreator = (*textureCreator)(nil)

type textureCreator struct {
	renderer *sdl.Renderer
}

func (creator textureCreator) Create(dimensions *primitives.Dimensions, color *primitives.Color) (graphics.Texture, error) {
	texture, err := creator.renderer.CreateTexture(sdl.PIXELFORMAT_RGBA8888, sdl.TEXTUREACCESS_STATIC, int32(dimensions.Width), int32(dimensions.Height))
	if err != nil {
		return nil, err
	}
	pixelsLength := dimensions.Height * dimensions.Width * 4
	pixels := make([]uint8, pixelsLength)
	for i := 0; i < pixelsLength; i = i + 4 {
		pixels[i] = color.R
		pixels[i+1] = color.G
		pixels[i+2] = color.B
		pixels[i+3] = color.A
	}
	err = texture.Update(nil, pixels, dimensions.Width*4)
	if err != nil {
		return nil, err
	}
	drawer := textureDrawer{
		sdlTexture: texture,
		renderer:   creator.renderer,
	}
	return newTexture(drawer, *dimensions), nil
}
