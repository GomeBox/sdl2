package graphics

import (
	"github.com/GomeBox/gome/primitives"
)

type texture struct {
	textureDrawer
	dimensions primitives.Dimensions
}

func newTexture(drawer textureDrawer, dimensions primitives.Dimensions) *texture {
	texture := new(texture)
	texture.textureDrawer = drawer
	texture.dimensions = dimensions
	return texture
}

func (texture *texture) Dimensions() primitives.Dimensions {
	return texture.dimensions
}
