package graphics

import (
	"github.com/GomeBox/gome/adapters/graphics"
	"github.com/GomeBox/gome/primitives"
)

var _ graphics.TextDrawer = (*TextDrawer)(nil)

type TextDrawer struct {
	textureDrawer textureDrawer
}

func NewTextDrawer(textureDrawer textureDrawer) *TextDrawer {
	drawer := new(TextDrawer)
	drawer.textureDrawer = textureDrawer
	return drawer
}

func (drawer *TextDrawer) Draw(position *primitives.Point) error {
	destRect := &primitives.Rectangle{
		Point:      primitives.Point{X: position.X, Y: position.Y},
		Dimensions: primitives.Dimensions{Height: drawer.textureDrawer.height, Width: drawer.textureDrawer.width}}
	return drawer.draw(destRect)
}

func (drawer *TextDrawer) DrawScaled(destRect *primitives.Rectangle) error {
	return drawer.draw(destRect)
}

func (drawer *TextDrawer) DrawF(position *primitives.PointF) error {
	destRect := &primitives.RectangleF{
		PointF:      primitives.PointF{X: position.X, Y: position.Y},
		DimensionsF: primitives.DimensionsF{Height: float32(drawer.textureDrawer.height), Width: float32(drawer.textureDrawer.width)}}
	return drawer.drawF(destRect)
}

func (drawer *TextDrawer) DrawScaledF(destRect *primitives.RectangleF) error {
	return drawer.drawF(destRect)
}

func (drawer *TextDrawer) draw(destRect *primitives.Rectangle) error {
	err := drawer.textureDrawer.Draw(nil, destRect)
	if err != nil {
		return err
	}
	return nil
}

func (drawer *TextDrawer) drawF(destRect *primitives.RectangleF) error {
	err := drawer.textureDrawer.DrawF(nil, destRect)
	if err != nil {
		return err
	}
	return nil
}
