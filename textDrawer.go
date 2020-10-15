package sdl2

import (
	"github.com/GomeBox/gome/primitives"
)

type TextDrawer struct {
	textureDrawer textureDrawer
}

func NewTextDrawer(textureDrawer textureDrawer) *TextDrawer {
	drawer := new(TextDrawer)
	drawer.textureDrawer = textureDrawer
	return drawer
}

func (drawer *TextDrawer) Draw(position primitives.Point) error {
	destRect := &primitives.Rectangle{
		Point:      primitives.Point{X: position.X, Y: position.Y},
		Dimensions: primitives.Dimensions{Height: drawer.textureDrawer.height, Width: drawer.textureDrawer.width}}
	return drawer.draw(destRect)
}

func (drawer *TextDrawer) DrawScaled(destRect primitives.Rectangle) error {
	return drawer.draw(&destRect)
}

func (drawer *TextDrawer) draw(destRect *primitives.Rectangle) error {
	err := drawer.textureDrawer.Draw(nil, destRect)
	if err != nil {
		return err
	}
	return nil
}
