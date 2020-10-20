package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
	"github.com/veandco/go-sdl2/sdl"
)

var _ graphics.TextureDrawer = (*textureDrawer)(nil)

type textureDrawer struct {
	sdlTexture *sdl.Texture
	renderer   *sdl.Renderer
}

func (drawer textureDrawer) Draw(source, dest *primitives.Rectangle) error {
	destRect := ConvertRect(dest)
	srcRect := ConvertRect(source)
	err := drawer.renderer.Copy(drawer.sdlTexture, srcRect, destRect)
	if err != nil {
		return err
	}
	return nil
}

func (drawer textureDrawer) DrawF(source *primitives.Rectangle, dest *primitives.RectangleF) error {
	destRect := ConvertRectF(dest)
	srcRect := ConvertRect(source)
	err := drawer.renderer.CopyF(drawer.sdlTexture, srcRect, destRect)
	if err != nil {
		return err
	}
	return nil
}

func ConvertRect(rect *primitives.Rectangle) *sdl.Rect {
	if rect == nil {
		return nil
	}
	return &sdl.Rect{X: int32(rect.X), Y: int32(rect.Y), H: int32(rect.Height), W: int32(rect.Width)}
}

func ConvertRectF(rect *primitives.RectangleF) *sdl.FRect {
	if rect == nil {
		return nil
	}
	return &sdl.FRect{X: rect.X, Y: rect.Y, H: rect.Height, W: rect.Width}
}
